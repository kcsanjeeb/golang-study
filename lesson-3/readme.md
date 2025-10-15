---
title: Variables 
description: Understanding variable declaration, initialization, and scoping in Go
date: 2024-01-19
tags: ['go', 'variables', 'declaration', 'scoping', 'data-types']
---

# Variables 

## Variable Declaration Methods

Go provides multiple ways to declare variables, each with specific use cases.

### 1. Package-Level Variable Declaration

Variables declared outside functions have package-level scope.

```go
package main

import "fmt"

// Package-level variable with explicit type
var myInteger1 int = 10

func main() {
    fmt.Println(myInteger1) // Accessible throughout package
}
```

**Characteristics:**
- Available to all functions in the package
- Must use `var` keyword (no short declaration `:=`)
- Can have explicit type or use type inference

### 2. Basic Variable Declaration with Explicit Type

```go
func main() {
    // Multiple string variables declared together
    var a, b, c string
    
    // Individual assignment
    a = "Hello"
    b = "World" 
    c = "!!"
    
    fmt.Println(a + " " + b + " " + c) // Hello World !!
    
    // Single variable with explicit type
    var myInt int = 10
    var myString string = "hello world !"
    fmt.Println(myInt)     // 10
    fmt.Println(myString) // hello world !
}
```

**Key Points:**
- `var variableName type` syntax
- Can declare first, assign later
- Type must be explicitly specified

### 3. Type Inference (Implicit Typing)

Go can infer the type from the initial value.

```go
func main() {
    // Type inferred as int from value 25
    var age = 25
    fmt.Println(age) // 25
    
    // Multiple variables with type inference
    var year, month = 2025, "Jan"
    fmt.Println(year, month) // 2025 Jan
}
```

**Note:** The compiler determines the type based on the assigned value.

### 4. Short Variable Declaration (Inside Functions)

The `:=` operator provides concise variable declaration inside functions.

```go
func main() {
    // Short declaration - type inferred
    postcode := 55800     // Equivalent to: var postcode int = 55800
    address := "shenzhen" // Equivalent to: var address string = "shenzhen"
    fmt.Println(postcode, address) // 55800 shenzhen
    
    // Multiple short declarations
    otp_code, language := 1234, "Chinese"
    // Equivalent to:
    // var otp_code int = 1234
    // var language string = "Chinese"
    fmt.Println(language, otp_code) // Chinese 1234
}
```

**Important Rules for Short Declaration:**
- Only works inside functions
- Must initialize at declaration
- Type is automatically inferred
- At least one variable must be new in redeclaration

## Variable Scoping

### Package Scope vs Function Scope

```go
package main

import "fmt"

// Package-level variable
var myInteger1 int = 10

func main() {
    // Function-level variable
    var localVar string = "I'm local"
    fmt.Println(myInteger1) // Can access package variable
    fmt.Println(localVar)   // Can access local variable
}

func something() {
    fmt.Println("Something: ", myInteger1) // Can access package variable
    // fmt.Println(localVar) // ERROR: localVar not accessible here
}
```

**Scope Rules:**
- **Package-level**: Accessible throughout the package
- **Function-level**: Only accessible within the function

## Multiple Variable Declaration

### Different Patterns

```go
func main() {
    // Method 1: Same type, individual declaration
    var a, b, c string
    a, b, c = "Hello", "World", "!!"
    
    // Method 2: Different types with explicit assignment
    var year int = 2025
    var month string = "Jan"
    
    // Method 3: Multiple variables with type inference
    var score, name = 95, "Alice"
    
    // Method 4: Short declaration for multiple variables
    x, y, z := 10, "hello", true
    
    fmt.Println(a, b, c, year, month, score, name, x, y, z)
}
```

## Best Practices

### 1. Choose Appropriate Declaration Style

```go
func main() {
    // Use var when type is important to specify
    var port int = 8080
    
    // Use short declaration for local, obvious types
    count := 10
    message := "Processing..."
    
    // Use explicit type for clarity with complex types
    var coordinates [2]float64 = [2]float64{40.7128, -74.0060}
}
```

### 2. Group Related Variables

```go
func main() {
    // Group related variables
    var (
        username string = "john_doe"
        age      int    = 30
        email    string = "john@example.com"
    )
    
    // Or use multiple short declaration
    firstName, lastName := "John", "Doe"
    
    fmt.Println(username, age, email, firstName, lastName)
}
```

### 3. Meaningful Variable Names

```go
func main() {
    // Good - descriptive names
    userCount := 150
    maxRetries := 3
    apiEndpoint := "https://api.example.com"
    
    // Avoid - unclear names
    // x := 150
    // y := 3
    // z := "https://api.example.com"
}
```

## Common Patterns and Examples

### Function Return Values

```go
func getUserInfo() (string, int, bool) {
    return "Alice", 25, true
}

func main() {
    // Capture multiple return values
    name, age, isActive := getUserInfo()
    fmt.Println(name, age, isActive)
    
    // Ignore some values using blank identifier
    username, _, _ := getUserInfo()
    fmt.Println(username)
}
```

### Swapping Variables

```go
func main() {
    a, b := 10, 20
    fmt.Printf("Before: a=%d, b=%d\n", a, b)
    
    // Swap without temporary variable
    a, b = b, a
    fmt.Printf("After: a=%d, b=%d\n", a, b)
}
```

## Zero Values

Variables declared without initialization get zero values:

```go
func main() {
    var i int     // 0
    var f float64 // 0.0
    var s string  // ""
    var b bool    // false
    
    fmt.Println(i, f, s, b) // 0 0  false
}
```

## Summary Table

| Declaration Type | Syntax | Scope | Type Specification |
|------------------|--------|-------|-------------------|
| Package-level | `var name type = value` | Package | Explicit |
| Function explicit | `var name type = value` | Function | Explicit |
| Type inference | `var name = value` | Function | Inferred |
| Short declaration | `name := value` | Function | Inferred |

---
