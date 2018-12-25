# Dynamic Webhook Admission Controller

This Repo contains a webapp, dockerfile and kube yaml files that will deploy an object of kind ValidatingWebhookConfiguration and its associated web application to a kube cluster.  

## Kube Cluster

In order to deploy this to a kube cluster, one must do the following:
1. clone the repo 
2. build the image 
3. create the ca bundle
      * openssl req -newkey rsa:2048 -nodes -keyout [keyname].key -out [csr name].csr
      * openssl x509 -signkey [keyname].key -in [csr name].csr -req -days 365 -out [cert name].crt
4. Get the base64 coded contents of the authoritative CA
      * cat [cert name].crt | openssl enc -base64 -A
5. Add the encoded contents to the ValidatingWebhookConfiguration yaml file's caBundle attribute
6. Create a kube secret containing both the key and the cert - where the key shall act as the name of the file in the mounted directory
      * kubectl-eks -n [namespace] create secret generic [secret name] --from-file=[contained key name]=[key name] --from-file=[cert name]=[contained cert name]
7. Update the volume mount and volumes portions of the web app's yaml to accurately reflect the secret  
8. Deploy, woo!


## Run Locally 
In main.go -
1. Point to the ca files on the local directory 
2. Have the server listen on 127.0.0.1 instead of 0.0.0.0
