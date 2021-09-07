package dbaccess

// import (
// 	"encoding/json"
// 	"text/template"
// 	"bytes"
// 	"os"
// )

func AddFavorite(user_id int, movie_id int) (error) {
	query_string := "INSERT INTO favorites(movie_id, user_id) Values (:movie_id, :user_id)"
	db, err := connect()
	defer db.Close()
	if err != nil {
		return err
	}

	_, err = db.NamedExec(query_string, &Favorite{user_id, movie_id})
	if err != nil{
		return err
	}

	return nil
}
