package runbooks

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

type Runbook struct {
	Name     string   `json:"name"`
	Title    string   `json:"title"`
	Category string   `json:"category,omitempty"`
	Severity string   `json:"severity,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	Path     string   `json:"path"`
	Summary  string   `json:"summary,omitempty"`
	Content  string   `json:"-"`
}
type frontmatter struct {
	Title    string   `yaml:"title"`
	Category string   `yaml:"category"`
	Severity string   `yaml:"severity"`
	Tags     []string `yaml:"tags"`
}

func Discover(roots ...string) ([]Runbook, error) {
	var result []Runbook
	for _, root := range roots {
		entries, err := os.ReadDir(root)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, err
		}
		for _, entry := range entries {
			if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
				continue
			}
			path := filepath.Join(root, entry.Name())
			book, err := read(path)
			if err != nil {
				return nil, err
			}
			result = append(result, book)
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })
	return result, nil
}
func read(path string) (Runbook, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Runbook{}, err
	}
	text := string(data)
	meta := frontmatter{}
	body := text
	if strings.HasPrefix(text, "---\n") {
		parts := strings.SplitN(strings.TrimPrefix(text, "---\n"), "\n---\n", 2)
		if len(parts) == 2 {
			if err := yaml.Unmarshal([]byte(parts[0]), &meta); err != nil {
				return Runbook{}, fmt.Errorf("parse runbook frontmatter %s: %w", path, err)
			}
			body = parts[1]
		}
	}
	title := meta.Title
	summary := ""
	scanner := bufio.NewScanner(strings.NewReader(body))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if title == "" && strings.HasPrefix(line, "# ") {
			title = strings.TrimSpace(strings.TrimPrefix(line, "# "))
			continue
		}
		if summary == "" && line != "" && !strings.HasPrefix(line, "#") {
			summary = line
		}
	}
	name := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	return Runbook{Name: name, Title: title, Category: meta.Category, Severity: meta.Severity, Tags: meta.Tags, Path: path, Summary: summary, Content: text}, nil
}
func Find(all []Runbook, name string) (Runbook, error) {
	for _, book := range all {
		if book.Name == name {
			return book, nil
		}
	}
	return Runbook{}, fmt.Errorf("runbook %q not found", name)
}
func Search(all []Runbook, query string) []Runbook {
	query = strings.ToLower(strings.TrimSpace(query))
	var result []Runbook
	for _, book := range all {
		haystack := strings.ToLower(strings.Join(append([]string{book.Name, book.Title, book.Category, book.Severity}, book.Tags...), " "))
		if strings.Contains(haystack, query) {
			result = append(result, book)
		}
	}
	return result
}
