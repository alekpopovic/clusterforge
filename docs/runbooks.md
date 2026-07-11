# Runbook discovery

Use `cf runbook list`, `show <name>`, `search <query>`, and `open <name>` to
discover Markdown under `docs/incident-response` and `docs/dr`. Discovery reads
optional title/category/severity/tags frontmatter and never executes commands.

`open` deliberately prints the local path; it does not start an editor, browser
or background process. Review every step, environment and approval boundary
before acting, especially commands marked destructive or high-risk.
