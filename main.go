/*
 *           QueenayerLinkTree
 *     Copyright (c) Synertry 2025.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

// Package main is the entry point for the QueenayerLinkTree application.
// It initializes the command-line interface and embeds the web assets.
package main

import (
	"embed"
	"github.com/Synertry/QueenayerLinkTree/cmd"
)

// dist is an embedded filesystem containing the web assets from the web/dist directory.
// These assets are compiled into the binary and served by the web server.
//
//go:embed all:web/dist
var dist embed.FS

// main is the entry point of the application.
// It calls the Execute function from the cmd package, passing the embedded filesystem.
func main() {
	cmd.Execute(&dist)
}
