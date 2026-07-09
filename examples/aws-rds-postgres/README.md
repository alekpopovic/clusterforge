# aws-rds-postgres

Example private RDS PostgreSQL instance with AWS-managed master password.

Run local validation:

```bash
terraform init
terraform validate
terraform plan -refresh=false
```

Do not apply with fake credentials. Review RDS cost, backup retention,
deletion protection, and secret access before using this in a real account.
