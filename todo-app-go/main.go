package main

import (
	"fmt"
	"todoApp/internal/db"
	"todoApp/internal/json"
)

func main() {
	todos := db.SelectAll()

	for _, t := range todos {
		fmt.Printf("Todo: %+v\n", t)
	}

	fmt.Println("Serialize Todos data to json data")

	json.SerializeTodos(todos)
}
