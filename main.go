package main

import (
	"database/sql"
	"fmt"
	"todo_app/app/controllers"
	"todo_app/app/models"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func main() {

	fmt.Println(models.Db)

	controllers.StartMainServer()

}
