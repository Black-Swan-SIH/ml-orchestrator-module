package storage

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func TempSave(file io.Reader, filename string) (string, error) {
	tempDir := "temp_uploads"
	err := os.MkdirAll(tempDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	tempPath := filepath.Join(tempDir, filename)
	out, err := os.Create(tempPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}

	return tempPath, nil
}

func isVenvActivated(venv string) bool {
	venvPython := filepath.Join(venv, "Scripts", "python") // For Windows
	if runtime.GOOS != "windows" {
		venvPython = filepath.Join(venv, "bin", "python") // For Linux/Mac
	}
	out, err := exec.Command(venvPython, "--version").Output()
	return err == nil && len(out) > 0
}
