/*
 *           QueenayerLinkTree
 *     Copyright (c) Synertry 2025.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package cmd

import (
	"fmt"

	"github.com/Synertry/GoSysUtils/Path"
	"github.com/Synertry/QueenayerLinkTree/internal/app/SelfUpdate"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command which updates the application binary.
// It replaces the current binary with a new version provided as an argument.
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates this binary to the latest version",
	Long: `Ensures that you have the latest version of this binary.
If not it will replace the binary with the latest version.

The command requires a path to the new binary file as an argument.
A backup of the current binary will be created before updating.

Example:
  QueenayerLinkTree update path/to/new/binary`,
	Version:      versionCmd.Version,
	Args:         cobra.MatchAll(cobra.ExactArgs(1), argPathValidator),
	RunE:         SelfUpdate.Update,
	SilenceUsage: true,
}

// argPathValidator ensures you have at least one path as an argument and checks it for existence.
// It validates that the path provided as an argument exists before attempting to use it for updating.
func argPathValidator(cmd *cobra.Command, args []string) error {
	// Run the custom validation logic
	for _, arg := range args {
		ok, err := Path.Check(arg)
		if err != nil {
			return err
		} else if !ok {
			// return error that arg path does not exist
			return fmt.Errorf("path %s does not exist", arg)
		}
	}
	return nil
}

// init adds the update command to the root command and sets up its flags.
func init() {
	rootCmd.AddCommand(updateCmd)

	// Add a patch flag to specify whether to use a patch file for updating
	updateCmd.Flags().BoolP("patch", "p", false, "Use patch file to update binary")
}
