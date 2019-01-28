/*
Copyright 2019 The Knative Authors.

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

package v1alpha1

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	duckv1alpha1 "github.com/knative/pkg/apis/duck/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

func TestCloudEventsIngressSourceStatusIsReady(t *testing.T) {
	tests := []struct {
		name string
		s    *CloudEventsIngressSourceStatus
		want bool
	}{{
		name: "uninitialized",
		s:    &CloudEventsIngressSourceStatus{},
		want: false,
	}, {
		name: "initialized",
		s: func() *CloudEventsIngressSourceStatus {
			s := &CloudEventsIngressSourceStatus{}
			s.InitializeConditions()
			return s
		}(),
		want: false,
	}, {
		name: "mark deployed",
		s: func() *CloudEventsIngressSourceStatus {
			s := &CloudEventsIngressSourceStatus{}
			s.InitializeConditions()
			s.MarkDeployed()
			return s
		}(),
		want: false,
	}, {
		name: "mark sink",
		s: func() *CloudEventsIngressSourceStatus {
			s := &CloudEventsIngressSourceStatus{}
			s.InitializeConditions()
			s.MarkSink("uri://example")
			return s
		}(),
		want: false,
	}, {
		name: "mark sink and deployed",
		s: func() *CloudEventsIngressSourceStatus {
			s := &CloudEventsIngressSourceStatus{}
			s.InitializeConditions()
			s.MarkSink("uri://example")
			s.MarkDeployed()
			return s
		}(),
		want: true,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.s.IsReady()
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("%s: unexpected condition (-want, +got) = %v", test.name, diff)
			}
		})
	}
}

func TestCloudEventsIngressSourceStatusGetCondition(t *testing.T) {
	tests := []struct {
		name      string
		s         *CloudEventsIngressSourceStatus
		condQuery duckv1alpha1.ConditionType
		want      *duckv1alpha1.Condition
	}{{
		name:      "uninitialized",
		s:         &CloudEventsIngressSourceStatus{},
		condQuery: CloudEventsIngressSourceConditionReady,
		want:      nil,
	}, {
		name: "initialized",
		s: func() *CloudEventsIngressSourceStatus {
			s := &CloudEventsIngressSourceStatus{}
			s.InitializeConditions()
			return s
		}(),
		condQuery: CloudEventsIngressSourceConditionReady,
		want: &duckv1alpha1.Condition{
			Type:   CloudEventsIngressSourceConditionReady,
			Status: corev1.ConditionUnknown,
		},
	}, {
		name: "mark deployed",
		s: func() *CloudEventsIngressSourceStatus {
			s := &CloudEventsIngressSourceStatus{}
			s.InitializeConditions()
			s.MarkDeployed()
			return s
		}(),
		condQuery: CloudEventsIngressSourceConditionReady,
		want: &duckv1alpha1.Condition{
			Type:   CloudEventsIngressSourceConditionReady,
			Status: corev1.ConditionUnknown,
		},
	}, {
		name: "mark sink",
		s: func() *CloudEventsIngressSourceStatus {
			s := &CloudEventsIngressSourceStatus{}
			s.InitializeConditions()
			s.MarkSink("uri://example")
			return s
		}(),
		condQuery: CloudEventsIngressSourceConditionReady,
		want: &duckv1alpha1.Condition{
			Type:   CloudEventsIngressSourceConditionReady,
			Status: corev1.ConditionUnknown,
		},
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.s.GetCondition(test.condQuery)
			ignoreTime := cmpopts.IgnoreFields(duckv1alpha1.Condition{},
				"LastTransitionTime", "Severity")
			if diff := cmp.Diff(test.want, got, ignoreTime); diff != "" {
				t.Errorf("unexpected condition (-want, +got) = %v", diff)
			}
		})
	}
}
