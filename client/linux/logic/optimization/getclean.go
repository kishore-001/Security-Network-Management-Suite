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
	Status      string           `json:"status"`
	Message     string           `json:"message"`
	DirsCleaned []string         `json:"dirs_cleaned"`
	BytesFreed  map[string]int64 `json:"bytes_freed"`
	Timestamp   string           `json:"timestamp"`
	User        string           `json:"user"`
}

// sizeOfDir recursively calculates the total size of files in a directory
func sizeOfDir(path string) (int64, error) {
	var total int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			total += info.Size()
		}
		return nil
	})
	return total, err
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

	usr, _ := user.Current()
	userCache := filepath.Join(usr.HomeDir, ".cache")

	dirs := []string{"/tmp", "/var/tmp", userCache}
	var cleaned []string
	var failed []string
	bytesFreed := make(map[string]int64)

	for _, dir := range dirs {
		sizeBefore, err := sizeOfDir(dir)
		if err != nil {
			failed = append(failed, fmt.Sprintf("%s (size calc error: %v)", dir, err))
			continue
		}

		err = cleanDir(dir)
		if err == nil {
			cleaned = append(cleaned, dir)
			bytesFreed[dir] = sizeBefore
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
		BytesFreed:  bytesFreed,
		Timestamp:   time.Now().UTC().Format("2006-01-02 15:04:05"),
		User:        usr.Username,
	}

	json.NewEncoder(w).Encode(result)
}

