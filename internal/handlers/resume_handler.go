package handlers

import (
	"ml-orchestrator-module/internal/config"
	_ "ml-orchestrator-module/internal/config"
	"ml-orchestrator-module/internal/python"
	"ml-orchestrator-module/storage"
	"net/http"
	"os"
	//"path/filepath" Dont keep it in prod
)

// ResumeDaddy daddy of resume
func ResumeDaddy(cfg *config.Config, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	tempPath, err := storage.TempSave(file, header.Filename)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempPath)

	result, err := python.GodFunction(tempPath, cfg)
	if err != nil {
		http.Error(w, "Failed to process resume: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
