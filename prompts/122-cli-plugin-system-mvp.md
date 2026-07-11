## Prompt 122 — CLI plugin system MVP

```text
Implement the first version of the ClusterForge CLI plugin system.

Goal:
Allow trusted local executable plugins to extend ClusterForge without modifying the core CLI.

Plugin model:
- Executable plugins named:
  cf-plugin-<name>
- Plugins are discovered from:
  1. configured plugin directories
  2. PATH
  3. .clusterforge/plugins
  4. .cf/plugins

Config:
clusterforge.yaml:
  plugins:
    enabled: true
    directories:
      - .clusterforge/plugins
    allow_path_plugins: true

CLI commands:
- cf plugin list
- cf plugin discover
- cf plugin info <name>
- cf plugin run <name> -- <args>
- cf plugin enable <name>
- cf plugin disable <name>

Plugin protocol:
- cf-plugin-foo --clusterforge-plugin-info
Should return JSON:
{
  "name": "foo",
  "version": "0.1.0",
  "description": "...",
  "commands": ["..."],
  "capabilities": ["generator", "policy", "template"]
}

Security:
- Plugins are trusted local code.
- Print plugin path before running.
- Add --no-plugins global flag.
- Add --allow-plugins flag for CI if default is restricted.
- Do not auto-install plugins from the internet.
- Do not execute plugins during normal commands unless explicitly configured.

Tests:
- Plugin discovery from temp directory.
- Disabled plugin is not executed.
- Plugin info JSON parsing.
- Plugin run passes args.
- --no-plugins disables plugin discovery.

Docs:
- docs/plugins.md
- Explain plugin security model.
- Explain plugin discovery.
- Explain example plugin.

Create example plugin:
- examples/plugins/cf-plugin-hello

Rules:
- Keep plugin system simple.
- Do not implement marketplace.
- Do not download remote code.
- Do not break existing CLI commands.

Run:
- gofmt
- go test ./...
```


---
