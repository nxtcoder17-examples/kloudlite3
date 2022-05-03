package v1

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ManagedServiceSpec defines the desired state of ManagedService
type ManagedServiceSpec struct {
	Type   string            `json:"type"`
	Inputs map[string]string `json:"inputs"`
}

// ManagedServiceStatus defines the observed state of ManagedService
type ManagedServiceStatus struct {
	Generation int64              `json:"generation,omitempty"`
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ManagedService is the Schema for the managedservices API
type ManagedService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagedServiceSpec   `json:"spec,omitempty"`
	Status ManagedServiceStatus `json:"status,omitempty"`
}

func (m *ManagedService) LogRef() string {
	return fmt.Sprintf("%s/%s/%s", m.Namespace, m.Kind, m.Name)
}

func (m *ManagedService) HasLabels() bool {
	if _, ok := m.Labels["msvc.kloudlite.io/type"]; !ok {
		return false
	}
	return true
}

func (m *ManagedService) EnsureLabels() {
	m.SetLabels(map[string]string{
		"msvc.kloudlite.io/type": m.Spec.Type,
	})
}

// +kubebuilder:object:root=true

// ManagedServiceList contains a list of ManagedService
type ManagedServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManagedService{}, &ManagedServiceList{})
}
