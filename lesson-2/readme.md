---
title: Primitive Data Types in Go
description: A Comprehensive Guide to Booleans, Strings, and Numbers
date: 2025-10-15
tags: ['go', 'variables', 'data-types', 'constants', 'type-system']
---

# Primitive Data Types in Go

## A Comprehensive Guide to Booleans, Strings, and Numbers

## Go's Type System

Go is a **statically typed** and **strongly typed** language, which means:

- **Statically Typed**: Every variable and constant must have a type known at compile time
- **Strongly Typed**: The compiler enforces type safety and won't compile code with type mismatches

```go
// This will cause compile-time errors if types don't match
var age int = "twenty"  // Error: cannot use "twenty" (type string) as type int
```


## Primitive Data Types

Primitive data types are the basic building blocks from which all other data types are constructed.

### Basic Types

#### Integer Types
```go
var a int = 42        // Platform-dependent size (32 or 64 bits)
var b int32 = 100     // 32-bit integer
var c int64 = 1000    // 64-bit integer
var d uint = 50       // Unsigned integer
var e uint8 = 255     // 8-bit unsigned integer (0-255)
```

#### Floating-Point Types
```go
var f float32 = 3.14    // 32-bit floating point
var g float64 = 2.71828 // 64-bit floating point (default for floating-point literals)
```

#### Boolean Type
```go
var isReady bool = true
var isFinished bool = false
```

#### String Type
```go
var name string = "Go Programming"
var message string = `This is a 
multi-line string`
```

#### Byte and Rune Types
```go
var ch byte = 'A'           // alias for uint8, represents ASCII characters
var unicodeChar rune = '世' // alias for int32, represents Unicode code points
```

### Zero Values

In Go, variables declared without an initial value are given their **zero value**:

```go
var i int     // 0
var f float64 // 0.0
var b bool    // false
var s string  // ""
```


## Constants

Constants are immutable values declared with the `const` keyword:

```go
const Pi = 3.14159
const MaxSize = 100
const AppName = "MyGoApp"

// Multiple constants
const (
    StatusOK = 200
    StatusNotFound = 404
    ServerPort = 8080
)
```


### Typed vs Untyped Constants

```go
const TypedConst int = 42        // Typed constant
const UntypedConst = 3.14        // Untyped constant (can be used with different numeric types)

var x int = UntypedConst         // Works: 3.14 is truncated to 3
var y float64 = UntypedConst     // Works: 3.14 remains as float64
```

## Variable Declaration

### Explicit Declaration
```go
var count int
var username string = "john"
var isActive bool = true
```

### Type Inference
```go
var count = 10           // Type inferred as int
var name = "Go"          // Type inferred as string
var price = 19.99        // Type inferred as float64
```

### Short Variable Declaration (Inside functions)
The **short variable declaration** uses the `:=` operator and is one of Go's most frequently used features for variable declaration inside functions.
#### Key Characteristics
1. **Function-scoped**: Only available inside function bodies
2. **Type inference**: Compiler automatically determines the type
3. **Concise syntax**: More compact than `var` declaration
4. **Mandatory initialization**: Must assign a value at declaration
```go
func main() {
    count := 10          // Type inferred as int
    message := "Hello"   // Type inferred as string
    ratio := 3.14        // Type inferred as float64
    
    // Multiple variables
    x, y := 10, 20
    name, age := "Alice", 30
}
```
### Multiple Variable Declaration
```go
var (
    name string = "John"
    age int = 25
    salary float64 = 50000.0
)

// Alternative syntax
var x, y, z int = 1, 2, 3
```