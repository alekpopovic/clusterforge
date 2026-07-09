# RFC 005: ECS Blue/Green Deployments

Goal: support safer ECS/Fargate production deployments using CodeDeploy and
ALB target groups.

Design:

- ECS service uses `deployment_controller = "CODE_DEPLOY"`
- ALB has blue and green target groups
- CodeDeploy application and deployment group control traffic shifting
- rollback is enabled for deployment failure
- IAM service role grants CodeDeploy ECS permissions

The default ECS service remains rolling `ECS` deployment. Blue/green is
explicitly opt-in.

Limitations: test listener support, canary/linear configs, and full AppSpec
automation are future work.
