package config_2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

// NetworkUpdateRequest represents the request for updating network configuration
type NetworkUpdateRequest struct {
	Method  string `json:"method"`            // "static" or "dynamic"
	IP      string `json:"ip,omitempty"`      // IP address (only for static)
	Subnet  string `json:"subnet,omitempty"`  // Subnet mask (only for static)
	Gateway string `json:"gateway,omitempty"` // Default gateway (only for static)
	DNS     string `json:"dns,omitempty"`     // DNS servers (comma-separated)
}

// NetworkUpdateResponse represents the response after updating network configuration
type NetworkUpdateResponse struct {
	Success   bool                  `json:"success"`
	Message   string                `json:"message"`
	Details   string                `json:"details,omitempty"`
	OldConfig *NetworkUpdateRequest `json:"old_config,omitempty"`
	NewConfig *NetworkUpdateRequest `json:"new_config"`
}

// HandleUpdateNetworkConfig handles POST requests to update network configuration
func HandleUpdateNetworkConfig(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Current Date and Time (UTC - YYYY-MM-DD HH:MM:SS formatted):",
		time.Now().UTC().Format("2006-01-02 15:04:05"))
	fmt.Println("Current User's Login: kishore-001")
	fmt.Println("Handling network configuration update request...")

	// Check for POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// Set content type for response
	w.Header().Set("Content-Type", "application/json")

	// Parse the request body
	var request NetworkUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		sendErrorResponse(w, "Failed to parse request body", err)
		return
	}

	// Validate the request
	if request.Method != "static" && request.Method != "dynamic" {
		sendErrorResponse(w, "Invalid method, must be 'static' or 'dynamic'", nil)
		return
	}

	if request.Method == "static" {
		// For static method, validate required fields
		if request.IP == "" {
			sendErrorResponse(w, "IP address is required for static configuration", nil)
			return
		}
	}

	// Get the active interface
	activeInterface, err := getActiveInterface()
	if err != nil {
		sendErrorResponse(w, "Failed to get active network interface", err)
		return
	}

	// Get the current network configuration for the active interface
	oldConfig, err := getCurrentNetworkConfig(activeInterface)
	if err != nil {
		sendErrorResponse(w, "Failed to get current network configuration", err)
		return
	}

	// Update the network configuration
	var details string
	var success bool
	if request.Method == "dynamic" {
		success, details = setDynamicIP(activeInterface)
	} else {
		// Fill in missing values with current values
		if request.Subnet == "" {
			request.Subnet = oldConfig.Subnet
		}
		if request.Gateway == "" {
			request.Gateway = oldConfig.Gateway
		}
		if request.DNS == "" {
			request.DNS = oldConfig.DNS
		}

		success, details = setStaticIP(activeInterface, request.IP, request.Subnet, request.Gateway, request.DNS)
	}

	message := "Failed to update network configuration"
	if success {
		message = "Network configuration updated successfully"
	}

	response := NetworkUpdateResponse{
		Success:   success,
		Message:   message,
		Details:   details,
		OldConfig: oldConfig,
		NewConfig: &request,
	}

	// Send the response
	json.NewEncoder(w).Encode(response)
}

// getCurrentNetworkConfig gets the current network configuration for the interface
func getCurrentNetworkConfig(iface string) (*NetworkUpdateRequest, error) {
	config := &NetworkUpdateRequest{}

	// Check if interface is using DHCP
	isUsingDHCP, err := isInterfaceUsingDHCP(iface)
	if err != nil {
		return nil, err
	}

	if isUsingDHCP {
		config.Method = "dynamic"
	} else {
		config.Method = "static"
	}

	// Get IP address and subnet
	ipOutput, err := exec.Command("ip", "-o", "-4", "addr", "show", iface).Output()
	if err == nil {
		ipLine := string(ipOutput)
		ipMatch := strings.Fields(ipLine)
		if len(ipMatch) > 3 {
			ipCIDR := ipMatch[3]
			ipParts := strings.Split(ipCIDR, "/")
			if len(ipParts) > 0 {
				config.IP = ipParts[0]
				if len(ipParts) > 1 {
					// Convert CIDR to subnet mask
					cidr := ipParts[1]
					config.Subnet = cidrToNetmask(cidr)
				}
			}
		}
	}

	// Get gateway
	gatewayOutput, err := exec.Command("ip", "route", "show", "dev", iface).Output()
	if err == nil {
		gatewayLines := strings.Split(string(gatewayOutput), "\n")
		for _, line := range gatewayLines {
			if strings.Contains(line, "default via") {
				gatewayFields := strings.Fields(line)
				if len(gatewayFields) > 2 {
					config.Gateway = gatewayFields[2]
				}
				break
			}
		}
	}

	// Get DNS servers
	dnsServers, err := getDNSServers()
	if err == nil && len(dnsServers) > 0 {
		config.DNS = strings.Join(dnsServers, ",")
	}

	return config, nil
}

// isInterfaceUsingDHCP checks if the interface is configured to use DHCP
func isInterfaceUsingDHCP(iface string) (bool, error) {
	// Check for DHCP client processes
	psOutput, err := exec.Command("ps", "-ef").Output()
	if err != nil {
		return false, err
	}

	psString := string(psOutput)

	// Check for common DHCP client processes
	if strings.Contains(psString, "dhclient "+iface) ||
		strings.Contains(psString, "dhclient -i "+iface) ||
		strings.Contains(psString, "dhcpcd "+iface) {
		return true, nil
	}

	// Also check NetworkManager if available
	_, err = exec.LookPath("nmcli")
	if err == nil {
		nmOutput, err := exec.Command("nmcli", "-g", "IP4.METHOD", "device", "show", iface).Output()
		if err == nil && strings.Contains(string(nmOutput), "auto") {
			return true, nil
		}
	}

	// Check ifconfig/dhcpcd config files (for non-NetworkManager systems)
	interfaceConfigs := []string{
		"/etc/network/interfaces",
		"/etc/sysconfig/network-scripts/ifcfg-" + iface,
	}

	for _, configFile := range interfaceConfigs {
		content, err := os.ReadFile(configFile)
		if err == nil {
			contentStr := string(content)
			if strings.Contains(contentStr, "dhcp") ||
				strings.Contains(contentStr, "BOOTPROTO=dhcp") {
				return true, nil
			}
		}
	}

	return false, nil
}

// getActiveInterface gets the active network interface
func getActiveInterface() (string, error) {
	// Run ip route command to get default route
	cmd := exec.Command("ip", "route", "show", "default")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get default route: %w", err)
	}

	// Parse output like: "default via 192.168.1.1 dev eth0"
	routeStr := string(output)
	routeParts := strings.Fields(routeStr)

	for i, part := range routeParts {
		if part == "dev" && i+1 < len(routeParts) {
			return routeParts[i+1], nil
		}
	}

	return "", fmt.Errorf("could not determine active interface")
}

// setDynamicIP sets the interface to use DHCP
func setDynamicIP(iface string) (bool, string) {
	// Check for NetworkManager
	_, err := exec.LookPath("nmcli")
	if err == nil {
		// Use NetworkManager to set dynamic IP
		cmd := exec.Command("nmcli", "con", "mod", iface,
			"ipv4.method", "auto",
			"ipv4.addresses", "",
			"ipv4.gateway", "",
			"ipv4.dns", "")

		output, err := cmd.CombinedOutput()
		if err != nil {
			return false, fmt.Sprintf("Failed to set dynamic IP using NetworkManager: %s: %s", err, output)
		}

		// Apply the connection
		reloadCmd := exec.Command("nmcli", "con", "up", iface)
		reloadOut, err := reloadCmd.CombinedOutput()
		if err != nil {
			return false, fmt.Sprintf("Failed to apply NetworkManager connection: %s: %s", err, reloadOut)
		}

		return true, "Successfully set dynamic IP using NetworkManager"
	}

	// Fall back to dhclient
	// First release the current lease
	releaseCmd := exec.Command("dhclient", "-r", iface)
	releaseCmd.Run() // Ignore errors

	// Start DHCP client
	cmd := exec.Command("dhclient", iface)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false, fmt.Sprintf("Failed to set dynamic IP using dhclient: %s: %s", err, output)
	}

	return true, "Successfully set dynamic IP using dhclient"
}

// setStaticIP sets a static IP configuration for the interface
func setStaticIP(iface, ip, subnet, gateway, dns string) (bool, string) {
	// Check for NetworkManager
	_, err := exec.LookPath("nmcli")
	if err == nil {
		// Use NetworkManager to set static IP
		dnsArgs := ""
		if dns != "" {
			dnsArgs = dns
		}

		// Build the command
		cmd := exec.Command("nmcli", "con", "mod", iface,
			"ipv4.method", "manual",
			"ipv4.addresses", fmt.Sprintf("%s/%s", ip, subnetToPrefix(subnet)),
			"ipv4.gateway", gateway)

		if dnsArgs != "" {
			cmd = exec.Command("nmcli", "con", "mod", iface,
				"ipv4.method", "manual",
				"ipv4.addresses", fmt.Sprintf("%s/%s", ip, subnetToPrefix(subnet)),
				"ipv4.gateway", gateway,
				"ipv4.dns", dnsArgs)
		}

		output, err := cmd.CombinedOutput()
		if err != nil {
			return false, fmt.Sprintf("Failed to set static IP using NetworkManager: %s: %s", err, output)
		}

		// Apply the connection
		reloadCmd := exec.Command("nmcli", "con", "up", iface)
		reloadOut, err := reloadCmd.CombinedOutput()
		if err != nil {
			return false, fmt.Sprintf("Failed to apply NetworkManager connection: %s: %s", err, reloadOut)
		}

		return true, "Successfully set static IP using NetworkManager"
	}

	// Fall back to ip command
	// First flush existing addresses
	flushCmd := exec.Command("ip", "addr", "flush", "dev", iface)
	flushCmd.Run() // Ignore errors

	// Add IP address
	addIPCmd := exec.Command("ip", "addr", "add", fmt.Sprintf("%s/%s", ip, subnetToPrefix(subnet)), "dev", iface)
	output, err := addIPCmd.CombinedOutput()
	if err != nil {
		return false, fmt.Sprintf("Failed to set static IP: %s: %s", err, output)
	}

	// Add gateway route
	if gateway != "" {
		// Delete existing default route first
		delRouteCmd := exec.Command("ip", "route", "del", "default")
		delRouteCmd.Run() // Ignore errors

		// Add new default gateway
		addRouteCmd := exec.Command("ip", "route", "add", "default", "via", gateway, "dev", iface)
		routeOut, err := addRouteCmd.CombinedOutput()
		if err != nil {
			return false, fmt.Sprintf("Failed to set gateway: %s: %s", err, routeOut)
		}
	}

	// Update DNS servers if provided
	if dns != "" {
		updatedDNS := false
		dnsServers := strings.Split(dns, ",")

		// Try updating resolv.conf
		resolvContent := ""
		for _, server := range dnsServers {
			server = strings.TrimSpace(server)
			if server != "" {
				resolvContent += fmt.Sprintf("nameserver %s\n", server)
			}
		}

		if resolvContent != "" {
			err := ioutil.WriteFile("/etc/resolv.conf", []byte(resolvContent), 0644)
			if err == nil {
				updatedDNS = true
			}
		}

		if !updatedDNS {
			return true, fmt.Sprintf("IP and gateway set successfully, but DNS servers could not be updated: %s", dns)
		}
	}

	return true, "Successfully set static IP configuration"
}

// subnetToPrefix converts a subnet mask to CIDR prefix length
func subnetToPrefix(subnet string) string {
	if subnet == "" {
		return "24" // Default to /24 if not specified
	}

	// If already in CIDR format, return the prefix
	if strings.Contains(subnet, "/") {
		parts := strings.Split(subnet, "/")
		if len(parts) > 1 {
			return parts[1]
		}
	}

	// Convert dotted decimal to CIDR
	var count int
	octets := strings.Split(subnet, ".")
	if len(octets) != 4 {
		return "24" // Default to /24 for invalid input
	}

	for _, octet := range octets {
		var octetInt int
		fmt.Sscanf(octet, "%d", &octetInt)
		for i := 7; i >= 0; i-- {
			if octetInt&(1<<uint(i)) != 0 {
				count++
			} else {
				// Break at the first 0 bit
				break
			}
		}
	}

	return fmt.Sprintf("%d", count)
}

// sendErrorResponse sends an error response in JSON format
func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusBadRequest)

	errorDetail := ""
	if err != nil {
		errorDetail = err.Error()
	}

	errorResp := NetworkUpdateResponse{
		Success: false,
		Message: message,
		Details: errorDetail,
	}

	json.NewEncoder(w).Encode(errorResp)
}
