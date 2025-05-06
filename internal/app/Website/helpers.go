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
	"runtime/debug"
)

// serverError handles internal server errors by logging detailed error information
// and returning a generic 500 Internal Server Error response to the client.
// This function logs the error message, HTTP method, URI, and a stack trace to
// help with debugging while keeping the details hidden from the client.
//
// Parameters:
//   - w: The HTTP response writer to send the error response
//   - r: The HTTP request that triggered the error
//   - err: The error that occurred
func serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		// Get the stack trace to help with debugging
		trace = string(debug.Stack())
	)

	// Log the error with detailed information for debugging
	Logger.Out.Error(
		err.Error(),
		slog.String("method", method),
		slog.String("uri", uri),
		slog.String("trace", trace),
	)

	// Send a generic 500 Internal Server Error response to the client
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError sends an appropriate HTTP error response to the client with the
// specified status code and its corresponding standard description.
// This function is used for client-side errors such as 400 Bad Request,
// 404 Not Found, etc., where the error is due to the client's request.
//
// Parameters:
//   - w: The HTTP response writer to send the error response
//   - status: The HTTP status code to send (e.g., http.StatusBadRequest, http.StatusNotFound)
func clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// responseWriter is a custom http.ResponseWriter that captures the status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code and passes it to the underlying ResponseWriter
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
