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
	"synertry": "https://github.com/Synertry",
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

// returnLostVisit show that the pages does not exist, but redirects to the main page after 5 seconds
func returnLostVisit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)

	html := `<!--
  ~           QueenayerLinkTree
  ~     Copyright (c) Synertry 2025.
  ~ Distributed under the Boost Software License, Version 1.0.
  ~     (See accompanying file LICENSE or copy at
  ~           https://www.boost.org/LICENSE_1_0.txt)
  -->

<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
<!--  <meta http-equiv="refresh" content="5;url=/">-->
  <title>Page Not Found</title>
  <style>
    body {
      font-family: Arial, sans-serif;
      text-align: center;
      padding: 50px;
      background-color: #000000;
    }
    .container {
      max-width: 600px;
      margin: 0 auto;
      background-color: #FFDEDE;
      padding: 30px;
      border-radius: 10px;
      box-shadow: 0 0 10px rgba(0,0,0,0.1);
    }
    h1 {
      color: #FF0B55;
    }
    p {
      margin: 20px 0;
      font-size: 18px;
    }
    a {
      color: #CF0F47;
      text-decoration: none;
    }
    a:hover {
      text-decoration: underline;
    }
  </style>
</head>
<body>
<div class="container">
  <h1>404 - Page Not Found</h1>
  <p>The page you are looking for does not exist.</p>
  <p>You will be redirected to the home page in <span id="counter">5</span> seconds.</p>
  <p>If you are not redirected automatically, <a href="/">click here</a>.</p>
</div>
</body>

<script>
    // Redirect to the home page after 5 seconds
    let timeLeft = 5;
    const counterElement = document.getElementById('counter');

    const countdown = setInterval(function () {
      timeLeft--;
      counterElement.textContent = timeLeft.toString();

      if (timeLeft <= 0) {
        clearInterval(countdown);
        window.location.href = '/';
      }
    }, 1000);
</script>
</html>`

	_, _ = w.Write([]byte(html))
}
