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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// FeedApplyConfiguration represents a declarative configuration of the Feed type for use
// with apply.
type FeedApplyConfiguration struct {
	Kind       *string `json:"kind,omitempty"`
	APIVersion *string `json:"apiVersion,omitempty"`
	Namespace  *string `json:"namespace,omitempty"`
	Name       *string `json:"name,omitempty"`
}

// FeedApplyConfiguration constructs a declarative configuration of the Feed type for use with
// apply.
func Feed() *FeedApplyConfiguration {
	return &FeedApplyConfiguration{}
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *FeedApplyConfiguration) WithKind(value string) *FeedApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *FeedApplyConfiguration) WithAPIVersion(value string) *FeedApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *FeedApplyConfiguration) WithNamespace(value string) *FeedApplyConfiguration {
	b.Namespace = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *FeedApplyConfiguration) WithName(value string) *FeedApplyConfiguration {
	b.Name = &value
	return b
}
