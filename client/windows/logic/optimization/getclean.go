package optimization

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"time"
)

type CleanResult struct {
	Status            string   `json:"status"`
	Message           string   `json:"message"`
	DirsCleaned       []string `json:"dirs_cleaned"`
	DirsFailed        []string `json:"dirs_failed"`
	RecycleBinCleared bool     `json:"recycle_bin_cleared"`
	Timestamp         string   `json:"timestamp"`
	User              string   `json:"user"`
}

// cleanDir removes all files and folders inside the given directory
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
		fullPath := filepath.Join(path, name)
		err := os.RemoveAll(fullPath)
		if err != nil {
			return err
		}
	}
	return nil
}

// clearRecycleBin tries to empty the Recycle Bin using PowerShell
func clearRecycleBin() error {
	cmd := exec.Command("PowerShell", "-Command", "Clear-RecycleBin -Force")
	return cmd.Run()
}

func HandleFileClean(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get current user info
	usr, _ := user.Current()

	// Fetch environment paths
	temp := os.TempDir()
	localAppData := os.Getenv("LOCALAPPDATA")
	appData := os.Getenv("APPDATA")
	winDir := os.Getenv("SystemRoot")
	userProfile := os.Getenv("USERPROFILE")

	// Directories to clean
	dirs := []string{
		temp,
		filepath.Join(localAppData, "Temp"),
		filepath.Join(localAppData, "Cache"),
		filepath.Join(appData, "Microsoft", "Windows", "Recent"),
		filepath.Join(winDir, "Temp"),
		filepath.Join(userProfile, "AppData", "Local", "Microsoft", "Windows", "INetCache"),
		filepath.Join(userProfile, "AppData", "Local", "CrashDumps"),
		filepath.Join(userProfile, "AppData", "LocalLow", "Temp"),
	}

	var cleanedDirs []string
	var failedDirs []string

	// Attempt to clean each directory
	for _, dir := range dirs {
		if err := cleanDir(dir); err == nil {
			cleanedDirs = append(cleanedDirs, dir)
		} else {
			failedDirs = append(failedDirs, fmt.Sprintf("%s (%v)", dir, err))
		}
	}

	// Attempt to clear Recycle Bin
	recycleBinCleared := true
	if err := clearRecycleBin(); err != nil {
		recycleBinCleared = false
	}

	// Final status and message
	status := "success"
	if len(failedDirs) > 0 || !recycleBinCleared {
		status = "partial"
	}
	message := fmt.Sprintf("Cleaned: %d directories. Failed: %d directories. Recycle Bin Cleared: %v",
		len(cleanedDirs), len(failedDirs), recycleBinCleared)

	// Response payload
	result := CleanResult{
		Status:            status,
		Message:           message,
		DirsCleaned:       cleanedDirs,
		DirsFailed:        failedDirs,
		RecycleBinCleared: recycleBinCleared,
		Timestamp:         time.Now().UTC().Format("2006-01-02 15:04:05"),
		User:              usr.Username,
	}

	// Return result
	json.NewEncoder(w).Encode(result)
}
