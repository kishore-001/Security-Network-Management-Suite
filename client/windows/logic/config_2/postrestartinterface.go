package config_2

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

// HandleRestartInterfaces handles the request to restart all enabled interfaces (Windows version)
func HandleRestartInterfaces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		sendError1(w, "Failed to get network interfaces", err)
		return
	}

	restartedInterfaces := []string{}

	for _, iface := range interfaces {
		// Skip loopback and virtual interfaces
		if iface.Flags&net.FlagLoopback != 0 ||
			strings.Contains(iface.Name, "vEthernet") ||
			strings.Contains(iface.Name, "Loopback") ||
			strings.Contains(iface.Name, "Virtual") {
			continue
		}

		if iface.Flags&net.FlagUp != 0 {
			err := restartInterfaceWindows(iface.Name)
			if err != nil {
				fmt.Printf("Failed to restart interface %s: %v\n", iface.Name, err)
				continue
			}
			restartedInterfaces = append(restartedInterfaces, iface.Name)
		}
	}

	response := map[string]interface{}{
		"status":     "success",
		"message":    fmt.Sprintf("Restarted %d interfaces", len(restartedInterfaces)),
		"interfaces": restartedInterfaces,
	}

	if len(restartedInterfaces) == 0 {
		response["status"] = "warning"
		response["message"] = "No interfaces were restarted"
	}

	json.NewEncoder(w).Encode(response)
}

// restartInterfaceWindows disables and re-enables a network adapter using PowerShell
func restartInterfaceWindows(interfaceName string) error {
	// Use PowerShell to disable the interface
	disableCmd := exec.Command("powershell", "-Command", fmt.Sprintf(`Disable-NetAdapter -Name "%s" -Confirm:$false`, interfaceName))
	disableOutput, err := disableCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to disable interface: %v, output: %s", err, string(disableOutput))
	}

	time.Sleep(1 * time.Second) // Wait for clean disable

	// Re-enable the interface
	enableCmd := exec.Command("powershell", "-Command", fmt.Sprintf(`Enable-NetAdapter -Name "%s" -Confirm:$false`, interfaceName))
	enableOutput, err := enableCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to enable interface: %v, output: %s", err, string(enableOutput))
	}

	fmt.Printf("Successfully restarted interface %s\n", interfaceName)
	return nil
}

// sendError1 sends an error response as JSON
func sendError1(w http.ResponseWriter, msg string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "error",
		"message": msg,
		"details": err.Error(),
	})
}
