# Go Commands Cheat Sheet 📌

This cheat sheet provides a list of commonly used Go commands related to module management, dependency handling, building, running, and debugging Go programs.

## 🔹 Installation and Dependencies

### 1. Install or Update a Go Package
```sh
$ go get -u github.com/gorilla/mux
```
- Installs or updates the `github.com/gorilla/mux` package.
- The `-u` flag updates dependencies to the latest minor or patch version.

### 2. Clean and Optimize Dependencies
```sh
$ go mod tidy
```
- Removes unused dependencies and ensures `go.mod` and `go.sum` are up to date.

### 3. Verify Dependencies
```sh
$ go mod verify
```
- Checks integrity and authenticity of modules in `go.sum`.

### 4. View the Dependency Graph
```sh
$ go mod graph
```
- Displays module dependencies in a tree format.

### 5. Check Why a Specific Dependency is Used
```sh
$ go mod why github.com/gorilla/mux
```
- Shows why a package is included in the project.

### 6. Create a `vendor` Directory
```sh
$ go mod vendor
```
- Copies dependencies into a `vendor` directory for offline builds.

## 🔹 Building and Running Go Applications

### 7. Build a Go Program
```sh
$ go build .
```
- Compiles the Go source code in the current directory into an executable binary.

### 8. Run a Go Program
```sh
$ go run main.go
```
- Compiles and runs `main.go` in a single command.

### 9. Compile and Execute Multiple Files
```sh
$ go run .
```
- Runs all `.go` files in the current directory.

## 🔹 Listing Packages and Modules

### 10. List All Packages in the Current Module
```sh
$ go list
```
- Displays the import path of the main package.

### 11. List All Available Packages
```sh
$ go list all
```
- Shows all Go packages, including standard and external dependencies.

### 12. List All Modules Used in the Project
```sh
$ go list -m all
```
- Lists all modules in the project, including indirect dependencies.

## 🔹 Code Formatting and Linting

### 13. Format Go Code
```sh
$ go fmt ./...
```
- Formats all `.go` files in the current directory and subdirectories.

### 14. Check Code for Errors
```sh
$ go vet ./...
```
- Reports potential issues in Go code.

## 🔹 Testing

### 15. Run Tests
```sh
$ go test ./...
```
- Runs all tests in the current module.

### 16. Run Tests with Detailed Output
```sh
$ go test -v ./...
```
- Runs tests with verbose output.

### 17. Run Benchmarks
```sh
$ go test -bench .
```
- Runs benchmark tests in the project.

## 🔹 Working with Executables

### 18. Install a Go Binary
```sh
$ go install ./...
```
- Compiles and installs the binary to `$GOBIN`.

### 19. Print Go Environment Variables
```sh
$ go env
```
- Displays Go-related environment variables.

### 20. Check Go Version
```sh
$ go version
```
- Displays the installed Go version.

## 🔹 Debugging

### 21. Debug a Go Program (Using Delve)
```sh
$ dlv debug
```
- Starts the Go debugger (requires Delve installation).

### 22. Profile CPU Usage
```sh
$ go test -cpuprofile cpu.prof -bench .
```
- Profiles CPU usage while running benchmarks.

---
This cheat sheet provides essential Go commands to enhance your development workflow! 🚀

