# A JSON Formatter Command Line Interface (CLI) in Go is a useful tool for pretty-printing JSON data, making it more readable. It can accept JSON input from a file or standard input (stdin) and output a formatted version.

---

## 🛠️ Go JSON Formatter CLI

This Go application will be a simple command-line interface tool that takes JSON input and outputs a formatted, pretty-printed version.

### 📁 Project Structure

```
json-formatter-cli/
├── main.go
└── go.mod
```

### 📋 Prerequisites

* Go installed (version 1.16 or higher recommended)

### 🚀 Step-by-Step Implementation

#### 1. Initialize Go Module

First, create a new directory for your project and initialize a Go module:

```bash
mkdir json-formatter-cli
cd json-formatter-cli
go mod init json-formatter-cli
```

#### 2. Create `main.go`

Now, create a `main.go` file inside the `json-formatter-cli` directory and paste the following code:

```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"encoding/json"
)

func main() {
	// Determine input source: file or stdin
	var inputReader io.Reader
	var inputFilename string

	if len(os.Args) > 1 {
		// If an argument is provided, treat it as a filename
		inputFilename = os.Args[1]
		file, err := os.Open(inputFilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", inputFilename, err)
			os.Exit(1)
		}
		defer file.Close()
		inputReader = file
	} else {
		// No argument provided, read from stdin
		inputReader = os.Stdin
	}

	// Read all content from the input source
	inputBytes, err := io.ReadAll(inputReader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Unmarshal the JSON data into a generic interface{}
	// This allows us to handle any valid JSON structure (object or array)
	var data interface{}
	err = json.Unmarshal(inputBytes, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Marshal the data back into JSON with indentation for pretty-printing
	// The prefix is for the beginning of each line, and indent for indentation.
	// We use an empty prefix and 4 spaces for indentation.
	formattedJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error formatting JSON: %v\n", err)
		os.Exit(1)
	}

	// Print the formatted JSON to standard output
	fmt.Println(string(formattedJSON))
}
```

### 📝 Code Explanation

* **`package main` and `func main()`:** This is the entry point of our Go CLI application.
* **`import` statements:** We import necessary packages:
    * `bytes`: Used for byte buffer operations (though `io.ReadAll` handles this cleanly now).
    * `fmt`: For formatted I/O (printing to console).
    * `io`: For input/output operations (reading from files or stdin).
    * `os`: For interacting with the operating system (accessing command-line arguments, stdin/stdout, exiting).
    * `encoding/json`: The core package for JSON encoding and decoding.
* **Input Source Determination:**
    * The `if len(os.Args) > 1` block checks if a command-line argument is provided. If so, it assumes the argument is a **filename** and attempts to open it.
    * If no argument is given, it defaults to reading from **`os.Stdin`** (standard input), which allows piping JSON data to the CLI.
* **Reading Input:**
    * `io.ReadAll(inputReader)` reads all content from the chosen input source (file or stdin) into a `[]byte` slice.
* **JSON Parsing (`json.Unmarshal`):**
    * `var data interface{}`: We declare a variable `data` of type `interface{}`. This is a powerful Go feature that allows `json.Unmarshal` to parse any valid JSON structure (it could be a JSON object, array, string, number, boolean, or null) without needing a predefined struct. The JSON decoder will automatically represent the JSON as `map[string]interface{}` for objects or `[]interface{}` for arrays.
    * `json.Unmarshal(inputBytes, &data)`: This decodes the raw JSON bytes into the `data` variable. If the JSON is invalid, an error is returned.
* **JSON Formatting (`json.MarshalIndent`):**
    * `json.MarshalIndent(data, "", " ")`: This is the magic function for pretty-printing.
        * The first argument (`data`) is the Go data structure to be encoded.
        * The second argument (`""`) is the **prefix** to be added to each line of the output. We use an empty string as we don't need a prefix.
        * The third argument (`" "`) is the **indent** string to use for indentation. Using ` ` (four spaces) is a common convention for JSON formatting.
    * This function returns the formatted JSON as a `[]byte` slice and an error if formatting fails.
* **Output:**
    * `fmt.Println(string(formattedJSON))`: The formatted `[]byte` slice is converted back to a string and printed to `os.Stdout` (standard output), which appears in the terminal.
* **Error Handling:** Basic error handling is included using `fmt.Fprintf(os.Stderr, ...)` to print errors to standard error and `os.Exit(1)` to exit with a non-zero status code, indicating an error.

### ▶️ How to Build and Run

1.  **Navigate to the project directory:**
    ```bash
    cd json-formatter-cli
    ```
2.  **Build the executable:**
    ```bash
    go build -o jsonfmt
    ```
    This command compiles `main.go` and creates an executable file named `jsonfmt` (or `jsonfmt.exe` on Windows) in the current directory.

### 🧪 How to Test

You can test the `jsonfmt` CLI in a few ways:

#### 1. Formatting from a File

Create a sample JSON file, e.g., `sample.json`:

```json
{"name":"Alice","age":30,"isStudent":false,"courses":[{"title":"Math","grade":"A"},{"title":"Physics","grade":"B"}],"address":null}
```

Now, run the `jsonfmt` command with the filename:

```bash
./jsonfmt sample.json
```

**Expected Output:**

```json
{
    "address": null,
    "age": 30,
    "courses": [
        {
            "grade": "A",
            "title": "Math"
        },
        {
            "grade": "B",
            "title": "Physics"
        }
    ],
    "isStudent": false,
    "name": "Alice"
}
```

#### 2. Formatting from Standard Input (stdin)

You can pipe JSON directly into the `jsonfmt` command:

```bash
echo '{"product":"Laptop","price":1200,"specs":{"cpu":"i7","ram":16},"features":["fast","light"]}' | ./jsonfmt
```

**Expected Output:**

```json
{
    "features": [
        "fast",
        "light"
    ],
    "price": 1200,
    "product": "Laptop",
    "specs": {
        "cpu": "i7",
        "ram": 16
    }
}
```

You can also paste JSON directly into the terminal after running `./jsonfmt` without any arguments, then press `Ctrl+D` (on Linux/macOS) or `Ctrl+Z` then `Enter` (on Windows) to signal end-of-input.

#### 3. Handling Invalid JSON

```bash
echo '{"key":"value", invalid}' | ./jsonfmt
```

**Expected Output (to stderr):**

```
Error parsing JSON: invalid character 'i' looking for beginning of object key string
```

---

### 🚀 Future Improvements

This is a functional basic JSON formatter, but you could enhance it with:

* **Custom Indentation:** Allow users to specify the number of spaces or tabs for indentation (e.g., `./jsonfmt -indent 2`).
* **Compact Output:** Add an option to output compact JSON without any extra whitespace.
* **Syntax Highlighting:** Integrate a library for colorizing the output JSON.
* **Error Reporting:** More detailed error messages, possibly highlighting the location of parsing errors.
* **Input Validation:** Add more checks for common JSON issues.

This CLI provides a great starting point for building more complex command-line tools in Go!