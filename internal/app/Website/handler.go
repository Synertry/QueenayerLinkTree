/*
 *           QueenayerLinkTree
 *     Copyright (c) Synertry 2025.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package Website

import "net/http"

// social handles the social media redirects
func social(w http.ResponseWriter, r *http.Request) {

	socialMap := map[string]string{
		"discord":  "https://discord.com/users/402581006185660423",
		"steam":    "https://steamcommunity.com/profiles/76561199206914943",
		"twitter":  "https://x.com/QueenAyer",
		"x":        "https://x.com/QueenAyer",
		"reddit":   "https://www.reddit.com/u/QueenAyer",
		"youtube":  "https://www.youtube.com@queenayer2671",
		"twitch":   "https://www.twitch.tv/queenayer",
		"synertry": "http://synertry.me",
	}

	if url, ok := socialMap[r.PathValue("social")]; ok {
		http.Redirect(w, r, url, http.StatusFound)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
