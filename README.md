# Appclacks Kubernetes operator

Kubernetes operator for the [Appclacks cloud platform](https://appclacks.com/).

It allows you configure health checks on Appclacks using Kubernetes Custom Resources Definitions.

## Getting started

### Modify the Appclacks secret

First, add in the file `config/deployment/secret.yaml` the required Appclacks credentials for the operator.

See the [authentication section](https://www.doc.appclacks.com/getting-started/#authentication) in the documentation to create a token.

### Create the required CRDs

Once the secret ready, you can create the CRDs for the operator and then deploy the operator. The operator will be deployed in a namespace named `appclacks`:

```
kubectl apply -f config/crd/bases
kubectl apply -f config/deployment
```

You should then be able to manage healthchecks on your account using the Appclacks CRD:

```
kubectl apply -f config/samples/healthchecks_v1alpha1_dnshealthcheck.yaml
```
