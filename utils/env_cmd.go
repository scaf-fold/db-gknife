package utils

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

func Shell(cmd string) ([]byte, error) {
	if cmd != "" {
		return exec.Command("sh", "-c", cmd).CombinedOutput()
	} else {
		return nil, errors.New("execute command not found")
	}
}

func GoWorkspace() (string, error) {
	n, err := Shell("echo $GOPATH")
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(string(n), "\n", ""), nil
}

func GetImportPath(path string) (string, error) {
	dir, err := GoWorkspace()
	if err != nil {
		return "", err
	}
	pt := fmt.Sprintf("%s%c%s%c%s", dir, filepath.Separator, "src", filepath.Separator, path)
	return pt, nil
}
