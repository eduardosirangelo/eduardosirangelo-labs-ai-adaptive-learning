# Kubernetes Manifests

Manifests and resources for deployment in a Kubernetes cluster:
- `deployments/` – Deployment + Service + ConfigMap
- `ingress/`     – Ingress or Gateway API
- `secrets/`     – Secret manifests (via Kustomize)

## How to apply

```bash
  cd infrastructure/k8s-manifests
```
```bash
  kubectl apply -k .
```

## Structure
- `base/` - generic configurations
- `overlays/` - customizations per environment (dev, prd)

## Best practices
- Do not version secrets in plain text.
- Use liveness/readiness probes in all pods.