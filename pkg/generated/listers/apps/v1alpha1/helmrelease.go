/*
Copyright The Clusternet Authors.

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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	appsv1alpha1 "github.com/clusternet/clusternet/pkg/apis/apps/v1alpha1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// HelmReleaseLister helps list HelmReleases.
// All objects returned here must be treated as read-only.
type HelmReleaseLister interface {
	// List lists all HelmReleases in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*appsv1alpha1.HelmRelease, err error)
	// HelmReleases returns an object that can list and get HelmReleases.
	HelmReleases(namespace string) HelmReleaseNamespaceLister
	HelmReleaseListerExpansion
}

// helmReleaseLister implements the HelmReleaseLister interface.
type helmReleaseLister struct {
	listers.ResourceIndexer[*appsv1alpha1.HelmRelease]
}

// NewHelmReleaseLister returns a new HelmReleaseLister.
func NewHelmReleaseLister(indexer cache.Indexer) HelmReleaseLister {
	return &helmReleaseLister{listers.New[*appsv1alpha1.HelmRelease](indexer, appsv1alpha1.Resource("helmrelease"))}
}

// HelmReleases returns an object that can list and get HelmReleases.
func (s *helmReleaseLister) HelmReleases(namespace string) HelmReleaseNamespaceLister {
	return helmReleaseNamespaceLister{listers.NewNamespaced[*appsv1alpha1.HelmRelease](s.ResourceIndexer, namespace)}
}

// HelmReleaseNamespaceLister helps list and get HelmReleases.
// All objects returned here must be treated as read-only.
type HelmReleaseNamespaceLister interface {
	// List lists all HelmReleases in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*appsv1alpha1.HelmRelease, err error)
	// Get retrieves the HelmRelease from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*appsv1alpha1.HelmRelease, error)
	HelmReleaseNamespaceListerExpansion
}

// helmReleaseNamespaceLister implements the HelmReleaseNamespaceLister
// interface.
type helmReleaseNamespaceLister struct {
	listers.ResourceIndexer[*appsv1alpha1.HelmRelease]
}
