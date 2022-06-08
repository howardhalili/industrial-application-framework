// +build !ignore_autogenerated

// Copyright 2020 Nokia
// Licensed under the BSD 3-Clause License.
// SPDX-License-Identifier: BSD-3-Clause

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/nokia/industrial-application-framework/consul-backup/pkg/k8sdynamic"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppPodFixIp) DeepCopyInto(out *AppPodFixIp) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppPodFixIp.
func (in *AppPodFixIp) DeepCopy() *AppPodFixIp {
	if in == nil {
		return nil
	}
	out := new(AppPodFixIp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppReporteData) DeepCopyInto(out *AppReporteData) {
	*out = *in
	if in.PrivateNetworkIpAddress != nil {
		in, out := &in.PrivateNetworkIpAddress, &out.PrivateNetworkIpAddress
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppReporteData.
func (in *AppReporteData) DeepCopy() *AppReporteData {
	if in == nil {
		return nil
	}
	out := new(AppReporteData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Consul) DeepCopyInto(out *Consul) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Consul.
func (in *Consul) DeepCopy() *Consul {
	if in == nil {
		return nil
	}
	out := new(Consul)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Consul) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsulList) DeepCopyInto(out *ConsulList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Consul, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsulList.
func (in *ConsulList) DeepCopy() *ConsulList {
	if in == nil {
		return nil
	}
	out := new(ConsulList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsulList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsulSpec) DeepCopyInto(out *ConsulSpec) {
	*out = *in
	out.Ports = in.Ports
	if in.PrivateNetworkAccess != nil {
		in, out := &in.PrivateNetworkAccess, &out.PrivateNetworkAccess
		*out = new(PrivateNetworkAccess)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsulSpec.
func (in *ConsulSpec) DeepCopy() *ConsulSpec {
	if in == nil {
		return nil
	}
	out := new(ConsulSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsulStatus) DeepCopyInto(out *ConsulStatus) {
	*out = *in
	if in.PrevSpec != nil {
		in, out := &in.PrevSpec, &out.PrevSpec
		*out = new(ConsulSpec)
		(*in).DeepCopyInto(*out)
	}
	in.AppReportedData.DeepCopyInto(&out.AppReportedData)
	if in.AppliedResources != nil {
		in, out := &in.AppliedResources, &out.AppliedResources
		*out = make([]k8sdynamic.ResourceDescriptor, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsulStatus.
func (in *ConsulStatus) DeepCopy() *ConsulStatus {
	if in == nil {
		return nil
	}
	out := new(ConsulStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	if in.AdditionalRoutes != nil {
		in, out := &in.AdditionalRoutes, &out.AdditionalRoutes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ports) DeepCopyInto(out *Ports) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ports.
func (in *Ports) DeepCopy() *Ports {
	if in == nil {
		return nil
	}
	out := new(Ports)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateNetworkAccess) DeepCopyInto(out *PrivateNetworkAccess) {
	*out = *in
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]Network, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdditionalRoutes != nil {
		in, out := &in.AdditionalRoutes, &out.AdditionalRoutes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AppPodFixIp != nil {
		in, out := &in.AppPodFixIp, &out.AppPodFixIp
		*out = new(AppPodFixIp)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateNetworkAccess.
func (in *PrivateNetworkAccess) DeepCopy() *PrivateNetworkAccess {
	if in == nil {
		return nil
	}
	out := new(PrivateNetworkAccess)
	in.DeepCopyInto(out)
	return out
}
