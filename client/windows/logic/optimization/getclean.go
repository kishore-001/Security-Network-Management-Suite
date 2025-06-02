package optimization

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

type CleanResult struct {
	Status      string   `json:"status"`
	Message     string   `json:"message"`
	DirsCleaned []string `json:"dirs_cleaned"`
	Timestamp   string   `json:"timestamp"`
	User        string   `json:"user"`
}

// cleanDir tries to remove all files/subfolders in a directory (but not the dir itself)
func cleanDir(path string) error {
	dir, err := os.Open(path)
	if err != nil {
		return err
	}
	defer dir.Close()

	names, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		fullpath := filepath.Join(path, name)
		err = os.RemoveAll(fullpath)
		if err != nil {
			return err
		}
	}
	return nil
}

func HandleFileClean(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get current user and paths
	usr, _ := user.Current()
	temp := os.TempDir()

	localAppData := os.Getenv("LOCALAPPDATA") // e.g., C:\Users\<User>\AppData\Local

	dirs := []string{
		temp,                                 // System temp
		filepath.Join(localAppData, "Temp"),  // App-specific temp
		filepath.Join(localAppData, "Cache"), // App-specific cache
	}

	var cleaned []string
	var failed []string

	for _, dir := range dirs {
		if err := cleanDir(dir); err == nil {
			cleaned = append(cleaned, dir)
		} else {
			failed = append(failed, fmt.Sprintf("%s (%v)", dir, err))
		}
	}

	status := "success"
	message := fmt.Sprintf("Cleaned: %v", cleaned)
	if len(failed) > 0 {
		status = "partial"
		message = fmt.Sprintf("Cleaned: %v. Failed: %v", cleaned, failed)
	}

	result := CleanResult{
		Status:      status,
		Message:     message,
		DirsCleaned: cleaned,
		Timestamp:   time.Now().UTC().Format("2006-01-02 15:04:05"),
		User:        usr.Username,
	}

	json.NewEncoder(w).Encode(result)
}
