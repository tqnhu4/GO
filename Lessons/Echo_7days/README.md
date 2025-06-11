# Echo (Go) 7-Day Learning Roadmap
This 7-day roadmap is designed to get you proficient with Echo, a high-performance, minimalist web framework for Go. It assumes you have a basic understanding of Go's syntax and core concepts (variables, functions, loops, structs, interfaces).

## 🚀 Day 1: Go & Echo Fundamentals – First Steps
Start by setting up your Go environment and running your very first Echo web server.

- Go Environment Setup:
  - Install Go: Download the latest version from go.dev/dl/.
  - Verify installation: Open your terminal and run go version.
  - Familiarize yourself with basic Go commands like go run (to execute a Go program), go build (to compile), go mod init (to initialize a new module), and go get (to download dependencies).
- Echo Installation:
  - Create a new Go module for your project: go mod init your-echo-project (replace your-echo-project with your desired name).
  - Install Echo: go get github.com/labstack/echo/v4
- Basic Echo Server: Understand how to create an Echo instance, define routes, and start the server.

```go
// main.go
package main

import (
	"net/http" // Standard Go HTTP package
	"github.com/labstack/echo/v4" // Echo framework
	"github.com/labstack/echo/v4/middleware" // Echo middleware
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Add default middleware: Logger and Recover (recovers from panics)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define a GET route for the root ("/") endpoint
	e.GET("/", func(c echo.Context) error {
		// c.String sends a plain text response
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	// Start the server on port 8080
	e.Logger.Fatal(e.Start(":8080")) // e.Logger.Fatal logs and exits on error
}
```

- Run your first Echo app:
  - Save the code as main.go in your project directory.
  - Run go run main.go in your terminal.
  - Open http://localhost:8080/ in your web browser or use curl http://localhost:8080/.
- Hands-on: Create a new Echo project. Implement a /greeting endpoint that returns a plain text message "Greetings from Echo!".


## 🛣️ Day 2: Routes, Path & Query Parameters
Learn how to define different API endpoints and extract data from the URL.

- Path Parameters: Extract dynamic values directly from the URL path. Echo automatically validates and converts types if specified.

```go
// main.go (continuing from Day 1)
package main

import (
	"net/http"
	"strconv" // Used for string to int conversion
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Path parameter: :name
	e.GET("/users/:name", func(c echo.Context) error {
		name := c.Param("name") // Get the value of the 'name' parameter
		return c.JSON(http.StatusOK, map[string]string{"user": name})
	})

	// Path parameter with a type conversion (optional, but good practice)
	e.GET("/products/:id", func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr) // Convert string to int
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid product ID")
		}
		return c.JSON(http.StatusOK, map[string]int{"product_id": id})
	})

	// Path parameter with multi-segment wildcard: *
	e.GET("/files/:filepath/*", func(c echo.Context) error {
		filepath := c.Param("filepath")
		action := c.Param("*") // This will capture anything after /filepath/
		return c.JSON(http.StatusOK, map[string]string{
			"filepath": filepath,
			"action":   action,
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
```

- Query Parameters: Extract values from the URL query string (e.g., ?key=value). They are optional by default.

```go
// main.go (continuing from Day 1)
package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Query parameters: /search?q=keyword&limit=10
	e.GET("/search", func(c echo.Context) error {
		query := c.QueryParam("q")         // Get "q" parameter, "" if not present
		limit := c.QueryParam("limit")     // Get "limit" parameter

		if limit == "" { // Provide a default if not present
			limit = "10"
		}

		return c.JSON(http.StatusOK, map[string]string{
			"query": query,
			"limit": limit,
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
```

- Hands-on:
  - Create a GET /books/:title endpoint that returns the book title.
  - Create a GET /articles?author=John&year=2023 endpoint that returns the author and year.
  - Test your endpoints using curl or Postman to see how parameters are handled.


## 📦 Day 3: Request Body & Data Binding
Learn how to handle data sent in the request body, typically for POST or PUT requests, and bind it to Go structs.

- Request Body & Binding: Echo provides Bind and Validate methods to parse JSON, XML, or form data into Go structs.  

```go
// main.go (continuing from Day 1)
package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

// Define a struct to bind JSON or form data to
type User struct {
	Name  string `json:"name" form:"name" query:"name"` // JSON, form-data, and query parameter binding
	Email string `json:"email" form:"email" query:"email" validate:"required,email"` // Added validation tag
	Age   int    `json:"age" form:"age" query:"age"`
}

func main() {
	e := echo.New()

	// Add a validator for struct fields
	e.Validator = &CustomValidator{validator: validator.New()}

	// POST endpoint to create a user from JSON or form data
	e.POST("/users", func(c echo.Context) error {
		user := new(User) // Create a pointer to a User struct
		if err := c.Bind(user); err != nil { // Binds JSON, form-data, or query parameters
			return c.String(http.StatusBadRequest, err.Error())
		}
		// Validate the bound struct
		if err := c.Validate(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{"message": "User created", "user": user})
	})

	e.Logger.Fatal(e.Start(":8080"))
}

// You need to install go-playground/validator for validation:
// go get github.com/go-playground/validator/v10

// CustomValidator structure to integrate go-playground/validator with Echo
type CustomValidator struct {
	validator *validator.Validate
}

// Validate method for Echo's Validator interface
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, return a custom error message
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
```

- Validation: Integrate a validation library (like go-playground/validator) to validate incoming data against rules defined in struct tags.
- Hands-on:
  - Create a POST /products endpoint that accepts a JSON request body for a Product (e.g., Name: string, Price: float, InStock: bool).
  - Add validation to ensure Name is required and Price is positive.
  - Test with valid and invalid JSON input to observe the binding and validation in action.

## Day 4: 🔒 Middleware & Route Grouping
Discover how to use Echo's powerful middleware system for common tasks and organize your routes efficiently.

- Middleware: Functions that run before or after route handlers. Used for logging, authentication, error handling, CORS, etc.
  - Built-in Middleware: Echo comes with many useful middleware (Logger, Recover, CORS, JWT, etc.).
  - Custom Middleware: Write your own middleware functions.  

```go
// main.go (continuing from Day 1)
package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Custom Middleware: Request Logger
func RequestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c) // Call the next middleware or route handler
			duration := time.Since(start)
			fmt.Printf("Request: %s %s | Status: %d | Latency: %s\n",
				c.Request().Method, c.Request().URL.Path, c.Response().Status, duration)
			return err
		}
	}
}

func main() {
	e := echo.New()

	// Apply custom middleware globally
	e.Use(RequestLogger())
	e.Use(middleware.Recover()) // Built-in recovery middleware

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
```

- Route Grouping: Organize routes with common prefixes or middleware.

```go

// main.go (continuing from Day 1)
package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

// Simple authentication middleware
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("X-API-Key") != "SUPER_SECRET" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid API Key")
		}
		return next(c)
	}
}

func main() {
	e := echo.New()

	// Public routes group
	publicGroup := e.Group("/public")
	{ // The curly braces are just for visual grouping, not strictly necessary for Go syntax
		publicGroup.GET("/info", func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]string{"message": "This is public info"})
		})
	}

	// Private routes group with AuthMiddleware
	privateGroup := e.Group("/private")
	privateGroup.Use(AuthMiddleware) // Apply AuthMiddleware to all routes in this group
	{
		privateGroup.GET("/dashboard", func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]string{"data": "Private dashboard data"})
		})
		privateGroup.POST("/settings", func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]string{"message": "Settings updated"})
		})
	}

	e.Logger.Fatal(e.Start(":8080"))
}

```

- Hands-on:
  - Create a custom middleware that logs the incoming request method, path, and response status code. Apply it globally.
  - Create a route group /admin and apply an AuthMiddleware (which checks for a hardcoded X-Admin-Token header) to it. Create an endpoint /admin/status.


## 📊 Day 5: Data Handling & Response Types
Explore various ways to handle incoming data and send different types of responses back to the client.

- Uploading Files: Handle single and multiple file uploads.  

```go
// main.go (continuing from Day 1)
package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Endpoint for single file upload
	e.POST("/upload", func(c echo.Context) error {
		fileHeader, err := c.FormFile("file") // "file" is the name of the input field in the form
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		src, err := fileHeader.Open()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		defer src.Close()

		// Create the destination directory if it doesn't exist
		uploadDir := "uploads"
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			os.Mkdir(uploadDir, 0755) // Create directory with read/write/execute permissions for owner, read/execute for others
		}

		destination := filepath.Join(uploadDir, fileHeader.Filename)
		dst, err := os.Create(destination)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "File uploaded successfully!", "filename": fileHeader.Filename})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
```
- Rendering HTML Templates: Serve dynamic HTML using Go's html/template package with Echo.

```go
// main.go (continuing from Day 1)
package main

import (
	"html/template"
	"io"
	"net/http"
	"github.com/labstack/echo/v4"
)

// Define a custom template renderer
type Template struct {
	templates *template.Template
}

// Implement Echo's Renderer interface
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	// Load templates
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")), // Load all .html files from 'templates'
	}
	e.Renderer = t // Set the custom renderer

	e.GET("/html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{ // Render "index.html"
			"title": "Echo HTML Example",
			"name":  "Echo User",
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}

// templates/index.html (create this file)
<!DOCTYPE html>
<html>
<head>
    <title>{{ .title }}</title>
</head>
<body>
    <h1>Hello, {{ .name }}!</h1>
    <p>This is an Echo HTML template example.</p>
</body>
</html>
```

- Serving Static Files: Serve static assets like CSS, JavaScript, and images efficiently.


```go
// main.go (continuing from Day 1)
package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Serve static files from the "static" directory under the "/static" URL path
	// Files will be accessed like http://localhost:8080/static/css/style.css
	e.Static("/static", "static")

	// You can also serve a single static file directly
	// e.File("/favicon.ico", "static/favicon.ico")

	e.Logger.Fatal(e.Start(":8080"))
}
// Create a 'static' directory and put some files in it, e.g., 'static/css/style.css', 'static/images/logo.png'
```

- Hands-on:
  - Create an endpoint that allows users to upload a profile picture.
  - Create a simple HTML template for a "dashboard" and an endpoint that renders it with dynamic user data.
  - Set up serving static files (e.g., a JavaScript file or an image).

## 💾 Day 6: Database Integration & RESTful API
Connect your Echo application to a database and build a full RESTful API with CRUD operations.

- Database Choice: Choose a database (e.g., SQLite for simplicity, PostgreSQL for more robust apps).
- ORM/Database Driver:
  - GORM: A popular ORM for Go (Object-Relational Mapper).
  - database/sql package for direct database interaction.
  - For this example, we'll use GORM with SQLite for simplicity.
- GORM Installation: go get gorm.io/gorm and go get gorm.io/driver/sqlite
- Defining Models: Go structs annotated for GORM.  

```go
// models/user.go (create a new directory 'models' and this file)
package models

import "gorm.io/gorm"

type User struct {
    gorm.Model // Provides ID, CreatedAt, UpdatedAt, DeletedAt
    Name  string `json:"name"`
    Email string `json:"email" gorm:"unique"`
}

// database/db.go (create a new directory 'database' and this file)
package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"your-echo-project/models" // Adjust import path
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	// Migrate the schema
	DB.AutoMigrate(&models.User{})
}
```

```go
// main.go (Integrate DB and API handlers)
package main

import (
	"net/http"
	"strconv" // For converting ID string to int

	"github.com/labstack/echo/v4"
	"your-echo-project/database" // Adjust import path
	"your-echo-project/models"   // Adjust import path
)

func main() {
	e := echo.New()
    
    database.InitDB() // Initialize the database connection
    // Create User (POST /users)
	e.POST("/users", func(c echo.Context) error {
		user := new(models.User)
		if err := c.Bind(user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		database.DB.Create(user)
		return c.JSON(http.StatusCreated, user)
	})

	// Get All Users (GET /users)
	e.GET("/users", func(c echo.Context) error {
		var users []models.User
		database.DB.Find(&users)
		return c.JSON(http.StatusOK, users)
	})

	// Get User by ID (GET /users/:id)
	e.GET("/users/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id")) // Convert ID from string to int
		var user models.User
		if result := database.DB.First(&user, id); result.Error != nil {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return c.JSON(http.StatusOK, user)
	})

	// Update User (PUT /users/:id)
	e.PUT("/users/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var user models.User
		if result := database.DB.First(&user, id); result.Error != nil {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}

		if err := c.Bind(&user); err != nil { // Bind updated data to existing user object
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		database.DB.Save(&user)
		return c.JSON(http.StatusOK, user)
	})

	// Delete User (DELETE /users/:id)
	e.DELETE("/users/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var user models.User
		if result := database.DB.First(&user, id); result.Error != nil {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		database.DB.Delete(&user)
		return c.NoContent(http.StatusNoContent) // No content on successful delete
	})

	e.Logger.Fatal(e.Start(":8080"))
}    
```

- Hands-on:
  - Set up a SQLite database and integrate GORM.
  - Define a Product model with Name, Description, Price fields.
  - Implement full CRUD operations for Products (create, read all, read by ID, update, delete).

## 🧪 Day 7: Testing, Error Handling & Deployment
Learn how to write tests for your Echo API, handle errors gracefully, and understand deployment concepts.

- Testing Echo Applications: Use Go's net/http/httptest package along with Echo to write unit and integration tests for your handlers.  

```go
// main_test.go (in the same package as main.go)
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert" // go get github.com/stretchr/testify
)

func TestHelloWorld(t *testing.T) {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, Echo!", rec.Body.String())
}

func TestCreateUserEndpoint(t *testing.T) {
	// For this test, you might want to mock the database interaction
	// or use an in-memory database like SQLite for testing
	// (Simplified for example purposes)
	e := echo.New()
	type User struct { Name string `json:"name"` Email string `json:"email"` }
	e.POST("/users", func(c echo.Context) error {
		user := new(User)
		if err := c.Bind(user); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		// Simulate database creation
		return c.JSON(http.StatusCreated, map[string]interface{}{"message": "User created", "user": user})
	})

	jsonStr := []byte(`{"name":"Test User", "email":"test@example.com"}`)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "User created", response["message"])
	assert.Equal(t, "Test User", response["user"].(map[string]interface{})["name"])
}
// Run tests with: go test -v
```

- Custom Error Handling: Implement more sophisticated error handling for specific HTTP status codes or custom error types.


```go
// main.go (Custom Error Handling Example)
package main

import (
	"errors"
	"net/http"
	"github.com/labstack/echo/v4"
)

// Define custom error types
var ErrUserNotFound = errors.New("user not found")
var ErrInvalidUserEmail = errors.New("invalid user email")

// Custom HTTP error handler function
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"

	// Check if it's an Echo HTTP error
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = fmt.Sprintf("%v", he.Message) // Convert interface{} to string
	} else if errors.Is(err, ErrUserNotFound) {
		code = http.StatusNotFound
		message = "User not found"
	} else if errors.Is(err, ErrInvalidUserEmail) {
		code = http.StatusBadRequest
		message = "Provided email is invalid"
	}

	// Send the JSON error response
	if !c.Response().Committed { // Check if response is already committed
		c.JSON(code, map[string]string{"error": message})
	}
	c.Logger().Error(err) // Log the error
}

func main() {
	e := echo.New()

	// Set the custom error handler
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.GET("/fetch-user/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "nonexistent" {
			return ErrUserNotFound // Return custom error
		}
		if id == "bademail" {
			return ErrInvalidUserEmail
		}
		return c.String(http.StatusOK, "User fetched: "+id)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
```

- Graceful Shutdown: Implement logic to gracefully shut down your Echo server on receiving termination signals (e.g., Ctrl+C).
- Deployment Strategies:
  - Building the executable: go build -o myapp main.go
  - Docker: Containerize your Echo application for easy deployment.
  - Cloud Platforms: Deploy to AWS (EC2, ECS, Lambda with API Gateway), Google Cloud (Compute Engine, Cloud Run), Heroku, DigitalOcean.
  - Reverse Proxies: Using Nginx or Caddy in front of your Go application.
- Hands-on:
  - Write tests for your Product CRUD API from Day 6. Test valid/invalid inputs and verify status codes.
  - Implement custom error handling for 404 Not Found (when a product ID doesn't exist) and 400 Bad Request (for validation errors).
  - Research how to build a Docker image for your Echo application.



