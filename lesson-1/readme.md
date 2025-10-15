---
title: Getting Started with Go
description: Installing Go and writing your first "Hello World" program
date: 2024-01-15
tags: ['go', 'beginners', 'installation', 'hello-world']
---

# Getting Started with Go

## Installing Go

### Download and Installation

1. **Visit the official Go website**: [golang.org/dl](https://golang.org/dl)

2. **Download the installer** for your operating system:
   - **Windows**: `.msi` installer
   - **macOS**: `.pkg` installer
   - **Linux**: `.tar.gz` archive

3. **Run the installer** and follow the setup instructions

4. **Verify installation** by opening a terminal and running:
```bash
go version
```
You should see output like: go version go1.21.0

5. Setting up Your Workspace
Go uses a specific directory structure for your projects:
```bash
# Create your Go workspace
mkdir -p ~/go/src
mkdir -p ~/go/bin
mkdir -p ~/go/pkg

# Set GOPATH environment variable (usually set automatically in newer versions)
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

6. **Your First Go Program**
Let's create the classic "Hello World" program in Go.
```bash
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

7. Understanding the Code
    1. Package Declaration `package main`
        * Every Go file starts with a package declaration
        * main package creates an executable program
        * Other packages (not main) create reusable libraries
    2. Import Statement
        * import brings in external packages
        * "fmt" is the format package for input/output operations
        * Multiple imports can be grouped
    3. Main Function
        * func declares a function
        * main is the entry point of the program
        * Execution starts from the main function
        * Curly braces {} define the function body
    4. Print Statement
        * fmt package provides formatted I/O
        * Println function prints text with a newline
        * String is enclosed in double quotes " "

8. Creating and Running the Program
    * Run the program: `go run main.go`
    * Build and Run: `go build -o hello-world main.go`, `./hello-world`

