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

type NetworkUpdateRequest struct {
	Method  string `json:"method"`
	IP      string `json:"ip,omitempty"`
	Subnet  string `json:"subnet,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	DNS     string `json:"dns,omitempty"`
}

type NetworkUpdateResponse struct {
	Success   bool                  `json:"success"`
	Message   string                `json:"message"`
	Details   string                `json:"details,omitempty"`
	OldConfig *NetworkUpdateRequest `json:"old_config,omitempty"`
	NewConfig *NetworkUpdateRequest `json:"new_config"`
}

func HandleUpdateNetworkConfig(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Current Date and Time (UTC):", time.Now().UTC().Format("2006-01-02 15:04:05"))
	fmt.Println("Handling network configuration update request on Windows...")

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	var request NetworkUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		sendErrorResponse(w, "Failed to parse request body", err)
		return
	}

	if request.Method != "static" && request.Method != "dynamic" {
		sendErrorResponse(w, "Invalid method, must be 'static' or 'dynamic'", nil)
		return
	}

	if request.Method == "static" && request.IP == "" {
		sendErrorResponse(w, "IP address is required for static configuration", nil)
		return
	}

	iface, err := getActiveInterface()
	if err != nil {
		sendErrorResponse(w, "Failed to determine active interface", err)
		return
	}

	oldConfig, err := getCurrentNetworkConfigWindows(iface)
	if err != nil {
		sendErrorResponse(w, "Failed to get current configuration", err)
		return
	}

	var success bool
	var details string
	if request.Method == "dynamic" {
		success, details = setDynamicIPWindows(iface)
	} else {
		if request.Subnet == "" {
			request.Subnet = oldConfig.Subnet
		}
		if request.Gateway == "" {
			request.Gateway = oldConfig.Gateway
		}
		if request.DNS == "" {
			request.DNS = oldConfig.DNS
		}
		success, details = setStaticIPWindows(iface, request.IP, request.Subnet, request.Gateway, request.DNS)
	}

	msg := "Failed to update network configuration"
	if success {
		msg = "Network configuration updated successfully"
	}

	resp := NetworkUpdateResponse{
		Success:   success,
		Message:   msg,
		Details:   details,
		OldConfig: oldConfig,
		NewConfig: &request,
	}
	json.NewEncoder(w).Encode(resp)
}

func getActiveInterface() (string, error) {
	cmd := exec.Command("powershell", "-Command", "Get-NetRoute -DestinationPrefix 0.0.0.0/0 | Sort-Object RouteMetric | Select-Object -First 1 -ExpandProperty InterfaceAlias")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func getCurrentNetworkConfigWindows(iface string) (*NetworkUpdateRequest, error) {
	cmd := exec.Command("powershell", "-Command", fmt.Sprintf(`(Get-NetIPAddress -InterfaceAlias "%s" -AddressFamily IPv4 | Where-Object {$_.IPAddress -ne $null} | Select-Object -First 1).IPAddress`, iface))
	ip, _ := cmd.Output()

	cmd = exec.Command("powershell", "-Command", fmt.Sprintf(`(Get-NetIPAddress -InterfaceAlias "%s" -AddressFamily IPv4 | Where-Object {$_.PrefixLength -ne $null} | Select-Object -First 1).PrefixLength`, iface))
	prefix, _ := cmd.Output()

	cmd = exec.Command("powershell", "-Command", fmt.Sprintf(`(Get-NetRoute -InterfaceAlias "%s" -DestinationPrefix 0.0.0.0/0 | Select-Object -First 1).NextHop`, iface))
	gw, _ := cmd.Output()

	cmd = exec.Command("powershell", "-Command", fmt.Sprintf(`(Get-DnsClientServerAddress -InterfaceAlias "%s" -AddressFamily IPv4).ServerAddresses -join ","`, iface))
	dns, _ := cmd.Output()

	return &NetworkUpdateRequest{
		Method:  "static", // Cannot detect DHCP reliably here
		IP:      strings.TrimSpace(string(ip)),
		Subnet:  prefixToNetmask(strings.TrimSpace(string(prefix))),
		Gateway: strings.TrimSpace(string(gw)),
		DNS:     strings.TrimSpace(string(dns)),
	}, nil
}

func setDynamicIPWindows(iface string) (bool, string) {
	resetCmd := fmt.Sprintf(`netsh interface ip set address name="%s" source=dhcp`, iface)
	resetDNS := fmt.Sprintf(`netsh interface ip set dns name="%s" source=dhcp`, iface)

	if err := exec.Command("cmd", "/C", resetCmd).Run(); err != nil {
		return false, "Failed to switch to DHCP (IP)"
	}
	if err := exec.Command("cmd", "/C", resetDNS).Run(); err != nil {
		return false, "IP set to DHCP, but DNS change failed"
	}
	return true, "Set to dynamic IP using netsh"
}

func setStaticIPWindows(iface, ip, subnet, gateway, dns string) (bool, string) {
	subnetMask := subnet
	cmd := fmt.Sprintf(`netsh interface ip set address name="%s" static %s %s %s`, iface, ip, subnetMask, gateway)
	if err := exec.Command("cmd", "/C", cmd).Run(); err != nil {
		return false, "Failed to set static IP"
	}

	if dns != "" {
		for i, d := range strings.Split(dns, ",") {
			dnsCmd := fmt.Sprintf(`netsh interface ip add dns name="%s" %s %s`, iface, d, func() string {
				if i == 0 {
					return "validate=no"
				}
				return "index=2"
			}())
			if err := exec.Command("cmd", "/C", dnsCmd).Run(); err != nil {
				return false, "IP set, DNS failed: " + err.Error()
			}
		}
	}
	return true, "Successfully applied static IP"
}

func prefixToNetmask(prefix string) string {
	p := 0
	fmt.Sscanf(prefix, "%d", &p)
	mask := net.CIDRMask(p, 32)
	return fmt.Sprintf("%d.%d.%d.%d", mask[0], mask[1], mask[2], mask[3])
}

func sendErrorResponse(w http.ResponseWriter, msg string, err error) {
	w.WriteHeader(http.StatusBadRequest)
	resp := NetworkUpdateResponse{
		Success: false,
		Message: msg,
	}
	if err != nil {
		resp.Details = err.Error()
	}
	json.NewEncoder(w).Encode(resp)
}
