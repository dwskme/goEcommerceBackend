package main

import (
	"database/sql"
	"log"

	"github.com/dwskme/goEcommerceBackend/cmd/api"
	"github.com/dwskme/goEcommerceBackend/config"
	"github.com/dwskme/goEcommerceBackend/db"
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
	if err != nil {
		log.Fatal(err)
	}
  initStorage(db)
	server := api.NewAPIServer(config.Envs.Port, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB){
  err:=db.Ping()
  if err != nil {
    log.Fatal(err)
  }
  log.Println("DB:Successfully connected!")
}
