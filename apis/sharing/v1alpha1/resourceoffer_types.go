package v1alpha1

import (
	object_references "github.com/liqotech/liqo/pkg/object-references"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ResourceOfferSpec defines the desired state of ResourceOffer
type ResourceOfferSpec struct {
	// ClusterId is the identifier of the cluster that is sending this ResourceOffer.
	// It is the uid of the first master node in you cluster.
	ClusterId string `json:"clusterId"`
	// Images is the list of the images already stored in the cluster.
	Images []corev1.ContainerImage `json:"images,omitempty"`
	// LimitRange contains the limits for every kind of resource (cpu, memory...).
	LimitRange corev1.LimitRangeSpec `json:"limitRange,omitempty"`
	// ResourceQuota contains the quantity of resources made available by the cluster.
	ResourceQuota corev1.ResourceQuotaSpec `json:"resourceQuota,omitempty"`
	// Labels contains the label to be added to the virtual node.
	Labels map[string]string `json:"labels,omitempty"`
	// Neighbors is a map where the key is the name of a virtual node (representing a foreign cluster) and the value are the resources allocatable on that node.
	Neighbors map[corev1.ResourceName]corev1.ResourceList `json:"neighbors,omitempty"`
	// Properties can contain any additional information about the cluster.
	Properties map[corev1.ResourceName]string `json:"properties,omitempty"`
	// Prices contains the possible prices for every kind of resource (cpu, memory, image).
	Prices        corev1.ResourceList    `json:"prices,omitempty"`
	KubeConfigRef corev1.SecretReference `json:"kubeConfigRef"`
	// Timestamp is the time instant when this ResourceOffer was created.
	Timestamp metav1.Time `json:"timestamp"`
	// TimeToLive is the time instant until this ResourceOffer will be valid.
	// If not refreshed, an ResourceOffer will expire after 30 minutes.
	TimeToLive metav1.Time `json:"timeToLive"`
}

// OfferPhase describes the phase of the ResourceOffer
type OfferPhase string

const (
	ResourceOfferAccepted OfferPhase = "Accepted"
	ResourceOfferRefused  OfferPhase = "Refused"
)

// ResourceOfferStatus defines the observed state of ResourceOffer
type ResourceOfferStatus struct {
	// ResourceOfferStatus is the status of this ResourceOffer.
	// When the offer is created it is checked by the operator, which sets this field to "Accepted" or "Refused" on tha base of cluster configuration.
	// If the ResourceOffer is accepted a virtual-kubelet for the foreign cluster will be created.
	// +kubebuilder:validation:Enum="";"Accepted";"Refused"
	ResourceOfferStatus OfferPhase `json:"resourceOfferStatus"`
	// VkCreated indicates if the virtual-kubelet for this ResourceOffer has been created or not.
	VkCreated bool `json:"vkCreated"`
	// VkReference is a reference to the deployment running the virtual-kubelet.
	VkReference object_references.DeploymentReference `json:"vkReference,omitempty"`
	// VnodeReference is a reference to the virtual node linked to this ResourceOffer
	VnodeReference object_references.NodeReference `json:"vnodeReference,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName="offer"

// ResourceOffer is the Schema for the resourceOffers API
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.resourceOfferStatus`
// +kubebuilder:printcolumn:name="Expiration",type=string,JSONPath=`.spec.timeToLive`
// +kubebuilder:printcolumn:name="VkCreated",type=boolean,JSONPath=`.status.vkCreated`
type ResourceOffer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceOfferSpec   `json:"spec,omitempty"`
	Status ResourceOfferStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceOfferList contains a list of ResourceOffer
type ResourceOfferList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceOffer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ResourceOffer{}, &ResourceOfferList{})
}
