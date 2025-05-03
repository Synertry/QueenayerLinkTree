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
	"path/filepath"
)

// neuteredFileSystem is a custom http.FileSystem implementation that
// prevents directory listing and only serves files with an index.html
// file in the directory.
// This enhances security by preventing users from browsing the directory structure.
// Source: https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings
type neuteredFileSystem struct {
	fs http.FileSystem
}

// Open implements the http.FileSystem interface for neuteredFileSystem.
// It opens the specified path from the underlying filesystem, but adds
// security by preventing directory listing. If the path is a directory,
// it checks if an index.html file exists in that directory. If not, it
// returns an error, effectively disabling directory browsing.
//
// Parameters:
//   - path: The file path to open
//
// Returns:
//   - An http.File if the file exists or if the directory contains an index.html file
//   - An error if the file doesn't exist or if a directory doesn't contain an index.html file
func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	// Try to open the file from the underlying filesystem
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	// Get file information
	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	// If the path is a directory, check for index.html
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			// Close the directory file handle to prevent resource leaks
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			// Return the original error (index.html not found)
			return nil, err
		}
	}

	// Return the file handle
	return f, nil
}
