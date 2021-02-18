package gopackages

import (
	"path/filepath"

	"golang.org/x/xerrors"
)

type Module struct {
	module  string
	rootDir string
}

func NewModule(in string) (*Module, error) {
	gmp, err := GetGoModPath(in)

	if err != nil {
		return nil, xerrors.Errorf("failed to call GetGoModPath: %w", err)
	}

	module, err := GetGoModule(gmp)

	if err != nil {
		return nil, xerrors.Errorf("failed to call GetGoModule: %w", err)
	}

	rootDir, err := filepath.Abs(filepath.Dir(gmp))

	if err != nil {
		return nil, xerrors.Errorf("failed to get absolute path for the directory of go.mod: %w", err)
	}

	return &Module{
		module:  module,
		rootDir: rootDir,
	}, nil
}

func (m *Module) GetImportPath(path string) (string, error) {
	abs, err := filepath.Abs(path)

	if err != nil {
		return "", xerrors.Errorf("failed to get absolute path for %s: %w", path, err)
	}

	rel, err := filepath.Rel(m.rootDir, abs)

	if err != nil {
		return "", xerrors.Errorf("failed to calculate relative path from %s to %s: %w", m.rootDir, abs, err)
	}

	return filepath.Join(m.module, rel), nil
}
