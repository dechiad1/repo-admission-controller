---
apiVersion: v1
kind: Service
metadata: 
  name: repo-whitelist
  namespace: repo-whitelist
  labels:
    name: repo-whitelist
spec:
  ports:
  -  name: webhook
     port: 443
     targetPort: 8080
  selector:
    name: repo-whitelist
---
apiVersion: apps/v1beta1
kind: Deployment
metadata: 
  name: repo-whitelist
  namespace: repo-whitelist
  labels: 
    name: repo-whitelist
spec:
  replicas: 1
  template:
    metadata:
      name: repo-whitelist
      labels:
        name: repo-whitelist
    spec: 
      containers:
        - name: webhook
          image: ${REPONAME}/repovac:${VERSION}
          imagePullPolicy: Always
          ports:
            - name: endpoint
              containerPort: 8080
          volumeMounts:
            - name: webhook-ca
              mountPath: /etc/certs
              readOnly: true
          securityContext:
            readOnlyRootFilesystem: true
      volumes: 
        - name: webhook-ca
          secret:
            secretName: repo-whitelist
