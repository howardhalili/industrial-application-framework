// Copyright 2020 Nokia
// Licensed under the BSD 3-Clause License.
// SPDX-License-Identifier: BSD-3-Clause

package template

import (
	"bytes"
	"github.com/nokia/industrial-application-framework/application-lib/pkg/config"
	copy2 "github.com/nokia/industrial-application-framework/application-lib/pkg/util/copy"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	templ "text/template"

	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

type Templater struct {
	Data          interface{}
	Namespace     string
	DeploymentDir string
	DirName       string
	SourceDir     string
	WorkDir       string
	Delimiters    config.TemplateConfig
}

var log = logf.Log.WithName("template_controller")

func NewTemplater(data interface{}, namespace string, deploymentDir string, dirName string, templateConfig config.TemplateConfig) (*Templater, error) {
	t := &Templater{
		Data:       data,
		Namespace:  namespace,
		DirName:    dirName,
		Delimiters: templateConfig,
	}
	if t.DeploymentDir = deploymentDir; t.DeploymentDir == "" {
		return nil, errors.New("deploymentDir is not set")
	}

	if (config.TemplateConfig{}) == templateConfig {
		t.Delimiters = config.TemplateConfig{
			LeftDelimiter:  "[[",
			RightDelimiter: "]]",
		}

	}

	t.SourceDir = filepath.Join(t.DeploymentDir, t.DirName)
	t.WorkDir = filepath.Join(t.DeploymentDir, t.DirName+"-generated")

	if err := t.copyDeploymentYamls(); err != nil {
		return nil, err
	}

	return t, nil
}

func (t *Templater) RunCrTemplater(joinSeparator string) (string, error) {
	out, err := t.runGoTemplate(joinSeparator)
	if err != nil {
		return "", errors.Wrap(err, "failed to run the Cr Templater")
	}

	return out, nil
}

func (t *Templater) copyDeploymentYamls() error {
	os.RemoveAll(t.WorkDir)

	if err := copy2.CopyDir(t.SourceDir, t.WorkDir); err != nil {
		return err
	}

	return nil
}

func (t *Templater) runGoTemplate(joinSeparator string) (string, error) {
	log.Info("Running go templates")

	return t.templateDir(t.WorkDir, joinSeparator)
}

func (t *Templater) templateFile(workDir string, file os.FileInfo, joinSeparator string) (string, error) {
	if strings.Contains(file.Name(), "yaml") {
		log.Info("templating", "file", file.Name())

		template, err := templ.New(file.Name()).Delims(t.Delimiters.LeftDelimiter, t.Delimiters.RightDelimiter).ParseFiles(filepath.Join(workDir, file.Name()))
		if err != nil {
			return "", err
		}

		buffer := bytes.NewBuffer([]byte{})

		err = template.Execute(buffer, t.Data)
		if err != nil {
			return "", errors.Wrap(err, "failed to execute template")
		}

		newFile, err := os.Create(filepath.Join(workDir, file.Name()))
		if err != nil {
			return "", errors.Wrap(err, "failed to create file: "+file.Name())
		}

		newFile.Write(buffer.Bytes())
		newFile.Close()

		return joinSeparator + string(buffer.Bytes()), nil
	}
	return "", nil
}

func (t *Templater) templateDir(workDir string, joinSeparator string) (string, error) {
	log.Info("templating", "dir", workDir)

	files, err := ioutil.ReadDir(workDir)
	if err != nil {

		return "", errors.Wrap(err, "failed to read dir")
	}

	out := ""
	for _, file := range files {
		res := ""
		if file.IsDir() {
			res, err = t.templateDir(filepath.Join(workDir, file.Name()), joinSeparator)
		} else {
			res, err = t.templateFile(workDir, file, joinSeparator)
		}
		if err != nil {
			return "", err
		}
		out += res
	}

	return out, nil
}
