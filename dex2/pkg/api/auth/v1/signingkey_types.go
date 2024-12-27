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

package v1

import (
	"gopkg.in/square/go-jose.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SigningKeySpec defines the desired state of SigningKey
type SigningKeySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Key for creating and verifying signatures. These may be nil.
	SigningKey    *JSONWebKey `json:"signingKey,omitempty"`
	SigningKeyPub *JSONWebKey `json:"signingKeyPub,omitempty"`
}

// SigningKeyStatus defines the observed state of SigningKey
type SigningKeyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// JSONWebKey
type JSONWebKey struct {
	jose.JSONWebKey
}

//
// DeepCopyInto creates a deep-copy of the Time value.  The underlying time.Time
// type is effectively immutable in the time API, so it is safe to
// copy-by-assign, despite the presence of (unexported) Pointer fields.
func (t *JSONWebKey) DeepCopyInto(out *JSONWebKey) {
	*out = *t
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced

// SigningKey is the Schema for the signingkeys API
type SigningKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SigningKeySpec   `json:"spec,omitempty"`
	Status SigningKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SigningKeyList contains a list of SigningKey
type SigningKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SigningKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SigningKey{}, &SigningKeyList{})
}
