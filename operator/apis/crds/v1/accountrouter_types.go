package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"operators.kloudlite.io/lib/constants"
	rApi "operators.kloudlite.io/lib/operator"
)

// AccountRouterSpec defines the desired state of AccountRouter
type AccountRouterSpec struct {
	AccountRef string `json:"accountRef"`
	// +kubebuilder:validation:Enum=ClusterIP;LoadBalancer
	ServiceType    string            `json:"serviceType"`
	DefaultSSLCert string            `json:"defaultSSLCert,omitempty"`
	NodeSelector   map[string]string `json:"nodeSelector,omitempty"`
	// +kubebuilder:default=100
	MaxBodySizeInMB int       `json:"maxBodySizeInMB,omitempty"`
	RateLimit       RateLimit `json:"rateLimit,omitempty"`
	Https           Https     `json:"https,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AccountRouter is the Schema for the accountrouters API
type AccountRouter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccountRouterSpec `json:"spec,omitempty"`
	Status rApi.Status       `json:"status,omitempty"`
}

func (r *AccountRouter) GetStatus() *rApi.Status {
	return &r.Status
}

func (r *AccountRouter) GetEnsuredLabels() map[string]string {
	return map[string]string{
		constants.AccountRef: r.Spec.AccountRef,
	}
}

func (r *AccountRouter) GetEnsuredAnnotations() map[string]string {
	return map[string]string{}
}

// +kubebuilder:object:root=true

// AccountRouterList contains a list of AccountRouter
type AccountRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccountRouter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AccountRouter{}, &AccountRouterList{})
}
