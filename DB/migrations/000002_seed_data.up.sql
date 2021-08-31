insert into movies (id, name, description, writer, release_date, cover_img)
values 
(1, 'movie 1', 'description for movie 1', 'writer 1', now(), 'http://example.com'),
(2, 'movie 2', 'description for movie 2', 'writer 2', now(), 'http://example.com'),
(3, 'movie 3', 'description for movie 3', 'writer 3', now(), 'http://example.com'),
(4, 'movie 4', 'description for movie 4', 'writer 4', now(), 'http://example.com'),
(5, 'movie 5', 'description for movie 5', 'writer 5', now(), 'http://example.com'),
(6, 'movie 6', 'description for movie 6', 'writer 6', now(), 'http://example.com'),
(7, 'movie 7', 'description for movie 7', 'writer 7', now(), 'http://example.com'),
(8, 'movie 8', 'description for movie 8', 'writer 8', now(), 'http://example.com'),
(9, 'movie 9', 'description for movie 9', 'writer 9', now(), 'http://example.com'),
(10, 'movie 10', 'description for movie 10', 'writer 10', now(), 'http://example.com');

insert into actors (id, "name", birth_day, nationality)
values 
(1, 'actor 1', TO_TIMESTAMP('2000-01-01', 'YYYY-MM-DD HH:MI:SS'), 'Japan'),
(2, 'actor 2', TO_TIMESTAMP('2000-01-01', 'YYYY-MM-DD HH:MI:SS'), 'Japan'),
(3, 'actor 3', TO_TIMESTAMP('2000-01-01', 'YYYY-MM-DD HH:MI:SS'), 'Japan'),
(4, 'actor 4', TO_TIMESTAMP('2000-01-01', 'YYYY-MM-DD HH:MI:SS'), 'Japan'),
(5, 'actor 5', TO_TIMESTAMP('2000-01-01', 'YYYY-MM-DD HH:MI:SS'), 'Japan');

insert into actor_x_movie (movie_id, actor_id, character_name)
values
(1,1, 'character_name_1_1'),
(2,1, 'character_name_2_1'),
(1,2, 'character_name_1_2'),
(2,3, 'character_name_2_3'),
(3,1, 'character_name_3_1'),
(4,1, 'character_name_4_1'),
(5,1, 'character_name_5_1'),
(6,1, 'character_name_6_1'),
(7,1, 'character_name_7_1'),
(8,1, 'character_name_8_1'),
(9,1, 'character_name_9_1'),
(10,1, 'character_name_10_1');

insert into genres (id, genres)
values
(1, 'genres 1'),
(2, 'genres 2'),
(3, 'genres 3');

insert into movies_x_genres (movie_id, genre_id)
values
(1, 1),
(1, 2),
(1, 3),
(2, 1),
(3, 2),
(4, 3),
(5, 1),
(6, 2),
(7, 3),
(8, 1),
(9, 2),
(10, 3);

insert into users(id,"name",email,is_active)
values
(1,'user 1', 'user1@example.com', true),
(2,'user 2', 'user2@example.com', true);