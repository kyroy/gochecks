package gotest

import (
	"os/exec"
	"path/filepath"
)

// Run executes go tests in the given directory and all subdirectories.
func Run(dir string) ([]byte, error) {
	cmd := exec.Command("go", "test", "-v", "-json", "-cover", "./...")
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	cmd.Dir = absDir
	output, err := cmd.CombinedOutput()
	return output, err
}
