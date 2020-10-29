package main

import (
	"database/sql"
	"fmt"
	"language_tournament/route"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// return c.String(http.StatusOK, "hello, World!")
	// })
	// e.Logger.Fatal(e.Start(":8080"))
	initdb()

	fmt.Println("hello")
	router := route.Init()
	router.Start(":8080")
}

func initdb() {
	connectionString := getConnectionString()
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func getConnectionString() string {
	host := "mysql"
	port := "3306"
	user := "mysql"
	pass := "mysql"
	dbname := "comparision_box"
	protocol := "tcp"

	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s", user, pass, protocol, host, port, dbname)
}
