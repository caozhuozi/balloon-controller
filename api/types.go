package api

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type BalloonSpec struct {
	// ReleaseTime
	ReleaseTime string `json:"release_time"`
}

type Balloon struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec BalloonSpec `json:"spec"`

	Status BalloonStatus `json:"status"`
}

type BalloonStatus struct {
	Status string `json:"status"`
}

type BalloonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Balloon `json:"items"`
}
