// Copyright 2020-2023 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "go.pinniped.dev/generated/1.17/apis/supervisor/idp/v1alpha1"
	scheme "go.pinniped.dev/generated/1.17/client/supervisor/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ActiveDirectoryIdentityProvidersGetter has a method to return a ActiveDirectoryIdentityProviderInterface.
// A group's client should implement this interface.
type ActiveDirectoryIdentityProvidersGetter interface {
	ActiveDirectoryIdentityProviders(namespace string) ActiveDirectoryIdentityProviderInterface
}

// ActiveDirectoryIdentityProviderInterface has methods to work with ActiveDirectoryIdentityProvider resources.
type ActiveDirectoryIdentityProviderInterface interface {
	Create(*v1alpha1.ActiveDirectoryIdentityProvider) (*v1alpha1.ActiveDirectoryIdentityProvider, error)
	Update(*v1alpha1.ActiveDirectoryIdentityProvider) (*v1alpha1.ActiveDirectoryIdentityProvider, error)
	UpdateStatus(*v1alpha1.ActiveDirectoryIdentityProvider) (*v1alpha1.ActiveDirectoryIdentityProvider, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ActiveDirectoryIdentityProvider, error)
	List(opts v1.ListOptions) (*v1alpha1.ActiveDirectoryIdentityProviderList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ActiveDirectoryIdentityProvider, err error)
	ActiveDirectoryIdentityProviderExpansion
}

// activeDirectoryIdentityProviders implements ActiveDirectoryIdentityProviderInterface
type activeDirectoryIdentityProviders struct {
	client rest.Interface
	ns     string
}

// newActiveDirectoryIdentityProviders returns a ActiveDirectoryIdentityProviders
func newActiveDirectoryIdentityProviders(c *IDPV1alpha1Client, namespace string) *activeDirectoryIdentityProviders {
	return &activeDirectoryIdentityProviders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the activeDirectoryIdentityProvider, and returns the corresponding activeDirectoryIdentityProvider object, and an error if there is any.
func (c *activeDirectoryIdentityProviders) Get(name string, options v1.GetOptions) (result *v1alpha1.ActiveDirectoryIdentityProvider, err error) {
	result = &v1alpha1.ActiveDirectoryIdentityProvider{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("activedirectoryidentityproviders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ActiveDirectoryIdentityProviders that match those selectors.
func (c *activeDirectoryIdentityProviders) List(opts v1.ListOptions) (result *v1alpha1.ActiveDirectoryIdentityProviderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ActiveDirectoryIdentityProviderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("activedirectoryidentityproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested activeDirectoryIdentityProviders.
func (c *activeDirectoryIdentityProviders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("activedirectoryidentityproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a activeDirectoryIdentityProvider and creates it.  Returns the server's representation of the activeDirectoryIdentityProvider, and an error, if there is any.
func (c *activeDirectoryIdentityProviders) Create(activeDirectoryIdentityProvider *v1alpha1.ActiveDirectoryIdentityProvider) (result *v1alpha1.ActiveDirectoryIdentityProvider, err error) {
	result = &v1alpha1.ActiveDirectoryIdentityProvider{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("activedirectoryidentityproviders").
		Body(activeDirectoryIdentityProvider).
		Do().
		Into(result)
	return
}

// Update takes the representation of a activeDirectoryIdentityProvider and updates it. Returns the server's representation of the activeDirectoryIdentityProvider, and an error, if there is any.
func (c *activeDirectoryIdentityProviders) Update(activeDirectoryIdentityProvider *v1alpha1.ActiveDirectoryIdentityProvider) (result *v1alpha1.ActiveDirectoryIdentityProvider, err error) {
	result = &v1alpha1.ActiveDirectoryIdentityProvider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("activedirectoryidentityproviders").
		Name(activeDirectoryIdentityProvider.Name).
		Body(activeDirectoryIdentityProvider).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *activeDirectoryIdentityProviders) UpdateStatus(activeDirectoryIdentityProvider *v1alpha1.ActiveDirectoryIdentityProvider) (result *v1alpha1.ActiveDirectoryIdentityProvider, err error) {
	result = &v1alpha1.ActiveDirectoryIdentityProvider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("activedirectoryidentityproviders").
		Name(activeDirectoryIdentityProvider.Name).
		SubResource("status").
		Body(activeDirectoryIdentityProvider).
		Do().
		Into(result)
	return
}

// Delete takes name of the activeDirectoryIdentityProvider and deletes it. Returns an error if one occurs.
func (c *activeDirectoryIdentityProviders) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("activedirectoryidentityproviders").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *activeDirectoryIdentityProviders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("activedirectoryidentityproviders").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched activeDirectoryIdentityProvider.
func (c *activeDirectoryIdentityProviders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ActiveDirectoryIdentityProvider, err error) {
	result = &v1alpha1.ActiveDirectoryIdentityProvider{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("activedirectoryidentityproviders").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
