/*
Copyright AppsCode Inc. and Contributors

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConversionOperation describes a cr to tf conversion operation.
type CustomResourceToTerraformConversionOperation struct {
	metav1.TypeMeta `json:",inline"`
	// Request describes the attributes for the admission request.
	// +optional
	Request *CustomResourceToTerraformConversionRequest `json:"request,omitempty"`
	// Response describes the attributes for the admission response.
	// +optional
	Response *CustomResourceToTerraformConversionResponse `json:"response,omitempty"`
}

// CustomResourceToTerraformConversionRequest describes the request for the cr to tf conversion operation.
type CustomResourceToTerraformConversionRequest struct {
	Object runtime.RawExtension `json:"object"`
}

// CustomResourceToTerraformConversionResponse describes the response for the cr to tf conversion operation.
type CustomResourceToTerraformConversionResponse struct {
	Terraform string `json:"terraform"`
	State     string `json:"state"`
}