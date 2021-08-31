drop table favorites cascade;
create table favorites (
  movie_id integer not null,
  user_id integer not null,
  constraint favorites_PKC primary key (movie_id,user_id)
) ;

drop table api_key cascade;
create table api_key (
  api_key varchar(200) not null,
  user_id integer not null,
  expired_on time with time zone,
  constraint api_key_PKC primary key (api_key)
) ;

drop table users cascade;
create table users (
  id integer not null,
  name varchar(100) not null,
  email varchar(100),
  is_active BOOLEAN default true not null,
  constraint users_PKC primary key (id)
) ;

drop table actor_x_movie cascade;
create table actor_x_movie (
  movie_id integer not null,
  actor_id integer not null,
  character_name varchar(100) not null,
  constraint actor_x_movie_PKC primary key (movie_id,actor_id)
) ;

drop table actors cascade;
create table actors (
  id integer not null,
  name varchar(100),
  birth_day time with time zone,
  nationality varchar(100),
  constraint actors_PKC primary key (id)
) ;

drop table movies_x_genres cascade;
create table movies_x_genres (
  movie_id integer not null,
  genre_id integer not null,
  constraint movies_x_genres_PKC primary key (movie_id,genre_id)
) ;

drop table genres cascade;
create table genres (
  id integer not null,
  genres varchar(50) not null,
  constraint genres_PKC primary key (id)
) ;

drop table movies cascade;
create table movies (
  id integer not null,
  name varchar(200),
  description text,
  writer varchar(200),
  release_date time with time zone,
  cover_img text,
  constraint movies_PKC primary key (id)
) ;