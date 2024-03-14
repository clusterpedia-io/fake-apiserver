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

package v1beta2

import (
	v1beta2 "k8s.io/api/flowcontrol/v1beta2"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&v1beta2.FlowSchema{}, func(obj interface{}) { SetObjectDefaults_FlowSchema(obj.(*v1beta2.FlowSchema)) })
	scheme.AddTypeDefaultingFunc(&v1beta2.FlowSchemaList{}, func(obj interface{}) { SetObjectDefaults_FlowSchemaList(obj.(*v1beta2.FlowSchemaList)) })
	scheme.AddTypeDefaultingFunc(&v1beta2.PriorityLevelConfiguration{}, func(obj interface{}) {
		SetObjectDefaults_PriorityLevelConfiguration(obj.(*v1beta2.PriorityLevelConfiguration))
	})
	scheme.AddTypeDefaultingFunc(&v1beta2.PriorityLevelConfigurationList{}, func(obj interface{}) {
		SetObjectDefaults_PriorityLevelConfigurationList(obj.(*v1beta2.PriorityLevelConfigurationList))
	})
	return nil
}

func SetObjectDefaults_FlowSchema(in *v1beta2.FlowSchema) {
	SetDefaults_FlowSchemaSpec(&in.Spec)
}

func SetObjectDefaults_FlowSchemaList(in *v1beta2.FlowSchemaList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_FlowSchema(a)
	}
}

func SetObjectDefaults_PriorityLevelConfiguration(in *v1beta2.PriorityLevelConfiguration) {
	if in.Spec.Limited != nil {
		SetDefaults_LimitedPriorityLevelConfiguration(in.Spec.Limited)
		if in.Spec.Limited.LimitResponse.Queuing != nil {
			SetDefaults_QueuingConfiguration(in.Spec.Limited.LimitResponse.Queuing)
		}
	}
	if in.Spec.Exempt != nil {
		SetDefaults_ExemptPriorityLevelConfiguration(in.Spec.Exempt)
	}
}

func SetObjectDefaults_PriorityLevelConfigurationList(in *v1beta2.PriorityLevelConfigurationList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_PriorityLevelConfiguration(a)
	}
}
