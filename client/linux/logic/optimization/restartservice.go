package optimization

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

type RestartServiceRequest struct {
	Name string `json:"name"`
}

type RestartServiceResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Service   string `json:"service"`
	Timestamp string `json:"timestamp"`
}

func HandleRestartService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RestartServiceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Name == "" {
		http.Error(w, "Invalid request, service name required", http.StatusBadRequest)
		return
	}

	// Stop the service
	cmd := exec.Command("systemctl", "stop", req.Name)
	output, err := cmd.CombinedOutput()

	resp := RestartServiceResponse{
		Service: req.Name,
	}

	if err != nil {
		resp.Status = "error"
		resp.Message = fmt.Sprintf("Failed to stop service: %v, output: %s", err, string(output))
	} else {
		resp.Status = "success"
		resp.Message = fmt.Sprintf("Service '%s' stopped successfully.", req.Name)
	}

	json.NewEncoder(w).Encode(resp)
}
