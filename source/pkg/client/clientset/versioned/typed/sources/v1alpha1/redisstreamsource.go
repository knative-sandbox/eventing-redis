/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "knative.dev/eventing-redis/source/pkg/apis/sources/v1alpha1"
	scheme "knative.dev/eventing-redis/source/pkg/client/clientset/versioned/scheme"
)

// RedisStreamSourcesGetter has a method to return a RedisStreamSourceInterface.
// A group's client should implement this interface.
type RedisStreamSourcesGetter interface {
	RedisStreamSources(namespace string) RedisStreamSourceInterface
}

// RedisStreamSourceInterface has methods to work with RedisStreamSource resources.
type RedisStreamSourceInterface interface {
	Create(ctx context.Context, redisStreamSource *v1alpha1.RedisStreamSource, opts v1.CreateOptions) (*v1alpha1.RedisStreamSource, error)
	Update(ctx context.Context, redisStreamSource *v1alpha1.RedisStreamSource, opts v1.UpdateOptions) (*v1alpha1.RedisStreamSource, error)
	UpdateStatus(ctx context.Context, redisStreamSource *v1alpha1.RedisStreamSource, opts v1.UpdateOptions) (*v1alpha1.RedisStreamSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.RedisStreamSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.RedisStreamSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RedisStreamSource, err error)
	RedisStreamSourceExpansion
}

// redisStreamSources implements RedisStreamSourceInterface
type redisStreamSources struct {
	client rest.Interface
	ns     string
}

// newRedisStreamSources returns a RedisStreamSources
func newRedisStreamSources(c *SourcesV1alpha1Client, namespace string) *redisStreamSources {
	return &redisStreamSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the redisStreamSource, and returns the corresponding redisStreamSource object, and an error if there is any.
func (c *redisStreamSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RedisStreamSource, err error) {
	result = &v1alpha1.RedisStreamSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("redisstreamsources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RedisStreamSources that match those selectors.
func (c *redisStreamSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RedisStreamSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RedisStreamSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("redisstreamsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested redisStreamSources.
func (c *redisStreamSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("redisstreamsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a redisStreamSource and creates it.  Returns the server's representation of the redisStreamSource, and an error, if there is any.
func (c *redisStreamSources) Create(ctx context.Context, redisStreamSource *v1alpha1.RedisStreamSource, opts v1.CreateOptions) (result *v1alpha1.RedisStreamSource, err error) {
	result = &v1alpha1.RedisStreamSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("redisstreamsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(redisStreamSource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a redisStreamSource and updates it. Returns the server's representation of the redisStreamSource, and an error, if there is any.
func (c *redisStreamSources) Update(ctx context.Context, redisStreamSource *v1alpha1.RedisStreamSource, opts v1.UpdateOptions) (result *v1alpha1.RedisStreamSource, err error) {
	result = &v1alpha1.RedisStreamSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("redisstreamsources").
		Name(redisStreamSource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(redisStreamSource).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *redisStreamSources) UpdateStatus(ctx context.Context, redisStreamSource *v1alpha1.RedisStreamSource, opts v1.UpdateOptions) (result *v1alpha1.RedisStreamSource, err error) {
	result = &v1alpha1.RedisStreamSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("redisstreamsources").
		Name(redisStreamSource.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(redisStreamSource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the redisStreamSource and deletes it. Returns an error if one occurs.
func (c *redisStreamSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("redisstreamsources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *redisStreamSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("redisstreamsources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched redisStreamSource.
func (c *redisStreamSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RedisStreamSource, err 
error) {
	result = &v1alpha1.RedisStreamSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("redisstreamsources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
