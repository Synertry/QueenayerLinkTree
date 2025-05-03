/*
 *           QueenayerLinkTree
 *     Copyright (c) Synertry 2025.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package Website

import (
	"net/http"
)

// routes defines and configures all HTTP routes for the application.
// It creates a new ServeMux and sets up handlers for serving static files
// and the main page. The function uses a neuteredFileSystem to prevent
// directory listing for security reasons.
//
// Routes defined:
//   - GET /: Serves the main page (index.html)
//   - GET /assets/: Serves static assets (CSS, JS, images, etc.)
//   - GET /assets: Returns 404 Not Found (prevents directory listing)
//
// Returns:
//   - A configured http.ServeMux with all application routes
func routes() *http.ServeMux {
	// Create a new ServeMux for routing HTTP requests
	mux := http.NewServeMux()

	// Create a file server with the neutered filesystem to prevent directory listing
	fileServer := http.FileServer(neuteredFileSystem{http.FS(dist)})

	// Return 404 for direct requests to /assets to prevent directory listing
	mux.Handle("GET /assets", http.NotFoundHandler())

	// Serve static files from the /assets/ directory
	mux.Handle("GET /assets/", fileServer)

	// Serve the main page (index.html) at the root URL
	mux.Handle("GET /", fileServer)

	return mux
}
