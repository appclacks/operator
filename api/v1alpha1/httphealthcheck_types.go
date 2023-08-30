/*
Copyright 2023.

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
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HTTPHealthcheckSpec defines the desired state of HTTPHealthcheck
type HTTPHealthcheckSpec struct {
	// +kubebuilder:validation:Optional
	Description string `json:"description"`
	Interval    string `json:"interval"`
	// +kubebuilder:validation:Optional
	Enabled bool   `json:"enabled"`
	Timeout string `json:"timeout"`

	ValidStatus []uint `json:"validStatus"`
	Target      string `json:"target"`
	Method      string `json:"method"`
	Port        uint   `json:"port"`
	Protocol    string `json:"protocol"`

	// +kubebuilder:validation:Optional
	Redirect bool `json:"redirect"`
	// +kubebuilder:validation:Optional
	Query map[string]string `json:"query,omitempty"`
	// +kubebuilder:validation:Optional
	Body string `json:"body,omitempty"`
	// +kubebuilder:validation:Optional
	BodyRegexp []string `json:"bodyRegexp,omitempty"`
	// +kubebuilder:validation:Optional
	Headers map[string]string `json:"headers,omitempty"`
	// +kubebuilder:validation:Optional
	Path string `json:"path,omitempty"`
	// +kubebuilder:validation:Optional
	Key string `json:"key,omitempty"`
	// +kubebuilder:validation:Optional
	Cert string `json:"cert,omitempty"`
	// +kubebuilder:validation:Optional
	Cacert string `json:"cacert,omitempty"`
	// +kubebuilder:validation:Optional
	Insecure bool `json:"insecure,omitempty"`
	// +kubebuilder:validation:Optional
	ServerName string `json:"serverName,omitempty"`
	// +kubebuilder:validation:Optional
	Host string `json:"host,omitempty"`
}

// HTTPHealthcheckStatus defines the observed state of HTTPHealthcheck
type HTTPHealthcheckStatus struct {
	State string `json:"state"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// HTTPHealthcheck is the Schema for the httphealthchecks API
type HTTPHealthcheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HTTPHealthcheckSpec   `json:"spec,omitempty"`
	Status HTTPHealthcheckStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HTTPHealthcheckList contains a list of HTTPHealthcheck
type HTTPHealthcheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HTTPHealthcheck `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HTTPHealthcheck{}, &HTTPHealthcheckList{})
}
