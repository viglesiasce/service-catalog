/*
Copyright 2017 The Kubernetes Authors.

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

package fake

import (
	v1alpha1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterServicePlans implements ClusterServicePlanInterface
type FakeClusterServicePlans struct {
	Fake *FakeServicecatalogV1alpha1
}

var clusterserviceplansResource = schema.GroupVersionResource{Group: "servicecatalog.k8s.io", Version: "v1alpha1", Resource: "clusterserviceplans"}

var clusterserviceplansKind = schema.GroupVersionKind{Group: "servicecatalog.k8s.io", Version: "v1alpha1", Kind: "ClusterServicePlan"}

func (c *FakeClusterServicePlans) Create(clusterServicePlan *v1alpha1.ClusterServicePlan) (result *v1alpha1.ClusterServicePlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterserviceplansResource, clusterServicePlan), &v1alpha1.ClusterServicePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterServicePlan), err
}

func (c *FakeClusterServicePlans) Update(clusterServicePlan *v1alpha1.ClusterServicePlan) (result *v1alpha1.ClusterServicePlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterserviceplansResource, clusterServicePlan), &v1alpha1.ClusterServicePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterServicePlan), err
}

func (c *FakeClusterServicePlans) UpdateStatus(clusterServicePlan *v1alpha1.ClusterServicePlan) (*v1alpha1.ClusterServicePlan, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterserviceplansResource, "status", clusterServicePlan), &v1alpha1.ClusterServicePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterServicePlan), err
}

func (c *FakeClusterServicePlans) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterserviceplansResource, name), &v1alpha1.ClusterServicePlan{})
	return err
}

func (c *FakeClusterServicePlans) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterserviceplansResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterServicePlanList{})
	return err
}

func (c *FakeClusterServicePlans) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterServicePlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterserviceplansResource, name), &v1alpha1.ClusterServicePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterServicePlan), err
}

func (c *FakeClusterServicePlans) List(opts v1.ListOptions) (result *v1alpha1.ClusterServicePlanList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterserviceplansResource, clusterserviceplansKind, opts), &v1alpha1.ClusterServicePlanList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterServicePlanList{}
	for _, item := range obj.(*v1alpha1.ClusterServicePlanList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterServicePlans.
func (c *FakeClusterServicePlans) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterserviceplansResource, opts))
}

// Patch applies the patch and returns the patched clusterServicePlan.
func (c *FakeClusterServicePlans) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterServicePlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterserviceplansResource, name, data, subresources...), &v1alpha1.ClusterServicePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterServicePlan), err
}