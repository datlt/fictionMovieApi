SELECT 
	movies.*,
	array_to_json(array_agg(genres.*)),
	array_to_json(array_agg(actors.*))
FROM movies 
LEFT JOIN movies_x_genres on movies.id = movies_x_genres.movie_id
LEFT JOIN genres on genres.id = movies_x_genres.genre_id
LEFT JOIN actor_x_movie on actor_x_movie.movie_id = movies.id
LEFT JOIN actors on actors.id = actor_x_movie.actor_id
LEFT JOIN favorites on movies.id = favorites.movie_id
{{if or .Id .Search .Movie_name .Actor .Genre -}}
WHERE
	True
	{{if .Id -}}
	AND movies.id = :id
	{{- end}}
	{{if .Search -}}
	AND (
		movies.name LIKE CONCAT('%',CAST(:search AS TEXT),'%') OR
		movies.description LIKE CONCAT('%',CAST(:search AS TEXT),'%') OR
		movies.writer LIKE CONCAT('%',CAST(:search AS TEXT),'%') OR
		genres.genres LIKE CONCAT('%',CAST(:search AS TEXT),'%') OR
		actors.name LIKE CONCAT('%',CAST(:search AS TEXT),'%')
	)
	{{- end}}
	{{if .Movie_name -}}
	AND movies.name like CONCAT('%',CAST(:movie_name AS TEXT),'%')
	{{- end}}
	{{if .Actor -}}
	AND actors.name LIKE CONCAT('%',CAST(:actor AS TEXT),'%')
	{{- end}}
	{{if .Genre -}}
	AND genres.genres LIKE CONCAT('%',CAST(:genre AS TEXT),'%')
	{{- end}}
{{- end}}
GROUP BY (
	movies.id,
	movies.name,
	movies.description,
	movies.writer,
	movies.release_date,
	movies.cover_img
)
ORDER BY count(favorites.user_id)
LIMIT :limit
OFFSET :offset