# Validatting Admission Controller

This Repo contains a app, Dockerfile and associate scripts/yaml files that will deploy an object of kind ValidatingWebhookConfiguration and its associated web application to a Kube Cluster.  

## Kube Cluster

In order to deploy this to a kube cluster, one must do the following:
1. Clone the repo 
2. Build the image (./build uses a default registry or ./build [registey-name])
3. Deploy (cd deployments; ./deploy)
