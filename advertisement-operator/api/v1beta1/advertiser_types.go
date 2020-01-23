/*

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

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type Resource struct {
	Image corev1.ContainerImage `json:"image"`
	Price resource.Quantity     `json:"price"`
}

type FreeResource struct {
	Cpu      resource.Quantity `json:"cpu"`
	CpuPrice resource.Quantity `json:"cpuPrice"`
	Ram      resource.Quantity `json:"ram"`
	RamPrice resource.Quantity `json:"ramPricePerMB"`
}

// AdvertiserSpec defines the desired state of Advertiser
type AdvertiserSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Advertiser. Edit Advertiser_types.go to remove/update
	ClusterId    string       `json:"clusterId"`
	Resources    []Resource   `json:"resources"`
	Availability FreeResource `json:"availability"`
	Timestamp    metav1.Time  `json:"timestamp"`
	Validity     metav1.Time  `json:"validity"`
}

// AdvertiserStatus defines the observed state of Advertiser
type AdvertiserStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// Advertiser is the Schema for the advertisers API
type Advertiser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AdvertiserSpec   `json:"spec,omitempty"`
	Status AdvertiserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AdvertiserList contains a list of Advertiser
type AdvertiserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Advertiser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Advertiser{}, &AdvertiserList{})
}