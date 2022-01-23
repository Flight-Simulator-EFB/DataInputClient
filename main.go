package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//Marshall JSON into Objects

	jsonString, _ := os.ReadFile("./waypoints.json")
	data := []Waypoint{}
	x := json.Unmarshal(jsonString, &data)

	fmt.Println(string(jsonString), x)

	//Open a connection to the database
	db, err := sql.Open("mysql", "root:@tcp(:3306)/EFB")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//Make a Query to the database
	db.Exec("INSERT INTO ")
	//Close the connection
	db.Close()
}
