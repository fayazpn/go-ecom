package main

import (
	"database/sql"
	"log"

	"github.com/fayazpn/ecom/cmd/api"
	"github.com/fayazpn/ecom/config"
	"github.com/fayazpn/ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	// connect the db
	initStorage(db)

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8081", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

// confirm the database connection
func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected")

}
