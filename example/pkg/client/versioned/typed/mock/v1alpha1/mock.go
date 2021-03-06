/*

MIT License

Copyright (c) 2019 Fumihiro Ito

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "go.f110.dev/rules_k8s_controller/example/pkg/api/mock/v1alpha1"
	scheme "go.f110.dev/rules_k8s_controller/example/pkg/client/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MocksGetter has a method to return a MockInterface.
// A group's client should implement this interface.
type MocksGetter interface {
	Mocks(namespace string) MockInterface
}

// MockInterface has methods to work with Mock resources.
type MockInterface interface {
	Create(ctx context.Context, mock *v1alpha1.Mock, opts v1.CreateOptions) (*v1alpha1.Mock, error)
	Update(ctx context.Context, mock *v1alpha1.Mock, opts v1.UpdateOptions) (*v1alpha1.Mock, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Mock, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.MockList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Mock, err error)
	MockExpansion
}

// mocks implements MockInterface
type mocks struct {
	client rest.Interface
	ns     string
}

// newMocks returns a Mocks
func newMocks(c *MockV1alpha1Client, namespace string) *mocks {
	return &mocks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mock, and returns the corresponding mock object, and an error if there is any.
func (c *mocks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Mock, err error) {
	result = &v1alpha1.Mock{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mocks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Mocks that match those selectors.
func (c *mocks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MockList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MockList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mocks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mocks.
func (c *mocks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mocks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a mock and creates it.  Returns the server's representation of the mock, and an error, if there is any.
func (c *mocks) Create(ctx context.Context, mock *v1alpha1.Mock, opts v1.CreateOptions) (result *v1alpha1.Mock, err error) {
	result = &v1alpha1.Mock{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mocks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mock).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a mock and updates it. Returns the server's representation of the mock, and an error, if there is any.
func (c *mocks) Update(ctx context.Context, mock *v1alpha1.Mock, opts v1.UpdateOptions) (result *v1alpha1.Mock, err error) {
	result = &v1alpha1.Mock{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mocks").
		Name(mock.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mock).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the mock and deletes it. Returns an error if one occurs.
func (c *mocks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mocks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mocks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mocks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched mock.
func (c *mocks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Mock, err error) {
	result = &v1alpha1.Mock{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mocks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
