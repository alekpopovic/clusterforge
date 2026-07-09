# RFC 006: CLI Plugin Architecture

## Goals

- external orchestrators
- custom templates
- organization policy packs
- custom app renderers

## Non-Goals

- remote code execution by default
- untrusted plugins
- plugin marketplace

## Plugin Types

- generator plugin
- policy plugin
- orchestrator plugin
- template pack

## Mechanisms

- local template packs first
- executable plugins named `cf-plugin-*` later
- local plugin directory with explicit enablement
- Go plugins are not recommended because portability is poor

## Security

Plugins are trusted code. Users must explicitly install and enable them. CI
should disable plugins unless allowed.

Recommended first implementation: template packs before executable plugins.
