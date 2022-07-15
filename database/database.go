package database

import (
	"fmt"
	"log"

	"database/sql"

	config "github.com/Questee29/testBookService/configs"
	_ "github.com/go-sql-driver/mysql"
)

func New() (*sql.DB, error) {

	config, err := config.LoadConfig("app", ".")
	if err != nil {
		return nil, err
	}
	//connectionInfo
	psqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DBName)
	log.Println(psqlInfo)
	//open connection
	db, err := sql.Open(config.Database.DbDriver, psqlInfo)
	if err != nil {
		log.Println("OPEN")
		return nil, err
	}

	//make migrations

	// if err := Migrate(db); err != nil {
	// 	return nil, nil
	// }

	//ping

	if err := db.Ping(); err != nil {
		log.Println("PING")
		return nil, err
	}
	log.Println("Successfully connected to database!")
	return db, nil
}
