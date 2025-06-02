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

// isUserService returns true if the process looks like a user-level service or app server
func isUserService(proc *process.Process) bool {
	name, _ := proc.Name()
	cmd, _ := proc.Cmdline()

	// Skip kernel threads and system daemons with no proper names
	if strings.HasPrefix(name, "[") || name == "" {
		return false
	}

	// Skip known system/daemon processes
	systemProcesses := []string{
		"systemd", "kthreadd", "rcu_sched", "ksoftirqd", "kworker",
		"bioset", "kswapd", "watchdog", "migration", "jbd2",
		"khugepaged", "ksmd",
	}
	for _, sys := range systemProcesses {
		if strings.Contains(name, sys) {
			return false
		}
	}

	// Common user-level services/servers keywords
	serviceKeywords := []string{
		"nginx", "apache", "httpd", "node", "python", "uwsgi", "gunicorn",
		"php", "ruby", "rails", "mysql", "postgres", "mongod", "redis", "memcached",
		"java", "go", "dotnet", "pm2", "serve", "flask", "django", "express",
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

	procs, _ := process.Processes()
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
		Message:   "List of detected user-level services and application servers",
		Services:  services,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}

	json.NewEncoder(w).Encode(result)
}
