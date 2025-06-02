package optimization

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

type ServiceInfo struct {
	PID     int32  `json:"pid"`
	User    string `json:"user"`
	Name    string `json:"name"`
	Cmdline string `json:"cmdline"`
}

type ServiceListResult struct {
	Status    string        `json:"status"`
	Message   string        `json:"message"`
	Services  []ServiceInfo `json:"services"`
	Timestamp string        `json:"timestamp"`
}

// isUserService filters out system processes and focuses on user-level servers/apps
func isUserService(proc *process.Process) bool {
	name, _ := proc.Name()
	cmd, _ := proc.Cmdline()

	// Skip empty names
	if name == "" {
		return false
	}

	// Windows built-in system processes to skip
	systemProcesses := []string{
		"System", "Idle", "smss.exe", "csrss.exe", "wininit.exe", "services.exe",
		"lsass.exe", "svchost.exe", "winlogon.exe", "fontdrvhost.exe", "dwm.exe",
		"explorer.exe", "conhost.exe", "taskhostw.exe", "SearchIndexer.exe",
	}
	nameLower := strings.ToLower(name)
	for _, sys := range systemProcesses {
		if nameLower == strings.ToLower(sys) {
			return false
		}
	}

	// Detect common user services
	serviceKeywords := []string{
		"nginx", "apache", "httpd", "node", "python", "uwsgi", "gunicorn",
		"php", "ruby", "rails", "mysql", "postgres", "mongod", "redis",
		"memcached", "java", "go", "dotnet", "pm2", "serve", "flask", "django", "express",
		"powershell", "cmd.exe",
	}
	for _, kw := range serviceKeywords {
		if strings.Contains(strings.ToLower(name), kw) || strings.Contains(strings.ToLower(cmd), kw) {
			return true
		}
	}

	return false
}

func HandleListService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	procs, err := process.Processes()
	if err != nil {
		http.Error(w, "Failed to fetch processes", http.StatusInternalServerError)
		return
	}

	services := []ServiceInfo{}

	for _, proc := range procs {
		if isUserService(proc) {
			name, _ := proc.Name()
			cmd, _ := proc.Cmdline()
			username, _ := proc.Username()

			svc := ServiceInfo{
				PID:     proc.Pid,
				User:    username,
				Name:    name,
				Cmdline: cmd,
			}
			services = append(services, svc)
		}
	}

	result := ServiceListResult{
		Status:    "success",
		Message:   "List of detected user-level services and application servers (Windows)",
		Services:  services,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}

	json.NewEncoder(w).Encode(result)
}
