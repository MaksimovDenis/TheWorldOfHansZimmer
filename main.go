package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// Background images
var imageURLs []string

// Soundtracks
var interstellarSoundrackUrls []string
var batmanSoundrackUrls []string
var duneSoundrackUrls []string
var inceptionSoundrackUrls []string
var piratesOfTheCaribbeanSoundrackUrls []string

func interstellar(w http.ResponseWriter, r *http.Request) {
	imageInterstellar := imageURLs[3]
	tmpl, _ := template.ParseFiles("templates/header.html", "templates/interstellar.html", "templates/player.html")
	tmpl.ExecuteTemplate(w, "interstellar", imageInterstellar)
}

func batman(w http.ResponseWriter, r *http.Request) {
	imageBatman := imageURLs[0]
	tmpl, _ := template.ParseFiles("templates/header.html", "templates/batman.html", "templates/player.html")
	tmpl.ExecuteTemplate(w, "batman", imageBatman)
}

func dune(w http.ResponseWriter, r *http.Request) {
	image := imageURLs[1]
	tmpl, _ := template.ParseFiles("templates/header.html", "templates/dune.html", "templates/player.html")
	tmpl.ExecuteTemplate(w, "dune", image)
}

func inception(w http.ResponseWriter, r *http.Request) {
	image := imageURLs[2]
	tmpl, _ := template.ParseFiles("templates/header.html", "templates/inception.html", "templates/player.html")
	tmpl.ExecuteTemplate(w, "inception", image)
}

func piratesOfTheCaribbean(w http.ResponseWriter, r *http.Request) {
	image := imageURLs[4]
	tmpl, _ := template.ParseFiles("templates/header.html", "templates/piratesOfTheCaribbean.html", "templates/player.html")
	tmpl.ExecuteTemplate(w, "piratesOfTheCaribbean", image)
}

func trackInterstellarHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(interstellarSoundrackUrls)
}

func trackBatmanHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(batmanSoundrackUrls)
}

func trackDuneHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(duneSoundrackUrls)
}

func trackInceptionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inceptionSoundrackUrls)
}

func trackPiratesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(piratesOfTheCaribbeanSoundrackUrls)
}

func handlRequest() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", interstellar)
	http.HandleFunc("/batman", batman)
	http.HandleFunc("/dune", dune)
	http.HandleFunc("/inception", inception)
	http.HandleFunc("/piratesOfTheCaribbean", piratesOfTheCaribbean)
	http.HandleFunc("/interstellarSountrack", trackInterstellarHandler)
	http.HandleFunc("/batmanSountrack", trackBatmanHandler)
	http.HandleFunc("/duneSountrack", trackDuneHandler)
	http.HandleFunc("/inceptionSountrack", trackInceptionHandler)
	http.HandleFunc("/piratesSountrack", trackPiratesHandler)
	http.ListenAndServe(":8080", nil)
}

func main() {
	imageURLs, interstellarSoundrackUrls, batmanSoundrackUrls, duneSoundrackUrls, inceptionSoundrackUrls, piratesOfTheCaribbeanSoundrackUrls = YandexStorage()
	fmt.Println(imageURLs)
	fmt.Println()
	fmt.Println(interstellarSoundrackUrls)
	fmt.Println()
	fmt.Println(batmanSoundrackUrls)
	fmt.Println()
	fmt.Println(duneSoundrackUrls)
	fmt.Println()
	fmt.Println(inceptionSoundrackUrls)
	fmt.Println()
	fmt.Println(piratesOfTheCaribbeanSoundrackUrls)
	handlRequest()
}
