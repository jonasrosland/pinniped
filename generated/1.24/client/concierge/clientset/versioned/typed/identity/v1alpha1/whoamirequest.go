// Copyright 2020-2023 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "go.pinniped.dev/generated/1.24/apis/concierge/identity/v1alpha1"
	scheme "go.pinniped.dev/generated/1.24/client/concierge/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
)

// WhoAmIRequestsGetter has a method to return a WhoAmIRequestInterface.
// A group's client should implement this interface.
type WhoAmIRequestsGetter interface {
	WhoAmIRequests() WhoAmIRequestInterface
}

// WhoAmIRequestInterface has methods to work with WhoAmIRequest resources.
type WhoAmIRequestInterface interface {
	Create(ctx context.Context, whoAmIRequest *v1alpha1.WhoAmIRequest, opts v1.CreateOptions) (*v1alpha1.WhoAmIRequest, error)
	WhoAmIRequestExpansion
}

// whoAmIRequests implements WhoAmIRequestInterface
type whoAmIRequests struct {
	client rest.Interface
}

// newWhoAmIRequests returns a WhoAmIRequests
func newWhoAmIRequests(c *IdentityV1alpha1Client) *whoAmIRequests {
	return &whoAmIRequests{
		client: c.RESTClient(),
	}
}

// Create takes the representation of a whoAmIRequest and creates it.  Returns the server's representation of the whoAmIRequest, and an error, if there is any.
func (c *whoAmIRequests) Create(ctx context.Context, whoAmIRequest *v1alpha1.WhoAmIRequest, opts v1.CreateOptions) (result *v1alpha1.WhoAmIRequest, err error) {
	result = &v1alpha1.WhoAmIRequest{}
	err = c.client.Post().
		Resource("whoamirequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(whoAmIRequest).
		Do(ctx).
		Into(result)
	return
}
