# Validating Admission Controller Sample

This project contains a webapp, Dockerfile to build it and the needed scripts/yaml to deploy the controller and the ValidatingWebhookConfiguration into a Kubernetes Cluster.  This controller only allows pods to be deployed that come from a specific container repository.  

## Usage

In order to deploy this to a Kubernetes cluster, one must do the following:
1. Clone the Repository
2. Build and publish the controller images
   ```
   ./build (tags and pushed to the default registry of registry1.lab-1.cloud.local)
   ./build [registry-name] (option to specifcy the registry you want to tag and push to).
   ```
3. Deploy the controller and ValidationWebhookConfiguration into a Kubernetes Cluster
   ```
    cd deployments
    ./deploy (used the default registry of registry1.lab-1.cloud.local to pull the image from)
    ./deploy [registry-name] (option to specify the registry you want to pull the image from)
    ```
    You can add an argument of -saferepo [reponame] to the kube deployment template if you want have the controller allow from a specific repository.
