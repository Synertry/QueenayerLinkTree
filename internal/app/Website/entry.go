/*
 *           QueenayerLinkTree
 *     Copyright (c) Synertry 2025.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

// Package Website provides functionality for serving the link tree website.
// It handles HTTP routing, file serving, and error handling for the web server.
package Website

import (
	"embed"
	"fmt"
	"github.com/Synertry/QueenayerLinkTree/Build"
	"github.com/Synertry/QueenayerLinkTree/internal/pkg/Logger"
	"io/fs"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// dist is a filesystem that holds the web assets to be served.
// It's initialized in the Serve function with a subset of the embedded filesystem.
var dist fs.FS

// Serve is the entry point for the website server.
// It initializes the server, sets up logging and error handling, and starts
// the HTTP server on the specified port using the provided embedded filesystem.
//
// Parameters:
//   - port: The port number on which the server will listen
//   - embed: A pointer to an embedded filesystem containing the web assets
//
// Returns:
//   - An error if the server fails to start or if there's an issue with the embedded filesystem
func Serve(port int, embed *embed.FS) (err error) {
	// -------------------------------------------------- INIT --------------------------------------------------
	// Track execution time of the function
	defer Logger.TimeTrack(time.Now(), "Processing "+Build.App)

	// Handle panics by creating a crash log
	defer func() {
		if rec := recover(); rec != nil {
			if os.WriteFile(Build.App+"-crash.log", []byte(fmt.Sprintf("%v", rec)), 0666) != nil {
				log.Fatalf("Failed to create %s-crash.log", Build.App)
			}
			Logger.Out.Error("Application crashed, check %s-crash.log in same directory.\n", Build.App)
		}
	}()

	// Write logs to file when the function returns
	defer func() {
		if errL := os.WriteFile(Build.App+".log", Logger.LogBuffer.Bytes(), 0664); errL != nil {
			log.Panic("Error creating main log ", Build.App, ".log: ", errL.Error())
		}
	}()

	Logger.Out.Info("Starting " + Build.App)

	// -------------------------------------------------- MAIN --------------------------------------------------

	// Extract the web/dist directory from the embedded filesystem
	dist, err = fs.Sub(embed, "web/dist")
	if err != nil {
		Logger.Out.Error("Failed to move into embedded web/dist directory", slog.String("error", err.Error()))
		return
	}

	// Start the HTTP server with the configured routes
	Logger.Out.Info("Starting website server", slog.Int("port", port))
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), routes())
	if err != nil {
		Logger.Out.Error("Failed to start website server", slog.String("error", err.Error()))
	}

	// -------------------------------------------------- FINISH --------------------------------------------------
	Logger.Out.Info("Finished " + Build.App)
	return
}
