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

// InterfaceInfo represents information about a network interface
type InterfaceInfo struct {
	Mode   string `json:"mode"`
	Status string `json:"status"`
}

// NetworkConfigResponse represents the response format you specified
type NetworkConfigResponse struct {
	IPMethod  string                   `json:"ip_method"`
	IPAddress string                   `json:"ip_address"`
	Gateway   string                   `json:"gateway"`
	Subnet    string                   `json:"subnet"`
	DNS       string                   `json:"dns"`
	Interface map[string]InterfaceInfo `json:"interface"`
}

// HandleNetworkConfig provides network configuration as an HTTP response
// in the exact format specified
func HandleNetworkConfig(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Current Date and Time (UTC - YYYY-MM-DD HH:MM:SS formatted):",
		time.Now().UTC().Format("2006-01-02 15:04:05"))
	fmt.Println("Current User's Login: kishore-001")
	fmt.Println("Handling network configuration request...")

	// Check for GET method
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Get all network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		sendError(w, "Failed to get network interfaces", err)
		return
	}

	// Get default route to determine active interface and gateway
	defaultRoute, err := getDefaultRoute()
	if err != nil {
		defaultRoute = DefaultRoute{} // Continue with empty route info
	}

	// Get DNS servers
	dnsServers, err := getDNSServers()
	if err != nil {
		dnsServers = []string{} // Continue with empty DNS info
	}

	// Prepare response in the requested format
	response := NetworkConfigResponse{
		IPMethod:  "unknown",
		IPAddress: "",
		Gateway:   defaultRoute.Gateway,
		Subnet:    "",
		DNS:       strings.Join(dnsServers, ", "),
		Interface: make(map[string]InterfaceInfo),
	}

	// Process each interface
	interfaceCount := 1
	for _, iface := range interfaces {
		// Skip loopback, non-up interfaces, and virtual interfaces
		if iface.Flags&net.FlagLoopback != 0 ||
			iface.Flags&net.FlagUp == 0 ||
			strings.Contains(iface.Name, "docker") ||
			strings.Contains(iface.Name, "veth") ||
			strings.Contains(iface.Name, "br-") {
			continue
		}

		// Determine if this interface is active
		isActive := iface.Name == defaultRoute.Interface
		status := "inactive" // Default to inactive
		if isActive {
			status = "active"
			// Get IP method for active interface only
			response.IPMethod = getIPMethod(iface.Name)
		}

		// Add interface to the map
		response.Interface[fmt.Sprintf("%d", interfaceCount)] = InterfaceInfo{
			Mode:   iface.Name,
			Status: status, // Now will be either "active" or "inactive"
		}
		interfaceCount++

		// Get IP address and subnet for active interface
		if isActive {
			addrs, err := iface.Addrs()
			if err == nil {
				for _, addr := range addrs {
					if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
						if ipnet.IP.To4() != nil { // Only consider IPv4 for now
							response.IPAddress = ipnet.IP.String()

							// Calculate subnet in CIDR notation
							ones, _ := ipnet.Mask.Size()
							response.Subnet = fmt.Sprintf("%s/%d",
								ipnet.IP.Mask(ipnet.Mask).String(), ones)
							break
						}
					}
				}
			}
		}
	}

	// Convert response.IPMethod to "static" or "dynamic" as requested
	if response.IPMethod == "dhcp" {
		response.IPMethod = "dynamic"
	} else if response.IPMethod == "unknown" {
		response.IPMethod = "static" // Default to static if unknown
	}

	fmt.Println("Sending network configuration response...")
	// Send the network configuration using NewEncoder().Encode()
	json.NewEncoder(w).Encode(response)
}

// DefaultRoute represents the default gateway route
type DefaultRoute struct {
	Interface string
	Gateway   string
}

// getDefaultRoute gets the default gateway and interface
func getDefaultRoute() (DefaultRoute, error) {
	route := DefaultRoute{}

	// Run ip route command to get default route
	cmd := exec.Command("ip", "route", "show", "default")
	output, err := cmd.Output()
	if err != nil {
		return route, err
	}

	// Parse output like: "default via 192.168.1.1 dev eth0"
	routeStr := string(output)
	if routeStr == "" {
		return route, fmt.Errorf("no default route found")
	}

	parts := strings.Fields(routeStr)
	if len(parts) < 5 || parts[0] != "default" || parts[1] != "via" {
		return route, fmt.Errorf("unexpected route format: %s", routeStr)
	}

	route.Gateway = parts[2]

	// Find the dev part
	for i, part := range parts {
		if part == "dev" && i+1 < len(parts) {
			route.Interface = parts[i+1]
			break
		}
	}

	return route, nil
}

// getDNSServers reads DNS servers from /etc/resolv.conf
func getDNSServers() ([]string, error) {
	servers := []string{}

	// Run grep to extract nameservers from resolv.conf
	cmd := exec.Command("grep", "^nameserver", "/etc/resolv.conf")
	output, err := cmd.Output()
	if err != nil {
		return servers, err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			servers = append(servers, parts[1])
		}
	}

	return servers, nil
}

// getIPMethod determines if the interface uses DHCP or static IP
func getIPMethod(ifaceName string) string {
	// Check if dhclient is running for this interface
	cmd := exec.Command("ps", "-ef")
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}

	if strings.Contains(string(output), "dhclient "+ifaceName) ||
		strings.Contains(string(output), "dhclient -i "+ifaceName) ||
		strings.Contains(string(output), "dhcpcd "+ifaceName) {
		return "dhcp"
	}

	// Also check for NetworkManager DHCP
	cmd = exec.Command("nmcli", "-g", "IP4.ADDRESS", "device", "show", ifaceName)
	if _, err := cmd.Output(); err == nil {
		// If we can get the IP from NetworkManager, check if DHCP is enabled
		cmd = exec.Command("nmcli", "-g", "IP4.METHOD", "device", "show", ifaceName)
		output, err := cmd.Output()
		if err == nil && strings.Contains(string(output), "auto") {
			return "dhcp"
		}
	}

	// Default to static if no DHCP client found
	return "static"
}

// sendError sends an error response in JSON format
func sendError(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	errorResp := map[string]string{
		"error":   message,
		"details": err.Error(),
	}
	json.NewEncoder(w).Encode(errorResp)
}
