package main

import (
	"github.com/gocql/gocql"
)

func main() {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "vhs"
	cluster.Consistency = gocql.LocalOne //TODO: CHANGE TO USE A TOTAL OF 3 NODES

	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	x := menu(map[string]string{
		"1": "Deliveries",
		"2": "Movies",
		"3": "Customers",
		"q": "quit",
	})

	// deliveries
	y := menu(map[string]string{
		"1": "Pending for today",
		"2": "Late deliveries",
		"q": "Go back",
	})

	// movies
	z := menu(map[string]string{
		"1": "List by title",
		"2": "List by genre",
		"3": "List by year",
		"4": "List by promo code",
		"5": "View details",
		"q": "Go back",
	})

	// customers
	w := menu(map[string]string{
		"1": "Check rentals",
		"2": "View customer",
		"q": "Go back",
	})

	// customer
	println(x, y, z, w)

	/*
		if err := session.Query("INSERT INTO movies (movie_id, movie_title, movie_year, movie_genre, movie_promo) VALUES (?, ?, ?, ?, ?)",
			1, "Dummy movie", 1999, "Test", "C").Exec(); err != nil {
			panic(err)
		}

		var title string
		if err := session.Query("SELECT movie_title FROM movies WHERE movie_id = ?", 1).Consistency(gocql.One).Scan(&title); err != nil {
			panic(err)
		}

		println("OK, the title is: ", title)
	*/
}
