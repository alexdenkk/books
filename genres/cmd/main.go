package main

import (
	"alexdenkk/books/genres/app"
	"alexdenkk/books/genres/model"
	"alexdenkk/books/genres/pkg/db"
	"flag"
	"log"
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

	if err != nil {
		log.Fatal(err)
	}

	log.Println("db connected")

	// migration
	model.Migrate(appDB)
	log.Println("migration complete")

	// create app
	genresApp := app.New(appDB, []byte(jwt), addr)
	log.Println("app initialized")

	// run
	log.Fatal(genresApp.Run())
}

func parseFlags() {
	flag.StringVar(&addr, "address", ":8004", "address and(or) port for app")
	flag.StringVar(&jwt, "jwt", "fuck", "jwt sign key for user tokens")
	flag.StringVar(&dbName, "db-name", "db", "database name")
	flag.StringVar(&dbPort, "db-port", "5432", "database port")
	flag.StringVar(&dbUser, "db-user", "alexdenkk", "database user")
	flag.StringVar(&dbPswd, "db-password", "", "database password")
	flag.Parse()
}
