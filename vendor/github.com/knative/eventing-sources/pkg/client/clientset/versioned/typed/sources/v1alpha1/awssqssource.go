/*
Copyright 2018 The Knative Authors

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

package v1alpha1

import (
	v1alpha1 "github.com/knative/eventing-sources/pkg/apis/sources/v1alpha1"
	scheme "github.com/knative/eventing-sources/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AwsSqsSourcesGetter has a method to return a AwsSqsSourceInterface.
// A group's client should implement this interface.
type AwsSqsSourcesGetter interface {
	AwsSqsSources(namespace string) AwsSqsSourceInterface
}

// AwsSqsSourceInterface has methods to work with AwsSqsSource resources.
type AwsSqsSourceInterface interface {
	Create(*v1alpha1.AwsSqsSource) (*v1alpha1.AwsSqsSource, error)
	Update(*v1alpha1.AwsSqsSource) (*v1alpha1.AwsSqsSource, error)
	UpdateStatus(*v1alpha1.AwsSqsSource) (*v1alpha1.AwsSqsSource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsSqsSource, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsSqsSourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSqsSource, err error)
	AwsSqsSourceExpansion
}

// awsSqsSources implements AwsSqsSourceInterface
type awsSqsSources struct {
	client rest.Interface
	ns     string
}

// newAwsSqsSources returns a AwsSqsSources
func newAwsSqsSources(c *SourcesV1alpha1Client, namespace string) *awsSqsSources {
	return &awsSqsSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsSqsSource, and returns the corresponding awsSqsSource object, and an error if there is any.
func (c *awsSqsSources) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSqsSource, err error) {
	result = &v1alpha1.AwsSqsSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssqssources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsSqsSources that match those selectors.
func (c *awsSqsSources) List(opts v1.ListOptions) (result *v1alpha1.AwsSqsSourceList, err error) {
	result = &v1alpha1.AwsSqsSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssqssources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsSqsSources.
func (c *awsSqsSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awssqssources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsSqsSource and creates it.  Returns the server's representation of the awsSqsSource, and an error, if there is any.
func (c *awsSqsSources) Create(awsSqsSource *v1alpha1.AwsSqsSource) (result *v1alpha1.AwsSqsSource, err error) {
	result = &v1alpha1.AwsSqsSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awssqssources").
		Body(awsSqsSource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsSqsSource and updates it. Returns the server's representation of the awsSqsSource, and an error, if there is any.
func (c *awsSqsSources) Update(awsSqsSource *v1alpha1.AwsSqsSource) (result *v1alpha1.AwsSqsSource, err error) {
	result = &v1alpha1.AwsSqsSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awssqssources").
		Name(awsSqsSource.Name).
		Body(awsSqsSource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsSqsSources) UpdateStatus(awsSqsSource *v1alpha1.AwsSqsSource) (result *v1alpha1.AwsSqsSource, err error) {
	result = &v1alpha1.AwsSqsSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awssqssources").
		Name(awsSqsSource.Name).
		SubResource("status").
		Body(awsSqsSource).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsSqsSource and deletes it. Returns an error if one occurs.
func (c *awsSqsSources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssqssources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsSqsSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssqssources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsSqsSource.
func (c *awsSqsSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSqsSource, err error) {
	result = &v1alpha1.AwsSqsSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awssqssources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
