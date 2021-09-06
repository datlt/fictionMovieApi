module fictionMovieApi/app

go 1.17

replace fictionMovieApi/dbaccess => ./dbaccess

require (
	fictionMovieApi/dbaccess v0.0.0-00010101000000-000000000000 // indirect
	github.com/jmoiron/sqlx v1.3.4 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/lib/pq v1.10.2 // indirect
)
