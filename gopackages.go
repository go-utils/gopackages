package gopackages

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"
	"golang.org/x/xerrors"
)

var (
	NotFound        = xerrors.New("not found")
	PkgNameNotFound = xerrors.New("package name was not found")
)

// GetGoModPath - Search for go.mod by breadth-first-search directly under the Git repository
func GetGoModPath(in string) (string, error) {
	base := filepath.Clean(in)
	prev := base

	for {
		goMod := filepath.Join(base, "go.mod")

		_, err := os.Stat(goMod)
		if err == nil {
			return goMod, nil
		}

		base, err = filepath.Abs(filepath.Join(base, ".."))
		if err != nil {
			return "", NotFound
		}

		if prev == base {
			return "", NotFound
		}

		prev = base
	}
}

// GetGoModule - Get the Go root package name from go.mod
func GetGoModule(goMod string) (string, error) {
	d, err := ioutil.ReadFile(goMod)
	if err != nil {
		return "", xerrors.Errorf("error in ioutil ReadFile method: %w", err)
	}

	f, err := modfile.Parse("", d, nil)
	if err != nil {
		return "", xerrors.Errorf("failed to mod file parsed: %w", err)
	}

	if len(f.Module.Mod.Path) == 0 {
		return "", PkgNameNotFound
	}

	return f.Module.Mod.Path, nil
}
