# aws-vpc-endpoints-private-eks

Example VPC endpoints for private EKS-style networking. The example creates a
network and endpoints needed for private nodes to pull images from ECR and
write CloudWatch logs without routing that traffic through a NAT gateway.

Run local validation:

```bash
terraform init
terraform validate
terraform plan -refresh=false
```

Do not apply with fake credentials. In a real cluster root, pass the node or
cluster security group IDs into `allowed_security_group_ids` so interface
endpoints accept HTTPS traffic only from approved workloads.
