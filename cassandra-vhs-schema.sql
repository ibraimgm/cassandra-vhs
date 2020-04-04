CREATE KEYSPACE vhs
WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 3};


CREATE TABLE vhs.pending_deliveries (
  quarter_bucket TEXT,
  delivery_date TIMESTAMP,
  rental_id INT,
  movie_id INT,
  movie_name TEXT,
  PRIMARY KEY (quarter_bucket, delivery_date, rental_id, movie_id)
) WITH CLUSTERING ORDER BY (delivery_date DESC, rental_id ASC, movie_id ASC);

CREATE TABLE vhs.rentals (
  customer_id INT,
  rental_date TIMESTAMP,
  rental_id INT,
  total DOUBLE,
  PRIMARY KEY (customer_id, rental_date, rental_id)
) WITH CLUSTERING ORDER BY (rental_date DESC, rental_id ASC);

CREATE TABLE vhs.customer (
  customer_id INT,
  customer_name TEXT,
  PRIMARY KEY (customer_id, customer_name)
) WITH CLUSTERING ORDER BY (customer_name ASC);

CREATE TABLE vhs.movie_by_title (
  movie_title TEXT,
  movie_id INT,
  PRIMARY KEY (movie_title, movie_id)
);

CREATE TABLE vhs.movie_by_year (
  movie_year INT,
  movie_id INT,
  PRIMARY KEY (movie_year, movie_id)
);

CREATE TABLE vhs.movie_by_genre (
  movie_genre TEXT,
  movie_id INT,
  PRIMARY KEY (movie_genre, movie_id)
);

CREATE TABLE vhs.movie_by_promo (
  movie_promo TEXT,
  movie_id INT,
  PRIMARY KEY (movie_promo, movie_id)
);

CREATE TABLE vhs.movies (
  movie_id INT,
  movie_title TEXT,
  movie_genre TEXT,
  movie_year INT,
  movie_promo TEXT,
  PRIMARY KEY (movie_id)
);
