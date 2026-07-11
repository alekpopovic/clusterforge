package plugins

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

const prefix = "cf-plugin-"

type Plugin struct {
	Name     string
	Path     string
	Disabled bool
}

type Info struct {
	Name         string   `json:"name"`
	Version      string   `json:"version"`
	Description  string   `json:"description"`
	Commands     []string `json:"commands"`
	Capabilities []string `json:"capabilities"`
}

type DiscoverOptions struct {
	Directories      []string
	AllowPathPlugins bool
	DisabledNames    []string
	NoPlugins        bool
}

func Discover(options DiscoverOptions) ([]Plugin, error) {
	if options.NoPlugins {
		return nil, nil
	}

	disabled := make(map[string]bool, len(options.DisabledNames))
	for _, name := range options.DisabledNames {
		disabled[strings.TrimSpace(name)] = true
	}

	paths := append([]string{}, options.Directories...)
	if options.AllowPathPlugins {
		paths = append(paths, filepath.SplitList(os.Getenv("PATH"))...)
	}
	paths = append(paths, ".clusterforge/plugins", ".cf/plugins")

	found := map[string]Plugin{}
	for _, directory := range paths {
		if strings.TrimSpace(directory) == "" {
			continue
		}
		entries, err := os.ReadDir(directory)
		if err != nil {
			if os.IsNotExist(err) || os.IsPermission(err) {
				continue
			}
			return nil, fmt.Errorf("read plugin directory %s: %w", directory, err)
		}
		for _, entry := range entries {
			if entry.IsDir() || !strings.HasPrefix(entry.Name(), prefix) {
				continue
			}
			name := strings.TrimPrefix(entry.Name(), prefix)
			if name == "" {
				continue
			}
			if _, exists := found[name]; exists {
				continue
			}
			path, err := filepath.Abs(filepath.Join(directory, entry.Name()))
			if err != nil {
				return nil, fmt.Errorf("resolve plugin %s: %w", entry.Name(), err)
			}
			stat, err := os.Stat(path)
			if err != nil || stat.Mode().Perm()&0o111 == 0 {
				continue
			}
			found[name] = Plugin{Name: name, Path: path, Disabled: disabled[name]}
		}
	}

	result := make([]Plugin, 0, len(found))
	for _, plugin := range found {
		result = append(result, plugin)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })
	return result, nil
}

func Find(all []Plugin, name string) (Plugin, error) {
	for _, plugin := range all {
		if plugin.Name == name {
			return plugin, nil
		}
	}
	return Plugin{}, fmt.Errorf("plugin %q was not discovered", name)
}

func ReadInfo(plugin Plugin) (Info, error) {
	if plugin.Disabled {
		return Info{}, fmt.Errorf("plugin %q is disabled", plugin.Name)
	}
	command := exec.Command(plugin.Path, "--clusterforge-plugin-info")
	var stdout, stderr bytes.Buffer
	command.Stdout = &stdout
	command.Stderr = &stderr
	if err := command.Run(); err != nil {
		return Info{}, fmt.Errorf("plugin %q info failed: %w: %s", plugin.Name, err, strings.TrimSpace(stderr.String()))
	}
	var info Info
	if err := json.Unmarshal(stdout.Bytes(), &info); err != nil {
		return Info{}, fmt.Errorf("parse plugin %q info: %w", plugin.Name, err)
	}
	if info.Name == "" || info.Version == "" {
		return Info{}, fmt.Errorf("plugin %q info requires name and version", plugin.Name)
	}
	if info.Name != plugin.Name {
		return Info{}, fmt.Errorf("plugin %q reported name %q", plugin.Name, info.Name)
	}
	return info, nil
}

func Run(plugin Plugin, args []string, stdin io.Reader, stdout, stderr io.Writer) error {
	if plugin.Disabled {
		return fmt.Errorf("plugin %q is disabled", plugin.Name)
	}
	command := exec.Command(plugin.Path, args...)
	command.Stdin = stdin
	command.Stdout = stdout
	command.Stderr = stderr
	if err := command.Run(); err != nil {
		return fmt.Errorf("plugin %q failed: %w", plugin.Name, err)
	}
	return nil
}
