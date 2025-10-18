---
title: Panic and Recovery in Go
description: Understanding panic, recovery mechanisms, and proper error handling
date: 2024-02-14
tags: ['go', 'panic', 'recover', 'error-handling', 'defer']
---

# Panic and Recovery in Go

## 🚨 Panic

### What is Panic?
Panic occurs when a Go program reaches an unrecoverable state. It's similar to exceptions in other languages but should be used sparingly.

### Causes of Panic:
- **Runtime errors**: Array out-of-bounds, nil pointer dereference
- **Explicit calls**: Using `panic()` function
- **Critical failures**: When program cannot continue

### Panic Behavior:
- Stops normal execution immediately
- Begins unwinding the call stack
- Executes deferred functions during unwinding
- Terminates program if not recovered

### Panic Example with Deferred Functions:
```go
package main

import "fmt"

func function1() {
    defer func() {
        fmt.Println("Function 1 deferred function called.")
    }()
    function2()
}

func function2() {
    defer func() {
        fmt.Println("Function 2 deferred function called.")
    }()
    panic("Function 2 panicked!")
}

func main() {
    defer func() {
        fmt.Println("Function Main deferred function called.")
    }()
    function1()
}
```

### Execution Flow:
1. **main()** registers defer → calls function1()
2. **function1()** registers defer → calls function2()
3. **function2()** registers defer → calls panic()
4. **Stack unwinding begins**:
   - function2's defer executes
   - function1's defer executes  
   - main's defer executes
5. **Program terminates** with panic message

### Output:
```
Function 2 deferred function called.
Function 1 deferred function called.
Function Main deferred function called.
panic: Function 2 panicked!
...
```

## 🛡️ Recovery

### The recover() Function
`recover()` allows you to regain control after a panic and prevent program termination.

### Key Rules:
- **Must be called within a deferred function**
- **Returns the value passed to panic()**
- **Returns nil if no panic occurred**

### Basic Recovery Example:
```go
package main

import "fmt"

func panicExample() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    panic("Something went wrong!")
}

func main() {
    fmt.Println("Program has started")
    panicExample()
    fmt.Println("Exited normally!")
}
```

### Output:
```
Program has started
Recovered from panic: Something went wrong!
Exited normally!
```

## 🎯 Proper Use of Recovery

### Recovery Best Practices:
- **Use for graceful shutdowns**, not as exception handling
- **Log the panic** before recovering for debugging
- **Return meaningful errors** instead of silent recovery
- **Use at package boundaries** to prevent panics from escaping

### Practical Recovery Pattern:
```go
package main

import (
    "fmt"
    "log"
)

func safeOperation() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("operation failed: %v", r)
            log.Printf("Recovered from panic: %v", r)
        }
    }()
    
    // Simulate risky operation
    riskyOperation()
    return nil
}

func riskyOperation() {
    // This could panic in real scenarios
    panic("database connection failed")
}

func main() {
    if err := safeOperation(); err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Operation completed successfully")
    }
}
```

## ⚠️ Important Considerations

### When to Use Panic:
- **True unrecoverable errors** (corrupted state, missing critical files)
- **Programming errors** (should never happen in correct code)
- **Startup failures** (missing configuration, failed dependencies)

### When to Use Recovery:
- **Third-party library panics** that you can't control
- **Graceful server shutdown** in HTTP handlers
- **Background goroutines** to prevent one from taking down entire application

### Anti-Patterns to Avoid:
```go
// ❌ DON'T: Using recover like try-catch
func badExample() {
    defer func() {
        recover() // Silent recovery - dangerous!
    }()
    panic("error")
}

// ✅ DO: Proper error handling with recovery
func goodExample() error {
    defer func() {
        if r := recover(); r != nil {
            // Convert panic to meaningful error
            return fmt.Errorf("recovered from: %v", r)
        }
    }()
    // Your code here
}
```

## 🔧 Real-World Example

### HTTP Server with Panic Recovery:
```go
package main

import (
    "fmt"
    "net/http"
)

func safeHandler(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if r := recover(); r != nil {
                fmt.Printf("Recovered from panic in handler: %v\n", r)
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            }
        }()
        handler(w, r)
    }
}

func riskyHandler(w http.ResponseWriter, r *http.Request) {
    // This might panic under certain conditions
    panic("database query failed")
}

func main() {
    http.HandleFunc("/", safeHandler(riskyHandler))
    fmt.Println("Server starting on :8080")
    http.ListenAndServe(":8080", nil)
}
```

## 📋 Key Takeaways

| Concept | Purpose | When to Use |
|---------|---------|-------------|
| **Panic** | Signal unrecoverable error | Critical failures, programming errors |
| **Recover** | Regain control after panic | Graceful shutdown, third-party libraries |
| **Defer** | Ensure cleanup runs | Always with recover, resource cleanup |

### Remember:
- **Panic is for catastrophic failures**, not routine errors
- **Recovery should be strategic**, not blanket coverage  
- **Always prefer returning errors** over using panic/recover
- **Use defer for cleanup** that must run even during panics
