---
title: The init Function
description: Package initialization and special init function behavior
date: 2024-02-11
tags: ['go', 'init', 'package-initialization', 'special-functions']
---

# The init Function

## What is the init Function?

The `init` function is a **special reserved function** in Go that executes automatically during package initialization. It's used for package-level setup and configuration.

### Key Characteristics
- **No parameters**: Accepts no arguments
- **No return values**: Returns nothing
- **Automatic execution**: Called implicitly when package is imported
- **Multiple allowed**: Can have multiple init functions per package

## Execution Order

### Init Function Sequence
1. **Package-level init functions** in imported packages
2. **Current package init functions** in source order
3. **Main function** execution

### Execution Rules
- Init functions run **before main()**
- Multiple init functions execute in **source code order**
- Imported package init functions run **first**

## Practical Examples

### Basic Multiple Init Functions
~~~go
package main

import "fmt"

func init() {
    fmt.Println("Init 1 called")
}

func init() {
    fmt.Println("Init 2 called")
}

func init() {
    fmt.Println("Init 3 called")
}

func main() {
    fmt.Println("Main function called")
}
~~~

**Output:**
* Init 1 called
* Init 2 called
* Init 3 called
* Main function called

### Package Initialization with Underscore Import
```go
// main.go
package main

import (
    "fmt"
    _ "github.com/kcsanjeeb/golang-study/lesson-29/init"
)

func init() {
    fmt.Println("Main package init called")
}

func main() {
    fmt.Println("Main function called")
}
```

```go
// init/init.go
package init

import "fmt"

func init() {
    fmt.Println("Init package init function called")
}
```

**Output:**
- Init package init function called
- Main package init called
- Main function called


## Common Use Cases

### 1. Database Connection Setup
```go
package database

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq" // Driver registration via init
)

var DB *sql.DB

func init() {
    var err error
    DB, err = sql.Open("postgres", "connection-string")
    if err != nil {
        panic(fmt.Sprintf("Failed to connect to database: %v", err))
    }
    
    fmt.Println("Database connection established")
}
```

### 2. Configuration Loading
```go
package config

import (
    "encoding/json"
    "os"
)

var AppConfig Config

type Config struct {
    Port     string `json:"port"`
    Database string `json:"database"`
}

func init() {
    file, err := os.Open("config.json")
    if err != nil {
        panic("Config file not found")
    }
    defer file.Close()
    
    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&AppConfig); err != nil {
        panic("Invalid config format")
    }
}
```

### 3. Registration Patterns
```go
package plugins

import "fmt"

var plugins = make(map[string]Plugin)

type Plugin interface {
    Name() string
}

func Register(p Plugin) {
    plugins[p.Name()] = p
}

func GetPlugin(name string) Plugin {
    return plugins[name]
}
```

```go
package myplugin

import "github.com/kcsanjeeb/golang-study/lesson-29/plugins"

type MyPlugin struct{}

func (p MyPlugin) Name() string {
    return "myplugin"
}

func init() {
    plugins.Register(MyPlugin{})
    fmt.Println("MyPlugin registered")
}
```

## Advanced Patterns

### Conditional Initialization
```go
package main

import (
    "fmt"
    "os"
)

func init() {
    if os.Getenv("DEBUG") == "true" {
        fmt.Println("Debug mode enabled")
        // Initialize debug handlers
    }
}
```

### Init with Build Tags
```go
// +build development

package config

func init() {
    fmt.Println("Development configuration loaded")
}
```

```go
// +build production

package config

func init() {
    fmt.Println("Production configuration loaded")
}
```

## Important Considerations

### Execution Order Dependencies
```go
package main

import "fmt"

var globalVar = initializeVar()

func initializeVar() string {
    fmt.Println("Global variable initialization")
    return "value"
}

func init() {
    fmt.Println("Init function - can use globalVar:", globalVar)
}

func main() {
    fmt.Println("Main function")
}
```

**Output:**
* Global variable initialization
* Init function - can use globalVar: value
* Main function


### Best Practices

1. **Keep init functions simple** - avoid complex logic
2. **Use for essential setup** only (DB connections, config loading)
3. **Avoid side effects** that aren't obvious to package users
4. **Document init behavior** if it has significant impact
5. **Consider explicit initialization** for complex setups

### Anti-Patterns to Avoid

```go
// ❌ DON'T: Complex business logic in init
func init() {
    // Processing user data, making API calls, etc.
}

// ✅ DO: Simple setup and registration
func init() {
    // Register handlers, load config, setup connections
}
```

## Real-World Example

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/lib/pq" // Driver registers itself via init
)

var db *sql.DB

func init() {
    // Database initialization
    var err error
    db, err = sql.Open("postgres", "user=postgres dbname=mydb sslmode=disable")
    if err != nil {
        log.Fatal("Database connection failed:", err)
    }
    
    // Configuration validation
    if err := db.Ping(); err != nil {
        log.Fatal("Database ping failed:", err)
    }
    
    fmt.Println("Database initialized successfully")
}

func main() {
    // Use the initialized database
    rows, err := db.Query("SELECT * FROM users")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    
    fmt.Println("Application running with database")
}
```

## Key Takeaways

- **Init functions** run automatically before main()
- **Execution order**: Imported packages → Current package → main()
- **Multiple init functions** execute in source order
- **Use for setup** not for business logic
- **Underscore imports** trigger package initialization without direct usage

