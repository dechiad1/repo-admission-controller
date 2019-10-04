---
kind: ValidatingWebhookConfiguration
apiVersion: admissionregistration.k8s.io/v1beta1
metadata:
  name: repo-whitelist-webhook
webhooks:
  - name: repo-whitelist.symettrical.dev
    rules:
      - apiGroups:
          - ""
        apiVersions:
          - "v1"
        operations:
          - "CREATE"
        resources:
          - "pods"
    failurePolicy: Fail
    clientConfig:
      caBundle: ${CA_BUNDLE}
      service:
        namespace: repo-whitelist
        name: repo-whitelist
