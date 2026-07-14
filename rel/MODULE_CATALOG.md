# Module Catalog

| Module path | Purpose | Stability | Providers | Example path | Test status | Registry-ready |
| --- | --- | --- | --- | --- | --- | --- |
| `modules/core/naming` | Naming helpers | stable | none | `examples/core-naming` | Terraform test | yes |
| `modules/core/tags` | Common tags | stable | none | `examples/core-metadata` | Terraform test | yes |
| `modules/core/labels` | Common labels | stable | none | `examples/core-metadata` | Terraform test | yes |
| `modules/cloud/aws/network` | AWS VPC/subnets/NAT | beta | aws | `examples/aws-network` | plan-mode test | yes |
| `modules/cloud/aws/tfstate-backend` | S3/DynamoDB backend resources | beta | aws | `examples/aws-tfstate-backend` | plan-mode test | yes |
| `modules/cloud/aws/dns` | Route53 zones/records | beta | aws | `examples/aws-route53-dns` | plan-mode partial | no |
| `modules/cloud/aws/irsa-role` | EKS IRSA role | beta | aws | `examples/aws-eks-karpenter` | plan-mode test | yes |
| `modules/orchestrators/kubernetes/eks` | EKS cluster | beta | aws | `examples/aws-eks-minimal` | validation | yes |
| `modules/orchestrators/ecs/cluster` | ECS cluster | beta | aws | `examples/ecs-cluster-minimal` | validation | yes |
| `modules/workloads/kubernetes/app` | Kubernetes web app | beta | kubernetes | `examples/kubernetes-basic-app` | validation | yes |
| `modules/workloads/ecs/service` | ECS Fargate service | beta | aws | `examples/ecs-fargate-app` | validation | yes |
| `modules/cloud/azure/network` | Azure VNet/subnets | experimental | azurerm | `examples/azure-aks-minimal` | format only | no |
| `modules/orchestrators/kubernetes/aks` | AKS cluster | experimental | azurerm | `examples/azure-aks-minimal` | format only | no |
| `modules/cloud/gcp/network` | GCP VPC/subnet | experimental | google | `examples/gcp-gke-minimal` | format only | no |
| `modules/orchestrators/kubernetes/gke` | GKE cluster | experimental | google | `examples/gcp-gke-minimal` | format only | no |
| `modules/orchestrators/kubernetes/k3s` | K3s user-data | experimental | none | `examples/k3s-cloud-init` | format only | no |
| `modules/orchestrators/kubernetes/rke2` | RKE2 user-data | experimental | none | `examples/rke2-cloud-init` | format only | no |
| `modules/platform/ecs/codedeploy-blue-green` | ECS CodeDeploy blue/green | experimental | aws | `examples/ecs-blue-green` | format only | no |

Modules not listed here are early platform/workload modules and should be
treated as beta or experimental until their README, examples, and tests are
reviewed for registry readiness.
