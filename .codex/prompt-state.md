# Prompt State

Track numbered prompt execution here. Update this file whenever a prompt is
completed or intentionally skipped.

## Current Position

| Field | Value |
| --- | --- |
| Last executed prompt | `080-v0-2-0-planning-and-milestone-board` |
| Next prompt to execute | `081-v0-2-release-gate-review` |
| Prompt directory | `prompts/` |
| Last updated | 2026-07-09 |

## Rules

- `Last executed prompt` is the highest prompt that was actually implemented,
  verified, committed, and pushed.
- `Next prompt to execute` is the next numbered prompt that should be run by
  default.
- Splitting prompt files does not count as executing those prompts.
- If prompts are executed out of order, record the exception in `Notes`.
- Do not mark a prompt complete if real cloud, smoke, or integration evidence
  was required but not collected.

## Notes

- Prompts `000` through `080` have repository artifacts from prior execution.
- Prompts `081` through `120` have been split into files but have not been
  executed as implementation prompts yet.
