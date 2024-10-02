package main

import (
	"fmt"
	"todoApp/db"
)

func main() {
	todo := db.SelectById(1)
	fmt.Printf("Todo: %+v\n", todo)
}
