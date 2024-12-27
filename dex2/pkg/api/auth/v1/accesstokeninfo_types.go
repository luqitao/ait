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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AccessTokenInfoSpec defines the desired state of AccessTokenInfo
type AccessTokenInfoSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Expiration  *int64 `json:"expiration,omitempty"`
	Timestamp   *int64 `json:"timestamp"`
	Description string `json:"description"`
}

// AccessTokenInfoStatus defines the observed state of UserSession
type AccessTokenInfoStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:subresource:status

// AccessTokenInfo is the Schema for the accesstokeninfos API
type AccessTokenInfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccessTokenInfoSpec   `json:"spec,omitempty"`
	Status AccessTokenInfoStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessTokenInfoList contains a list of AccessTokenInfo
type AccessTokenInfoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessTokenInfo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AccessTokenInfo{}, &AccessTokenInfoList{})
}
