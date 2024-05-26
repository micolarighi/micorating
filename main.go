package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

type RatingList struct {
	Title string
	Rating int
}

func rating (w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	rtlist := map[string][]RatingList {
		"RatingList" : {
			{Title : "Boruto Two Blue Vortex", Rating : 9},
			{Title : "One Punch Man", Rating : 9},
			{Title : "Boruto Naruto Next Generation", Rating : 8},
		},
	} 

	tmpl.Execute(w, rtlist)
}

func addRating (w http.ResponseWriter, req *http.Request) {	
	title := req.PostFormValue("title")
	r := req.PostFormValue("rating")
	rating, err := strconv.Atoi(r)
	if err != nil {
		log.Printf("Cannot convert value to int")
	}
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "rating-list-element", RatingList{Title: title, Rating: rating})
}

func main() {
	http.HandleFunc("/rating", rating)
	http.HandleFunc("/add-rating/", addRating)
	port := ":8090"
	log.Printf("Server is listening on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server %s\n", err)
		os.Exit(1)
	}

} 
