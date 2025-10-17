---
title: Defer Statement in Go
description: Deferred function execution and its three trigger conditions
date: 2024-02-04
tags: ['go', 'defer', 'cleanup', 'execution-flow']
---

# Defer Statement in Go

## What is Defer?

The `defer` statement postpones the execution of a function until the surrounding function completes. Deferred functions are commonly used for cleanup tasks, resource management, and panic handling.

## Three Conditions for Deferred Function Execution

Deferred functions execute when ANY of these conditions occur:

### 1. Surrounding Function Returns
When the function containing the defer statement completes normally with a return statement.

### 2. End of Function Body
When execution reaches the closing brace `}` of the function, even without an explicit return.

### 3. Function Panics
When the function panics, deferred functions run before the panic propagates upward.
```go
package main

import "fmt"

func main() {
    fmt.Println("Start of main function")
    
    // This deferred function will run even if we panic
    defer fmt.Println("Deferred: This runs BEFORE panic propagates")
    
    fmt.Println("About to panic...")
    
    // Cause a panic
    panic("Something went wrong!")
    
    // This line never executes
    fmt.Println("This will never print")
    
    // Deferred function executes here automatically before panic propagates
}
```

## Key Defer Behavior

- **Arguments evaluated immediately** - Values are captured when defer is called
- **LIFO order** - Multiple defers execute in reverse order (Last In, First Out)
- **Execute after return values set** - Deferred functions run after return values are determined but before returning to caller
- **Common uses** - Resource cleanup, file closing, mutex unlocking, panic recovery

### Simple Explanation:
* Defer statements are called in reverse order (Last In, First Out)
* They execute after the main function finishes
* No need to manually call them - Go handles it automatically

### LIFO Order 
```go 
package main

import "fmt"

func main() {
    fmt.Println("=== Starting Database Transaction ===")
    
    // Simulate multiple cleanup operations
    defer fmt.Println("5. Transaction COMMITTED")      // Last in - First out (5th)
    defer fmt.Println("4. Releasing connection")       // Fourth in - Second out
    defer fmt.Println("3. Closing temporary files")    // Third in - Third out  
    defer fmt.Println("2. Unlocking user records")     // Second in - Fourth out
    defer fmt.Println("1. Logging operation complete") // First in - Last out (1st)
    
    fmt.Println("Working on database operations...")
    fmt.Println("About to finish function...")
    
    // All deferred functions execute HERE in REVERSE order
}
```