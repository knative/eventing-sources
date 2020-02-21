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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "knative.dev/eventing-contrib/gitlab/pkg/apis/sources/v1alpha1"
)

// FakeGitLabSources implements GitLabSourceInterface
type FakeGitLabSources struct {
	Fake *FakeSourcesV1alpha1
	ns   string
}

var gitlabsourcesResource = schema.GroupVersionResource{Group: "sources.eventing.triggermesh.dev", Version: "v1alpha1", Resource: "gitlabsources"}

var gitlabsourcesKind = schema.GroupVersionKind{Group: "sources.eventing.triggermesh.dev", Version: "v1alpha1", Kind: "GitLabSource"}

// Get takes name of the gitLabSource, and returns the corresponding gitLabSource object, and an error if there is any.
func (c *FakeGitLabSources) Get(name string, options v1.GetOptions) (result *v1alpha1.GitLabSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gitlabsourcesResource, c.ns, name), &v1alpha1.GitLabSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GitLabSource), err
}

// List takes label and field selectors, and returns the list of GitLabSources that match those selectors.
func (c *FakeGitLabSources) List(opts v1.ListOptions) (result *v1alpha1.GitLabSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gitlabsourcesResource, gitlabsourcesKind, c.ns, opts), &v1alpha1.GitLabSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GitLabSourceList{ListMeta: obj.(*v1alpha1.GitLabSourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.GitLabSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gitLabSources.
func (c *FakeGitLabSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gitlabsourcesResource, c.ns, opts))

}

// Create takes the representation of a gitLabSource and creates it.  Returns the server's representation of the gitLabSource, and an error, if there is any.
func (c *FakeGitLabSources) Create(gitLabSource *v1alpha1.GitLabSource) (result *v1alpha1.GitLabSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gitlabsourcesResource, c.ns, gitLabSource), &v1alpha1.GitLabSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GitLabSource), err
}

// Update takes the representation of a gitLabSource and updates it. Returns the server's representation of the gitLabSource, and an error, if there is any.
func (c *FakeGitLabSources) Update(gitLabSource *v1alpha1.GitLabSource) (result *v1alpha1.GitLabSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gitlabsourcesResource, c.ns, gitLabSource), &v1alpha1.GitLabSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GitLabSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGitLabSources) UpdateStatus(gitLabSource *v1alpha1.GitLabSource) (*v1alpha1.GitLabSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(gitlabsourcesResource, "status", c.ns, gitLabSource), &v1alpha1.GitLabSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GitLabSource), err
}

// Delete takes name of the gitLabSource and deletes it. Returns an error if one occurs.
func (c *FakeGitLabSources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(gitlabsourcesResource, c.ns, name), &v1alpha1.GitLabSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGitLabSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gitlabsourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GitLabSourceList{})
	return err
}

// Patch applies the patch and returns the patched gitLabSource.
func (c *FakeGitLabSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GitLabSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gitlabsourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.GitLabSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GitLabSource), err
}
