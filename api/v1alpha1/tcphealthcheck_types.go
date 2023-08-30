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

// TCPHealthcheckSpec defines the desired state of TCPHealthcheck
type TCPHealthcheckSpec struct {
	// +kubebuilder:validation:Optional
	Description string `json:"description,omitempty"`
	Interval    string `json:"interval"`
	// +kubebuilder:validation:Optional
	Enabled bool   `json:"enabled"`
	Timeout string `json:"timeout"`
	Target  string `json:"target"`
	Port    uint   `json:"port"`
	// +kubebuilder:validation:Optional
	ShouldFail bool `json:"shouldFail"`
}

// TCPHealthcheckStatus defines the observed state of TCPHealthcheck
type TCPHealthcheckStatus struct {
	State string `json:"state"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TCPHealthcheck is the Schema for the tcphealthchecks API
type TCPHealthcheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TCPHealthcheckSpec   `json:"spec,omitempty"`
	Status TCPHealthcheckStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TCPHealthcheckList contains a list of TCPHealthcheck
type TCPHealthcheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TCPHealthcheck `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TCPHealthcheck{}, &TCPHealthcheckList{})
}
