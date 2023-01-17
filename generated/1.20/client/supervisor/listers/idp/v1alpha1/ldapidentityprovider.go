// Copyright 2020-2023 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "go.pinniped.dev/generated/1.20/apis/supervisor/idp/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LDAPIdentityProviderLister helps list LDAPIdentityProviders.
// All objects returned here must be treated as read-only.
type LDAPIdentityProviderLister interface {
	// List lists all LDAPIdentityProviders in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LDAPIdentityProvider, err error)
	// LDAPIdentityProviders returns an object that can list and get LDAPIdentityProviders.
	LDAPIdentityProviders(namespace string) LDAPIdentityProviderNamespaceLister
	LDAPIdentityProviderListerExpansion
}

// lDAPIdentityProviderLister implements the LDAPIdentityProviderLister interface.
type lDAPIdentityProviderLister struct {
	indexer cache.Indexer
}

// NewLDAPIdentityProviderLister returns a new LDAPIdentityProviderLister.
func NewLDAPIdentityProviderLister(indexer cache.Indexer) LDAPIdentityProviderLister {
	return &lDAPIdentityProviderLister{indexer: indexer}
}

// List lists all LDAPIdentityProviders in the indexer.
func (s *lDAPIdentityProviderLister) List(selector labels.Selector) (ret []*v1alpha1.LDAPIdentityProvider, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LDAPIdentityProvider))
	})
	return ret, err
}

// LDAPIdentityProviders returns an object that can list and get LDAPIdentityProviders.
func (s *lDAPIdentityProviderLister) LDAPIdentityProviders(namespace string) LDAPIdentityProviderNamespaceLister {
	return lDAPIdentityProviderNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LDAPIdentityProviderNamespaceLister helps list and get LDAPIdentityProviders.
// All objects returned here must be treated as read-only.
type LDAPIdentityProviderNamespaceLister interface {
	// List lists all LDAPIdentityProviders in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LDAPIdentityProvider, err error)
	// Get retrieves the LDAPIdentityProvider from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LDAPIdentityProvider, error)
	LDAPIdentityProviderNamespaceListerExpansion
}

// lDAPIdentityProviderNamespaceLister implements the LDAPIdentityProviderNamespaceLister
// interface.
type lDAPIdentityProviderNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LDAPIdentityProviders in the indexer for a given namespace.
func (s lDAPIdentityProviderNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LDAPIdentityProvider, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LDAPIdentityProvider))
	})
	return ret, err
}

// Get retrieves the LDAPIdentityProvider from the indexer for a given namespace and name.
func (s lDAPIdentityProviderNamespaceLister) Get(name string) (*v1alpha1.LDAPIdentityProvider, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ldapidentityprovider"), name)
	}
	return obj.(*v1alpha1.LDAPIdentityProvider), nil
}
