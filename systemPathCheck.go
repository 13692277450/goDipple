package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// isInSystemPath checks if the current executable's directory is in system PATH
func isInSystemPath() bool {
	exePath, err := os.Executable()
	if err != nil {
		return false
	}
	exeDir := filepath.Dir(exePath)

	systemPath := os.Getenv("PATH")
	paths := strings.Split(systemPath, ";")

	for _, path := range paths {
		if strings.EqualFold(path, exeDir) {
			return true
		}
	}
	return false
}

// addToSystemPath adds current executable's directory to system PATH
func addToSystemPath() error {
	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	exeDir := filepath.Dir(exePath)

	// This requires admin privileges
	cmd := exec.Command("powershell", "-Command",
		`$path = [Environment]::GetEnvironmentVariable('PATH', 'Machine');
		if (-not $path.Contains('`+exeDir+`')) {
			[Environment]::SetEnvironmentVariable('PATH', $path + ';' + '`+exeDir+`', 'Machine')
		}`)

	return cmd.Run()
}
