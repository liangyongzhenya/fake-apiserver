//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by defaulter-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/admissionregistration/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&v1.MutatingWebhookConfiguration{}, func(obj interface{}) {
		SetObjectDefaults_MutatingWebhookConfiguration(obj.(*v1.MutatingWebhookConfiguration))
	})
	scheme.AddTypeDefaultingFunc(&v1.MutatingWebhookConfigurationList{}, func(obj interface{}) {
		SetObjectDefaults_MutatingWebhookConfigurationList(obj.(*v1.MutatingWebhookConfigurationList))
	})
	scheme.AddTypeDefaultingFunc(&v1.ValidatingWebhookConfiguration{}, func(obj interface{}) {
		SetObjectDefaults_ValidatingWebhookConfiguration(obj.(*v1.ValidatingWebhookConfiguration))
	})
	scheme.AddTypeDefaultingFunc(&v1.ValidatingWebhookConfigurationList{}, func(obj interface{}) {
		SetObjectDefaults_ValidatingWebhookConfigurationList(obj.(*v1.ValidatingWebhookConfigurationList))
	})
	return nil
}

func SetObjectDefaults_MutatingWebhookConfiguration(in *v1.MutatingWebhookConfiguration) {
	for i := range in.Webhooks {
		a := &in.Webhooks[i]
		SetDefaults_MutatingWebhook(a)
		if a.ClientConfig.Service != nil {
			SetDefaults_ServiceReference(a.ClientConfig.Service)
		}
		for j := range a.Rules {
			b := &a.Rules[j]
			SetDefaults_Rule(&b.Rule)
		}
	}
}

func SetObjectDefaults_MutatingWebhookConfigurationList(in *v1.MutatingWebhookConfigurationList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_MutatingWebhookConfiguration(a)
	}
}

func SetObjectDefaults_ValidatingWebhookConfiguration(in *v1.ValidatingWebhookConfiguration) {
	for i := range in.Webhooks {
		a := &in.Webhooks[i]
		SetDefaults_ValidatingWebhook(a)
		if a.ClientConfig.Service != nil {
			SetDefaults_ServiceReference(a.ClientConfig.Service)
		}
		for j := range a.Rules {
			b := &a.Rules[j]
			SetDefaults_Rule(&b.Rule)
		}
	}
}

func SetObjectDefaults_ValidatingWebhookConfigurationList(in *v1.ValidatingWebhookConfigurationList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ValidatingWebhookConfiguration(a)
	}
}
