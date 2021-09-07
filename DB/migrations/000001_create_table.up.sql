create table if not exists favorites (
  movie_id integer not null,
  user_id integer not null,
  constraint favorites_PKC primary key (movie_id,user_id)
) ;

create table if not exists api_key (
  api_key varchar(200) not null,
  user_id integer not null,
  expired_on timestamp without time zone,
  constraint api_key_PKC primary key (api_key)
) ;

create table if not exists users (
  id integer not null,
  name varchar(100) not null,
  email varchar(100),
  is_active BOOLEAN default true not null,
  constraint users_PKC primary key (id)
) ;

create table if not exists actor_x_movie (
  movie_id integer not null,
  actor_id integer not null,
  character_name varchar(100) not null,
  constraint actor_x_movie_PKC primary key (movie_id,actor_id)
) ;

create table if not exists actors (
  id integer not null,
  name varchar(100),
  birth_day timestamp with time zone,
  nationality varchar(100),
  constraint actors_PKC primary key (id)
) ;

create table if not exists movies_x_genres (
  movie_id integer not null,
  genre_id integer not null,
  constraint movies_x_genres_PKC primary key (movie_id,genre_id)
) ;

create table if not exists genres (
  id integer not null,
  genres varchar(50) not null,
  constraint genres_PKC primary key (id)
) ;

create table if not exists movies (
  id integer not null,
  name varchar(200),
  description text,
  writer varchar(200),
  release_date timestamp with time zone,
  cover_img text,
  constraint movies_PKC primary key (id)
) ;