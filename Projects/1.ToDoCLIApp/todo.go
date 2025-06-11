package main

import (
	"encoding/json"
	"os"
)

type Todo struct {
	Task string
	Done bool
}

func SaveTodos(todos []Todo, filename string) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func LoadTodos(filename string) ([]Todo, error) {
	var todos []Todo
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return todos, nil // Return empty list if file doesn't exist
		}
		return nil, err
	}
	err = json.Unmarshal(data, &todos)
	return todos, err
}
