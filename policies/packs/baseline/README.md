# Baseline Policy Pack

Blocking intent:

- no plaintext secrets in committed tfvars
- no auto-approve in prod
- plan file required for production apply
- destroy blocked in production by default

Current implementation combines CLI built-ins with existing static scanners.
