/*
Copyright The Kubernetes Authors.
Copyright 2020 Authors of Arktos - file modified.

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

package v1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PodLister helps list Pods.
type PodLister interface {
	// List lists all Pods in the indexer.
	List(selector labels.Selector) (ret []*v1.Pod, err error)
	// Pods returns an object that can list and get Pods.
	Pods(namespace string) PodNamespaceLister
	PodsWithMultiTenancy(namespace string, tenant string) PodNamespaceLister
	PodListerExpansion
}

// podLister implements the PodLister interface.
type podLister struct {
	indexer cache.Indexer
}

// NewPodLister returns a new PodLister.
func NewPodLister(indexer cache.Indexer) PodLister {
	return &podLister{indexer: indexer}
}

// List lists all Pods in the indexer.
func (s *podLister) List(selector labels.Selector) (ret []*v1.Pod, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Pod))
	})
	return ret, err
}

// Pods returns an object that can list and get Pods.
func (s *podLister) Pods(namespace string) PodNamespaceLister {
	return podNamespaceLister{indexer: s.indexer, namespace: namespace, tenant: "system"}
}

func (s *podLister) PodsWithMultiTenancy(namespace string, tenant string) PodNamespaceLister {
	return podNamespaceLister{indexer: s.indexer, namespace: namespace, tenant: tenant}
}

// PodNamespaceLister helps list and get Pods.
type PodNamespaceLister interface {
	// List lists all Pods in the indexer for a given tenant/namespace.
	List(selector labels.Selector) (ret []*v1.Pod, err error)
	// Get retrieves the Pod from the indexer for a given tenant/namespace and name.
	Get(name string) (*v1.Pod, error)
	PodNamespaceListerExpansion
}

// podNamespaceLister implements the PodNamespaceLister
// interface.
type podNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
	tenant    string
}

// List lists all Pods in the indexer for a given namespace.
func (s podNamespaceLister) List(selector labels.Selector) (ret []*v1.Pod, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.tenant, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Pod))
	})
	return ret, err
}

// Get retrieves the Pod from the indexer for a given namespace and name.
func (s podNamespaceLister) Get(name string) (*v1.Pod, error) {
	key := s.tenant + "/" + s.namespace + "/" + name
	if s.tenant == "system" {
		key = s.namespace + "/" + name
	}
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("pod"), name)
	}
	return obj.(*v1.Pod), nil
}
