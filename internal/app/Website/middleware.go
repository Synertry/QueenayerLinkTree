/*
 *           QueenayerLinkTree
 *     Copyright (c) Synertry 2025.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package Website

import (
	"github.com/Synertry/QueenayerLinkTree/internal/pkg/Logger"
	"log/slog"
	"net/http"
	"time"
)

// logRequest is a middleware that logs HTTP requests with details including
// method, URI, status code, and response time.
//
// Parameters:
//   - next: The HTTP handler to wrap
//
// Returns:
//   - An HTTP handler function that logs the request and calls the next handler
func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Record start time
		start := time.Now()

		// Create a custom response writer to capture the status code
		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK, // Default to 200 OK
		}

		// Process the request with the next handler
		next.ServeHTTP(w, r)

		// Calculate duration
		duration := time.Since(start)

		// Log the request details
		Logger.Out.Info("Request",
			slog.String("method", r.Method),
			slog.String("uri", r.URL.RequestURI()),
			slog.String("remote_addr", r.RemoteAddr),
			slog.String("user_agent", r.UserAgent()),
			slog.Int("status", rw.statusCode),
			slog.String("duration", duration.String()),
		)
	})
}
