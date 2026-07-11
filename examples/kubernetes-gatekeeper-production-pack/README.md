# Gatekeeper production policy pack

This minimal production example demonstrates an approval annotation constraint.
The module overwrites constraint enforcement with `dryrun` by default. Expand it
with reviewed organization policies before use and set `enforce=true` only after
an observed rollout.
