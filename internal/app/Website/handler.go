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

// redirectMap is a map of links to redirect to
// mainly used to hold the social media links of the treelink site
var redirectMap = map[string]string{
	"discord":  "https://discord.com/users/402581006185660423",
	"steam":    "https://steamcommunity.com/profiles/76561199206914943",
	"twitter":  "https://x.com/QueenAyer",
	"x":        "https://x.com/QueenAyer",
	"reddit":   "https://www.reddit.com/u/QueenAyer",
	"youtube":  "https://www.youtube.com@queenayer2671",
	"twitch":   "https://www.twitch.tv/queenayer",
	"synertry": "http://synertry.me",
}

// pathServe handles the redirects ion the base path
func pathServe(w http.ResponseWriter, r *http.Request) {

	path := r.PathValue("path")
	if path == "" {
		clientError(w, http.StatusBadRequest)
		return
	}

	// check if the path is available in the FS dist
	// if yes, then serve the static file

	file, err := dist.Open(path)
	if err != nil {
		if url, ok := redirectMap[path]; ok {
			http.Redirect(w, r, url, http.StatusFound)
			return
		}
		// no file found and no redirect
		clientError(w, http.StatusNotFound)
		return
	}

	// check if the file is nil
	if file == nil {
		clientError(w, http.StatusNotFound)
		return
	}

	if err = file.Close(); err != nil {
		serverError(w, r, err)
		return
	}
	http.FileServerFS(dist).ServeHTTP(w, r)
}
