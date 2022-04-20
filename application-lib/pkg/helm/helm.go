// Copyright 2020 Nokia
// Licensed under the BSD 3-Clause License.
// SPDX-License-Identifier: BSD-3-Clause

package helm

import (
	"context"
	"os/exec"
	"time"

	"github.com/pkg/errors"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

type Helm struct {
	namespace string
	WorkDir   string
	timeout   time.Duration
}

const (
	defaultExecutionTimeout time.Duration = 30 * time.Second
	ReleaseName                           = "app-release"
	FlagNamespace                         = "--namespace"
)

var log = logf.Log.WithName("helm_controller")

func NewHelm(namespace string, deploymentDir string, executionTimeout *time.Duration) *Helm {
	actualTimeout := defaultExecutionTimeout
	if executionTimeout != nil {
		actualTimeout = *executionTimeout
	}

	return &Helm{
		namespace: namespace,
		WorkDir:   deploymentDir + "/app-deployment-generated",
		timeout:   actualTimeout,
	}
}

func (h *Helm) execCommand(args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "helm", args...)
	cmd.Dir = h.WorkDir

	out, err := cmd.CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		return "", errors.New("command timed out")
	}
	if err != nil {
		return "", errors.Wrapf(err, "command failed, output: %v", string(out))
	}

	log.Info("cmd output", "output", string(out))
	log.Info("command successfully executed")

	return string(out), nil
}

func (h *Helm) getRelease() (string, error) {
	out, err := h.execCommand("list", "-q", FlagNamespace, h.namespace)

	return string(out), err
}

func (h *Helm) install() error {
	_, err := h.execCommand("install", ReleaseName, FlagNamespace, h.namespace, ".")

	return err
}

func (h *Helm) upgrade() error {
	_, err := h.execCommand("upgrade", ReleaseName, FlagNamespace, h.namespace, ".")

	return err
}

func (h *Helm) Deploy() error {
	if release, err := h.getRelease(); err == nil {
		if release == "" {
			return h.install()
		} else {
			return h.upgrade()
		}
	} else {
		return err
	}
}

func (h *Helm) Undeploy() error {
	_, err := h.execCommand("uninstall", ReleaseName, FlagNamespace, h.namespace)

	return err
}
