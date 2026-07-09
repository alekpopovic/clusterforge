## Prompt 76 — Plugin architecture RFC

```text
Create a plugin architecture RFC for ClusterForge CLI.

Goal:
Design how third parties could add orchestrators, generators, policy checks, or templates.

Create:
- docs/rfcs/006-cli-plugin-architecture.md

Cover:
1. Plugin goals
   - external orchestrators
   - custom templates
   - organization policy packs
   - custom app renderers

2. Non-goals for first version
   - remote code execution by default
   - untrusted plugins
   - plugin marketplace

3. Plugin types:
   - generator plugin
   - policy plugin
   - orchestrator plugin
   - template pack

4. Possible mechanisms:
   - executable plugins named cf-plugin-*
   - local plugin directory
   - YAML-based template packs
   - Go plugin not recommended for portability

5. Security:
   - plugins are trusted code
   - require explicit install
   - show plugin source/path
   - disable plugins in CI unless allowed

6. CLI commands:
   - cf plugin list
   - cf plugin install
   - cf plugin enable
   - cf plugin disable
   - cf plugin run

7. First implementation recommendation:
   - template packs first
   - executable plugins later

Do not implement plugin system yet.
Only RFC and roadmap update.
```

---
