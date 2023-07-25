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

// S3UserClaimSpec defines the desired state of S3UserClaim
type S3UserClaimSpec struct {
	// +kubebuilder:validation:Optional
	S3User string `json:"s3User,omitempty"`

	// +kubebuilder:validation:Optional
	S3UserClass string `json:"s3UserClass,omitempty"`

	// +kubebuilder:validation:Required
	ReadonlySecret string `json:"readonlySecret"`

	// +kubebuilder:validation:Required
	AdminSecret string `json:"adminSecret"`

	// +kubebuilder:validation:Optional
	Quota *UserQuota `json:"quota,omitempty"`
}

// S3UserClaimStatus defines the observed state of S3UserClaim
type S3UserClaimStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// S3UserClaim is the Schema for the s3userclaims API
type S3UserClaim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   S3UserClaimSpec   `json:"spec,omitempty"`
	Status S3UserClaimStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// S3UserClaimList contains a list of S3UserClaim
type S3UserClaimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3UserClaim `json:"items"`
}

func init() {
	SchemeBuilder.Register(&S3UserClaim{}, &S3UserClaimList{})
}

func (suc *S3UserClaim) GetS3UserClass() string {
	return suc.Spec.S3UserClass
}
