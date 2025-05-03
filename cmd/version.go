/*
 *           QueenayerLinkTree
 *     Copyright (c) Synertry 2025.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package cmd

import (
	"github.com/spf13/cobra"
)

// versionCmd represents the version command which displays version information.
// It prints the version number, revision, and timestamp of the application.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of QueenayerLinkTree",
	Long: `Displays detailed version information about the QueenayerLinkTree application.
This includes the version number, revision hash, and build timestamp.

The command tries to get the version number from Git if available,
otherwise it will print the version number from the compiled binary.

Example:
  QueenayerLinkTree version`,
	Version: rootCmd.Version,
	Run: func(cmd *cobra.Command, args []string) { // https://github.com/spf13/cobra/issues/724#issuecomment-612015666
		// Call the root command with the --version flag to display version information
		root := cmd.Root()
		root.SetArgs([]string{"--version"})
		err := root.Execute()
		if err != nil {
			return
		}
	},
}

// init adds the version command to the root command.
func init() {
	rootCmd.AddCommand(versionCmd)
}
