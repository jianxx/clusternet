/*
Copyright 2021 The Clusternet Authors.

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

package localizer

import (
	"reflect"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"

	appsapi "github.com/clusternet/clusternet/pkg/apis/apps/v1alpha1"
)

func TestApplyOverrides(t *testing.T) {
	nameSpaceYaml := `
metadata:
    namespace: test2
`

	removeLabelYaml := `
metadata:
  labels:
    key: null
`

	helmDataYaml := `
address:
  planet: Earth
  street: 234 Spouter Inn Ct.
hole: black
`

	tests := []struct {
		name          string
		original      []byte
		originalChart []byte
		overrides     []appsapi.OverrideConfig
		want          []byte
		wantChart     []byte
	}{
		{
			name: "Helm",
			originalChart: []byte(`{
				"apiVersion": "apps.clusternet.io/v1alpha1",
				"kind": "HelmChart",
				"metadata": {
					"labels": {
						"apps.clusternet.io/config.group": "apps.clusternet.io"
					},
					"name": "cert-manager",
					"namespace": "clusternet-system"
				},
				"spec": {
					"chart": "cert-manager",
					"repo": "https://charts.bitnami.com/bitnami",
					"targetNamespace": "kube-system",
					"version": "0.5.8"
				}
			}`),
			original: []byte(`{
				"kind": "Guess",
				"address": {
					"city": "Nantucket",
					"street": "123 Spouter Inn Ct."
				},
				"boat": "pequod",
				"details": {
					"friends": ["Tashtego"]
				},
				"name": "Ishmael"
			}`),
			overrides: []appsapi.OverrideConfig{
				{
					Name:  "empty override",
					Type:  appsapi.HelmType,
					Value: ``,
				},
				{
					Name:  "add/update value - json format",
					Type:  appsapi.HelmType,
					Value: `{"address":{"country":"US","state":"MA"},"boat":"fighter"}`,
				},
				{
					Name:  "empty override with whitespaces",
					Type:  appsapi.HelmType,
					Value: `   `,
				},
				{
					Name:  "add/update value - yaml format",
					Type:  appsapi.HelmType,
					Value: helmDataYaml,
				},
				{
					Name:          "add label and annotation",
					Type:          appsapi.HelmType,
					OverrideChart: true,
					Value:         `{"metadata":{"labels":{"foo":"bar"},"annotations":{"foo":"bar"}}}`,
				},
				{
					Name:          "update chart repo and target namespace",
					Type:          appsapi.HelmType,
					OverrideChart: true,
					Value:         `{"spec":{"repo":"https://clusternet.github.io/charts","targetNamespace":"kube-public"}}`,
				},
				{
					Name:          "update chart name and version",
					Type:          appsapi.HelmType,
					OverrideChart: true,
					Value:         `{"spec":{"version":"0.6.1","chart":"my-cert-manager"}}`,
				},
			},
			wantChart: []byte(`{
				"apiVersion": "apps.clusternet.io/v1alpha1",
				"kind": "HelmChart",
				"metadata": {
					"annotations": {
						"foo": "bar"
					},
					"labels": {
						"apps.clusternet.io/config.group": "apps.clusternet.io",
						"foo": "bar"
					},
					"name": "cert-manager",
					"namespace": "clusternet-system"
				},
				"spec": {
					"chart": "my-cert-manager",
					"repo": "https://clusternet.github.io/charts",
					"targetNamespace": "kube-public",
					"version": "0.6.1"
				}
			}`),
			want: []byte(`{
				"kind": "Guess",
				"address": {
					"city": "Nantucket",
					"country": "US",
					"planet": "Earth",
					"state": "MA",
					"street": "234 Spouter Inn Ct."
                },
				"boat": "fighter",
				"details": {
					"friends": ["Tashtego"]
				},
				"hole": "black",
				"name": "Ishmael"
			}`),
		},
		{
			name:          "Helm with Empty Original",
			originalChart: []byte(``),
			original:      []byte(``),
			overrides: []appsapi.OverrideConfig{
				{
					Name:  "empty override",
					Type:  appsapi.HelmType,
					Value: `  `,
				},
				{
					Name:  "initial override",
					Type:  appsapi.HelmType,
					Value: `{"kind":"Guess","address":{"city":"Nantucket","street":"123 Spouter Inn Ct."},"boat":"pequod","details":{"friends":["Tashtego"]},"name":"Ishmael"}`,
				},
				{
					Name:  "add/update value - json format",
					Type:  appsapi.HelmType,
					Value: `{"address":{"country":"US","state":"MA"},"boat":"fighter"}`,
				},
				{
					Name:  "empty override with whitespaces",
					Type:  appsapi.HelmType,
					Value: `   `,
				},
				{
					Name:  "add/update value - yaml format",
					Type:  appsapi.HelmType,
					Value: helmDataYaml,
				},
				{
					Name:          "empty override with whitespaces for chart spec",
					Type:          appsapi.HelmType,
					OverrideChart: true,
					Value:         `  `,
				},
				{
					Name:          "add label and annotation",
					Type:          appsapi.HelmType,
					OverrideChart: true,
					Value:         `{"kind": "HelmChart","metadata":{"labels":{"foo":"bar"},"annotations":{"foo":"bar"}}}`,
				},
				{
					Name:          "update chart repo and target namespace",
					Type:          appsapi.HelmType,
					OverrideChart: true,
					Value:         `{"spec":{"repo":"https://clusternet.github.io/charts", "targetNamespace":"kube-public"}}`,
				},
				{
					Name:          "update chart name and version",
					Type:          appsapi.HelmType,
					OverrideChart: true,
					Value:         `{"spec":{"version":"0.6.1","chart":"my-cert-manager"}}`,
				},
			},
			wantChart: []byte(`{
				"kind": "HelmChart",
				"metadata": {
					"annotations": {
						"foo": "bar"
					},
					"labels": {
						"foo": "bar"
					}
				},
				"spec": {
					"chart": "my-cert-manager",
					"repo": "https://clusternet.github.io/charts",
					"targetNamespace": "kube-public",
					"version": "0.6.1",
				}
			}`),
			want: []byte(`{
				"kind": "Guess",
				"address": {
					"city": "Nantucket",
					"country": "US",
					"planet": "Earth",
					"state": "MA",
					"street": "234 Spouter Inn Ct."
				},
				"boat": "fighter",
				"details": {
					"friends": ["Tashtego"]
				},
				"hole": "black",
				"name": "Ishmael"
			}`),
		},
		{
			name: "HelmChart default overrideConfig",
			original: []byte(`{
				"apiVersion": "apps.clusternet.io/v1alpha1",
				"kind": "HelmChart",
				"metadata": {
					"labels": {
						"apps.clusternet.io/config.group": "apps.clusternet.io"
					},
					"name": "cert-manager",
					"namespace": "clusternet-system"
				},
				"spec": {
					"chart": "cert-manager",
					"repo": "https://charts.bitnami.com/bitnami",
					"targetNamespace": "kube-system",
					"version": "0.5.8"
				}
			}`),
			originalChart: []byte(``),
			overrides:     defaultChartOverrideConfigs,
			want: []byte(`{
				"apiVersion": "apps.clusternet.io/v1alpha1",
				"kind": "HelmChart",
				"metadata": {
					"labels": {
						"apps.clusternet.io/config.group": "apps.clusternet.io",
					},
					"name": "cert-manager",
					"namespace": "clusternet-system"
				},
				"spec": {
					"repo": "https://charts.bitnami.com/bitnami",
					"version": "0.5.8"
				}
			}`),
			wantChart: []byte(``),
		},
		{
			name: "JSONPatch and MergePatch",
			original: []byte(`{
				"apiVersion": "v1",
				"kind": "Pod",
				"metadata": {
					"name": "pod",
					"labels": {"app": "nginx"},
					"uid": "1234-678",
					"managedFields": [{
					  "manager": "kubectl-client-side-apply",
					  "operation": "Update",
					  "apiVersion": "apps.clusternet.io/v1alpha1",
					  "time": "2022-07-12T09:18:34Z",
					  "fieldsType": "FieldsV1",
					  "fieldsV1": {
						"f:metadata": {
						  "f:annotations": {
							".": {},
							"f:kubectl.kubernetes.io/last-applied-configuration": {}
						  }
						},
						"f:spec": {
						  ".": {},
						  "f:chart": {},
						  "f:repo": {},
						  "f:targetNamespace": {},
						  "f:version": {}
						}
					  }
					}, {
					  "manager": "clusternet-hub",
					  "operation": "Update",
					  "apiVersion": "apps.clusternet.io/v1alpha1",
					  "time": "2022-07-12T09:24:39Z",
					  "fieldsType": "FieldsV1",
					  "fieldsV1": {
						"f:status": {
						  ".": {},
						  "f:phase": {}
						}
					  },
					  "subresource": "status"
					}]
				},
				"spec": {
					"containers": [{
						"name":  "nginx",
						"image": "nginx:latest"
					}]
				},
				"status": {
					"a": "b",
					"some-value": [{
						"key1":  "value1",
						"key2": "value2"
					}]
				}
			}`),
			overrides: []appsapi.OverrideConfig{
				{
					Name:  "empty override with whitespaces",
					Type:  appsapi.MergePatchType,
					Value: `   `,
				},
				{
					Name:  "add namespace - json format",
					Type:  appsapi.MergePatchType,
					Value: `{"metadata":{"namespace":"test"}}`,
				},
				{
					Name:  "empty override with whitespaces",
					Type:  appsapi.JSONPatchType,
					Value: `   `,
				},
				{
					Name:  "add namespace - yaml format",
					Type:  appsapi.MergePatchType,
					Value: nameSpaceYaml,
				},
				{
					Name:  "replace container image - 1",
					Type:  appsapi.JSONPatchType,
					Value: `[{"op": "replace", "path": "/spec/containers/0/image", "value":"nginx:1.21.1"}]`,
				},
				{
					Name:  "add and update labels - json format",
					Type:  appsapi.MergePatchType,
					Value: `{"metadata":{"labels":{"foo":"bar","xyz":"def","key":"value"}}}`,
				},
				{
					Name:  "remove labels - json format",
					Type:  appsapi.MergePatchType,
					Value: `{"metadata":{"labels":{"xyz":null}}}`,
				},
				{
					Name:  "remove labels - yaml format",
					Type:  appsapi.MergePatchType,
					Value: removeLabelYaml,
				},
				{
					Name:  "replace container image - 2",
					Type:  appsapi.JSONPatchType,
					Value: `[{"op":"replace","path":"/spec/containers/0/image","value":"nginx:1.20.1"}]`,
				},
				{
					Name:  "inject new container - json format",
					Type:  appsapi.JSONPatchType,
					Value: `[{"op":"add","path": "/spec/containers/1","value":{"name":"injected-container","image":"redis:6.2.5"}}]`,
				},
			},
			want: []byte(`{
				"apiVersion": "v1",
				"kind": "Pod",
				"metadata": {
					"name": "pod",
					"namespace": "test2",
					"labels": {
						"app": "nginx",
						"foo": "bar",
					},
				},
				"spec": {
					"containers": [{
							"name": "nginx",
							"image": "nginx:1.20.1"
						},
						{
							"name": "injected-container",
							"image": "redis:6.2.5"
						}
					]
				}
			}`),
		},
		{
			name: "FieldPatch",
			original: []byte(`{
				"apiVersion": "v1",
				"kind": "ConfigMap",
				"metadata": {
					"name": "example-configmap",
					"namespace": "default"
				},
				"data": {
					"server.conf": "{\n    \"db_address\": \"127.0.0.1:3306\",\n    \"region_id\": \"us-west-1\",\n    \"zone_id\": \"zone-1\"\n}"
				}
			}`),
			overrides: []appsapi.OverrideConfig{
				{
					Name:  "empty override with whitespaces",
					Type:  appsapi.MergePatchType,
					Value: `   `,
				},
				{
					Name:        "overridge db address - json format",
					Type:        appsapi.FieldJSONPatchType,
					FieldPath:   "/data/server.conf",
					FieldFormat: appsapi.JSONFormat,
					Value:       `[{"op":"replace","path": "/db_address","value":"localhost:3306"}]`,
				},
			},
			want: []byte(`{
				"apiVersion": "v1",
				"kind": "ConfigMap",
				"metadata": {
					"name": "example-configmap",
					"namespace": "default"
				},
				"data": {
					"server.conf": "{\n    \"db_address\": \"localhost:3306\",\n    \"region_id\": \"us-west-1\",\n    \"zone_id\": \"zone-1\"\n}"
				}
			}`),
		},
		{
			name: "FieldPatch",
			original: []byte(`{
				"apiVersion": "v1",
				"kind": "ConfigMap",
				"metadata": {
					"name": "example-configmap",
					"namespace": "default"
				},
				"data": {
					"server.conf": "db_address: 127.0.0.1:5432\nregion_id: us-west-1\nzone_id: zone-1\n"
				}
			}`),
			overrides: []appsapi.OverrideConfig{
				{
					Name:  "empty override with whitespaces",
					Type:  appsapi.MergePatchType,
					Value: `   `,
				},
				{
					Name:        "overridge db address - json format",
					Type:        appsapi.FieldMergePatchType,
					FieldPath:   "/data/server.conf",
					FieldFormat: appsapi.YAMLFormat,
					Value:       "region_id: us-west-2\nzone_id: zone-2\n",
				},
			},
			want: []byte(`{
				"apiVersion": "v1",
				"kind": "ConfigMap",
				"metadata": {
					"name": "example-configmap",
					"namespace": "default"
				},
				"data": {
					"server.conf": "db_address: 127.0.0.1:5432\nregion_id: us-west-2\nzone_id: zone-2\n"
				}
			}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotChart, err := applyOverrides(tt.original, tt.originalChart, tt.overrides)
			if err != nil {
				t.Errorf("applyOverrides() error = %v", err)
				return
			}

			gotObj := &unstructured.Unstructured{Object: map[string]interface{}{}}
			wantObj := &unstructured.Unstructured{Object: map[string]interface{}{}}
			gotChartObj := &unstructured.Unstructured{Object: map[string]interface{}{}}
			wantChartObj := &unstructured.Unstructured{Object: map[string]interface{}{}}

			if err = yaml.Unmarshal(gotChart, &gotChartObj); err != nil {
				t.Fatalf("error decoding: %v", err)
			}
			if err = yaml.Unmarshal(tt.wantChart, &wantChartObj); err != nil {
				t.Fatalf("error decoding: %v", err)
			}
			if err = yaml.Unmarshal(got, &gotObj); err != nil {
				t.Fatalf("error decoding: %v", err)
			}
			if err = yaml.Unmarshal(tt.want, &wantObj); err != nil {
				t.Fatalf("error decoding: %v", err)
			}

			if !reflect.DeepEqual(gotChartObj, wantChartObj) {
				t.Errorf("applyOverrides() gotChart %s, wantChart %s", gotChartObj, wantChartObj)
			}
			if !reflect.DeepEqual(gotObj, wantObj) {
				t.Errorf("applyOverrides() got %s, want %s", gotObj, wantObj)
			}
		})
	}
}

func TestApplyJSONPatch(t *testing.T) {
	tests := []struct {
		name          string
		cur           []byte
		overrideBytes []byte
		want          []byte
		wantErr       bool
	}{
		{
			name:          "remove nonexistent key (/spec/chart)",
			cur:           []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"spec":{"version":"1.8.0"}}`),
			overrideBytes: []byte(defaultChartOverrideConfigs[0].Value),
			want:          []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"spec":{"version":"1.8.0"}}`),
			wantErr:       false,
		},
		{
			name:          "remove nonexistent key (/spec/targetNamespace)",
			cur:           []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"spec":{"version":"1.8.0"}}`),
			overrideBytes: []byte(defaultChartOverrideConfigs[1].Value),
			want:          []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"spec":{"version":"1.8.0"}}`),
			wantErr:       false,
		},
		{
			name:          "remove key (/spec/version)",
			cur:           []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"spec":{"version":"1.8.0"}}`),
			overrideBytes: []byte(`[{"path":"/spec/version","op":"remove"}]`),
			want:          []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"spec":{}}`),
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := applyJSONPatch(tt.cur, tt.overrideBytes, "")
			if (err != nil) != tt.wantErr {
				t.Errorf("applyJSONPatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyJSONPatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApplyFieldPatch(t *testing.T) {
	tests := []struct {
		name           string
		cur            []byte
		overrideBytes  []byte
		fieldPatchType appsapi.OverrideType
		fieldPath      string
		fieldFormat    appsapi.FieldFormatType
		want           []byte
		wantErr        bool
	}{
		{
			name:           "json patch add mysql timeout to configmap in json",
			cur:            []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"mydata":"{\"mysql\":{\"address\":\"127.0.0.1:3306\"}}"}`),
			overrideBytes:  []byte(`[{"path":"/mysql/timeout","op":"add","value": 5}]`),
			fieldPatchType: appsapi.FieldJSONPatchType,
			fieldPath:      "/mydata",
			fieldFormat:    appsapi.JSONFormat,
			want:           []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"mydata":"{\n    \"mysql\": {\n        \"address\": \"127.0.0.1:3306\",\n        \"timeout\": 5\n    }\n}"}`),
			wantErr:        false,
		},
		{
			name:           "merge patch mysql timeout to configmap in json",
			cur:            []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"mydata":"{\"mysql\":{\"address\":\"127.0.0.1:3306\"}}"}`),
			overrideBytes:  []byte(`{"mysql":{"timeout":5}}`),
			fieldPatchType: appsapi.FieldMergePatchType,
			fieldPath:      "/mydata",
			fieldFormat:    appsapi.JSONFormat,
			want:           []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"mydata":"{\n    \"mysql\": {\n        \"address\": \"127.0.0.1:3306\",\n        \"timeout\": 5\n    }\n}"}`),
			wantErr:        false,
		},
		{
			name:           "merge patch add mysql timeout to configmap in yaml",
			cur:            []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"mydata":"mysql:\n  address: 127.0.0.1:3306\n"}`),
			overrideBytes:  []byte(`{"mysql":{"timeout":5}}`),
			fieldPatchType: appsapi.FieldMergePatchType,
			fieldPath:      "/mydata",
			fieldFormat:    appsapi.YAMLFormat,
			want:           []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"mydata":"mysql:\n  address: 127.0.0.1:3306\n  timeout: 5\n"}`),
			wantErr:        false,
		},
		{
			name:           "json patch add mysql timeout to configmap in yaml",
			cur:            []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"mydata":"mysql:\n  address: 127.0.0.1:3306\n"}`),
			overrideBytes:  []byte(`[{"path":"/mysql/timeout","op":"add","value": 5}]`),
			fieldPatchType: appsapi.FieldJSONPatchType,
			fieldPath:      "/mydata",
			fieldFormat:    appsapi.YAMLFormat,
			want:           []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"mydata":"mysql:\n  address: 127.0.0.1:3306\n  timeout: 5\n"}`),
			wantErr:        false,
		},
		{
			name:           "add mysql timeout to configmap in yaml when field not exist",
			cur:            []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"mydata":"mysql:\n  address: 127.0.0.1:3306\n"}`),
			overrideBytes:  []byte(`{"mysql":{"timeout":5}}`),
			fieldPatchType: appsapi.FieldMergePatchType,
			fieldPath:      "/field-not-exist",
			fieldFormat:    appsapi.YAMLFormat,
			want:           nil,
			wantErr:        true,
		},
		{
			name:           "invalid yaml field value",
			cur:            []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"mydata":"mysql:  address: 127.0.0.1:3306\n"}`),
			overrideBytes:  []byte(`{"mysql":{"timeout":5}}`),
			fieldPatchType: appsapi.FieldMergePatchType,
			fieldPath:      "/mydata",
			fieldFormat:    appsapi.YAMLFormat,
			want:           nil,
			wantErr:        true,
		},
		{
			name:           "invalid json field value",
			cur:            []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"mydata":"{\"mysql\"\"\"\":{\"address\":\"127.0.0.1:3306\"}}"}`),
			overrideBytes:  []byte(`{"mysql":{"timeout":5}}`),
			fieldPatchType: appsapi.FieldMergePatchType,
			fieldPath:      "/mydata",
			fieldFormat:    appsapi.JSONFormat,
			want:           nil,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []byte
			var err error
			if tt.fieldPatchType == appsapi.FieldJSONPatchType {
				got, err = applyFieldJSONPatch(tt.cur, tt.fieldPath, tt.fieldFormat, tt.overrideBytes)
			}
			if tt.fieldPatchType == appsapi.FieldMergePatchType {
				got, err = applyFieldMergePatch(tt.cur, tt.fieldPath, tt.fieldFormat, tt.overrideBytes)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("applyFieldPatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyFieldPatch() got = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func TestFindField(t *testing.T) {
	testCases := []struct {
		name      string
		cur       []byte
		fieldPath string
		result    string
		wantErr   bool
	}{
		{
			name:      "find metadata.labels",
			cur:       []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"spec":{"version":"1.8.0"}}`),
			fieldPath: "/metadata/labels",
			result:    "",
			wantErr:   true,
		},
		{
			name:      "find metadata.labels.another-label",
			cur:       []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"spec":{"version":"1.8.0"}}`),
			fieldPath: "/metadata/labels/another-label",
			result:    "another-value",
		},
		{
			name:      "find some field in string array",
			cur:       []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"spec":{"commands":["sleep", "999999"],"version":"1.8.0"}}`),
			fieldPath: "/spec/commands/1",
			result:    "999999",
		},
		{
			name:      "find some field in object array",
			cur:       []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"spec":{"containers":[{"image":"nginx:latest"}, {"image":"redis:latest"}],"version":"1.8.0"}}`),
			fieldPath: "/spec/containers/1/image",
			result:    "redis:latest",
		},
		{
			name:      "error field path",
			cur:       []byte(`{"metadata":{"labels":{"another-label":"another-value","some-label":"some-value"}},"spec":{"containers":[{"image":"nginx:latest"}, {"image":"redis:latest"}],"version":"1.8.0"}}`),
			fieldPath: "/spec/containers/error/image",
			result:    "",
			wantErr:   true,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findField(tt.cur, tt.fieldPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("findField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.result {
				t.Errorf("findField() got = %v, want %v", got, tt.result)
			}
		})
	}
}
