package v1alpha1

import (
	advtypes "github.com/liqotech/liqo/apis/sharing/v1alpha1"
	"github.com/liqotech/liqo/pkg/crdClient"
	object_references "github.com/liqotech/liqo/pkg/object-references"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ResourceRequestSpec defines the desired state of ResourceRequest
type ResourceRequestSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foreign Cluster Identity
	ClusterIdentity ClusterIdentity `json:"clusterIdentity"`
	// Namespace where Liqo is deployed
	Namespace string `json:"namespace"`
	// KubeConfig file (with Advertisement creation role) secret reference
	KubeConfigRef *v1.ObjectReference `json:"kubeConfigRef,omitempty"`
	// Local auth service address
	AuthUrl string `json:"authUrl"`
}

// ResourceRequestStatus defines the observed state of ResourceRequest
type ResourceRequestStatus struct {
	BroadcasterRef      *object_references.DeploymentReference `json:"broadcasterRef,omitempty"`
	AdvertisementStatus advtypes.AdvPhase                      `json:"advertisementStatus,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceRequest is the Schema for the ResourceRequests API
type ResourceRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceRequestSpec   `json:"spec,omitempty"`
	Status ResourceRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceRequestList contains a list of ResourceRequest
type ResourceRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ResourceRequest{}, &ResourceRequestList{})

	if err := AddToScheme(scheme.Scheme); err != nil {
		panic(err)
	}
	crdClient.AddToRegistry("resourcerequests", &ResourceRequest{}, &ResourceRequestList{}, nil, schema.GroupResource{
		Group:    v1.SchemeGroupVersion.Group,
		Resource: "resourcerequests",
	})
}
