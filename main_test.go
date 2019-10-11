package main

import (
	"encoding/json"
	"log"
	"testing"

	"k8s.io/api/admission/v1beta1"
	v1 "k8s.io/api/core/v1"
)

func TestHandleAdmission(t *testing.T) {
	// Create an Admission Review to populate and pass to handleAdmission
	review := &v1beta1.AdmissionReview{}

	// First we need an AdmissionRequest
	request := &v1beta1.AdmissionRequest{}

	// Create a pod, podspec, containerlist and container where the image is from an approved registry
	pod := &v1.Pod{
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Image: "registry1.lab-1.cloud.local/api-warning-controller:v1.2.7",
				},
			},
		},
	}

	// Set the AdmissionsReview Request to our crafted AdmissionRequest and set the Pod's raw data in the AdmissionRequest
	review.Request = request
	review.Request.Object.Raw, _ = json.Marshal(pod)

	// Convert the AdmissionRequest to a byte slice so it can be passed to the function for testing
	var data []byte
	data, _ = json.Marshal(review)

	// Get the response back as an AdmissionReview and print the output for a few test repos, some approved and others not.
	testSafeRepos := []string{"registry1.lab-1.cloud.local", "registry2.lab-1.cloud.local"}
	for _, safeRepo := range testSafeRepos {
		log.Printf("Testing with safeRepo set to: " + safeRepo)
		reviewAnswer := handleAdmission(data, safeRepo)
		log.Printf("Admission status is Allowed:%t with Message:%s.", reviewAnswer.Response.Allowed, reviewAnswer.Response.Result.Reason)
	}
}
