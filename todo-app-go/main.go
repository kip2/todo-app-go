package main

import (
	"fmt"
	"todoApp/internal/db"
)

func main() {
	todos := db.SelectAll()

	for _, t := range todos {
		fmt.Printf("Todo: %+v\n", t)
	}
}
