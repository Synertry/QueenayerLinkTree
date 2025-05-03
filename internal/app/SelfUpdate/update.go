/*
 *           QueenayerLinkTree
 *     Copyright (c) Synertry 2025.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

// Package SelfUpdate provides functionality for updating the application binary.
// It handles the process of replacing the current binary with a new version
// while creating a backup of the original binary.
package SelfUpdate

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Synertry/QueenayerLinkTree/internal/pkg/Self"

	"github.com/Synertry/GoSysUtils/Path"
	"github.com/Synertry/GoSysUtils/Str"
	"github.com/fynelabs/selfupdate"
	"github.com/spf13/cobra"
)

// Update replaces the current application binary with a new version.
// It takes a cobra.Command and a slice of string arguments, where the first argument
// should be the path to the new binary file. The function creates a backup of the
// current binary before replacing it with the new one.
//
// The backup is stored in a "Backup" directory in the same location as the executable,
// with a timestamp appended to the filename to ensure uniqueness.
//
// Returns an error if:
// - No replacement binary is specified
// - The specified file cannot be opened
// - The backup directory cannot be created
// - The update process fails
func Update(_ *cobra.Command, args []string) (err error) {
	// Ensure a replacement binary is specified
	if len(args) == 0 {
		return errors.New("please specify replacement binary")
	}

	// Open the new binary file
	var file *os.File
	file, err = os.Open(args[0])
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", args[0], err)
	}

	// Ensure the file is closed when the function returns
	defer func(c io.Closer, err *error) {
		if cerr := c.Close(); cerr != nil && *err == nil {
			*err = fmt.Errorf("failed to close file %s: %w", args[0], cerr)
		}
	}(file, &err)

	// Create the backup directory if it doesn't exist
	backupDir := filepath.Join(Self.PathExecutableDir, "Backup")
	var dir bool
	dir, err = Path.CheckDir(backupDir)
	if err != nil {
		return fmt.Errorf("failed to check if %s is a directory: %w", backupDir, err)
	}
	if !dir {
		err = os.Mkdir(backupDir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", backupDir, err)
		}
	}

	// Apply the update, creating a backup of the current binary
	err = selfupdate.Apply(file, selfupdate.Options{
		OldSavePath: filepath.Join(backupDir, Str.Concat(
			strings.TrimSuffix(filepath.Base(args[0]), filepath.Ext(args[0])),
			".",
			time.Now().Format("2006-01-02_15-04-05"),
			filepath.Ext(args[0]),
		)),
	})
	if err != nil {
		err = fmt.Errorf("binary update failed: %w", err)
	}
	return
}
