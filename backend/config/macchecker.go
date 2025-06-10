package config

import (
	generaldb "backend/db/gen/general"
	"net/http"
	"strings"
)

// MacCheckerMiddleware checks if the MAC address is whitelisted
func MacCheckerMiddleware(queries *generaldb.Queries) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mac := r.Header.Get("X-MAC-Address")
			if mac == "" {
				http.Error(w, "MAC address missing in header", http.StatusBadRequest)
				return
			}

			mac = strings.ToUpper(strings.TrimSpace(mac)) // Normalize

			isAllowed, err := queries.IsMacWhitelisted(r.Context(), mac)
			if err != nil {
				http.Error(w, "Error checking MAC access", http.StatusInternalServerError)
				return
			}

			if !isAllowed {
				http.Error(w, "MAC address not whitelisted", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
