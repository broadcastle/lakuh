package utils

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

// FullPath takes a relative path and converts is to a absolute path.
func FullPath(p string) (string, error) {

	if p == "" {
		return "", errors.New("empty path")
	}

	if filepath.IsAbs(p) {
		return filepath.Clean(p), nil
	}

	if p[0] == '~' {
		home, err := homedir.Dir()
		if err != nil {
			return "", err
		}

		return path.Join(home, p[1:]), nil
	}

	return filepath.Abs(p)

}

// CopyFile copies a file from src to dst.
func CopyFile(src, dst string) error {

	// Get some source information.
	si, err := os.Stat(src)
	if err != nil {
		return err
	}

	// Check if src is actually a file.
	if !si.Mode().IsRegular() {
		return fmt.Errorf("can not copy %s (%q)", si.Name(), si.Mode().String())
	}

	// Get some destination information
	di, err := os.Stat(dst)
	if err != nil {

		if !os.IsNotExist(err) {
			return err
		}

	} else {

		// Check destination
		if !(di.Mode().IsRegular()) {
			return fmt.Errorf("not a file %s (%q)", di.Name(), di.Mode().String())
		}

		if os.SameFile(si, di) {
			return nil
		}

	}

	// Open Input
	input, err := os.Open(src)
	if err != nil {
		return err
	}

	defer input.Close()

	// Create output file
	output, err := os.Create(dst)
	if err != nil {
		return err
	}

	// Copy data.
	if _, err := io.Copy(output, input); err != nil {
		return err
	}

	if err := output.Sync(); err != nil {
		return err
	}

	return output.Close()

}
