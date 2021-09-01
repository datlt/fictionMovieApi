package main

import (
	"fictionMovieApi/dbaccess"
	"fmt"
	"log"
	"net/http"
    "encoding/json"
)

func handler(w http.ResponseWriter, r *http.Request) {
    movies, err := dbaccess.AllMovie()
    if err != nil {
        http.Error(w, http.StatusText(500), http.StatusInternalServerError)
        fmt.Println(err)
        return
    }
    json_res, err := json.Marshal(movies)
    fmt.Println("movie Json: ", movies)
    if err != nil {
        http.Error(w, http.StatusText(500), http.StatusInternalServerError)
        fmt.Println(err)
        return
    }
	fmt.Fprint(w, string(json_res))
}

func main() {
	http.HandleFunc("/movies", handler)
    fmt.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
