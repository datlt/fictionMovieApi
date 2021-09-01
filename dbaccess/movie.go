package dbaccess

type Movie struct {
	ID          int
	Name        string
	Description string
	Writer      string
	ReleaseDate string
	CoverImg    string
}

func AllMovie() ([]Movie, error) {
	db, err := connect()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	movies := make([]Movie, 0)

	rows, err := db.Query("SELECT * FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		movie := Movie{}
		err := rows.Scan(&movie.ID, &movie.Name, &movie.Description, &movie.Writer, &movie.ReleaseDate, &movie.CoverImg)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}
