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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/cloud-ark/kubeplus/platform-operator/pkg/apis/workflowcontroller/v1alpha1"
	scheme "github.com/cloud-ark/kubeplus/platform-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ResourceMonitorsGetter has a method to return a ResourceMonitorInterface.
// A group's client should implement this interface.
type ResourceMonitorsGetter interface {
	ResourceMonitors(namespace string) ResourceMonitorInterface
}

// ResourceMonitorInterface has methods to work with ResourceMonitor resources.
type ResourceMonitorInterface interface {
	Create(ctx context.Context, resourceMonitor *v1alpha1.ResourceMonitor, opts v1.CreateOptions) (*v1alpha1.ResourceMonitor, error)
	Update(ctx context.Context, resourceMonitor *v1alpha1.ResourceMonitor, opts v1.UpdateOptions) (*v1alpha1.ResourceMonitor, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ResourceMonitor, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ResourceMonitorList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ResourceMonitor, err error)
	ResourceMonitorExpansion
}

// resourceMonitors implements ResourceMonitorInterface
type resourceMonitors struct {
	client rest.Interface
	ns     string
}

// newResourceMonitors returns a ResourceMonitors
func newResourceMonitors(c *WorkflowsV1alpha1Client, namespace string) *resourceMonitors {
	return &resourceMonitors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the resourceMonitor, and returns the corresponding resourceMonitor object, and an error if there is any.
func (c *resourceMonitors) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ResourceMonitor, err error) {
	result = &v1alpha1.ResourceMonitor{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourcemonitors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ResourceMonitors that match those selectors.
func (c *resourceMonitors) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ResourceMonitorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ResourceMonitorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourcemonitors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested resourceMonitors.
func (c *resourceMonitors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("resourcemonitors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a resourceMonitor and creates it.  Returns the server's representation of the resourceMonitor, and an error, if there is any.
func (c *resourceMonitors) Create(ctx context.Context, resourceMonitor *v1alpha1.ResourceMonitor, opts v1.CreateOptions) (result *v1alpha1.ResourceMonitor, err error) {
	result = &v1alpha1.ResourceMonitor{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("resourcemonitors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourceMonitor).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a resourceMonitor and updates it. Returns the server's representation of the resourceMonitor, and an error, if there is any.
func (c *resourceMonitors) Update(ctx context.Context, resourceMonitor *v1alpha1.ResourceMonitor, opts v1.UpdateOptions) (result *v1alpha1.ResourceMonitor, err error) {
	result = &v1alpha1.ResourceMonitor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourcemonitors").
		Name(resourceMonitor.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourceMonitor).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the resourceMonitor and deletes it. Returns an error if one occurs.
func (c *resourceMonitors) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourcemonitors").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *resourceMonitors) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourcemonitors").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched resourceMonitor.
func (c *resourceMonitors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ResourceMonitor, err error) {
	result = &v1alpha1.ResourceMonitor{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("resourcemonitors").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
