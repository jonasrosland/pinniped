// Copyright 2020 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	loginv1alpha1 "go.pinniped.dev/generated/1.17/apis/concierge/login/v1alpha1"
	versioned "go.pinniped.dev/generated/1.17/client/concierge/clientset/versioned"
	internalinterfaces "go.pinniped.dev/generated/1.17/client/concierge/informers/externalversions/internalinterfaces"
	v1alpha1 "go.pinniped.dev/generated/1.17/client/concierge/listers/login/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TokenCredentialRequestInformer provides access to a shared informer and lister for
// TokenCredentialRequests.
type TokenCredentialRequestInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TokenCredentialRequestLister
}

type tokenCredentialRequestInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewTokenCredentialRequestInformer constructs a new informer for TokenCredentialRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTokenCredentialRequestInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTokenCredentialRequestInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredTokenCredentialRequestInformer constructs a new informer for TokenCredentialRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTokenCredentialRequestInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LoginV1alpha1().TokenCredentialRequests().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LoginV1alpha1().TokenCredentialRequests().Watch(options)
			},
		},
		&loginv1alpha1.TokenCredentialRequest{},
		resyncPeriod,
		indexers,
	)
}

func (f *tokenCredentialRequestInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTokenCredentialRequestInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tokenCredentialRequestInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&loginv1alpha1.TokenCredentialRequest{}, f.defaultInformer)
}

func (f *tokenCredentialRequestInformer) Lister() v1alpha1.TokenCredentialRequestLister {
	return v1alpha1.NewTokenCredentialRequestLister(f.Informer().GetIndexer())
}
