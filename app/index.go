package main

import (
	"fictionMovieApi/dbaccess"
	"fmt"
	"log"
	"net/http"
    "encoding/json"
    "github.com/julienschmidt/httprouter"
    "strconv"
    "errors"
)

func Auth(w http.ResponseWriter, r *http.Request)(dbaccess.User, error){
    //TODO save api_key in Redis
    api_key, ok := r.URL.Query()["api_key"]
    if(!ok || len(api_key) < 1) {
        http.Error(w, http.StatusText(400), http.StatusBadRequest)
        log.Println("dont have api key")
        return dbaccess.User{}, errors.New("empty api_key")
    }

    users, err := dbaccess.GetUser(api_key[0])
    if err != nil {
        http.Error(w, http.StatusText(500), http.StatusInternalServerError)
        log.Println(err)
        return dbaccess.User{}, err
    }

    if len(users) != 1 {
        http.Error(w, http.StatusText(400), http.StatusBadRequest)
        log.Println("Api key not found")
        return dbaccess.User{}, errors.New("api_key not valid")
    }

    return users[0], nil
}

func queryAndReturnMovie(queryParam dbaccess.MovieQueryParam, w http.ResponseWriter){
    movies, err := dbaccess.QueryMovies(queryParam)
    if err != nil {
        http.Error(w, http.StatusText(500), http.StatusInternalServerError)
        log.Println(err)
        return
    }
    json_res, err := json.Marshal(movies)
    if err != nil {
        http.Error(w, http.StatusText(500), http.StatusInternalServerError)
        log.Println(err)
        return
    }
	fmt.Fprint(w, string(json_res))
}

func SearchMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    _, err := Auth(w, r)
    if(err != nil){
        return
    }
    var queryParam = dbaccess.MovieQueryParam{}
    getParams := r.URL.Query()

    search, ok := getParams["search"]
    if ok && len(search) >= 1 {
        queryParam.Search = search[0]
    } 

    movie_name, ok := getParams["movie_name"]
    if ok && len(movie_name) >= 1 {
        queryParam.Movie_name = movie_name[0]
    } 

    actor, ok := getParams["actor"]
    if ok && len(actor) >= 1 {
        queryParam.Actor = actor[0]
    } 

    genre, ok := getParams["genre"]
    if ok && len(genre) >= 1 {
        queryParam.Genre = genre[0]
    } 

    page, ok := getParams["page"]
    if ok && len(page) >= 1 {
        page_num, err := strconv.Atoi(page[0])
        if err == nil {
            queryParam.Page = page_num
        }
    } 

    queryAndReturnMovie(queryParam, w)
}

func GetMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
    _, err := Auth(w, r)
    if(err != nil){
        return
    }
    id_string := ps.ByName("id")
    id, err := strconv.Atoi(id_string)
    if err != nil {
        http.Error(w, http.StatusText(400), http.StatusBadRequest)
        log.Println("Bad Request")
        return
    }
    var queryParam = dbaccess.MovieQueryParam{}
    queryParam.Id = id

    queryAndReturnMovie(queryParam, w)
}

func GetFavorites(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
    user, err := Auth(w, r)
    if(err != nil){
        return
    }
    var queryParam = dbaccess.MovieQueryParam{}
    queryParam.User_id = user.ID

    getParams := r.URL.Query()
    page, ok := getParams["page"]
    if ok && len(page) >= 1 {
        page_num, err := strconv.Atoi(page[0])
        if err == nil {
            queryParam.Page = page_num
        }
    } 
    queryAndReturnMovie(queryParam, w)
}

func AddFavorite(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
    user, err := Auth(w, r)
    if(err != nil){
        return
    }

    id_string := ps.ByName("id")
    id, err := strconv.Atoi(id_string)
    if err != nil {
        http.Error(w, http.StatusText(400), http.StatusBadRequest)
        log.Println("Bad Request")
        return
    }

    err = dbaccess.AddFavorite(user.ID, id) 
    if err != nil {
        http.Error(w, http.StatusText(500), http.StatusInternalServerError)
        return
    } else {
        fmt.Fprint(w, "Favorite Added")
    }

}

func main() {
    router := httprouter.New()
	router.GET("/movies", SearchMovies)
    router.GET("/movies/:id", GetMovie)
    router.GET("/favorites", GetFavorites)
    router.POST("/favorite/:id", AddFavorite)

    log.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
