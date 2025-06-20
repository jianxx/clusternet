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
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"

	appsv1alpha1 "github.com/clusternet/clusternet/pkg/apis/apps/v1alpha1"
	applyconfigurationappsv1alpha1 "github.com/clusternet/clusternet/pkg/generated/applyconfiguration/apps/v1alpha1"
	scheme "github.com/clusternet/clusternet/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ManifestsGetter has a method to return a ManifestInterface.
// A group's client should implement this interface.
type ManifestsGetter interface {
	Manifests(namespace string) ManifestInterface
}

// ManifestInterface has methods to work with Manifest resources.
type ManifestInterface interface {
	Create(ctx context.Context, manifest *appsv1alpha1.Manifest, opts v1.CreateOptions) (*appsv1alpha1.Manifest, error)
	Update(ctx context.Context, manifest *appsv1alpha1.Manifest, opts v1.UpdateOptions) (*appsv1alpha1.Manifest, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*appsv1alpha1.Manifest, error)
	List(ctx context.Context, opts v1.ListOptions) (*appsv1alpha1.ManifestList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *appsv1alpha1.Manifest, err error)
	Apply(ctx context.Context, manifest *applyconfigurationappsv1alpha1.ManifestApplyConfiguration, opts v1.ApplyOptions) (result *appsv1alpha1.Manifest, err error)
	ManifestExpansion
}

// manifests implements ManifestInterface
type manifests struct {
	*gentype.ClientWithListAndApply[*appsv1alpha1.Manifest, *appsv1alpha1.ManifestList, *applyconfigurationappsv1alpha1.ManifestApplyConfiguration]
}

// newManifests returns a Manifests
func newManifests(c *AppsV1alpha1Client, namespace string) *manifests {
	return &manifests{
		gentype.NewClientWithListAndApply[*appsv1alpha1.Manifest, *appsv1alpha1.ManifestList, *applyconfigurationappsv1alpha1.ManifestApplyConfiguration](
			"manifests",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *appsv1alpha1.Manifest { return &appsv1alpha1.Manifest{} },
			func() *appsv1alpha1.ManifestList { return &appsv1alpha1.ManifestList{} },
		),
	}
}
