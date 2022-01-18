package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	//Marshall JSON into Objects

	jsonString, _ := os.ReadFile("./waypoints.json")
	data := []Waypoint{}
	x := json.Unmarshal(jsonString, &data)

	fmt.Println(string(jsonString), x)

	//Open a connection to the database
	db, err := sql.Open("mysql", "")
	if err != nil {

	}
	//Make a Query to the database
	db.Exec("INSERT INTO ")
	//Close the connection
	db.Close()
}
