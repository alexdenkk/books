package main

import (
	"alexdenkk/books/app"
	"alexdenkk/books/model"
	"alexdenkk/books/pkg/db"
	"log"
	"flag"
)

var (
	addr   string
	jwt    string
	dbName string
	dbPort string
	dbUser string
	dbPswd string
)

func main() {
	parseFlags()

	// connecting to db
	appDB, err := db.Connect(
		dbName,
		dbPort,
		dbUser,
		dbPswd,
	)

	model.Migrate(appDB)

	if err != nil {
		log.Fatal(err)
	}

	// create app
	booksApp := app.New(appDB, []byte(jwt), addr)

	// run
	log.Fatal(booksApp.Run())
}

func parseFlags() {
	flag.StringVar(&addr, "address", ":8080", "address and(or) port for app")
	flag.StringVar(&jwt, "jwt", "fuck", "jwt sign key for generating user tokens")
	flag.StringVar(&dbName, "db-name", "db", "database name")
	flag.StringVar(&dbPort, "db-port", "5432", "database port")
	flag.StringVar(&dbUser, "db-user", "alexdenkk", "database user")
	flag.StringVar(&dbPswd, "db-password", "", "database password")

	flag.Parse()
}
