/*
Copyright 2021-2022 Red Hat, Inc.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	// DisplayName refers to the name that an application will be deployed with in App Studio.
	// Required.
	DisplayName string `json:"displayName"`

	// AppModelRepository refers to the git repository that will store the application model (a devfile)
	// Can be the same as GitOps repository.
	// A repository will be generated if this field is left blank.
	// Optional
	AppModelRepository ApplicationGitRepository `json:"appModelRepository,omitempty"`

	// GitOpsRepository refers to the git repository that will store the gitops resources.
	// Can be the same as App Model Repository.
	// A repository will be generated if this field is left blank.
	// Optional
	GitOpsRepository ApplicationGitRepository `json:"gitOpsRepository,omitempty"`

	// Description refers to a brief description of the application.
	// Optional
	Description string `json:"description,omitempty"`
}

// ApplicationGitRepository defines a git repository for a given Application resource (either appmodel or gitops)
type ApplicationGitRepository struct {
	// URL refers to the repository URL that should be used. If not specified, a GitOps repository under the
	// $GITHUB_ORG (defaults to redhat-appstudio-appdata) organization on GitHub will be generated by HAS.
	// Example: https://github.com/devfile-test/myrepo
	// +required
	URL string `json:"url"`

	// Branch corresponds to the branch in the repository that should be used
	// Example: devel
	// +optional
	Branch string `json:"branch,omitempty"`

	// Context corresponds to the context within the repository that should be used
	// Example: folderA/folderB/gitops
	// +optional
	Context string `json:"context,omitempty"`
}

// ApplicationStatus defines the observed state of Application
type ApplicationStatus struct {
	// Conditions is an array of the Application's status conditions
	Conditions []metav1.Condition `json:"conditions"`

	// Devfile corresponds to the devfile representation of the Application resource
	Devfile string `json:"devfile,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Application is the Schema for the applications API
// +kubebuilder:resource:path=applications,shortName=hasapp;ha;app
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.conditions[-1].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[-1].reason"
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationSpec   `json:"spec"`
	Status ApplicationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ApplicationList contains a list of Application
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
