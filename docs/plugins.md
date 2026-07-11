# Trusted local CLI plugins

ClusterForge plugins are local executables named `cf-plugin-<name>`. They can
extend local workflows without being compiled into the core CLI, but they run
with the same user permissions and environment as `cf`. Treat every plugin as
trusted code and review it before enabling or running it.

ClusterForge does not download plugins, provide a marketplace, execute plugins
during unrelated commands, or sandbox plugin processes.

## Configuration

Plugins are disabled unless explicitly enabled in `clusterforge.yaml`:

```yaml
plugins:
  enabled: true
  directories:
    - .clusterforge/plugins
  allow_path_plugins: true
```

Configured relative directories resolve relative to the configuration file.
Discovery checks sources in this order:

1. configured plugin directories;
2. directories on `PATH`, when `allow_path_plugins` is true;
3. `.clusterforge/plugins`;
4. `.cf/plugins`.

The first executable with a given name wins. Files without the
`cf-plugin-` prefix and non-executable files are ignored.

## Commands

```bash
cf plugin discover
cf plugin list
cf plugin info hello
cf plugin run hello -- greet Ada
cf plugin disable hello
cf plugin enable hello
```

`discover` and `list` do not execute plugins. `info` explicitly starts the
plugin with `--clusterforge-plugin-info` and expects JSON containing `name`,
`version`, `description`, `commands`, and `capabilities`. `run` prints the
resolved executable path before passing all arguments after the plugin name to
the executable.

Use the global `--no-plugins` flag to disable discovery and execution for a
command. When the `CI` environment variable is set, plugin operations are
blocked unless `--allow-plugins` is passed. Only use that override for reviewed
executables supplied by the trusted CI workspace; never run plugins from an
untrusted pull request or writable shared path.

## Example plugin

The repository includes `examples/plugins/cf-plugin-hello`. Try it from a
temporary project configuration:

```bash
chmod +x examples/plugins/cf-plugin-hello
mkdir -p .clusterforge/plugins
cp examples/plugins/cf-plugin-hello .clusterforge/plugins/
cf plugin info hello
cf plugin run hello -- greet Ada
```

The example implements the metadata protocol and a single `greet` action. A
production plugin should validate inputs, avoid printing secrets, return
non-zero on failure, and document every capability it uses.
