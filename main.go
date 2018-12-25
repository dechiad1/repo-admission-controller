package main

import (
  "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"k8s.io/api/admission/v1beta1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func handleAdmission(data []byte) *v1beta1.AdmissionReview {
	review := &v1beta1.AdmissionReview{}
	err := json.Unmarshal(data, review)
	if err != nil {
		panic(err)
	}

	raw := review.Request.Object.Raw
	pod := &v1.Pod{}
	err = json.Unmarshal(raw, pod)
	if err != nil {
		panic(err)
	}

	reviewStatus := v1beta1.AdmissionResponse{
		Allowed: true,
		Result: &metav1.Status{
			Message: "Welcome!",
		},
	}
	for _, container := range pod.Spec.Containers {
		if !strings.Contains(container.Image, "dechiada/") {
			reviewStatus.Allowed = false
			reviewStatus.Result = &metav1.Status {
				Reason: "can only pull registries in danny's repo",
			}
			log.Println("Blocking " + container.Image)
		} else {
			log.Println("Allowing " + container.Image )
		}
	}

	review.Response = &reviewStatus
	return review
}

func serve(w http.ResponseWriter, r *http.Request) {
		log.Printf("request from %s, %s", r.Host, r.URL.Path)
		var bodyBytes []byte

		if r.Body != nil {
			bodyBytes,_ = ioutil.ReadAll(r.Body)
		} else {
			//log this! body should always be present
			log.Println("no body!")
		}

		review := handleAdmission(bodyBytes)
		resp, err := json.Marshal(review)
		if err != nil {
			panic(err)
		}

		if _, err := w.Write(resp); err != nil {
			panic(err)
		}
}

func main() {
	http.HandleFunc("/", serve)

	cert := "/etc/certs/danny.crt"
	key := "/etc/certs/danny.key"

	log.Fatal(http.ListenAndServeTLS("0.0.0.0:8080", cert, key, nil))
}
