package dbaccess

import (
	"encoding/json"
	"text/template"
	"bytes"
	"os"
	"fmt"
)

type MovieQueryParam struct {
	Id          int
	Search      string
	Movie_name  string
	Actor       string
	Genre 		string
	Limit    	int
	Offset      int
	Page      	int
}

func QueryMovies(queryParam MovieQueryParam) ([]Movie, error) {
	if queryParam.Page <= 0 {
		queryParam.Page = 1
	}
	queryParam.Limit = PAGE_SIZE
	queryParam.Offset = PAGE_SIZE*(queryParam.Page -1)
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	query_template, err := template.ParseFiles(wd + "/dbaccess/query/getmovie.go.sql")
	if err != nil {
		return nil, err
	}

	var query_string bytes.Buffer
	if err := query_template.Execute(&query_string, queryParam); err!= nil {
		return nil, err
	}

	db, err := connect()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	
	rows, err := db.NamedQuery(query_string.String(), &queryParam)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	movies := make([]Movie, 0)
	for rows.Next() {
		movie := Movie{}
		genre_json := ""
		actor_json := ""
		err := rows.Scan(&movie.ID, &movie.Name, &movie.Description, &movie.Writer, &movie.ReleaseDate, &movie.CoverImg, &genre_json, &actor_json)
		if err != nil {
			return nil, err
		}
		json.Unmarshal([]byte(genre_json), &movie.Genres)
		json.Unmarshal([]byte(actor_json), &movie.Actors)
		movies = append(movies, movie)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}
