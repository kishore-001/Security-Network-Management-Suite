package optimization

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os/exec"
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

// getRestartableServices retrieves active services managed by systemd
func getServices() (map[string]bool, error) {
	cmd := exec.Command("systemctl", "list-units", "--type=service", "--state=running", "--no-pager", "--no-legend")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	services := make(map[string]bool)
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) > 0 {
			unitName := fields[0]
			if strings.HasSuffix(unitName, ".service") {
				service := strings.TrimSuffix(unitName, ".service")
				services[service] = true
			}
		}
	}

	return services, nil
}

func HandleListService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get systemctl-managed restartable services
	restartableServices, err := getServices()
	if err != nil {
		http.Error(w, "Failed to list systemd services", http.StatusInternalServerError)
		return
	}

	procs, err := process.Processes()
	if err != nil {
		http.Error(w, "Failed to fetch processes", http.StatusInternalServerError)
		return
	}

	services := []ServiceInfo{}

	for _, proc := range procs {
		name, _ := proc.Name()
		cmdline, _ := proc.Cmdline()
		username, _ := proc.Username()

		// Check if this process corresponds to a restartable systemd service
		nameLower := strings.ToLower(name)
		if restartableServices[nameLower] {
			svc := ServiceInfo{
				PID:     proc.Pid,
				User:    username,
				Name:    name,
				Cmdline: cmdline,
			}
			services = append(services, svc)
		}
	}

	result := ServiceListResult{
		Status:    "success",
		Message:   "List of restartable user-level services (Linux)",
		Services:  services,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}

	json.NewEncoder(w).Encode(result)
}
