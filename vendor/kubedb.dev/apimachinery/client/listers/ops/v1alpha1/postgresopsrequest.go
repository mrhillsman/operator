/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubedb.dev/apimachinery/apis/ops/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PostgresOpsRequestLister helps list PostgresOpsRequests.
type PostgresOpsRequestLister interface {
	// List lists all PostgresOpsRequests in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PostgresOpsRequest, err error)
	// PostgresOpsRequests returns an object that can list and get PostgresOpsRequests.
	PostgresOpsRequests(namespace string) PostgresOpsRequestNamespaceLister
	PostgresOpsRequestListerExpansion
}

// postgresOpsRequestLister implements the PostgresOpsRequestLister interface.
type postgresOpsRequestLister struct {
	indexer cache.Indexer
}

// NewPostgresOpsRequestLister returns a new PostgresOpsRequestLister.
func NewPostgresOpsRequestLister(indexer cache.Indexer) PostgresOpsRequestLister {
	return &postgresOpsRequestLister{indexer: indexer}
}

// List lists all PostgresOpsRequests in the indexer.
func (s *postgresOpsRequestLister) List(selector labels.Selector) (ret []*v1alpha1.PostgresOpsRequest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PostgresOpsRequest))
	})
	return ret, err
}

// PostgresOpsRequests returns an object that can list and get PostgresOpsRequests.
func (s *postgresOpsRequestLister) PostgresOpsRequests(namespace string) PostgresOpsRequestNamespaceLister {
	return postgresOpsRequestNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PostgresOpsRequestNamespaceLister helps list and get PostgresOpsRequests.
type PostgresOpsRequestNamespaceLister interface {
	// List lists all PostgresOpsRequests in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PostgresOpsRequest, err error)
	// Get retrieves the PostgresOpsRequest from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PostgresOpsRequest, error)
	PostgresOpsRequestNamespaceListerExpansion
}

// postgresOpsRequestNamespaceLister implements the PostgresOpsRequestNamespaceLister
// interface.
type postgresOpsRequestNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PostgresOpsRequests in the indexer for a given namespace.
func (s postgresOpsRequestNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PostgresOpsRequest, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PostgresOpsRequest))
	})
	return ret, err
}

// Get retrieves the PostgresOpsRequest from the indexer for a given namespace and name.
func (s postgresOpsRequestNamespaceLister) Get(name string) (*v1alpha1.PostgresOpsRequest, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("postgresopsrequest"), name)
	}
	return obj.(*v1alpha1.PostgresOpsRequest), nil
}