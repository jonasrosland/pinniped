// Copyright 2020-2023 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "go.pinniped.dev/generated/1.23/apis/supervisor/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOIDCClients implements OIDCClientInterface
type FakeOIDCClients struct {
	Fake *FakeConfigV1alpha1
	ns   string
}

var oidcclientsResource = schema.GroupVersionResource{Group: "config.supervisor.pinniped.dev", Version: "v1alpha1", Resource: "oidcclients"}

var oidcclientsKind = schema.GroupVersionKind{Group: "config.supervisor.pinniped.dev", Version: "v1alpha1", Kind: "OIDCClient"}

// Get takes name of the oIDCClient, and returns the corresponding oIDCClient object, and an error if there is any.
func (c *FakeOIDCClients) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OIDCClient, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(oidcclientsResource, c.ns, name), &v1alpha1.OIDCClient{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OIDCClient), err
}

// List takes label and field selectors, and returns the list of OIDCClients that match those selectors.
func (c *FakeOIDCClients) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OIDCClientList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(oidcclientsResource, oidcclientsKind, c.ns, opts), &v1alpha1.OIDCClientList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OIDCClientList{ListMeta: obj.(*v1alpha1.OIDCClientList).ListMeta}
	for _, item := range obj.(*v1alpha1.OIDCClientList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested oIDCClients.
func (c *FakeOIDCClients) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(oidcclientsResource, c.ns, opts))

}

// Create takes the representation of a oIDCClient and creates it.  Returns the server's representation of the oIDCClient, and an error, if there is any.
func (c *FakeOIDCClients) Create(ctx context.Context, oIDCClient *v1alpha1.OIDCClient, opts v1.CreateOptions) (result *v1alpha1.OIDCClient, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(oidcclientsResource, c.ns, oIDCClient), &v1alpha1.OIDCClient{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OIDCClient), err
}

// Update takes the representation of a oIDCClient and updates it. Returns the server's representation of the oIDCClient, and an error, if there is any.
func (c *FakeOIDCClients) Update(ctx context.Context, oIDCClient *v1alpha1.OIDCClient, opts v1.UpdateOptions) (result *v1alpha1.OIDCClient, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(oidcclientsResource, c.ns, oIDCClient), &v1alpha1.OIDCClient{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OIDCClient), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOIDCClients) UpdateStatus(ctx context.Context, oIDCClient *v1alpha1.OIDCClient, opts v1.UpdateOptions) (*v1alpha1.OIDCClient, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(oidcclientsResource, "status", c.ns, oIDCClient), &v1alpha1.OIDCClient{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OIDCClient), err
}

// Delete takes name of the oIDCClient and deletes it. Returns an error if one occurs.
func (c *FakeOIDCClients) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(oidcclientsResource, c.ns, name, opts), &v1alpha1.OIDCClient{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOIDCClients) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(oidcclientsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OIDCClientList{})
	return err
}

// Patch applies the patch and returns the patched oIDCClient.
func (c *FakeOIDCClients) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OIDCClient, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(oidcclientsResource, c.ns, name, pt, data, subresources...), &v1alpha1.OIDCClient{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OIDCClient), err
}
