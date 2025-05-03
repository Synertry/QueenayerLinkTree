/*
 *           QueenayerLinkTree
 *     Copyright (c) Synertry 2025.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package Self

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	PathExecutable    string
	PathExecutableDir string
)

func init() {
	var err error
	PathExecutable, err = os.Executable()
	if err != nil {
		log.Panicln(fmt.Errorf("failed to get executable path: %w", err))
	}
	PathExecutableDir = filepath.Dir(PathExecutable)
}
