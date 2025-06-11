# 📚 Bookstore API

## Table of Contents
- Introduction
- Features
- Technologies Used
- Project Structure
  - Getting Started
  - Prerequisites
  - Installation
  - Running the Application
- API Endpoints
- Error Handling
- Future Enhancements

## Introduction ##
This project is a RESTful API for a Bookstore built with Go. It provides a set of endpoints to manage book records, including operations to create, read, update, and delete books. The API is designed with a layered architecture (handlers, services, repositories) to ensure maintainability, testability, and separation of concerns.

## Features
- Create Book: Add new book records to the database.
- Get All Books: Retrieve a list of all available books.
- Get Book by ID: Fetch details of a specific book using its unique ID.
- Update Book: Modify existing book details.
- Delete Book: Remove a book record from the database.
- Standardized Error Handling: Consistent JSON error responses for API clients.

## Technologies Used
- Go: The primary programming language.
- Gorilla Mux: A powerful HTTP router and URL matcher for building web services.
- GORM: An excellent ORM (Object-Relational Mapper) library for Go, used for database interactions.
- PostgreSQL: The relational database used to store book information.
- Godotenv (Optional): For loading environment variables from a .env file during local development.

## Project Structure
The project follows a clean and modular structure to enhance organization and scalability:

```text
bookstore-api/
├── main.go               // Application entry point
├── config/               // Database connection and general application configurations
│   └── config.go
├── models/               // Structs defining data models (e.g., Book)
│   └── book.go
├── handlers/             // HTTP request handlers (controllers)
│   └── book_handler.go
├── repository/           // Data Access Layer (interfaces and GORM implementations)
│   └── book_repository.go
├── services/             // Business Logic Layer (implements core business rules)
│   └── book_service.go
├── router/               // API routing definitions
│   └── router.go
├── utils/                // Utility functions (e.g., standardized error handling)
│   └── error_handler.go
├── go.mod                // Go module dependencies
├── go.sum
└── .env.example          // Example environment variables file
```

## Getting Started
Follow these steps to set up and run the Bookstore API on your local machine.

### Prerequisites
Before you begin, ensure you have the following installed:

- Go: Version 1.18 or higher.
- PostgreSQL: A running PostgreSQL instance.
### Installation
- Clone the repository:
```bash
git clone git@github.com:tqnhu4/go-projects.git
cd 16.Bookstore
```

- Install Go modules:
```bash
go mod tidy
```
- Set up your PostgreSQL database:
  - Create a new database (e.g., bookstore):
  ```sql
  CREATE DATABASE bookstore;
  ```

  - CREATE DATABASE bookstore;

- Create a .env file:
Create a file named .env in the root directory of the project, based on .env.example, and fill in your database credentials:

```text
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_postgres_user
DB_PASSWORD=your_postgres_password
DB_NAME=bookstore
PORT=8080
```

  - Replace your_postgres_user and your_postgres_password with your actual PostgreSQL username and password.
  - You can change the PORT if 8080 is already in use.

### Running the Application
Once everything is set up, you can run the application:  

```bash
go run main.go
```

You should see output indicating that the database is connected and the server is starting:

```text
Database connected successfully!
Database auto-migration complete.
Server starting on port 8080...
```

The API will now be accessible at http://localhost:8080 (or your chosen port).

### 📚 API Endpoints


| Method | Endpoint       | Description              | Request Body (JSON)                                                                                      | Response (JSON)                             |
|--------|----------------|--------------------------|----------------------------------------------------------------------------------------------------------|---------------------------------------------|
| GET    | `/books`       | Get all books            | None                                                                                                     | `[{"id": 1, "title": "...", ...}]`          |
| GET    | `/books/{id}`  | Get a book by ID         | None                                                                                                     | `{"id": 1, "title": "...", ...}`            |
| POST   | `/books`       | Create a new book        | `{"title": "...", "author": "...", "isbn": "...", "published_date": "YYYY-MM-DDTHH:MM:SSZ", "price": 0.0, "stock": 0}` | `{"id": 1, "title": "...", ...}`            |
| PUT    | `/books/{id}`  | Update an existing book  | `{"title": "...", "author": "...", "isbn": "...", "published_date": "YYYY-MM-DDTHH:MM:SSZ", "price": 0.0, "stock": 0}` | `{"id": 1, "title": "...", ...}`            |
| DELETE | `/books/{id}`  | Delete a book by ID      | None                                                                                                     | `204 No Content` (on success)               |

Example POST /books request body:

```json
{
    "title": "The Hitchhiker's Guide to the Galaxy",
    "author": "Douglas Adams",
    "isbn": "978-0345391803",
    "published_date": "1979-10-12T00:00:00Z",
    "price": 12.99,
    "stock": 100
}
```
### Error Handling
The API provides standardized JSON error responses to make error handling on the client side more predictable.

Example Error Response:
```json
{
    "message": "Book not found",
    "status_code": 404,
    "code": "BOOK_NOT_FOUND"
}
```
Common error codes and their corresponding HTTP status codes:

- 400 Bad Request:
  - INVALID_ID_FORMAT: Provided ID is not a valid number.
  - INVALID_REQUEST_BODY: Request body is malformed or invalid JSON.
  - MISSING_REQUIRED_FIELDS: Essential fields (e.g., title, author, ISBN) are missing.
- 404 Not Found:
  - BOOK_NOT_FOUND: No book found with the given ID.
- 409 Conflict:
  - DUPLICATE_ISBN: A book with the provided ISBN already exists.
- 500 Internal Server Error:
  - INTERNAL_ERROR: An unexpected server-side error occurred.


### Future Enhancements
- Advanced Validation: Implement more robust input validation for all fields.
- Authentication & Authorization: Secure API endpoints using JWT or other methods.
- Pagination & Filtering: Add query parameters for retrieving books with pagination and filtering options (e.g., by author, published year).
- Unit & Integration Tests: Comprehensive testing to ensure reliability.
- Structured Logging: Implement better logging with libraries like logrus or zap.
- Graceful Shutdown: Ensure the server shuts down gracefully.
- Dockerization: Provide Dockerfile and docker-compose.yml for easy deployment.
- CORS Configuration: Configure Cross-Origin Resource Sharing if the API is consumed by a different domain.  
