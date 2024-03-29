/*
Copyright 2024.

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
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	DefaultReplicas int32  = 1
	DefaultPort     int32  = 3306
	DefaultImage    string = "mysql:8.3.0"
	DefaultSize     string = "500Mi"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MySQLSpec defines the desired state of MySQL
type MySQLSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Specify the name of the deployment operator
	Name string `json:"name,omitempty"`
	// Storage indicates the size of the Persistent Volume Claim.
	StorageSize resource.Quantity `json:"storage_size"`
	// Number of mysql instances in the cluster.
	Replicas int32 `json:"replicas,omitempty"`
	// Port specifies port for MySQL server.
	Port int32 `json:"port,omitempty"`
	// Image allows to specify mysql image.
	Image string `json:"image"`
	// Image version allows to specify mysql tag.
	Version string `json:"version,omitempty"`
	// Mysql database name.
	Database string `json:"database"`
	// Password is the name of Kubernetes secret containing the password.
	Password string `json:"password"`
	// Test controller message.
	Message string `json:"message,omitempty"`
}

// MySQLStatus defines the observed state of MySQL
type MySQLStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	State   string `json:"state,omitempty"`
	Message string `json:"message,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MySQL is the Schema for the mysqls API
type MySQL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MySQLSpec   `json:"spec,omitempty"`
	Status MySQLStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MySQLList contains a list of MySQL
type MySQLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQL `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MySQL{}, &MySQLList{})
}

// WithDefaults fills cluster missing fields with their default values.
func (c *MySQL) WithDefaults() {
	if c.Spec.Replicas == 0 {
		c.Spec.Replicas = DefaultReplicas
	}

	if c.Spec.Port == 0 {
		c.Spec.Port = DefaultPort
	}

	if c.Spec.Image == "" {
		c.Spec.Image = DefaultImage
	}

	if c.Spec.StorageSize.IsZero() {
		c.Spec.StorageSize = resource.MustParse(DefaultSize)
	}
}
