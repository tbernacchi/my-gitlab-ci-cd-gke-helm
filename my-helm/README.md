# My Application Helm Chart

This Helm chart deploys a microservices application consisting of a Python frontend and a Go backend on Kubernetes.

## Chart Structure

```
my-helm/
├── Chart.yaml         # Chart metadata
├── values.yaml        # Default configuration values
├── NOTES.txt         # Usage notes
├── auth-gke.sh       # GKE authentication script
└── templates/        # Kubernetes templates
    ├── 001-my-app-frontend.yml
    ├── 002-my-app-backend.yml
    └── 003-my-app-ingress.yml
```

## Prerequisites

- Kubernetes 1.19+
- Helm 3.0+
- GitLab registry access configured
- GKE cluster access

## Configuration

The following table lists the configurable parameters of the chart and their default values:

| Parameter | Description | Default |
|-----------|-------------|---------|
| `version.front` | Frontend version | `2.0` |
| `version.back` | Backend version | `1.0` |
| `frontend.image` | Frontend Docker image | `index.docker.io/tadeuuuuu/my-images-2020/frontend-teste` |
| `frontend.port` | Frontend service port | `80` |
| `frontend.containerport` | Frontend container port | `8000` |
| `frontend.targetport` | Frontend target port | `8000` |
| `backend.image` | Backend Docker image | `index.docker.io/tadeuuuuu/my-images-2020/backend-teste` |
| `backend.port` | Backend service port | `80` |
| `backend.containerport` | Backend container port | `8080` |
| `backend.targetport` | Backend target port | `8080` |

## Components

### Frontend Service
- Python-based application
- Deployed as a Kubernetes Deployment
- Exposed internally via ClusterIP Service
- Environment variables configured for backend communication

### Backend Service
- Go-based application
- Deployed as a Kubernetes Deployment
- Exposed internally via ClusterIP Service
- Accessible only within the cluster

### Ingress Configuration
- Routes external traffic to services
- Paths:
  - `/*`: Routes to frontend service
  - `/healthz`: Health check endpoint

## Installation

Install the chart:
```bash
helm install my-app ./my-helm \
  --namespace your-namespace \
  --create-namespace
```

## Testing and Validation

Before deploying to production, you can validate and test your chart using these commands:

### Lint the Chart
Check for syntax errors and best practices:
```bash
helm lint ./my-helm
```

### Validate Templates
Validate template rendering without installing:
```bash
# Basic template validation
helm template ./my-helm

# Template with specific values
helm template ./my-helm --set version.front=2.1

# Template with values file
helm template ./my-helm -f custom-values.yaml
```

### Dry Run Installation
Simulate an installation:
```bash
# Basic dry-run
helm install my-app ./my-helm --dry-run

# Detailed dry-run with debug info
helm install my-app ./my-helm --dry-run --debug

# Dry-run with custom values
helm install my-app ./my-helm --dry-run --set version.front=2.1
```

### Check Kubernetes Resources
Validate the Kubernetes manifests:
```bash
# Show all resources that would be created
helm template ./my-helm | kubectl apply --dry-run=client -f -

# Validate against server
helm template ./my-helm | kubectl apply --dry-run=server -f -

# Check specific namespace
helm template ./my-helm | kubectl apply -n your-namespace --dry-run=server -f -
```

### Debug Chart Installation
If you need to troubleshoot:
```bash
# Enable debug logs
helm install my-app ./my-helm --debug

# Show computed values
helm get values my-app

# Show all resources in a release
helm get manifest my-app
```

## Upgrading

To upgrade the release:
```bash
helm upgrade my-app ./my-helm \
  --namespace your-namespace \
  --set version.front=<new-version> \
  --set version.back=<new-version>
```

## Uninstallation

To remove the deployment:
```bash
helm uninstall my-app -n your-namespace
```

## Development

To modify the chart:
1. Update values in `values.yaml` for default configurations
2. Modify templates in `templates/` directory for Kubernetes resources
3. Update version in `Chart.yaml`
4. Test the chart:
```bash
helm lint .
helm template . --debug
```
