package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

var imageURLs []string

var trackURLs = []string{
	"https://raw.githubusercontent.com/himalayasingh/music-player-1/master/music/2.mp3",
	"https://raw.githubusercontent.com/himalayasingh/music-player-1/master/music/1.mp3",
	"https://raw.githubusercontent.com/himalayasingh/music-player-1/master/music/3.mp3",
	"https://raw.githubusercontent.com/himalayasingh/music-player-1/master/music/4.mp3",
	"https://raw.githubusercontent.com/himalayasingh/music-player-1/master/music/5.mp3",
}

func interstellar(w http.ResponseWriter, r *http.Request) {
	imageInterstellar := imageURLs[4]
	tmpl, _ := template.ParseFiles("templates/header.html", "templates/interstellar.html", "templates/player.html")
	tmpl.ExecuteTemplate(w, "interstellar", imageInterstellar)
}

func batman(w http.ResponseWriter, r *http.Request) {
	imageBatman := imageURLs[1]
	tmpl, _ := template.ParseFiles("templates/header.html", "templates/batman.html", "templates/player.html")
	tmpl.ExecuteTemplate(w, "batman", imageBatman)
}

func dune(w http.ResponseWriter, r *http.Request) {
	image := imageURLs[2]
	tmpl, _ := template.ParseFiles("templates/header.html", "templates/dune.html", "templates/player.html")
	tmpl.ExecuteTemplate(w, "dune", image)
}

func inception(w http.ResponseWriter, r *http.Request) {
	image := imageURLs[3]
	tmpl, _ := template.ParseFiles("templates/header.html", "templates/inception.html", "templates/player.html")
	tmpl.ExecuteTemplate(w, "inception", image)
}

func piratesOfTheCaribbean(w http.ResponseWriter, r *http.Request) {
	image := imageURLs[5]
	tmpl, _ := template.ParseFiles("templates/header.html", "templates/piratesOfTheCaribbean.html", "templates/player.html")
	tmpl.ExecuteTemplate(w, "piratesOfTheCaribbean", image)
}

func theLastSamurai(w http.ResponseWriter, r *http.Request) {
	image := imageURLs[6]
	tmpl, _ := template.ParseFiles("templates/header.html", "templates/theLastSamurai.html", "templates/player.html")
	tmpl.ExecuteTemplate(w, "theLastSamurai", image)
}

func trackURLSHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trackURLs)
}

func handlRequest() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", interstellar)
	http.HandleFunc("/batman", batman)
	http.HandleFunc("/dune", dune)
	http.HandleFunc("/inception", inception)
	http.HandleFunc("/piratesOfTheCaribbean", piratesOfTheCaribbean)
	http.HandleFunc("/theLastSamurai", theLastSamurai)
	http.HandleFunc("/track-urls", trackURLSHandler)
	http.ListenAndServe(":8080", nil)
}

func main() {
	imageURLs = YandexStorage()
	fmt.Println(imageURLs)
	handlRequest()
}
