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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ChatgptSpec defines the desired state of Chatgpt
type ChatgptSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Chatgpt. Edit chatgpt_types.go to remove/update
	Name      string    `json:"name,omitempty"`
	DbConfig  DbConfig  `json:"dbConfig,omitempty"`
	BotConfig BotConfig `json:"botConfig,omitempty"`
}

type DbConfig struct {
	Host              string                   `json:"host,omitempty"`
	Port              int                      `json:"port,omitempty"`
	Username          string                   `json:"username,omitempty"`
	PasswordSecretRef corev1.SecretKeySelector `json:"passwordSecretRef,omitempty"`
}

type BotConfig struct {
	// Default value for the 'NewDialogTimeout' field is 600 seconds
	// +kubebuilder:default:=600
	NewDialogTimeout int `json:"newDialogTimeout,omitempty"`
	// +kubebuilder:default:=true
	EnabeMsgStreming bool `json:"enableMsgStreaming,omitempty"`
	// +kubebuilder:default:=5
	NchatPerPage int `json:"nChatPerPage,omitempty"`
	// +kubebuilder:default:=0.002
	ChatgptPriceToken string `json:"chatgptPriceToken,omitempty"`
	// +kubebuilder:default:=0.02
	GptPriceToken string `json:"gptPriceToken,omitempty"`
	// +kubebuilder:default:=0.006
	WhisperPRice   string         `json:"whisperPrice,omitempty"`
	BotCredentials BotCredentials `json:"botCredentials,omitempty"`
}

type BotCredentials struct {
	TelgramToken corev1.SecretKeySelector `json:"telgramToken,omitempty"`
	OpenaiToken  corev1.SecretKeySelector `json:"openaiToken,omitempty"`
}

// ChatgptStatus defines the observed state of Chatgpt
type ChatgptStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Chatgpt is the Schema for the chatgpts API
type Chatgpt struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ChatgptSpec   `json:"spec,omitempty"`
	Status ChatgptStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ChatgptList contains a list of Chatgpt
type ChatgptList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Chatgpt `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Chatgpt{}, &ChatgptList{})
}
