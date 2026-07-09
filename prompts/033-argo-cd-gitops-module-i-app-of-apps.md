## Prompt 33 — Argo CD GitOps module i app-of-apps

```text
Improve Argo CD GitOps support.

Existing or target module:
- modules/platform/kubernetes/argocd

Goal:
Make Terraform install Argo CD and optionally bootstrap an app-of-apps pattern.

Tasks:
1. Ensure modules/platform/kubernetes/argocd installs Argo CD via Helm.
2. Add inputs:
   - namespace: string, default "argocd"
   - create_namespace: bool, default true
   - chart_version: string, default ""
   - values: list(string), default []
   - enable_app_of_apps: bool, default false
   - app_of_apps_name: string, default "cluster-apps"
   - app_of_apps_repo_url: string, default ""
   - app_of_apps_path: string, default "apps"
   - app_of_apps_revision: string, default "HEAD"
   - app_of_apps_destination_namespace: string, default "argocd"
   - app_of_apps_project: string, default "default"

3. If enable_app_of_apps=true:
   - Create an Argo CD Application manifest using kubernetes_manifest.
   - Validate that repo_url is not empty.
   - Do not include Git credentials in Terraform.

4. Outputs:
   - namespace
   - release_name
   - app_of_apps_name

5. Update bootstrap module:
   - Ensure enable_argocd wires through app-of-apps inputs.

6. Create example:
   - examples/kubernetes-argocd-bootstrap

7. Docs:
   - docs/gitops.md
   - Explain Terraform boundary:
     Terraform creates cluster and platform.
     Argo CD manages frequent app deployments.
   - Explain app-of-apps pattern.
   - Explain repo layout recommendation:
     gitops/
       apps/
       environments/
       projects/

Rules:
- Do not put repo credentials in Terraform.
- Do not make GitOps mandatory.
- Keep plain Terraform workload modules available for smaller use cases.

Run:
- terraform fmt -recursive
- validation where possible

Final response:
- Summarize GitOps changes.
- Include example usage.
- Mention security limitations.
```

---
