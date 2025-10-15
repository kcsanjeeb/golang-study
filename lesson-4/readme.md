---
title: Constants 
description: Understanding typed and untyped constants, their usage, and limitations in Go
date: 2024-01-20
tags: ['go', 'constants', 'typed', 'untyped', 'compile-time']
---

# Constants 

## What are Constants?

Constants are **immutable values** that cannot be changed after declaration. In Go, constants can only hold values that the compiler can infer to be constant at **compile time**.

### Key Characteristics
- Declared with `const` keyword
- Values must be determinable at compile time
- Cannot use runtime values or function calls that execute at runtime
- Can be typed or untyped

## Basic Constant Declaration

### Single Constant Declaration
~~~go
package main

import "fmt"

const x = 10       // Untyped constant
const y int32 = 15 // Typed constant (explicit type)

func main() {
    fmt.Printf("x: %T = %v\n", x, x) // x: int = 10
    fmt.Printf("y: %T = %v\n", y, y) // y: int32 = 15
}
~~~

### Multiple Constant Declaration
~~~go
const (
    g               = 10        // Untyped
    h         int32 = 20        // Typed
    lesson          = "lesson 4" // Untyped string
    isRunning       = false     // Untyped boolean
    character       = 'a'       // Untyped rune
    isTrue          = 1 > 2     // Untyped boolean (compile-time expression)
)
~~~

## Typed vs Untyped Constants

### Untyped Constants
- No explicit type specified
- Can be used with different types (flexible)
- Compiler infers type based on context

### Typed Constants
- Explicit type specified
- Can only be assigned to variables of the same type
- More restrictive but type-safe

## Constant Usage and Type Compatibility

### Untyped Constant Flexibility
~~~go
func main() {
    var a int
    a = x // Can use untyped constant with int
    fmt.Println(a) // 10
    
    var b float64
    b = x // Can use same untyped constant with float64
    fmt.Println(b) // 10.0
    
    // This demonstrates the flexibility of untyped constants
}
~~~

### Typed Constant Restrictions
~~~go
func main() {
    var c int
    // c = y // ERROR: cannot use y (type int32) as type int in assignment
    c = int(y) // Need explicit type conversion
    fmt.Println(c) // 15
    
    // var d string
    // d = x // ERROR: cannot use x (untyped int constant 10) as string value
}
~~~

## Compile-Time vs Runtime Constants

### Allowed: Compile-Time Evaluable
~~~go
func main() {
    const z = len("string") // Allowed - len on string literal is compile-time
    fmt.Println(z) // 6
    
    const k = complex(10.2, 100.9) // Allowed - complex with numeric literals
    fmt.Println(k) // (10.2 + 100.9i)
    
    const l = imag(k) // Allowed - imag with constant complex
    fmt.Println(l) // 100.9
    
    // Compile-time expressions
    const expr1 = 10 + 20
    const expr2 = 1 < 2
    const expr3 = "hello" + " world"
}
~~~

### Not Allowed: Runtime Dependencies
~~~go
func main() {
    var a int = 10
    var b float64 = 20.5
    
    // const z = a + int(b) // ERROR: a and b are variables (runtime)
    // const z = sum(a, b)  // ERROR: function call at runtime
    
    // These won't work because they depend on runtime values
}
~~~

## Constant Values and Expressions

### Valid Constant Types and Values
~~~go
const (
    // Boolean values
    isActive = true
    isReady  = false
    
    // Numerical values
    count     = 100
    price     = 99.99
    complexNum = 3 + 4i
    
    // String and rune
    greeting  = "Hello, World!"
    char      = 'A'
    
    // Compile-time expressions
    comparison = 10 > 5
    calculation = (10 + 20) * 2
    stringConcat = "Hello" + " " + "Go"
    
    // Built-in functions with constant arguments
    strLen = len("constant")
    realPart = real(3 + 4i)
    imagPart = imag(3 + 4i)
)
~~~