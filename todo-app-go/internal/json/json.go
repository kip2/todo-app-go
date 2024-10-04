package json

import (
	"encoding/json"
	"fmt"
	"os"
	"todoApp/internal/db"
)

func SerializeTodos(todos []db.Todo) {
	filename := "test.json"
	err := SaveToJson(filename, todos)

	if err != nil {
		fmt.Println("Error saving to JSON:", err)
	} else {
		fmt.Println("Successfully saved to users.json")
	}
}

func SaveToJson(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}
