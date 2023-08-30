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

// DNSHealthcheckSpec defines the desired state of DNSHealthcheck
type DNSHealthcheckSpec struct {
	// +kubebuilder:validation:Optional
	Description string `json:"description,omitempty"`
	Interval    string `json:"interval"`
	// +kubebuilder:validation:Optional
	Enabled bool   `json:"enabled"`
	Timeout string `json:"timeout"`
	Domain  string `json:"domain,omitempty"`
	// +kubebuilder:validation:Optional
	ExpectedIPs []string `json:"expectedIPs,omitempty"`
}

// DNSHealthcheckStatus defines the observed state of DNSHealthcheck
type DNSHealthcheckStatus struct {
	State string `json:"state"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DNSHealthcheck is the Schema for the dnshealthchecks API
type DNSHealthcheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DNSHealthcheckSpec   `json:"spec,omitempty"`
	Status DNSHealthcheckStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DNSHealthcheckList contains a list of DNSHealthcheck
type DNSHealthcheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNSHealthcheck `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DNSHealthcheck{}, &DNSHealthcheckList{})
}
