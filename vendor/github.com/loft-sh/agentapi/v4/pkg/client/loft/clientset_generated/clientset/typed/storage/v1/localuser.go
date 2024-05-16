// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/loft-sh/agentapi/v4/pkg/apis/loft/storage/v1"
	scheme "github.com/loft-sh/agentapi/v4/pkg/client/loft/clientset_generated/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LocalUsersGetter has a method to return a LocalUserInterface.
// A group's client should implement this interface.
type LocalUsersGetter interface {
	LocalUsers() LocalUserInterface
}

// LocalUserInterface has methods to work with LocalUser resources.
type LocalUserInterface interface {
	Create(ctx context.Context, localUser *v1.LocalUser, opts metav1.CreateOptions) (*v1.LocalUser, error)
	Update(ctx context.Context, localUser *v1.LocalUser, opts metav1.UpdateOptions) (*v1.LocalUser, error)
	UpdateStatus(ctx context.Context, localUser *v1.LocalUser, opts metav1.UpdateOptions) (*v1.LocalUser, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.LocalUser, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.LocalUserList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.LocalUser, err error)
	LocalUserExpansion
}

// localUsers implements LocalUserInterface
type localUsers struct {
	client rest.Interface
}

// newLocalUsers returns a LocalUsers
func newLocalUsers(c *StorageV1Client) *localUsers {
	return &localUsers{
		client: c.RESTClient(),
	}
}

// Get takes name of the localUser, and returns the corresponding localUser object, and an error if there is any.
func (c *localUsers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.LocalUser, err error) {
	result = &v1.LocalUser{}
	err = c.client.Get().
		Resource("localusers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LocalUsers that match those selectors.
func (c *localUsers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.LocalUserList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.LocalUserList{}
	err = c.client.Get().
		Resource("localusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested localUsers.
func (c *localUsers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("localusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a localUser and creates it.  Returns the server's representation of the localUser, and an error, if there is any.
func (c *localUsers) Create(ctx context.Context, localUser *v1.LocalUser, opts metav1.CreateOptions) (result *v1.LocalUser, err error) {
	result = &v1.LocalUser{}
	err = c.client.Post().
		Resource("localusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(localUser).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a localUser and updates it. Returns the server's representation of the localUser, and an error, if there is any.
func (c *localUsers) Update(ctx context.Context, localUser *v1.LocalUser, opts metav1.UpdateOptions) (result *v1.LocalUser, err error) {
	result = &v1.LocalUser{}
	err = c.client.Put().
		Resource("localusers").
		Name(localUser.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(localUser).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *localUsers) UpdateStatus(ctx context.Context, localUser *v1.LocalUser, opts metav1.UpdateOptions) (result *v1.LocalUser, err error) {
	result = &v1.LocalUser{}
	err = c.client.Put().
		Resource("localusers").
		Name(localUser.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(localUser).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the localUser and deletes it. Returns an error if one occurs.
func (c *localUsers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("localusers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *localUsers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("localusers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched localUser.
func (c *localUsers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.LocalUser, err error) {
	result = &v1.LocalUser{}
	err = c.client.Patch(pt).
		Resource("localusers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}