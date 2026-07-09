# Codex Workspace

This repository is ready for Codex-assisted development. Start with
`AGENTS.md` for binding repository rules, then use `.codex/` for task-specific
context, agent profiles, and repeatable workflows.

## Entry Points

- `AGENTS.md`: required rules for all contributors and AI agents.
- `.codex/README.md`: Codex workspace map.
- `.codex/agents/`: focused agent profiles for Terraform, CLI, docs, security,
  and release work.
- `.codex/workflows/`: repeatable task flows.
- `.codex/checklists/`: quick safety and completion checklists.
- `.codex/prompt-state.md`: last executed prompt and next prompt to run.
- `prompts/`: individual project prompts from `000` through `120`.

## Default Codex Flow

1. Read `AGENTS.md`.
2. Read the relevant `.codex/agents/*.md` profile.
3. Make scoped changes directly from `main`, following repository git rules.
4. Run the smallest useful validation first, then broader validation when the
   change touches shared behavior.
5. If executing numbered prompts, update `.codex/prompt-state.md` with the last
   completed prompt and the next prompt to run.
6. Commit and push from `main`.

## Safety Defaults

- Do not commit credentials, kubeconfigs, private keys, tfstate, tfplan, or
  local tfvars.
- Do not run real cloud integration tests unless explicitly requested and the
  documented environment gates are set.
- Do not claim smoke tests passed unless evidence exists.
