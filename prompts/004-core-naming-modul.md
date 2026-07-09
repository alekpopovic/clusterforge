## Prompt 4 — Core naming modul

```text
Implement modules/core/naming as a real reusable Terraform module.

Purpose:
Generate consistent names for cloud resources, Kubernetes resources, platform components, and workloads.

Inputs:
- project: string
- environment: string
- component: string
- name: string
- separator: string, default "-"
- max_length: number, default 63
- lowercase: bool, default true
- extra_parts: list(string), default []
- suffix: string, default ""

Validation:
- project, environment, component, name must not be empty.
- separator must be "-", "_" or "".
- max_length must be between 16 and 128.

Behavior:
- Build a base name from:
  project, environment, component, name, extra_parts, suffix
- Remove empty parts.
- Join using separator.
- Lowercase when lowercase=true.
- Replace unsupported characters with separator.
- Trim duplicated separators.
- Trim leading/trailing separators.
- Truncate to max_length.
- Ensure output is deterministic.

Outputs:
- name
- full_name
- parts
- labels_safe_name
- dns_safe_name

labels_safe_name:
- suitable for Kubernetes labels.
- max 63 chars.
- lowercase.
- alphanumeric plus dash where possible.

dns_safe_name:
- suitable for DNS-ish names.
- lowercase.
- max 63 chars.
- no underscores.

Add README.md with:
- explanation
- input table
- output table
- at least 3 usage examples:
  1. AWS resource name
  2. Kubernetes app name
  3. platform component name

Add tests if this repo has a Terraform testing structure; otherwise create examples/core-naming with a minimal root module that uses it.

Run terraform fmt -recursive.
```

---
