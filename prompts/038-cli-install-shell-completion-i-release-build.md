## Prompt 38 — CLI install, shell completion i release build

```text
Improve ClusterForge CLI developer and release experience.

Tasks:
1. Add CLI command:
   - cf completion bash
   - cf completion zsh
   - cf completion fish
   - cf completion powershell

2. Add CLI command:
   - cf version
   Print:
   - version
   - commit
   - date
   - Go version if practical

3. Add build metadata support:
   - version injected via ldflags
   - commit injected via ldflags
   - date injected via ldflags

4. Add install script:
   - scripts/install-cli.sh
   It should:
   - Build cli/cf
   - Install to /usr/local/bin or user-provided destination
   - Avoid requiring sudo unless needed
   - Print installed version

5. Add release workflow:
   - .github/workflows/release-cli.yml
   Trigger:
   - tags like v*
   Build:
   - linux amd64
   - linux arm64
   - darwin amd64
   - darwin arm64
   - windows amd64
   Upload artifacts.

6. Update docs/cli.md:
   - Install from source.
   - Build manually.
   - Shell completion.
   - Version command.
   - Release artifacts.

Rules:
- Do not publish anything automatically beyond GitHub release artifacts.
- Do not require secrets unless necessary.
- Keep cross-compilation simple.

Run:
- cd cli && go test ./...
- cd cli && go build -o cf .

Final response:
- List CLI improvements.
- Show build command.
- Mention artifacts produced by workflow.
```

---
