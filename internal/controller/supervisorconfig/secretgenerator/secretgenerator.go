// Copyright 2020 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package secretgenerator provides a controller that can ensure existence of a generated secret.
package secretgenerator

import (
	"context"
	"crypto/rand"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	corev1informers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
	"k8s.io/klog/v2"

	pinnipedcontroller "go.pinniped.dev/internal/controller"
	"go.pinniped.dev/internal/controllerlib"
	"go.pinniped.dev/internal/plog"
)

const (
	symmetricKeySecretType    = "secrets.pinniped.dev/symmetric"
	symmetricKeySecretDataKey = "key"

	symmetricKeySize = 32 // TODO: what should this be?
)

// generateKey is stubbed out for the purpose of testing. The default behavior is to generate a symmetric key.
//nolint:gochecknoglobals
var generateKey = generateSymmetricKey

func generateSymmetricKey() ([]byte, error) {
	b := make([]byte, symmetricKeySize)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

type controller struct {
	owner            *appsv1.Deployment
	client           kubernetes.Interface
	secrets          corev1informers.SecretInformer
	onCreateOrUpdate func(secret []byte)
}

// New instantiates a new controllerlib.Controller which will ensure existence of a generated secret.
func New(
	owner *appsv1.Deployment,
	client kubernetes.Interface,
	secrets corev1informers.SecretInformer,
	onCreateOrUpdate func(secret []byte),
) controllerlib.Controller {
	c := controller{
		owner:            owner,
		client:           client,
		secrets:          secrets,
		onCreateOrUpdate: onCreateOrUpdate,
	}
	filter := pinnipedcontroller.SimpleFilter(func(obj metav1.Object) bool {
		return metav1.IsControlledBy(obj, owner)
	}, nil)
	return controllerlib.New(
		controllerlib.Config{Name: owner.Name + "-secret-generator", Syncer: &c},
		controllerlib.WithInformer(secrets, filter, controllerlib.InformerOption{}),
		controllerlib.WithInitialEvent(controllerlib.Key{
			Namespace: owner.Namespace,
			Name:      owner.Name + "-keys",
		}),
	)
}

// Sync implements controllerlib.Syncer.Sync().
func (c *controller) Sync(ctx controllerlib.Context) error {
	secret, err := c.secrets.Lister().Secrets(ctx.Key.Namespace).Get(ctx.Key.Name)
	isNotFound := k8serrors.IsNotFound(err)
	if !isNotFound && err != nil {
		return fmt.Errorf("failed to list secret %s/%s: %w", ctx.Key.Namespace, ctx.Key.Name, err)
	}

	secretNeedsUpdate := isNotFound || !c.isValid(secret)
	if !secretNeedsUpdate {
		plog.Debug("secret is up to date", "secret", klog.KObj(secret))
		c.onCreateOrUpdate(secret.Data[symmetricKeySecretDataKey])
		return nil
	}

	newSecret, err := c.generateSecret(ctx.Key.Namespace, ctx.Key.Name)
	if err != nil {
		return fmt.Errorf("failed to generate secret: %w", err)
	}

	if isNotFound {
		err = c.createSecret(ctx.Context, newSecret)
	} else {
		err = c.updateSecret(ctx.Context, &newSecret, ctx.Key.Name)
	}
	if err != nil {
		return fmt.Errorf("failed to create/update secret %s/%s: %w", newSecret.Namespace, newSecret.Name, err)
	}

	c.onCreateOrUpdate(newSecret.Data[symmetricKeySecretDataKey])

	return nil
}

func (c *controller) isValid(secret *corev1.Secret) bool {
	if secret.Type != symmetricKeySecretType {
		return false
	}

	data, ok := secret.Data[symmetricKeySecretDataKey]
	if !ok {
		return false
	}
	if len(data) != symmetricKeySize {
		return false
	}

	return true
}

func (c *controller) generateSecret(namespace, name string) (*corev1.Secret, error) {
	symmetricKey, err := generateKey()
	if err != nil {
		return nil, err
	}

	deploymentGVK := schema.GroupVersionKind{
		Group:   appsv1.SchemeGroupVersion.Group,
		Version: appsv1.SchemeGroupVersion.Version,
		Kind:    "Deployment",
	}
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(c.owner, deploymentGVK),
			},
		},
		Type: symmetricKeySecretType,
		Data: map[string][]byte{
			symmetricKeySecretDataKey: symmetricKey,
		},
	}, nil
}

func (c *controller) createSecret(ctx context.Context, newSecret *corev1.Secret) error {
	_, err := c.client.CoreV1().Secrets(newSecret.Namespace).Create(ctx, newSecret, metav1.CreateOptions{})
	return err
}

func (c *controller) updateSecret(ctx context.Context, newSecret **corev1.Secret, secretName string) error {
	secrets := c.client.CoreV1().Secrets((*newSecret).Namespace)
	return retry.RetryOnConflict(retry.DefaultBackoff, func() error {
		currentSecret, err := secrets.Get(ctx, secretName, metav1.GetOptions{})
		isNotFound := k8serrors.IsNotFound(err)
		if !isNotFound && err != nil {
			return fmt.Errorf("failed to get secret: %w", err)
		}

		if isNotFound {
			if err := c.createSecret(ctx, *newSecret); err != nil {
				return fmt.Errorf("failed to create secret: %w", err)
			}
			return nil
		}

		if c.isValid(currentSecret) {
			*newSecret = currentSecret
			return nil
		}

		currentSecret.Type = (*newSecret).Type
		currentSecret.Data = (*newSecret).Data

		_, err = secrets.Update(ctx, currentSecret, metav1.UpdateOptions{})
		return err
	})
}
