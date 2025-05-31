# Terraform IaC

Contains modules and configurations for provisioning infrastructure on AWS or GCP:
- VPC, subnets, security (SG/NSG)
- Kubernetes (EKS/GKE)
- RDS (PostgreSQL), Redis, managed Kafka

## How to use
```bash
  cd infrastructure/terraform
```
```bash
  terraform init
```
```bash
  terraform plan # review changes
```
```bash
  terraform apply # apply infrastructure
```

## Organization
- `modules/` – reusable blocks (vpc, eks, rds)
- `envs/` – workspaces by environment (dev, qa, prd)

## Tips
- Always review plan before apply.
- Use terraform workspace select <env>.