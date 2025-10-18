---
title: Dependency Management with Go Modules
description: Complete guide to managing dependencies using Go modules
date: 2024-02-13
tags: ['go', 'modules', 'dependencies', 'go-mod', 'go-sum']
---

# Dependency Management with Go Modules

## 📦 What are Go Modules?

Go modules are the standard dependency management system that helps you track and manage external packages your project uses.

### Key Components:
- **go.mod**: Defines your module and its dependencies
- **go.sum**: Security lock file with cryptographic hashes
- **One module per repository** is the recommended practice

## 🚀 Initializing a Go Module

### Command:
```bash
go mod init github.com/your-username/project-name
```

### Example:
```bash
go mod init github.com/kcsanjeeb/first-module
```

### Generated go.mod:
```go
module github.com/kcsanjeeb/first-module

go 1.25.1
```

## 📥 Adding Dependencies

### Using `go get`:
```bash
go get github.com/fatih/color
```

### What happens:
- Downloads the module and its dependencies
- Updates `go.mod` with required modules
- Generates `go.sum` with security hashes

### Example Project:

#### main.go
```go
package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	str := color.CyanString("Hello World!")
	fmt.Println(str)
}
```

#### Updated go.mod
```go
module github.com/kcsanjeeb/first-module

go 1.25.1

require (
	github.com/fatih/color v1.18.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.25.0 // indirect
)
```

## 🔒 Understanding go.sum

The `go.sum` file acts as a security lock file:

### Purpose:
- Stores cryptographic hashes for each module version
- Ensures same versions are downloaded every time
- Prevents tampering with dependencies

### Example go.sum content:
```
github.com/fatih/color v1.18.0 h1:S8gINlzdQ840/4pfAwic/ZE0djQEH3wM94VfqLTZcOM=
github.com/fatih/color v1.18.0/go.mod h1:4FelSpRwEGDpQ12mAdzqdOukCy4u8WUtOY6lkT/6HfU=
```

## 📋 Listing Dependencies

### View all imported packages:
```bash
go list all
```

### View only modules (with -m flag):
```bash
go list -m all
```

### Sample output shows:
- Your main module
- Direct dependencies
- Transitive dependencies (dependencies of your dependencies)
- Standard library packages

## 🧹 Cleaning Up with go mod tidy

### What `go mod tidy` does:
1. **Adds missing dependencies** that are imported but not in go.mod
2. **Removes unused dependencies** that are in go.mod but not used
3. **Updates go.sum** with correct hashes

### Before cleanup:
```go
package main

func main() {
    // Commented out dependency usage
    // str := color.CyanString("Hello World!")
    // fmt.Println(str)
}
```

### Run cleanup:
```bash
go mod tidy
```

### Result:
- Unused `color` dependency removed from go.mod
- go.sum updated accordingly

## 💡 Best Practices

### 1. Always commit both files:
```bash
git add go.mod go.sum
```

### 2. Run tidy before commits:
```bash
go mod tidy
```

### 3. Check for updates:
```bash
go list -m -u all
```

### 4. Vendor dependencies (optional):
```bash
go mod vendor
```

## 🎯 Quick Reference

| Command | Purpose |
|---------|---------|
| `go mod init` | Create new module |
| `go get package` | Add dependency |
| `go mod tidy` | Clean up dependencies |
| `go list -m all` | List all modules |
| `go mod vendor` | Create vendor directory |

