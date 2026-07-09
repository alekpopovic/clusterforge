# Prompt Execution Workflow

Numbered prompts live under `prompts/`.

1. Read the target prompt file.
2. Check whether the requested work already exists.
3. Implement only the missing or outdated pieces.
4. Preserve previous user changes.
5. Run validation appropriate to the change.
6. Commit and push from `main`.

Do not execute real cloud operations from a prompt unless the user explicitly
requests it and required safety gates are present.
