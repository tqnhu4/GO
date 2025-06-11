# go-tutorials
# 7-Day Go Language Learning Roadmap for Beginners

Welcome to your accelerated Go (Golang) learning journey!  
This roadmap is meticulously crafted to guide absolute beginners through the fundamentals of Go within seven days, emphasizing **practical application** and **rapid skill acquisition**.

Each day focuses on a distinct set of core Go concepts, complete with illustrative examples and a culminating practical project to solidify your understanding.

---

## 📅 Day 1: Go Fundamentals & Setup

Today is all about getting started with Go — from setting up your development environment to writing your very first Go program and understanding its basic building blocks.

### ✅ Content:

- **Introduction to Go:**
  - What is Go?
  - Key features: Concurrency, performance, simplicity.
  - Why it's gaining popularity.

- **Installation and Environment Setup:**
  - Install Go on **Windows**, **macOS**, or **Linux**.
  - Configure your `GOPATH` and understand its role.
  - Recommended IDEs:
    - [VS Code](https://code.visualstudio.com/) with the Go extension
    - [GoLand](https://www.jetbrains.com/go/)

- **Your First Go Program:**
  - Writing and executing a simple `"Hello, World!"` application.

- **Basic Syntax and Structure:**
  - The role of `package main` and `func main()`.
  - Comments in Go: `//` for single-line, `/* */` for multi-line.

- **Variables and Basic Data Types:**
  - Declaring variables: `var` and `:=`
  - Common types:
    - Integers: `int`, `int32`, `int64`
    - Floats: `float32`, `float64`
    - Booleans: `bool`
    - Strings: `string`

---

### 💻 Example Code:

```go
package main // Declares the package as 'main', making it an executable program

import "fmt" // Imports the "fmt" package for formatted I/O (e.g., printing)

func main() { // The entry point of the program
    // Your first Go program
    fmt.Println("Hello, Go World!")

    // Variable declaration and initialization
    var name string = "Alice" // Explicit type declaration
    age := 30                 // Short variable declaration (type inferred)
    height := 1.75            // float64 by default
    isStudent := false

    // Formatted output
    fmt.Printf("Name: %s, Age: %d, Height: %.2f, Is Student: %t\n", name, age, height, isStudent)
}
```
---

## 📅 Day 2: Operators & Control Flow

Today, you'll delve into performing operations and controlling the execution flow of your Go programs using conditional statements and loops.

### ✅ Content:

- **Arithmetic Operators:** 
  - Addition, subtraction, multiplication, division, modulo.

- **Assignment Operators:** 
  - Simple assignment (=), compound assignments (+=, -=, etc.).

- **Comparison Operators:** Greater than, less than, equal to, not equal to, greater than or equal to, less than or equal to.
- **Logical Operators:** && (AND), || (OR), ! (NOT).
- **Conditional Statements (if, else if, else):** Decision-making based on conditions.
- **Looping Constructs (for):** Go's single, versatile for loop for all iteration needs (equivalent to for, while, and do-while in other languages).

---

### 💻 Example Code:

```go
package main

import "fmt"

func main() {
	// Arithmetic operators example
	a := 15
	b := 4
	fmt.Printf("Sum: %d\n", a+b)
	fmt.Printf("Difference: %d\n", a-b)
	fmt.Printf("Product: %d\n", a*b)
	fmt.Printf("Division: %f\n", float64(a)/float64(b)) // Type conversion for float division
	fmt.Printf("Modulo: %d\n", a%b)

	// Conditional statements example
	score := 82
	if score >= 90 {
		fmt.Println("Excellent!")
	} else if score >= 70 {
		fmt.Println("Good!")
	} else if score >= 50 {
		fmt.Println("Average.")
	} else {
		fmt.Println("Needs improvement.")
	}

	// For loop example (as a traditional for loop)
	fmt.Println("\nNumbers from 0 to 4:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// For loop example (as a while loop)
	count := 0
	fmt.Println("\nCount increments:")
	for count < 3 {
		fmt.Printf("Count: %d\n", count)
		count++
	}
}
```

---

## 📅 Day 3: Composite Data Types (Arrays, Slices, Maps)

Today you'll explore Go's fundamental data structures for grouping and organizing data.

### ✅ Content:

- **Arithmetic Operators:** 
  - Addition, subtraction, multiplication, division, modulo.

- **Arrays:** Fixed-size, ordered collections of elements of the same type.
  - Declaration, initialization, accessing elements.
- **Slices:** Dynamic, flexible views into arrays.
  - Creating slices, appending elements, slicing existing slices.
  - Understanding len() and cap().
- **Maps:** Unordered collections of key-value pairs (hash maps/dictionaries).
  - Creating, adding, accessing, updating, and deleting entries.
  - Checking for key existence.  

### 💻 Example Code:

```go
package main

import "fmt"

func main() {
	// Array example
	var numbers [5]int // Declares an array of 5 integers
	numbers[0] = 10
	numbers[1] = 20
	fmt.Println("Array:", numbers)
	fmt.Println("First element of array:", numbers[0])

	// Slice example
	fruits := []string{"apple", "orange", "banana"} // Creates a slice
	fmt.Println("\nInitial slice:", fruits)
	fruits = append(fruits, "mango") // Append elements
	fmt.Println("Slice after appending:", fruits)
	fmt.Println("Slice length:", len(fruits))
	fmt.Println("Slice capacity:", cap(fruits))

	// Slicing an existing slice
	someFruits := fruits[1:3] // Elements from index 1 (inclusive) to 3 (exclusive)
	fmt.Println("Sliced fruits (orange, banana):", someFruits)

	// Map example
	studentGrades := map[string]int{
		"Alice": 95,
		"Bob":   88,
	}
	fmt.Println("\nStudent Grades Map:", studentGrades)
	fmt.Println("Alice's grade:", studentGrades["Alice"])

	// Add a new entry
	studentGrades["Charlie"] = 77
	fmt.Println("Map after adding Charlie:", studentGrades)

	// Check if a key exists
	grade, exists := studentGrades["Bob"]
	if exists {
		fmt.Printf("Bob's grade: %d (exists: %t)\n", grade, exists)
	}

	// Delete an entry
	delete(studentGrades, "Bob")
	fmt.Println("Map after deleting Bob:", studentGrades)
}
```

----

## 📅 Day 4: Functions & Error Handling

Today you'll learn to modularize your code with functions and gracefully handle errors, which is a core part of Go's philosophy.

### ✅ Content:


- **Functions:**
  - Defining and calling functions.
  - Parameters and return values (single and multiple return values).
  - Named return values.
- **Error Handling:**
  - Understanding Go's idiomatic error handling: returning error as the last return value.
  - Checking for errors.
  - Using _ (blank identifier) to ignore unwanted return values (including errors when appropriate, but generally not recommended for errors).
- **defer Statement:** Scheduling a function call to be executed just before the surrounding function returns.

### 💻 Example Code:

```go
package main

import (
	"errors"
	"fmt"
)

// Simple function
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Function with single return value
func calculateSum(num1, num2 int) int {
	return num1 + num2
}

// Function with multiple return values and error handling
func divide(numerator, denominator float64) (float64, error) {
	if denominator == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, nil // nil means no error
}

func main() {
	greet("Go Learner")

	sum := calculateSum(10, 15)
	fmt.Printf("Sum of 10 and 15: %d\n", sum)

	// Error handling example
	result, err := divide(10.0, 2.0)
	if err != nil {
		fmt.Printf("Error during division: %s\n", err.Error())
	} else {
		fmt.Printf("Result of division (10/2): %.2f\n", result)
	}

	result, err = divide(10.0, 0.0) // This will cause an error
	if err != nil {
		fmt.Printf("Error during division: %s\n", err.Error())
	} else {
		fmt.Printf("Result of division (10/0): %.2f\n", result)
	}

	// Defer statement example
	fmt.Println("\nStarting main function")
	defer fmt.Println("Main function finished") // This will run last
	fmt.Println("Doing some work...")
}
```

---

## 📅 Day 5: Structs, Pointers & Methods

Today you'll learn about creating custom data types, working with memory addresses, and defining behaviors for your custom types.

### ✅ Content:


- **Structs:** Custom composite data types for grouping related fields.
  - Defining structs.
  - Creating struct instances and accessing fields.
  - Nested structs.
- **Pointers:** Variables that store memory addresses of other variables.
  - Declaring pointers, dereferencing (*), and taking address (&).
  - When to use pointers (passing by reference, modifying values).
- **Methods:** Functions associated with a specific type (structs).
  - Receiver functions (value receivers vs. pointer receivers).

### 💻 Example Code:

```go
package main

import "fmt"

// Define a Person struct
type Person struct {
	Name string
	Age  int
	City string
}

// Define a method for the Person struct (value receiver)
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// Define a method with a pointer receiver to modify the struct
func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Printf("%s just had a birthday! New age: %d\n", p.Name, p.Age)
}

func main() {
	// Create a struct instance
	person1 := Person{
		Name: "Bob",
		Age:  25,
		City: "New York",
	}
	fmt.Println("Person 1:", person1)
	person1.Greet()

	// Accessing struct fields
	fmt.Printf("%s lives in %s.\n", person1.Name, person1.City)

	// Pointers example
	x := 10
	ptrX := &x // ptrX now holds the memory address of x
	fmt.Printf("\nValue of x: %d, Address of x: %p\n", x, ptrX)
	fmt.Printf("Value at ptrX: %d\n", *ptrX) // Dereference ptrX to get the value

	*ptrX = 20 // Change the value at the memory address
	fmt.Printf("New value of x: %d\n", x)

	// Using a method with a pointer receiver
	fmt.Println("\nBefore birthday:", person1)
	person1.HaveBirthday() // Go automatically handles &person1 for method calls
	fmt.Println("After birthday:", person1)
}
```

---

## 📅 Day 6: Concurrency (Goroutines & Channels)

Go's superpower! Today you'll get a taste of Go's built-in concurrency primitives.

### ✅ Content:


- **Concurrency vs. Parallelism:** Understanding the difference.
  - Goroutines: Lightweight threads managed by the Go runtime.
  - Starting a goroutine using the go keyword.
  - The main goroutine.
- **Channels:** Type-safe conduits for communicating between goroutines.
  - Creating channels (make(chan type)).
  - Sending (<-chan) and receiving (chan<-) data from channels.
  - Buffered vs. unbuffered channels.
- **sync.WaitGroup:** A mechanism to wait for a collection of goroutines to finish.

### 💻 Example Code:

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// A function to be run as a goroutine
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine finishes
	fmt.Printf("Worker %d starting...\n", id)
	time.Sleep(time.Second) // Simulate some work
	fmt.Printf("Worker %d finished.\n", id)
}

// A function to send messages through a channel
func sender(ch chan string) {
	ch <- "Hello from sender!" // Send a message to the channel
}

// A function to receive messages from a channel
func receiver(ch chan string) {
	msg := <-ch // Receive a message from the channel
	fmt.Printf("Received: %s\n", msg)
}

func main() {
	fmt.Println("--- Goroutines with WaitGroup ---")
	var wg sync.WaitGroup // Create a WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)       // Increment the counter for each goroutine
		go worker(i, &wg) // Start a goroutine
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All workers finished.")

	fmt.Println("\n--- Channels ---")
	messageChannel := make(chan string) // Create an unbuffered channel of strings

	go sender(messageChannel)    // Start a goroutine to send
	go receiver(messageChannel) // Start a goroutine to receive

	// Give goroutines time to execute (in a real app, you'd use sync primitives)
	time.Sleep(100 * time.Millisecond)

	// Buffered channel example
	bufferedChannel := make(chan int, 2) // Buffered channel with capacity 2
	bufferedChannel <- 1
	bufferedChannel <- 2
	fmt.Println("\nReceived from buffered channel:", <-bufferedChannel)
	fmt.Println("Received from buffered channel:", <-bufferedChannel)
	close(bufferedChannel) // It's good practice to close channels when done sending
}
```

## 📅 Day 7: Practical Project & Next Steps

Today is about consolidating your knowledge by building a small, practical Go application and understanding where to go next.



### ✅ Content:


- **Building a Simple Command-Line Application:** Apply all the concepts learned into a coherent project.
  - Example Project: A simple Task Manager CLI (Command Line Interface).
    - Features: Add task, list tasks, mark task as complete.
    - This will involve structs, slices, functions, basic I/O, and possibly error handling.
- **Recommended Next Steps:**
  - Explore Go's standard library.
  - Learn about packages and modules for organizing larger projects.
  - Dive deeper into interfaces and composition.
  - Practice building more complex applications.
  - Consider contributing to open-source Go projects.

### Example Project: Simple Task Manager CLI
This example will demonstrate a basic task manager where you can add and list tasks. Expanding on this would be a great practice!  


```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task struct to represent a single task
type Task struct {
	ID        int
	Desc      string
	Completed bool
}

var tasks []Task // Global slice to store tasks
var nextID = 1    // Global counter for task IDs

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--- Simple Go Task Manager ---")
	for {
		fmt.Println("\nOptions:")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			addTask(reader)
		case "2":
			listTasks()
		case "3":
			fmt.Println("Exiting Task Manager. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

// addTask prompts the user for a task description and adds it to the list
func addTask(reader *bufio.Reader) {
	fmt.Print("Enter task description: ")
	desc, _ := reader.ReadString('\n')
	desc = strings.TrimSpace(desc)

	if desc == "" {
		fmt.Println("Task description cannot be empty.")
		return
	}

	newTask := Task{
		ID:        nextID,
		Desc:      desc,
		Completed: false,
	}
	tasks = append(tasks, newTask)
	nextID++
	fmt.Println("Task added successfully!")
}

// listTasks prints all current tasks
func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks yet. Add some!")
		return
	}

	fmt.Println("\n--- Your Tasks ---")
	for _, task := range tasks {
		status := "[ ]"
		if task.Completed {
			status = "[X]"
		}
		fmt.Printf("%s ID: %d - %s\n", status, task.ID, task.Desc)
	}
	fmt.Println("--------------------")
}

// Example of how you might extend this with a 'Mark Complete' function (not integrated into main loop)
func markTaskComplete(taskID int) {
	for i := range tasks {
		if tasks[i].ID == taskID {
			tasks[i].Completed = true
			fmt.Printf("Task %d marked as complete.\n", taskID)
			return
		}
	}
	fmt.Printf("Task with ID %d not found.\n", taskID)
}

// Example of how you might extend this with a 'Delete Task' function (not integrated into main loop)
func deleteTask(taskID int) {
    for i, task := range tasks {
        if task.ID == taskID {
            // Remove the task from the slice
            tasks = append(tasks[:i], tasks[i+1:]...)
            fmt.Printf("Task %d deleted.\n", taskID)
            return
        }
    }
    fmt.Printf("Task with ID %d not found.\n", taskID)
}
```

### 💡 Tips for Quick Learning and Immediate Application:
- Hands-on Practice is Crucial! The best way to learn Go is by writing code. Implement every example, and then try to modify it or create your own variations.
- Understand Go's Philosophy: Go emphasizes simplicity, efficiency, and concurrency. Try to think "the Go way" when designing your code.
- Read the Official Documentation: Go's documentation is excellent and concise. Refer to it often (e.g., go doc fmt.Println).
- Explore Go's Standard Library: Go has a rich standard library. Familiarize yourself with common packages like fmt, os, strings, strconv, net/http, etc.
- Error Handling: Pay close attention to Go's unique way of handling errors. Always check for errors!
- Community and Resources: Join Go communities on platforms like Reddit (r/golang), Discord, or Stack Overflow. There are also many great blogs and tutorials online.