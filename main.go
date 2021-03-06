package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"k8s.io/api/admission/v1beta1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func handleAdmission(data []byte, safeRepo string) *v1beta1.AdmissionReview {
	review := &v1beta1.AdmissionReview{}
	err := json.Unmarshal(data, review)
	if err != nil {
		panic(err)
	}

	raw := review.Request.Object.Raw
	pod := &v1.Pod{}
	err = json.Unmarshal(raw, pod)
	if err != nil {
		log.Println("invalid pod spec")
		return nil
	}

	reviewStatus := v1beta1.AdmissionResponse{
		Allowed: true,
		Result: &metav1.Status{
			Message: "Welcome!",
			Reason:  "Image is coming from a safe repository",
		},
	}

	for _, container := range pod.Spec.Containers {
		if !strings.Contains(container.Image, safeRepo) {
			reviewStatus.Allowed = false
			reviewStatus.Result = &metav1.Status{
				Reason: "Can only pull images from the safe Repository",
			}
			log.Println("Blocking " + container.Image)
		} else {
			log.Println("Allowing " + container.Image)
		}
	}

	review.Response = &reviewStatus
	return review
}

func serve(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from %s, %s", r.Host, r.URL.Path)
	var bodyBytes []byte

	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
		safeRepo := flag.Lookup("saferepo").Value.(flag.Getter).Get().(string) + "/"
		review := handleAdmission(bodyBytes, safeRepo)
		if review == nil {
			w.Write([]byte("No request body - body must contain valid kube pod creation spec"))
		} else {
			resp, err := json.Marshal(review)
			if err != nil {
				log.Println("Marshalling admission review results into a response")
				panic(err)
			}

			if _, err := w.Write(resp); err != nil {
				log.Println("Writing response failed")
				log.Println("writing response failed")
				panic(err)
			}
		}
	} else {
		//log this! body should always be present
		log.Println("No body!")
		w.Write([]byte("No request body - body must contain valid kube pod creation spec"))
	}

}

func parseFlags() {
	var saferepo string
	flag.StringVar(&saferepo, "saferepo", "registry1.lab-1.cloud.local", "Repository to white list.")
	flag.Parse()
}

func main() {
	parseFlags()
	log.Println("Controller starting and allowing images from registry: " + flag.Lookup("saferepo").Value.(flag.Getter).Get().(string))
	http.HandleFunc("/", serve)

	cert := "/etc/certs/tls.crt"
	key := "/etc/certs/tls.key"

	log.Fatal(http.ListenAndServeTLS("0.0.0.0:8080", cert, key, nil))
}
