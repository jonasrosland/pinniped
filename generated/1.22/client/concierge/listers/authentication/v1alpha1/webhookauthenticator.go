// Copyright 2020-2023 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "go.pinniped.dev/generated/1.22/apis/concierge/authentication/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WebhookAuthenticatorLister helps list WebhookAuthenticators.
// All objects returned here must be treated as read-only.
type WebhookAuthenticatorLister interface {
	// List lists all WebhookAuthenticators in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WebhookAuthenticator, err error)
	// Get retrieves the WebhookAuthenticator from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.WebhookAuthenticator, error)
	WebhookAuthenticatorListerExpansion
}

// webhookAuthenticatorLister implements the WebhookAuthenticatorLister interface.
type webhookAuthenticatorLister struct {
	indexer cache.Indexer
}

// NewWebhookAuthenticatorLister returns a new WebhookAuthenticatorLister.
func NewWebhookAuthenticatorLister(indexer cache.Indexer) WebhookAuthenticatorLister {
	return &webhookAuthenticatorLister{indexer: indexer}
}

// List lists all WebhookAuthenticators in the indexer.
func (s *webhookAuthenticatorLister) List(selector labels.Selector) (ret []*v1alpha1.WebhookAuthenticator, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WebhookAuthenticator))
	})
	return ret, err
}

// Get retrieves the WebhookAuthenticator from the index for a given name.
func (s *webhookAuthenticatorLister) Get(name string) (*v1alpha1.WebhookAuthenticator, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("webhookauthenticator"), name)
	}
	return obj.(*v1alpha1.WebhookAuthenticator), nil
}
