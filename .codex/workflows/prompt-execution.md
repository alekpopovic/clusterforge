# Prompt Execution Workflow

Numbered prompts live under `prompts/`.

1. Read `.codex/prompt-state.md`.
2. Execute the `Next prompt to execute` unless the user asks for a different
   prompt.
3. Read the target prompt file.
4. Check whether the requested work already exists.
5. Implement only the missing or outdated pieces.
6. Preserve previous user changes.
7. Run validation appropriate to the change.
8. Update `.codex/prompt-state.md`:
   - set `Last executed prompt` to the prompt that just completed
   - set `Next prompt to execute` to the next numbered prompt that has not been
     completed
   - add a short note with commit SHA or evidence when useful
9. Commit and push from `main`.

Do not execute real cloud operations from a prompt unless the user explicitly
requests it and required safety gates are present.
