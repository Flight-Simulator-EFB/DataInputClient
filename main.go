package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//Marshall JSON into Objects

	jsonString, _ := os.ReadFile("./waypoints.json")
	data := []Waypoint{}
	json.Unmarshal(jsonString, &data)

	//Open a connection to the database
	db, err := sql.Open("mysql", "root:XXXX@tcp(XXX.XXX.XXX.XXX:3306)/EFB")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	currentTime := time.Now()

	//Make a Query to the database for each item
	for _, waypoint := range data {
		insert, err := db.Query("INSERT INTO waypoints (name, type, latitude, longitude) VALUES (?, ?, ?, ?)",
			waypoint.Name,
			waypoint.NavType,
			waypoint.Coords[0],
			waypoint.Coords[1],
		)

		//Close the connection
		if err != nil {
			fmt.Println(err)
		}
		insert.Close()
	}

	fmt.Println("Total Execution Time:", time.Since(currentTime))
}
