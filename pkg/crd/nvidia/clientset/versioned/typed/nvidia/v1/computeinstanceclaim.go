/*
 * Copyright (c) 2022, NVIDIA CORPORATION.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	scheme "github.com/NVIDIA/k8s-dra-driver/pkg/crd/nvidia/clientset/versioned/scheme"
	v1 "github.com/NVIDIA/k8s-dra-driver/pkg/crd/nvidia/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ComputeInstanceClaimsGetter has a method to return a ComputeInstanceClaimInterface.
// A group's client should implement this interface.
type ComputeInstanceClaimsGetter interface {
	ComputeInstanceClaims(namespace string) ComputeInstanceClaimInterface
}

// ComputeInstanceClaimInterface has methods to work with ComputeInstanceClaim resources.
type ComputeInstanceClaimInterface interface {
	Create(ctx context.Context, computeInstanceClaim *v1.ComputeInstanceClaim, opts metav1.CreateOptions) (*v1.ComputeInstanceClaim, error)
	Update(ctx context.Context, computeInstanceClaim *v1.ComputeInstanceClaim, opts metav1.UpdateOptions) (*v1.ComputeInstanceClaim, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ComputeInstanceClaim, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ComputeInstanceClaimList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ComputeInstanceClaim, err error)
	ComputeInstanceClaimExpansion
}

// computeInstanceClaims implements ComputeInstanceClaimInterface
type computeInstanceClaims struct {
	client rest.Interface
	ns     string
}

// newComputeInstanceClaims returns a ComputeInstanceClaims
func newComputeInstanceClaims(c *DraV1Client, namespace string) *computeInstanceClaims {
	return &computeInstanceClaims{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeInstanceClaim, and returns the corresponding computeInstanceClaim object, and an error if there is any.
func (c *computeInstanceClaims) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ComputeInstanceClaim, err error) {
	result = &v1.ComputeInstanceClaim{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeinstanceclaims").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeInstanceClaims that match those selectors.
func (c *computeInstanceClaims) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ComputeInstanceClaimList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ComputeInstanceClaimList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeinstanceclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeInstanceClaims.
func (c *computeInstanceClaims) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computeinstanceclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a computeInstanceClaim and creates it.  Returns the server's representation of the computeInstanceClaim, and an error, if there is any.
func (c *computeInstanceClaims) Create(ctx context.Context, computeInstanceClaim *v1.ComputeInstanceClaim, opts metav1.CreateOptions) (result *v1.ComputeInstanceClaim, err error) {
	result = &v1.ComputeInstanceClaim{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computeinstanceclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeInstanceClaim).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a computeInstanceClaim and updates it. Returns the server's representation of the computeInstanceClaim, and an error, if there is any.
func (c *computeInstanceClaims) Update(ctx context.Context, computeInstanceClaim *v1.ComputeInstanceClaim, opts metav1.UpdateOptions) (result *v1.ComputeInstanceClaim, err error) {
	result = &v1.ComputeInstanceClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeinstanceclaims").
		Name(computeInstanceClaim.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeInstanceClaim).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the computeInstanceClaim and deletes it. Returns an error if one occurs.
func (c *computeInstanceClaims) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeinstanceclaims").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeInstanceClaims) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeinstanceclaims").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched computeInstanceClaim.
func (c *computeInstanceClaims) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ComputeInstanceClaim, err error) {
	result = &v1.ComputeInstanceClaim{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computeinstanceclaims").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}