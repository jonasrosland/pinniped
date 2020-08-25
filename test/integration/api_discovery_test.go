/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

package integration

import (
	"testing"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/suzerain-io/pinniped/test/library"
)

func TestGetAPIResourceList(t *testing.T) {
	library.SkipUnlessIntegration(t)
	library.SkipUnlessClusterHasCapability(t, library.ClusterSigningKeyIsAvailable)

	client := library.NewPinnipedClientset(t)

	groups, resources, err := client.Discovery().ServerGroupsAndResources()
	require.NoError(t, err)

	groupName := "pinniped.dev"
	actualGroup := findGroup(groupName, groups)
	require.NotNil(t, actualGroup)

	expectedGroup := &metav1.APIGroup{
		Name: "pinniped.dev",
		Versions: []metav1.GroupVersionForDiscovery{
			{
				GroupVersion: "pinniped.dev/v1alpha1",
				Version:      "v1alpha1",
			},
		},
		PreferredVersion: metav1.GroupVersionForDiscovery{
			GroupVersion: "pinniped.dev/v1alpha1",
			Version:      "v1alpha1",
		},
	}
	require.Equal(t, expectedGroup, actualGroup)

	actualPinnipedResources := findResources("pinniped.dev/v1alpha1", resources)
	require.NotNil(t, actualPinnipedResources)
	actualCrdPinnipedResources := findResources("crd.pinniped.dev/v1alpha1", resources)
	require.NotNil(t, actualPinnipedResources)

	expectedCredentialRequestAPIResource := metav1.APIResource{
		Name: "credentialrequests",
		Kind: "CredentialRequest",
		Verbs: metav1.Verbs([]string{
			"create",
		}),
		Namespaced: false,

		// This is currently an empty string in the response; maybe it should not be
		// empty? Seems like no harm in keeping it like this for now, but feel free
		// to update in the future if there is a compelling reason to do so.
		SingularName: "",
	}

	expectedCredentialIssuerConfigResource := metav1.APIResource{
		Name:         "credentialissuerconfigs",
		SingularName: "credentialissuerconfig",
		Namespaced:   true,
		Kind:         "CredentialIssuerConfig",
		Verbs: metav1.Verbs([]string{
			"delete", "deletecollection", "get", "list", "patch", "create", "update", "watch",
		}),
		ShortNames:         []string{"cic"},
		StorageVersionHash: "unknown: to be filled in automatically below",
	}

	require.Len(t, actualPinnipedResources.APIResources, 1)
	require.Equal(t, expectedCredentialRequestAPIResource, actualPinnipedResources.APIResources[0])

	require.Len(t, actualCrdPinnipedResources.APIResources, 1)
	actualAPIResource := actualCrdPinnipedResources.APIResources[0]
	// workaround because its hard to predict the storage version hash (e.g. "t/+v41y+3e4=")
	// so just don't worry about comparing that field
	expectedCredentialIssuerConfigResource.StorageVersionHash = actualAPIResource.StorageVersionHash
	require.Equal(t, expectedCredentialIssuerConfigResource, actualAPIResource)
}

func findGroup(name string, groups []*metav1.APIGroup) *metav1.APIGroup {
	for _, group := range groups {
		if group.Name == name {
			return group
		}
	}
	return nil
}

func findResources(groupVersion string, resources []*metav1.APIResourceList) *metav1.APIResourceList {
	for _, resource := range resources {
		if resource.GroupVersion == groupVersion {
			return resource
		}
	}
	return nil
}
