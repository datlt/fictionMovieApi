package dbaccess

type Movie struct {
	ID          int
	Name        string
	Description string
	Writer      string
	ReleaseDate string
	CoverImg    string
	Genres      []Genre
	Actors      []Actor
}

type Genre struct {
	Genres string
	ID int
}

type Actor struct {
	ID int
	Name string
	BirthDay string
	Nationality string
}