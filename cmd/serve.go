/*
 *           QueenayerLinkTree
 *     Copyright (c) Synertry 2025.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package cmd

import (
	"embed"
	"github.com/Synertry/QueenayerLinkTree/internal/app/Website"
	"github.com/spf13/cobra"
)

// dist holds the embedded filesystem containing web assets.
// It's initialized in the Execute function with the value passed from main.
var dist *embed.FS

// serveCmd represents the serve command which starts the web server.
// It serves the link tree website on the specified port.
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the link tree web server",
	Long: `Start a web server that serves the link tree website.
The server will listen on the specified port (default: 8080)
and serve the embedded web assets.

Example:
  QueenayerLinkTree serve --port 3000`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the port from the flags
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			return err
		}

		// Call the Website.Serve function with the port and embedded filesystem
		return Website.Serve(port, dist)
	},
}

// init adds the serve command to the root command and sets up its flags.
func init() {
	rootCmd.AddCommand(serveCmd)

	// Add a port flag to specify which port the server should listen on
	serveCmd.Flags().IntP("port", "p", 8080, "Port to serve the website on")
}
