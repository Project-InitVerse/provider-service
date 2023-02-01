package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// InventoryState define string type
type InventoryState string

// Define inventory type
const (
	InventoryStatePulled = InventoryState("PULLED")
	InventoryStateError  = InventoryState("ERROR")
)

// InventoryRequest is struct
// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type InventoryRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   InventoryRequestSpec   `json:"spec,omitempty"`
	Status InventoryRequestStatus `json:"status,omitempty"`
}

// InventoryRequestSpec stores name
type InventoryRequestSpec struct {
	Name string `json:"name"`
}

// InventoryRequestStatus stores state & msg
type InventoryRequestStatus struct {
	State   string `json:"state,omitempty"`
	Message string `json:"message,omitempty"`
}

// InventoryRequestList stores metadata and items list of storage class states
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type InventoryRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []InventoryRequest `json:"items"`
}

// Inventory is struct
// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Inventory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   InventorySpec   `json:"spec,omitempty"`
	Status InventoryStatus `json:"status,omitempty"`
}

// InventoryList stores metadata and items list of storage class states
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type InventoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Inventory `json:"items"`
}

// InventoryStatus stores state and list of msg
type InventoryStatus struct {
	State    InventoryState `json:"state,omitempty"`
	Messages []string       `json:"message,omitempty"`
}

// InventoryClusterStorage is struct
type InventoryClusterStorage struct {
	Class        string `json:"class,omitempty"`
	ResourcePair `json:",inline"`
}

// InventorySpec is struct
type InventorySpec struct {
	Storage []InventoryClusterStorage `json:"storage"`
}

// ResourcePair stores allocatable and allocated
type ResourcePair struct {
	Allocatable uint64 `json:"allocatable"`
	Allocated   uint64 `json:"allocated"`
}
