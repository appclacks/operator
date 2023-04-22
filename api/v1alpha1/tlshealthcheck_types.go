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

// TLSHealthcheckSpec defines the desired state of TLSHealthcheck
type TLSHealthcheckSpec struct {
	// +kubebuilder:validation:Optional
	Description string `json:"description"`
	Interval    string `json:"interval"`
	// +kubebuilder:validation:Optional
	Enabled bool   `json:"enabled"`
	Timeout string `json:"timeout"`

	Target string `json:"target"`
	Port   uint   `json:"port"`

	// +kubebuilder:validation:Optional
	Key string `json:"key"`
	// +kubebuilder:validation:Optional
	Cert string `json:"cert"`
	// +kubebuilder:validation:Optional
	Cacert string `json:"cacert"`
	// +kubebuilder:validation:Optional
	ServerName string `json:"serverName"`
	// +kubebuilder:validation:Optional
	Insecure bool `json:"insecure"`
	// +kubebuilder:validation:Optional
	ExpirationDelay string `json:"expirationDelay"`
}

// TLSHealthcheckStatus defines the observed state of TLSHealthcheck
type TLSHealthcheckStatus struct {
	State string `json:"state"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TLSHealthcheck is the Schema for the tlshealthchecks API
type TLSHealthcheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TLSHealthcheckSpec   `json:"spec,omitempty"`
	Status TLSHealthcheckStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TLSHealthcheckList contains a list of TLSHealthcheck
type TLSHealthcheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TLSHealthcheck `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TLSHealthcheck{}, &TLSHealthcheckList{})
}
