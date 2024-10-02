package main

import "todoApp/db"

func main() {
	name := "David"
	db.Insert(name)
}
