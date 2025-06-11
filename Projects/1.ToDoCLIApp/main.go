package main

import (
	"fmt"
	"os"
	"strconv"
)

const filename = "todo.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo [add|list|done|delete]")
		return
	}

	command := os.Args[1]
	todos, err := LoadTodos(filename)
	if err != nil {
		fmt.Println("Error loading todos:", err)
		return
	}

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("❗ Usage: todo add [task description]")
			return
		}
		task := os.Args[2]
		todos = append(todos, Todo{Task: task})
		fmt.Println("✅ Task added:", task)

	case "list":
		if len(todos) == 0 {
			fmt.Println("📭 No tasks found.")
			return
		}
		fmt.Println("📝 Todo List:")
		for i, todo := range todos {
			status := "❌"
			if todo.Done {
				status = "✅"
			}
			fmt.Printf("%d. %s [%s]\n", i+1, todo.Task, status)
		}

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("❗ Usage: todo done [task number]")
			return
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil || index < 1 || index > len(todos) {
			fmt.Println("⚠️ Invalid task number.")
			return
		}
		todos[index-1].Done = true
		fmt.Println("✅ Marked as done:", todos[index-1].Task)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("❗ Usage: todo delete [task number]")
			return
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil || index < 1 || index > len(todos) {
			fmt.Println("⚠️ Invalid task number.")
			return
		}
		removed := todos[index-1].Task
		todos = append(todos[:index-1], todos[index:]...)
		fmt.Println("🗑️ Task deleted:", removed)

	default:
		fmt.Println("🚫 Unknown command:", command)
	}

	if err := SaveTodos(todos, filename); err != nil {
		fmt.Println("Error saving todos:", err)
	}
}
