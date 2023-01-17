// Copyright 2020-2023 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "go.pinniped.dev/generated/1.25/apis/concierge/authentication/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// JWTAuthenticatorLister helps list JWTAuthenticators.
// All objects returned here must be treated as read-only.
type JWTAuthenticatorLister interface {
	// List lists all JWTAuthenticators in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.JWTAuthenticator, err error)
	// Get retrieves the JWTAuthenticator from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.JWTAuthenticator, error)
	JWTAuthenticatorListerExpansion
}

// jWTAuthenticatorLister implements the JWTAuthenticatorLister interface.
type jWTAuthenticatorLister struct {
	indexer cache.Indexer
}

// NewJWTAuthenticatorLister returns a new JWTAuthenticatorLister.
func NewJWTAuthenticatorLister(indexer cache.Indexer) JWTAuthenticatorLister {
	return &jWTAuthenticatorLister{indexer: indexer}
}

// List lists all JWTAuthenticators in the indexer.
func (s *jWTAuthenticatorLister) List(selector labels.Selector) (ret []*v1alpha1.JWTAuthenticator, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.JWTAuthenticator))
	})
	return ret, err
}

// Get retrieves the JWTAuthenticator from the index for a given name.
func (s *jWTAuthenticatorLister) Get(name string) (*v1alpha1.JWTAuthenticator, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("jwtauthenticator"), name)
	}
	return obj.(*v1alpha1.JWTAuthenticator), nil
}
