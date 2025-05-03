/*
 *           QueenayerLinkTree
 *     Copyright (c) Synertry 2025.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

// Package cmd provides the command-line interface for the QueenayerLinkTree application.
// It defines the root command and all subcommands using the Cobra library.
package cmd

import (
	"embed"
	"github.com/Synertry/GoSysUtils/Str"
	"github.com/Synertry/QueenayerLinkTree/Build"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands.
// It's the main entry point for the CLI application.
var rootCmd = &cobra.Command{
	Use:   "QueenayerLinkTree",
	Short: "A customizable link tree web application",
	Long: `QueenayerLinkTree is a web application that serves a customizable link tree,
allowing users to access various resources through a single page.

It provides a simple web server to host the link tree and includes
functionality for self-updating and version information.`,
	Version: Str.Concat(Build.App, " v", Build.Version, "\nRevision: ", Build.Revision, "\nTimestamp: ", Build.Timestamp),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
// It takes an embedded filesystem containing web assets and stores it in the dist variable.
func Execute(embed *embed.FS) {
	dist = embed

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// init is called automatically when the package is initialized.
// It sets up the root command and its flags.
func init() {
	cobra.MousetrapHelpText = "" // Enable start from explorer
	rootCmd.SetVersionTemplate(`{{.Version}}`)

	// Add a toggle flag for demonstration purposes.
	// This flag can be used to enable/disable certain features.
	rootCmd.Flags().BoolP("toggle", "t", false, "Toggle feature flag")
}
