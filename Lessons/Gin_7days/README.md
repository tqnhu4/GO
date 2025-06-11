# Gin (Go) 7-Day Learning Roadmap
This 7-day roadmap is designed to get you proficient with Gin, a popular web framework for Go, allowing you to build high-performance web APIs quickly. This roadmap assumes you have a basic understanding of Go's syntax and core concepts (variables, functions, loops, structs, interfaces).

## 🚀 Day 1: Go & Gin Fundamentals – Getting Started
Set up your Go environment and create your very first Gin web server.

- Go Environment Setup:
  - Install Go: Download from go.dev/dl/.
  - Set up your GOPATH (Go 1.11+ uses Go Modules by default, so GOPATH is less critical for projects but good to know).
  - Familiarize yourself with basic Go commands (go run, go build, go mod init, go get).
- Gin Installation:
  - go get github.com/gin-gonic/gin
- Basic Gin Server: Understand how to create a Gin Engine, define routes, and run the server.
   // main.go
package main

```go
import "github.com/gin-gonic/gin"

func main() {
    // Create a Gin router with default middleware: logger and recovery (crash-free)
    r := gin.Default()

    // Define a GET route for the root ("/") endpoint
    r.GET("/", func(c *gin.Context) {
        // c.JSON sends a JSON response with the specified status code
        c.JSON(200, gin.H{
            "message": "Hello, Gin!",
        })
    })

    // Run the server on port 8080
    // By default, it runs on :8080. You can specify a different port like r.Run(":8081")
    r.Run() // listen and serve on 0.0.0.0:8080
}
```

- Run your first Gin app:
  - Create a new Go module: go mod init your-project-name
  - Run the code: go run main.go
  - Open http://localhost:8080/ in your browser or use curl.
- Hands-on: Create a new Gin project. Implement a /hello endpoint that returns a JSON message "Hello, [Your Name]!".


## 🛣️ Day 2: Routes, Parameters & Data Binding
Learn how to define different routes, extract data from URLs, and bind request bodies to Go structs.

- Path Parameters: Extract dynamic values from the URL path.

```go
// main.go (continuing from Day 1)
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()

    // Path parameter: :name
    r.GET("/users/:name", func(c *gin.Context) {
        name := c.Param("name") // Get the value of the 'name' parameter
        c.JSON(200, gin.H{"user": name})
    })

    // Path parameter with multi-segment wildcard: *action
    r.GET("/files/:filepath/*action", func(c *gin.Context) {
        filepath := c.Param("filepath")
        action := c.Param("action") // This will capture anything after /filepath/
        c.JSON(200, gin.H{
            "filepath": filepath,
            "action":   action,
        })
    })

    r.Run()
}
```

- Query Parameters: Extract values from the URL query string.

```go
// main.go (continuing from Day 1)
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()

    // Query parameters: /search?q=keyword&limit=10
    r.GET("/search", func(c *gin.Context) {
        query := c.Query("q")        // Get "q" parameter, "" if not present
        limit := c.DefaultQuery("limit", "10") // Get "limit", default to "10" if not present
        c.JSON(200, gin.H{
            "query": query,
            "limit": limit,
        })
    })

    r.Run()
}
```

- Form & JSON Binding: Bind request data (from forms or JSON bodies) directly into Go structs.


```go
// main.go (continuing from Day 1)
package main

import "github.com/gin-gonic/gin"

// Define a struct to bind JSON or form data to
type User struct {
    Name  string `json:"name" form:"name" binding:"required"` // `binding:"required"` makes it mandatory
    Email string `json:"email" form:"email"`
    Age   int    `json:"age" form:"age"`
}

func main() {
    r := gin.Default()

    // POST endpoint to create a user
    r.POST("/users", func(c *gin.Context) {
        var user User
        // c.BindJSON for JSON, c.Bind for form data
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        c.JSON(201, gin.H{"message": "User created", "user": user})
    })

    r.Run()
}
```

- Hands-on:
  - Create a GET /products/:id endpoint.
  - Create a GET /products?category=electronics&min_price=100 endpoint.
  - Create a POST /register endpoint that accepts a JSON body for a User (with username, password, email) and binds it to a struct.


## 🔒 Day 3: Middleware & Grouping Routes
Learn how to use Gin's powerful middleware system for common tasks and organize your routes efficiently.

- Middleware: Functions that run before or after route handlers. Used for logging, authentication, error handling, etc.
  - Global Middleware: Applied to all routes.
  - Route-specific Middleware: Applied to individual routes.
  - Custom Middleware: Write your own middleware.  

```go
// main.go (continuing from Day 1)
package main

import (
    "fmt"
    "time"
    "github.com/gin-gonic/gin"
)

// Custom Middleware: Logger
func LoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now() // Start timer

        // Process request
        c.Next() // Call the next middleware or route handler

        // After request is processed
        latency := time.Since(t)
        status := c.Writer.Status()
        fmt.Printf("Request: %s %s | Status: %d | Latency: %s\n",
            c.Request.Method, c.Request.URL.Path, status, latency)
    }
}

func main() {
    r := gin.New() // gin.New() gives you a router without default middleware
    r.Use(LoggerMiddleware()) // Use our custom logger middleware globally
    r.Use(gin.Recovery())     // gin.Recovery middleware recovers from panics

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Welcome!"})
    })

    // Route with specific middleware
    r.GET("/ping", LoggerMiddleware(), func(c *gin.Context) { // LoggerMiddleware applied only here
        c.JSON(200, gin.H{"message": "pong"})
    })

    r.Run()
}
```

- Route Grouping: Organize routes with common prefixes or middleware.

```go
// main.go (continuing from Day 1)
package main

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Simulate authentication check
        if c.GetHeader("Authorization") != "Bearer my-secret-token" {
            c.AbortWithStatus(401) // Abort request with 401 Unauthorized
            return
        }
        c.Next() // Continue to the next handler
    }
}

func main() {
    r := gin.Default()

    // Public routes group
    public := r.Group("/public")
    {
        public.GET("/info", func(c *gin.Context) {
            c.JSON(200, gin.H{"data": "Public information"})
        })
    }

    // Authenticated routes group with AuthMiddleware
    private := r.Group("/private")
    private.Use(AuthMiddleware()) // Apply AuthMiddleware to all routes in this group
    {
        private.GET("/dashboard", func(c *gin.Context) {
            c.JSON(200, gin.H{"data": "Private dashboard data"})
        })
        private.POST("/settings", func(c *gin.Context) {
            c.JSON(200, gin.H{"message": "Settings updated"})
        })
    }

    r.Run()
}

```

- Hands-on:
  - Create a custom middleware that logs the incoming request method and path. Apply it globally.
  - Create a route group /admin and apply an AuthMiddleware (which just checks for a hardcoded X-Admin-Key header) to it. Create an endpoint /admin/dashboard.

## 📊 Day 4: Data Handling & Response Types
Explore various ways to handle data and send different types of responses.

- Uploading Files: Handle single and multiple file uploads.  

```go
// main.go (continuing from Day 1)
package main

import (
    "net/http"
    "path/filepath"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Endpoint for single file upload
    r.POST("/upload", func(c *gin.Context) {
        file, err := c.FormFile("file") // "file" is the name of the input field in the form
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Save the uploaded file
        destination := filepath.Join("uploads", file.Filename) // Make sure 'uploads' directory exists
        if err := c.SaveUploadedFile(file, destination); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully!", "filename": file.Filename})
    })

    r.Run()
}
```

- Rendering HTML Templates: Serve dynamic HTML using Go's html/template package with Gin.

```go
// main.go (continuing from Day 1)
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.LoadHTMLGlob("templates/*") // Load all HTML files from the "templates" directory

    r.GET("/html", func(c *gin.Context) {
        c.HTML(200, "index.html", gin.H{ // Render "index.html" template
            "title": "Gin HTML Example",
            "name":  "Gin User",
        })
    })

    r.Run()
}

// templates/index.html
<!DOCTYPE html>
<html>
<head>
    <title>{{ .title }}</title>
</head>
<body>
    <h1>Hello, {{ .name }}!</h1>
    <p>This is a Gin HTML template example.</p>
</body>
</html>
```

- Redirects: How to perform HTTP redirects.
- Static Files: Serve static assets like CSS, JavaScript, and images.


```go
// main.go (continuing from Day 1)
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()

    // Serve static files from the "static" directory under the "/assets" URL path
    r.Static("/assets", "./static") // Access like http://localhost:8080/assets/image.png

    // Serve a single static file directly
    r.StaticFile("/favicon.ico", "./static/favicon.ico")

    r.Run()
}
// Create a 'static' directory and put some files in it, e.g., 'static/style.css', 'static/image.png'
```

- Hands-on:
  - Create an endpoint that allows users to upload an image.
  - Create a simple HTML template and an endpoint that renders it with dynamic data.
  - Set up serving static files (e.g., a CSS file or an image).

## 💾 Day 5: Database Integration & RESTful API Design
Connect your Gin application to a database and build a full RESTful API.

- Database Choice: Pick a database (e.g., SQLite for simplicity, PostgreSQL for more robust apps).
- ORM/Database Driver:
  - GORM: A popular ORM for Go (similar to SQLAlchemy in Python, Active Record in Rails).
  - database/sql driver directly for basic operations.
  - For this example, we'll use GORM with SQLite.
- GORM Installation: go get gorm.io/gorm and go get gorm.io/driver/sqlite
- Defining Models: Go structs annotated for GORM.  


```go
// models/user.go
package models

import "gorm.io/gorm"

type User struct {
    gorm.Model // Provides ID, CreatedAt, UpdatedAt, DeletedAt
    Name  string `json:"name"`
    Email string `json:"email" gorm:"unique"`
}

// main.go (partial)
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "your-project-name/models" // Adjust path as needed
)

func setupDatabase() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&models.User{}) // Auto-migrate User model to database table
    return db
}

func main() {
    db := setupDatabase()
    r := gin.Default()

    // Dependency Injection for DB (using Gin's Context)
    r.Use(func(c *gin.Context) {
        c.Set("db", db) // Store DB instance in context
        c.Next()
    })

    // Create User (POST /users)
    r.POST("/users", func(c *gin.Context) {
        var input struct {
            Name  string `json:"name" binding:"required"`
            Email string `json:"email" binding:"required,email"`
        }
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        db := c.MustGet("db").(*gorm.DB)
        user := models.User{Name: input.Name, Email: input.Email}
        db.Create(&user)
        c.JSON(http.StatusCreated, user)
    })

    // Get Users (GET /users)
    r.GET("/users", func(c *gin.Context) {
        db := c.MustGet("db").(*gorm.DB)
        var users []models.User
        db.Find(&users)
        c.JSON(http.StatusOK, users)
    })

    // ... Implement GET /users/:id, PUT /users/:id, DELETE /users/:id ...

    r.Run()
}
```

- Building a RESTful API: Implement full CRUD (Create, Read, Update, Delete) operations for a resource (e.g., User or Product).
- Hands-on:
  - Set up a SQLite database and integrate GORM.
  - Define a Product model with Name, Description, Price fields.
  - Implement POST /products to create a product, GET /products to list all, GET /products/:id to get one, PUT /products/:id to update, and DELETE /products/:id to delete.

## 🧪 Day 6: Testing & Error Handling
Learn how to write tests for your Gin API and handle errors gracefully.

- Testing Gin Applications: Use net/http/httptest and Gin's TestEngine to write unit and integration tests.  

```go
// main_test.go
package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert" // go get github.com/stretchr/testify
)

func TestHelloWorld(t *testing.T) {
    gin.SetMode(gin.TestMode) // Set Gin to test mode (disables debug output)
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello, Gin!"})
    })

    // Create a request to send to the router
    req, _ := http.NewRequest("GET", "/", nil)
    // Create a ResponseRecorder (recorder) to record the response
    w := httptest.NewRecorder()
    // Serve the HTTP request to the recorder
    r.ServeHTTP(w, req)

    // Assertions
    assert.Equal(t, 200, w.Code) // Check status code
    var response map[string]string
    json.Unmarshal(w.Body.Bytes(), &response)
    assert.Equal(t, "Hello, Gin!", response["message"]) // Check response body
}

func TestCreateUser(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := gin.Default()
    // Define the endpoint (similar to Day 2 example)
    type User struct { Name string `json:"name"` }
    r.POST("/users", func(c *gin.Context) {
        var user User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        c.JSON(201, gin.H{"message": "User created", "user": user})
    })

    jsonStr := []byte(`{"name":"Test User"}`)
    req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, 201, w.Code)
    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    assert.Equal(t, "User created", response["message"])
    assert.Equal(t, "Test User", response["user"].(map[string]interface{})["name"])
}

// Run tests with: go test -v
```

- Custom Error Handling: Implement more sophisticated error handling for specific error types.

```go
// main.go (Custom Error Handling Example)
package main

import (
    "errors"
    "net/http"
    "github.com/gin-gonic/gin"
)

// Define custom error types
var ErrNotFound = errors.New("resource not found")
var ErrInvalidInput = errors.New("invalid input data")

func main() {
    r := gin.Default()

    r.GET("/resource/:id", func(c *gin.Context) {
        // Simulate an error
        id := c.Param("id")
        if id == "notfound" {
            c.Error(ErrNotFound) // Use c.Error to push an error to the context
            return
        }
        if id == "invalid" {
            c.Error(ErrInvalidInput)
            return
        }
        c.JSON(200, gin.H{"message": "Resource found"})
    })

    // Custom error handler middleware
    r.Use(func(c *gin.Context) {
        c.Next() // Process request first
        // After all handlers are done, check for errors
        if len(c.Errors) > 0 {
            for _, e := range c.Errors {
                if errors.Is(e.Err, ErrNotFound) {
                    c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
                } else if errors.Is(e.Err, ErrInvalidInput) {
                    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid input provided"})
                } else {
                    // Catch all other errors
                    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
                }
                // Only process the first error to avoid multiple responses
                break
            }
        }
    })

    r.Run()
}
```

- Logging: Configure structured logging for better observability.
- Hands-on:
  - Write tests for your Product CRUD API from Day 5. Test valid and invalid inputs, and verify status codes.
  - Implement custom error handling for 404 Not Found (e.g., when a product ID doesn't exist) and 400 Bad Request (for invalid input during creation/update).

## ⚙️ Day 7: Advanced Topics & Deployment
Explore more advanced Gin features and understand how to deploy your Go application.

- Concurrency in Go: How Go's goroutines and channels can be used to handle background tasks or complex processing within a Gin app (e.g., using a goroutine to send an email after a user registers).  

```go
// Example: Background Task with Goroutine
package main

import (
    "fmt"
    "time"
    "github.com/gin-gonic/gin"
)

func sendEmail(email string, subject string, body string) {
    fmt.Printf("Sending email to %s: Subject: %s, Body: %s\n", email, subject, body)
    time.Sleep(5 * time.Second) // Simulate email sending delay
    fmt.Printf("Email to %s sent!\n", email)
}

func main() {
    r := gin.Default()

    r.POST("/register", func(c *gin.Context) {
        email := c.PostForm("email") // Assuming form data
        if email == "" {
            c.JSON(400, gin.H{"error": "Email is required"})
            return
        }

        // Start sending email in a separate goroutine (non-blocking)
        go sendEmail(email, "Welcome!", "Thanks for registering!")

        c.JSON(200, gin.H{"message": "Registration successful, email sent in background"})
    })

    r.Run()
}
```

- Graceful Shutdown: How to gracefully shut down your Gin server to avoid dropping active requests.
- Deployment Strategies:
  - Building the executable: go build -o myapp main.go
  - Docker: Containerize your Gin application.
  - Cloud Platforms: Deploy to AWS (EC2, ECS, Lambda with API Gateway), Google Cloud (Compute Engine, Cloud Run), Heroku, DigitalOcean.
  - Reverse Proxies: Using Nginx or Caddy in front of your Go application.
- Project Refinement: Clean up your code, add more features to your RESTful API.
- Hands-on:
  - Implement a background task (using a goroutine) in your POST /users endpoint to simulate sending a welcome email after a user is created.
  - Research how to build a Docker image for your Gin application.
  - (Optional) Set up a simple Nginx configuration to act as a reverse proxy for your Gin app.