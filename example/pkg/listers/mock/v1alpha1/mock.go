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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "go.f110.dev/rules_k8s_controller/example/pkg/api/mock/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MockLister helps list Mocks.
// All objects returned here must be treated as read-only.
type MockLister interface {
	// List lists all Mocks in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Mock, err error)
	// Mocks returns an object that can list and get Mocks.
	Mocks(namespace string) MockNamespaceLister
	MockListerExpansion
}

// mockLister implements the MockLister interface.
type mockLister struct {
	indexer cache.Indexer
}

// NewMockLister returns a new MockLister.
func NewMockLister(indexer cache.Indexer) MockLister {
	return &mockLister{indexer: indexer}
}

// List lists all Mocks in the indexer.
func (s *mockLister) List(selector labels.Selector) (ret []*v1alpha1.Mock, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Mock))
	})
	return ret, err
}

// Mocks returns an object that can list and get Mocks.
func (s *mockLister) Mocks(namespace string) MockNamespaceLister {
	return mockNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MockNamespaceLister helps list and get Mocks.
// All objects returned here must be treated as read-only.
type MockNamespaceLister interface {
	// List lists all Mocks in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Mock, err error)
	// Get retrieves the Mock from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Mock, error)
	MockNamespaceListerExpansion
}

// mockNamespaceLister implements the MockNamespaceLister
// interface.
type mockNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Mocks in the indexer for a given namespace.
func (s mockNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Mock, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Mock))
	})
	return ret, err
}

// Get retrieves the Mock from the indexer for a given namespace and name.
func (s mockNamespaceLister) Get(name string) (*v1alpha1.Mock, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mock"), name)
	}
	return obj.(*v1alpha1.Mock), nil
}
