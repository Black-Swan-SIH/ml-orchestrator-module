package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
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

func Cleanify(input []byte) ([]byte, error) {

	strInput := string(input)

	strInput = strings.ReplaceAll(strInput, "Spacy Model is loading", "")
	strInput = strings.ReplaceAll(strInput, "files/res/pdf\\da.pdf", "")

	strInput = strings.TrimSpace(strInput)

	var data []map[string]interface{}
	if err := json.Unmarshal([]byte(strInput), &data); err != nil {
		return nil, fmt.Errorf("failed to parse input: %w", err)
	}

	output, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to restructure JSON: %w", err)
	}

	return output, nil
}
