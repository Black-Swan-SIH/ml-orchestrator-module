package python

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ml-orchestrator-module/internal/config"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func GodFunction(filePath string, cfg *config.Config) ([]byte, error) {
	venv := cfg.Venv
	scripp := cfg.Parser

	pythonExecutable := filepath.Join(venv, "Scripts", "python")
	if runtime.GOOS != "windows" {
		pythonExecutable = filepath.Join(venv, "bin", "python") // For linoox
	}

	if _, err := os.Stat(pythonExecutable); os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to find Python executable in the virtual environment: %s", pythonExecutable)
	}

	cmd := exec.Command(pythonExecutable, scripp, filePath)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to execute Python script: %s, error: %s", stderr.String(), err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(stdout.Bytes(), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Python output: %s", err)
	}

	return stdout.Bytes(), nil
}
