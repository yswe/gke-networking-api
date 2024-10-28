/*
Copyright 2024 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/GoogleCloudPlatform/gke-networking-api/apis/network/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSubnetworks implements SubnetworkInterface
type FakeSubnetworks struct {
	Fake *FakeNetworkingV1
}

var subnetworksResource = v1.SchemeGroupVersion.WithResource("subnetworks")

var subnetworksKind = v1.SchemeGroupVersion.WithKind("Subnetwork")

// Get takes name of the subnetwork, and returns the corresponding subnetwork object, and an error if there is any.
func (c *FakeSubnetworks) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Subnetwork, err error) {
	emptyResult := &v1.Subnetwork{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(subnetworksResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Subnetwork), err
}

// List takes label and field selectors, and returns the list of Subnetworks that match those selectors.
func (c *FakeSubnetworks) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SubnetworkList, err error) {
	emptyResult := &v1.SubnetworkList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(subnetworksResource, subnetworksKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.SubnetworkList{ListMeta: obj.(*v1.SubnetworkList).ListMeta}
	for _, item := range obj.(*v1.SubnetworkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested subnetworks.
func (c *FakeSubnetworks) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(subnetworksResource, opts))
}

// Create takes the representation of a subnetwork and creates it.  Returns the server's representation of the subnetwork, and an error, if there is any.
func (c *FakeSubnetworks) Create(ctx context.Context, subnetwork *v1.Subnetwork, opts metav1.CreateOptions) (result *v1.Subnetwork, err error) {
	emptyResult := &v1.Subnetwork{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(subnetworksResource, subnetwork, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Subnetwork), err
}

// Update takes the representation of a subnetwork and updates it. Returns the server's representation of the subnetwork, and an error, if there is any.
func (c *FakeSubnetworks) Update(ctx context.Context, subnetwork *v1.Subnetwork, opts metav1.UpdateOptions) (result *v1.Subnetwork, err error) {
	emptyResult := &v1.Subnetwork{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(subnetworksResource, subnetwork, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Subnetwork), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSubnetworks) UpdateStatus(ctx context.Context, subnetwork *v1.Subnetwork, opts metav1.UpdateOptions) (result *v1.Subnetwork, err error) {
	emptyResult := &v1.Subnetwork{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(subnetworksResource, "status", subnetwork, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Subnetwork), err
}

// Delete takes name of the subnetwork and deletes it. Returns an error if one occurs.
func (c *FakeSubnetworks) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(subnetworksResource, name, opts), &v1.Subnetwork{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSubnetworks) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(subnetworksResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.SubnetworkList{})
	return err
}

// Patch applies the patch and returns the patched subnetwork.
func (c *FakeSubnetworks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Subnetwork, err error) {
	emptyResult := &v1.Subnetwork{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(subnetworksResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Subnetwork), err
}