package python

import (
	"bytes"
	"fmt"
	"log/slog"
	"ml-orchestrator-module/internal/config"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func GodFunction(filePath string, cfg *config.Config) ([]byte, error) {
	venv := cfg.Venv
	scripp := cfg.Parser
	pythonExecutable := filepath.Join(venv, "Scripts", "python.exe")
	if runtime.GOOS != "windows" {
		pythonExecutable = filepath.Join(venv, "bin", "python")
	}
	if _, err := os.Stat(pythonExecutable); os.IsNotExist(err) {
		return nil, fmt.Errorf("Python executable not found: %s", pythonExecutable)
	}

	// Set up command
	cmd := exec.Command(pythonExecutable, scripp, filePath)

	// Set PYTHONPATH for the virtual environment
	cmd.Env = append(os.Environ(), "PYTHONPATH="+filepath.Join(venv, "Lib", "site-packages"))

	// Set working directory
	cmd.Dir = filepath.Dir(scripp)

	// Capture stdout and stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("Python script failed: %s, error: %s", stderr.String(), err)
	}

	// Parse output
	/*var result map[string]interface{}
	err = json.Unmarshal(stdout.Bytes(), &result)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse script output: %s", err)
	}*/
	slog.Info("Below is output. ")
	slog.AnyValue(stdout.Bytes())
	return stdout.Bytes(), nil
}
