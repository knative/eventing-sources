/*
Copyright 2020 The Knative Authors

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

// Code generated by injection-gen. DO NOT EDIT.

package gitlabsource

import (
	context "context"

	v1 "k8s.io/api/core/v1"
	v1alpha1 "knative.dev/eventing-contrib/gitlab/pkg/apis/sources/v1alpha1"
	gitlabsource "knative.dev/eventing-contrib/gitlab/pkg/client/injection/reconciler/sources/v1alpha1/gitlabsource"
	reconciler "knative.dev/pkg/reconciler"
)

// TODO: PLEASE COPY AND MODIFY THIS FILE AS A STARTING POINT

// newReconciledNormal makes a new reconciler event with event type Normal, and
// reason GitLabSourceReconciled.
func newReconciledNormal(namespace, name string) reconciler.Event {
	return reconciler.NewEvent(v1.EventTypeNormal, "GitLabSourceReconciled", "GitLabSource reconciled: \"%s/%s\"", namespace, name)
}

// Reconciler implements controller.Reconciler for GitLabSource resources.
type Reconciler struct {
	// TODO: add additional requirements here.
}

// Check that our Reconciler implements Interface
var _ gitlabsource.Interface = (*Reconciler)(nil)

// Optionally check that our Reconciler implements Finalizer
//var _ gitlabsource.Finalizer = (*Reconciler)(nil)

// ReconcileKind implements Interface.ReconcileKind.
func (r *Reconciler) ReconcileKind(ctx context.Context, o *v1alpha1.GitLabSource) reconciler.Event {
	o.Status.InitializeConditions()

	// TODO: add custom reconciliation logic here.

	o.Status.ObservedGeneration = o.Generation
	return newReconciledNormal(o.Namespace, o.Name)
}

// Optionally, use FinalizeKind to add finalizers. FinalizeKind will be called
// when the resource is deleted.
//func (r *Reconciler) FinalizeKind(ctx context.Context, o *v1alpha1.GitLabSource) reconciler.Event {
//	// TODO: add custom finalization logic here.
//	return nil
//}
