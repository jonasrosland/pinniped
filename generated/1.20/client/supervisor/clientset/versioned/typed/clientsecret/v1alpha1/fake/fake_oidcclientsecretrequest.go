// Copyright 2020-2023 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "go.pinniped.dev/generated/1.20/apis/supervisor/clientsecret/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// FakeOIDCClientSecretRequests implements OIDCClientSecretRequestInterface
type FakeOIDCClientSecretRequests struct {
	Fake *FakeClientsecretV1alpha1
	ns   string
}

var oidcclientsecretrequestsResource = schema.GroupVersionResource{Group: "clientsecret.supervisor.pinniped.dev", Version: "v1alpha1", Resource: "oidcclientsecretrequests"}

var oidcclientsecretrequestsKind = schema.GroupVersionKind{Group: "clientsecret.supervisor.pinniped.dev", Version: "v1alpha1", Kind: "OIDCClientSecretRequest"}

// Create takes the representation of a oIDCClientSecretRequest and creates it.  Returns the server's representation of the oIDCClientSecretRequest, and an error, if there is any.
func (c *FakeOIDCClientSecretRequests) Create(ctx context.Context, oIDCClientSecretRequest *v1alpha1.OIDCClientSecretRequest, opts v1.CreateOptions) (result *v1alpha1.OIDCClientSecretRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(oidcclientsecretrequestsResource, c.ns, oIDCClientSecretRequest), &v1alpha1.OIDCClientSecretRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OIDCClientSecretRequest), err
}
