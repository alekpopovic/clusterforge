package templatepacks

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type GitSource struct {
	URL string
	Ref string
}

func ParseGitSource(source string) (GitSource, error) {
	if !strings.HasPrefix(source, "git::") {
		return GitSource{}, fmt.Errorf("not a git source")
	}
	raw := strings.TrimPrefix(source, "git::")
	parts := strings.SplitN(raw, "?", 2)
	if strings.TrimSpace(parts[0]) == "" {
		return GitSource{}, fmt.Errorf("git source URL is required")
	}
	result := GitSource{URL: parts[0]}
	if len(parts) == 2 {
		for _, value := range strings.Split(parts[1], "&") {
			pair := strings.SplitN(value, "=", 2)
			if len(pair) == 2 && pair[0] == "ref" {
				result.Ref = pair[1]
			}
		}
	}
	if result.Ref == "" {
		return GitSource{}, fmt.Errorf("git source must include ?ref=<tag-or-commit>")
	}
	return result, nil
}

func WeakRef(ref string) bool {
	switch strings.ToLower(strings.TrimSpace(ref)) {
	case "main", "master", "develop", "development", "head":
		return true
	default:
		return false
	}
}

func CachePath(root, name, version string) string {
	return filepath.Join(root, ".cf", "cache", "template-packs", name, version)
}

func Fetch(source, destination string) error {
	if strings.HasPrefix(source, "git::") {
		parsed, err := ParseGitSource(source)
		if err != nil {
			return err
		}
		return fetchGit(parsed, destination)
	}
	archivePath := strings.TrimPrefix(source, "archive::")
	if isArchive(archivePath) {
		return extractArchive(archivePath, destination)
	}
	localPath := strings.TrimPrefix(source, "path::")
	return copyDirectory(localPath, destination)
}

func fetchGit(source GitSource, destination string) error {
	temporary := destination + ".tmp"
	if err := os.RemoveAll(temporary); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(destination), 0o755); err != nil {
		return err
	}
	command := exec.Command("git", "clone", "--depth", "1", "--branch", source.Ref, "--", source.URL, temporary)
	if output, err := command.CombinedOutput(); err != nil {
		return fmt.Errorf("fetch template pack git source: %w: %s", err, strings.TrimSpace(string(output)))
	}
	if err := os.RemoveAll(filepath.Join(temporary, ".git")); err != nil {
		return err
	}
	return replaceDirectory(temporary, destination)
}

func copyDirectory(source, destination string) error {
	info, err := os.Stat(source)
	if err != nil {
		return fmt.Errorf("read local template pack: %w", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("local template pack %s is not a directory", source)
	}
	sourceAbs, err := filepath.Abs(source)
	if err != nil {
		return err
	}
	destinationAbs, err := filepath.Abs(destination)
	if err != nil {
		return err
	}
	if destinationAbs == sourceAbs || strings.HasPrefix(destinationAbs, sourceAbs+string(os.PathSeparator)) {
		return fmt.Errorf("template pack cache must not be inside source directory %s", source)
	}
	temporary := destination + ".tmp"
	if err := os.RemoveAll(temporary); err != nil {
		return err
	}
	err = filepath.WalkDir(source, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		relative, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		target := filepath.Join(temporary, relative)
		if entry.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		if entry.Type()&os.ModeSymlink != 0 {
			return fmt.Errorf("template pack symlinks are not supported: %s", path)
		}
		input, err := os.Open(path)
		if err != nil {
			return err
		}
		output, err := os.OpenFile(target, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		if err != nil {
			input.Close()
			return err
		}
		_, copyErr := io.Copy(output, input)
		inputCloseErr := input.Close()
		closeErr := output.Close()
		if copyErr != nil {
			return copyErr
		}
		if inputCloseErr != nil {
			return inputCloseErr
		}
		return closeErr
	})
	if err != nil {
		_ = os.RemoveAll(temporary)
		return err
	}
	return replaceDirectory(temporary, destination)
}

func isArchive(path string) bool {
	return strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".tar.gz") || strings.HasSuffix(path, ".tgz")
}

func extractArchive(source, destination string) error {
	temporary := destination + ".tmp"
	if err := os.RemoveAll(temporary); err != nil {
		return err
	}
	if err := os.MkdirAll(temporary, 0o755); err != nil {
		return err
	}
	var err error
	if strings.HasSuffix(source, ".zip") {
		err = extractZip(source, temporary)
	} else {
		err = extractTarGz(source, temporary)
	}
	if err != nil {
		_ = os.RemoveAll(temporary)
		return err
	}
	return replaceDirectory(temporary, destination)
}

func safeTarget(root, name string) (string, error) {
	target := filepath.Join(root, filepath.Clean(name))
	rootWithSeparator := filepath.Clean(root) + string(os.PathSeparator)
	if target != filepath.Clean(root) && !strings.HasPrefix(target, rootWithSeparator) {
		return "", fmt.Errorf("archive entry escapes destination: %s", name)
	}
	return target, nil
}

func extractZip(source, destination string) error {
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		target, err := safeTarget(destination, file.Name)
		if err != nil {
			return err
		}
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(target, 0o755); err != nil {
				return err
			}
			continue
		}
		if file.Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("archive symlinks are not supported: %s", file.Name)
		}
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return err
		}
		input, err := file.Open()
		if err != nil {
			return err
		}
		output, err := os.OpenFile(target, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		if err != nil {
			input.Close()
			return err
		}
		_, err = io.Copy(output, input)
		input.Close()
		output.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func extractTarGz(source, destination string) error {
	file, err := os.Open(source)
	if err != nil {
		return err
	}
	defer file.Close()
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()
	reader := tar.NewReader(gzipReader)
	for {
		header, err := reader.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		target, err := safeTarget(destination, header.Name)
		if err != nil {
			return err
		}
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0o755); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
				return err
			}
			output, err := os.OpenFile(target, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
			if err != nil {
				return err
			}
			_, copyErr := io.Copy(output, reader)
			closeErr := output.Close()
			if copyErr != nil {
				return copyErr
			}
			if closeErr != nil {
				return closeErr
			}
		default:
			return fmt.Errorf("unsupported archive entry %s", header.Name)
		}
	}
}

func replaceDirectory(temporary, destination string) error {
	if err := os.RemoveAll(destination); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(destination), 0o755); err != nil {
		return err
	}
	return os.Rename(temporary, destination)
}
