# 📝 Todo CLI App in Go

A simple command-line todo list application built with Go.  
Manage your tasks directly from the terminal: add, list, complete, and delete todos.

---

## 🚀 Features

- ✅ Add new tasks
- 📋 List all tasks with their status (done or not)
- ☑️ Mark tasks as done
- 🗑️ Delete tasks
- 💾 Persist data in a local JSON file (`todo.json`)
- 🧱 Written in clean, idiomatic Go

---

## 📦 Requirements

- Go 1.16 or later

---

## 🛠 Installation

## Build the app:

go build -o todo

If get error:
- go mod init todo-cli
- go mod tidy

Add a new task

./todo add "Read Go documentation"

List all tasks

./todo list

Mark a task as done 

./todo done [task number]
# Example:
./todo done 1


Delete a task

./todo delete [task number]
# Example:
./todo delete 2

## Project Structure
todo-cli/
├── main.go        # Command-line interface logic
├── todo.go        # Task model and persistence functions
├── todo.json      # Auto-created storage for tasks
└── README.md

## 📌 Example

$ ./todo add "Buy groceries"
✅ Task added: Buy groceries

$ ./todo list
📝 Todo List:
1. Buy groceries [❌]

$ ./todo done 1
✅ Marked as done: Buy groceries

$ ./todo list
📝 Todo List:
1. Buy groceries [✅]

$ ./todo delete 1
🗑️ Task deleted: Buy groceries
