/*
Copyright 2021.

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

// CurioSpec defines the desired state of Curio
type CurioSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Curio. Edit curio_types.go to remove/update
	NameSpace                   string `json:"NAMESPACE,omitempty"`
	KumuluzeeEnvName            string `json:"KUMULUZEE_ENV_NAME,omitempty"`
	CurioIibLogLevel            string `json:"CURIO_IIB_LOG_LEVEL,omitempty"`
	CurioSiglaAplicacao         string `json:"CURIO_SIGLA_APLICACAO,omitempty"`
	CurioCacheConfiguracaoIib   string `json:"CURIO_CACHE_CONFIGURACAO_IIB,omitempty"`
	CurioCacheConfiguracaoIibId string `json:"CURIO_CACHE_CONFIGURACAO_IIB_ID,omitempty"`
	CurioOpProvedor             string `json:"CURIO_OP_PROVEDOR,omitempty"`
	KumuluzeeLogsLoggers0Name   string `json:"KUMULUZEE_LOGS_LOGGERS0_NAME,omitempty"`
	KumuluzeeLogsLoggers0Level  string `json:"KUMULUZEE_LOGS_LOGGERS0_LEVEL,omitempty"`
	CurioAplicacaoHost          string `json:"CURIO_APLICACAO_HOST,omitempty"`
	CurioModoDesenvolvimento    string `json:"CURIO_MODO_DESENVOLVIMENTO,omitempty"`
	CurioOpConsumidor           string `json:"CURIO_OP_CONSUMIDOR ,omitempty"`
	Versao                      int32  `json:"Versao,omitempty"`
}

// CurioStatus defines the observed state of Curio
type CurioStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Status string `json:"Status"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Curio is the Schema for the curios API
type Curio struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CurioSpec   `json:"spec,omitempty"`
	Status CurioStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CurioList contains a list of Curio
type CurioList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Curio `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Curio{}, &CurioList{})
}
