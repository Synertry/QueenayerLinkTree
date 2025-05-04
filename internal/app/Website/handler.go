/*
 *           QueenayerLinkTree
 *     Copyright (c) Synertry 2025.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package Website

import (
	"embed"
	"github.com/Synertry/GoSysUtils/Math"
	"github.com/Synertry/QueenayerLinkTree/internal/pkg/Logger"
	"log/slog"
	"net/http"
)

//go:embed httpStatusNotFound.html
var e404 embed.FS

// redirectMap is a map of links to redirect to
// mainly used to hold the social media links of the treelink site
var redirectMap = map[string]string{
	"discord":  "https://discord.com/users/402581006185660423",
	"steam":    "https://steamcommunity.com/profiles/76561199206914943",
	"twitter":  "https://x.com/QueenAyer",
	"x":        "https://x.com/QueenAyer",
	"reddit":   "https://www.reddit.com/u/QueenAyer",
	"youtube":  "https://www.youtube.com/@queenayer2671",
	"twitch":   "https://www.twitch.tv/queenayer",
	"synertry": "https://github.com/Synertry",
	//"profilepic": "https://cdn.discordapp.com/avatars/402581006185660423/4f3a0b1c2d5e7f8a9c4b8e6f8e6f8e6f.png", Discord avatar access does expire
	"profilepic-link": "https://mahoaku-anime.com/en/",
	"avatardeco":      "https://cdn.discordapp.com/avatar-decoration-presets/a_c86b11a49bb8057ce9c974a6f7ad658a.png?passthrough=true",
	"avatardeco-link": "https://discord.com/shop#itemSkuId=1333866045236314327",
	"conillo":         "",
	"conillo-link":    "https://conillo.tumblr.com",
}

var conilloURLs = []string{
	"https://media1.giphy.com/media/v1.Y2lkPTc5MGI3NjExN3RqdWFqNGU0bDRyb3R4NnM0ZWF5dGpobDFsYWFxbG9ndDVqYms2NSZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9cw/1NRau3x08osndXq1Gy/giphy.gif",
	"https://media0.giphy.com/media/v1.Y2lkPTc5MGI3NjExMzFrcHEyZW1waTdjdGxoN3Z4amVvYzg5c2drcmplY2s4bWFteDZrMiZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9cw/fCUBF6W8yAwTxkpqwe/giphy.gif",
	"https://media4.giphy.com/media/v1.Y2lkPTc5MGI3NjExbGFiY2J2MXBjbzdoNnJqZGYyM2djZXNqNW9haTF2NWQ2d2c0MXM4NyZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9cw/5eFud4suIAFRlOiumT/giphy.gif",
	"https://media1.giphy.com/media/v1.Y2lkPTc5MGI3NjExNGtoNjJ4angzbGcwZGlybzRrYXV3ODVyaHNrYmJrNjRuZG5icG4wbyZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9cw/55vOzVsWy21aetiKXL/giphy.gif",
	"https://media2.giphy.com/media/v1.Y2lkPTc5MGI3NjExMmxhYWlyZHplM2Q3amlobjg0Ymd6cnU1anJiOTQ1OWdzcmFpbjN6aiZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9cw/bcJUrhJZDZURtbDYxh/giphy.gif",
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
		if url, ok := redirectMapGet(path); ok {
			http.Redirect(w, r, url, http.StatusFound)
			return
		}
		// no file found and no redirect
		returnLostVisit(w, r)
		return
	}

	// check if the file is nil
	if file == nil {
		returnLostVisit(w, r)
		return
	}

	if err = file.Close(); err != nil {
		serverError(w, r, err)
		return
	}
	http.FileServerFS(dist).ServeHTTP(w, r)
}

// getConilloURL returns a random entry from conilloURLs
func getConilloURL() string {
	num := Math.GetRand().Int() % len(conilloURLs)
	return conilloURLs[num]
}

// redirectMapGet is a helper function to retrieve the URL from the redirect map
// this way we can inject a random conillo URL into the map before serving
// also a bit more error handling
func redirectMapGet(path string) (url string, ok bool) {
	redirectMap["conillo"] = getConilloURL()
	url, ok = redirectMap[path]
	return
}

// returnLostVisit show that the pages does not exist, but redirects to the main page after 5 seconds
func returnLostVisit(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)

	// read the embedded HTML file
	html, err := e404.ReadFile("httpStatusNotFound.html")
	if err != nil {
		serverError(w, nil, err)
		Logger.Out.Error("Failed to read httpStatusNotFound.html", slog.String("error", err.Error()))
		return
	}
	_, err = w.Write(html)
	if err != nil {
		serverError(w, nil, err)
		Logger.Out.Error("Failed to write httpStatusNotFound.html", slog.String("error", err.Error()))
		return
	}
}
