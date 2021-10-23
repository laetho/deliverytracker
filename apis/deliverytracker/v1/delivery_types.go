/*
Copyright 2021.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeliverySpec defines the desired state of Delivery
type DeliverySpec struct {
	Name        string                  `json:"name"`
	Repository  string                  `json:"repository"`
	Environment map[string]DeliveryInfo `json:"environment,omitempty"`
}

type DeliveryInfo struct {
	SemVer     string   `json:"semVer,omitempty"`
	GitHash    string   `json:"gitHash"`
	GitTags    []string `json:"gitTags,omitempty"`
	BehindHead int32    `json:"behindHead"`
}

// DeliveryStatus defines the observed state of Delivery
type DeliveryStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// Delivery is the Schema for the deliveries API
type Delivery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeliverySpec   `json:"spec,omitempty"`
	Status DeliveryStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
// DeliveryList contains a list of Delivery
type DeliveryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Delivery `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Delivery{}, &DeliveryList{})
}
