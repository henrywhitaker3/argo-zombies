# Argo-Zombies

A CLI that scans a kubernetes cluster to find resources that are not managed by ArgoCD.

## Installation

### CLI

Download the binary from the latest release and run:

```bash
argo-zombies detect
```

### Helm Chart

There is a helm chart available that sets up a CronJob to run detection:

```bash
helm repo add argo-zombies https://henrywhitaker3.github.io/argo-zombies
helm repo update
helm install argo-zombies argo-zombies/argo-zombies --version <version>
```

## Configuration

The config file (defaults to `.argo-zombies.yaml`) allows you to setup exclusions, so that resources are ignored by the detector:

```yaml
exclusions:
  resources: []
    # - name: bongo
    #   namespace: bongo
    #   kind: Secret
    #   version: v1
  namespaces: []
    # - bongo
  selectors: []
    # - labels: {}
    #   annotations: {}
  gvrs: []
    # - group: apiextensions.k8s.io
    #   version: v1
    #   resource: customresourcedefinitions
  bundles: []
    # - k3s
    # - longhorn
    # - aks
    # - ingress-nginx
    # - cert-manager
    # - datadog
```

You can update you helm values to pass these config values to the CronJob.
