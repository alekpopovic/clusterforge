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
8. Before staging files, update `.codex/prompt-state.md` with the current
   prompt state:
   - set `Last executed prompt` to the prompt that just completed
   - set `Next prompt to execute` to the next numbered prompt that has not been
     completed
   - add the current prompt result, validation summary, and evidence path when
     useful
9. Stage the prompt changes and prompt state together with `git add`.
10. Commit from `main`.
11. Use the prompt title as the commit message, without the prompt number. For
    example, `Prompt 081 — v0.2 release gate review` commits as
    `v0.2 release gate review`.
12. Push from `main`.

Do not execute real cloud operations from a prompt unless the user explicitly
requests it and required safety gates are present.
