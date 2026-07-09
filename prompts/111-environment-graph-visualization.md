## Prompt 111 — Environment graph visualization

```text
Add environment graph visualization support.

Goal:
Help users understand stack/module dependencies.

CLI:
- cf graph <env>
- cf graph <env> --stack <stack>
- cf graph <env> --format dot
- cf graph <env> --output graph.dot

Behavior:
- Use Terraform graph when available.
- For generated/known stack dependencies, also support a ClusterForge logical graph:
  network -> cluster -> platform -> apps
- Do not require terraform init for logical graph.
- For Terraform graph, run in selected workdir.

Docs:
- docs/graphs.md

Output:
- DOT format
- optional text summary

Tests:
- Logical graph generation.
- Stack graph.
- Output file writing.

Rules:
- Do not require Graphviz.
- Do not render images in CLI; only generate DOT unless easy.
- No apply.

Run:
- gofmt
- go test ./...
```

---
