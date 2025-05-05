package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type PageData struct {
	Videos []string
}

func main() {
	http.Handle("/videos/", http.StripPrefix("/videos/", http.FileServer(http.Dir("./videos"))))
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/Tv", serveTv)
	http.HandleFunc("/Movies", serveMovies)
	http.HandleFunc("/Playlist", servePlaylist)
	http.HandleFunc("/Videos", serveVideos)
	fmt.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {

        if r.Method == "GET" {

                http.ServeFile(w, r, "roryflixWebPages/index.html")
        }

}

func serveTv(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		
                http.ServeFile(w, r, "roryflixWebPages/roryflixTvImproved.html")		
	}

}

func serveMovies(w http.ResponseWriter, r *http.Request) {

        if r.Method == "GET" {

                http.ServeFile(w, r, "roryflixWebPages/roryflixMoviesImproved.html")
        }

}

func servePlaylist(w http.ResponseWriter, r *http.Request) {

        if r.Method == "GET" {

                http.ServeFile(w, r, "roryflixWebPages/roryflixPlaylistSaver.html")
        }

}


func serveVideos(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("./videos")
	if err != nil {
		http.Error(w, "Unable to read videos directory", 500)
		return
	}

	var videoFiles []string
	for _, file := range files {
		if !file.IsDir() {
			ext := filepath.Ext(file.Name())
			if ext == ".mp4" || ext == ".webm" || ext == ".ogg" {
				videoFiles = append(videoFiles, file.Name())
			}
		}
	}

	tmpl := `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>RoryFlix Video List</title>
<meta name="viewport" content="width=device-width, initial-scale=1">
<style>
/* Basic RoryFlix styles for navbar */
body {
  background-color: #141414;
  color: white;
  font-family: Arial, sans-serif;
  margin: 0;
  padding: 0;
}
header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px 50px;
  background: rgba(0, 0, 0, 0.8);
  position: fixed;
  width: 100%;
  z-index: 1000;
}
header .logo {
  font-size: 24px;
  font-weight: bold;
  color: #E50914;
}
header nav {
  display: flex;
  gap: 20px;
}
header nav a {
  color: #fff;
  text-decoration: none;
  font-size: 16px;
  transition: color 0.3s;
}
header nav a:hover {
  color: #E50914;
}
.container {
  padding-top: 100px;
  max-width: 800px;
  margin: auto;
}
ul {
  list-style: none;
  padding: 0;
}
li {
  margin: 15px 0;
}
a.video-link {
  color: #1db954;
  text-decoration: none;
  font-size: 20px;
}
a.video-link:hover {
  text-decoration: underline;
}
</style>
</head>
<body>

<header>
  <div class="logo">RoryFlix</div>
  <nav>

    <a href="/">Home</a>
    <a href="/Tv">TV Shows</a>
    <a href="/Movies">Movies</a>
    <a href="/Playlist">Playlist</a>
    <a href="/Videos">Videos</a>
    </nav>
</header>

<div class="container">
  <h1>Available Videos</h1>
  <ul>
    {{range .Videos}}
    <li><a class="video-link" href="/videos/{{.}}" target="_blank">{{.}}</a></li>
    {{end}}
  </ul>
</div>

</body>
</html>`

	t, err := template.New("page").Parse(tmpl)
	if err != nil {
		http.Error(w, "Template error", 500)
		return
	}

	data := PageData{Videos: videoFiles}
	t.Execute(w, data)
}

