## Prompt 86 — CLI end-to-end non-cloud tests

```text
Add CLI end-to-end tests that do not require cloud credentials.

Goal:
Test real CLI flows in temporary directories.

Create:
- cli/e2e/
  project_init_test.go
  env_create_test.go
  generate_test.go
  app_flow_test.go
  policy_test.go

Test flows:
1. New project:
   - cf project init demo
   - verify clusterforge.yaml
   - verify directories

2. Environment generation:
   - cf env create dev --cloud aws --orchestrator eks --region eu-central-1
   - cf generate dev
   - verify live/dev/aws-eks files

3. App flow:
   - cf app add api --image nginx:1.25 --port 80
   - cf app validate api
   - cf app render api --env dev
   - verify generated module call

4. Policy:
   - prod apply without plan file fails
   - prod destroy blocked by default

5. Doctor:
   - doctor returns useful output inside generated project

Rules:
- Use temp directories.
- Do not require Terraform binary unless test is explicitly skipped.
- Do not require cloud credentials.
- Do not write outside temp directory.
- Ensure tests are cross-platform where practical.

Run:
- cd cli && go test ./...
```

---
