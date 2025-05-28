package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n===== System Monitoring CLI Tool =====")
		fmt.Println("1. Hostname & Timezone")
		fmt.Println("2. Network Management")
		fmt.Println("3. Health Monitoring")
		fmt.Println("4. Restart Services")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			getSystemInfo()
		case "2":
			getNetworkInfo()
		case "3":
			getHealthMetrics()
		case "4":
			ListAndRestartRunningServices()
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func getSystemInfo() {
	hostname, _ := os.Hostname()
	fmt.Println("Hostname:", hostname)
	fmt.Println("Timezone:", time.Now().Location())
}

func runCommand(cmd string, args ...string) {
	out, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		fmt.Printf("Error running command: %v\n", err)
	}
	fmt.Println(string(out))
}

func getNetworkInfo() {
	fmt.Println("\n--- IP Configuration ---")
	runCommand("powershell", "Get-NetIPConfiguration")

	fmt.Println("\n--- DNS Servers ---")
	runCommand("powershell", "Get-DnsClientServerAddress")

	fmt.Println("\n--- Route Table ---")
	runCommand("powershell", "Get-NetRoute")

	fmt.Println("\n--- Firewall Rules ---")
	runCommand("powershell", "Get-NetFirewallRule | Select-Object Name,Enabled,Direction,Action")

	fmt.Println("\n--- Restart Network Adapter ---")
	runCommand("powershell", "Restart-NetAdapter -Name (Get-NetAdapter | Select-Object -First 1 -ExpandProperty Name)")

	fmt.Println("\n--- Interface Details ---")
	runCommand("powershell", "Get-NetAdapter | Format-List")
}

func getHealthMetrics() {
	fmt.Println("\n--- CPU Usage ---")
	c, _ := cpu.Percent(0, false)
	fmt.Printf("CPU Usage: %.2f%%\n", c[0])

	fmt.Println("\n--- RAM Usage ---")
	v, _ := mem.VirtualMemory()
	fmt.Printf("Used: %v MB (%.2f%%)\n", v.Used/1024/1024, v.UsedPercent)

	fmt.Println("\n--- Disk Usage ---")
	d, _ := disk.Usage("C:")
	fmt.Printf("Used: %v GB (%.2f%%)\n", d.Used/1024/1024/1024, d.UsedPercent)

	fmt.Println("\n--- Network I/O ---")
	n, _ := net.IOCounters(false)
	fmt.Printf("Bytes Sent: %v, Bytes Received: %v\n", n[0].BytesSent, n[0].BytesRecv)
}

func ListAndRestartRunningServices() {
	cmd := exec.Command("powershell", "-Command", `
		Get-Service | Where-Object {$_.Status -eq "Running"} | 
		Select-Object Name,CanStop
	`)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error fetching running services:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) <= 1 {
		fmt.Println("No running services found.")
		return
	}

	fmt.Println("Running Services:")
	restartableServices := make([]string, 0)

	for i, line := range lines[1:] { // Skip the header
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			name := fields[0]
			canStop := strings.ToLower(fields[1]) == "true"
			status := "🔒 Not Restartable"
			if canStop {
				status = "✅ Restartable"
				restartableServices = append(restartableServices, name)
			}
			fmt.Printf("%d. %s - %s\n", i+1, name, status)
		}
	}

	if len(restartableServices) == 0 {
		fmt.Println("\nNo restartable services available.")
		return
	}

	fmt.Print("\nDo you want to restart a service? (y/n): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	if input != "y" {
		return
	}

	fmt.Println("\nSelect a restartable service:")
	for i, svc := range restartableServices {
		fmt.Printf("%d. %s\n", i+1, svc)
	}

	fmt.Print("\nEnter the number of the service to restart: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var choice int
	_, err = fmt.Sscanf(input, "%d", &choice)
	if err != nil || choice < 1 || choice > len(restartableServices) {
		fmt.Println("Invalid selection.")
		return
	}

	selectedService := restartableServices[choice-1]
	restartCmd := exec.Command("powershell", "-Command", fmt.Sprintf(`Restart-Service -Name "%s" -Force`, selectedService))
	err = restartCmd.Run()
	if err != nil {
		fmt.Printf("Failed to restart service '%s': %v\n", selectedService, err)
	} else {
		fmt.Printf("Service '%s' restarted successfully.\n", selectedService)
	}
}
