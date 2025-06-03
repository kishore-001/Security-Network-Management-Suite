package log

import (
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
)

func HandleLog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// Optional: allow user to set number of lines via query param, default to 100
	numLines := 100
	if n := r.URL.Query().Get("lines"); n != "" {
		if parsed, err := strconv.Atoi(n); err == nil && parsed > 0 {
			numLines = parsed
		}
	}

	// Get the last numLines from journalctl
	cmd := exec.Command("journalctl", "-n", fmt.Sprintf("%d", numLines), "--no-pager")
	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, "Error fetching logs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

