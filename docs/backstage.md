# Backstage integration

Enable `backstage` metadata in `clusterforge.yaml`, then run
`cf backstage generate --output catalog-info.yaml`. Commit the reviewed YAML so
Backstage can discover it through the normal catalog location/provider flow.
ClusterForge does not call the Backstage API.

The project becomes a System, environments/clusters become Resources, apps
become Components, and configured teams become Groups. App-level owner/system/
lifecycle values override project defaults; otherwise ownership falls back to
the configured Backstage or organization owner. Generated entities contain
metadata only and must never include credentials or secret values.
