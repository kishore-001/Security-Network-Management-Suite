package log

import (
	"encoding/json"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type FileLogEntry struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

type EventLogEntry struct {
	Provider string `json:"provider"`
	Message  string `json:"message"`
}

type AllLogs struct {
	EventLogs []EventLogEntry `json:"event_logs"`
	FileLogs  []FileLogEntry  `json:"file_logs"`
}

func cleanString(s string) string {
	// Replace escape sequences with a space
	replacer := strings.NewReplacer(
		"\r\n", " ",
		"\n", " ",
		"\r", " ",
		"\t", " ",
	)
	return replacer.Replace(s)
}

func HandleAllSystemLogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	// 1. Get Event Logs from PowerShell
	psCommand := `Get-WinEvent -LogName System,Application,Security -MaxEvents 50 | Select-Object ProviderName, Message | ConvertTo-Json -Compress`
	cmd := exec.Command("powershell", "-Command", psCommand)
	output, err := cmd.Output()
	if err != nil {
		http.Error(w, "Failed to retrieve event logs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var eventLogs []EventLogEntry
	if err := json.Unmarshal(output, &eventLogs); err != nil {
		var single EventLogEntry
		if err2 := json.Unmarshal(output, &single); err2 == nil {
			eventLogs = []EventLogEntry{single}
		} else {
			http.Error(w, "Failed to parse event logs", http.StatusInternalServerError)
			return
		}
	}

	// Clean event log messages
	for i := range eventLogs {
		eventLogs[i].Message = cleanString(eventLogs[i].Message)
	}

	// 2. Read log files (excluding specific ones)
	logDirs := []string{
		`C:\Windows\Logs`,
		`C:\ProgramData\Microsoft\Windows\WER\ReportArchive`,
	}

	var fileLogs []FileLogEntry
	for _, dir := range logDirs {
		files, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			fullPath := filepath.Join(dir, file.Name())
			ext := filepath.Ext(fullPath)
			if ext != ".log" && ext != ".txt" {
				continue
			}

			// Exclude unwanted file
			if strings.EqualFold(file.Name(), "StorGroupPolicy.log") {
				continue
			}

			data, err := os.ReadFile(fullPath)
			if err != nil {
				continue
			}

			fileLogs = append(fileLogs, FileLogEntry{
				Path:    fullPath,
				Content: cleanString(string(data)),
			})
		}
	}

	// 3. Combine and return all logs
	logs := AllLogs{
		EventLogs: eventLogs,
		FileLogs:  fileLogs,
	}

	if err := json.NewEncoder(w).Encode(logs); err != nil {
		http.Error(w, "Failed to encode logs: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
