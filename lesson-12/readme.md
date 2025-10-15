---
title: Code Blocks and Shadowing 
description: Understanding variable scope and shadowing in Go
date: 2024-01-28
tags: ['go', 'scope', 'shadowing', 'code-blocks']
---

# Code Blocks and Shadowing

## Code Blocks

In Go, variables can be created at different scope levels:

### Universal Block
- Contains all Go predeclared identifiers
- Built-in types, functions, and constants
- Always available everywhere

### Package Block
- Variables declared at package level
- Accessible throughout the package
- Shared across all files in the package

### File Block
- Imports are file-scoped
- Variables declared in one file aren't automatically available in another

### Function Block
- Variables declared inside functions
- Local to that function only

### Inner Blocks
- Variables declared within `{}` blocks
- Limited to that specific block scope

## Shadowing

Shadowing occurs when a variable in inner scope has the same name as a variable in outer scope. The inner variable "shadows" the outer one.

## Example

### util.go
~~~go
package main

import "fmt"

func something() {
    fmt.Println(a) // Accesses package-level variable a = 10
}
~~~

### main.go
~~~go
package main

import "fmt"

var (
    a = 10 // Package-level variable
)

func main() {
    fmt.Println(a) // 10 (package-level)
    {
        a := 15 // Inner block variable - SHADOWS package-level 'a'
        fmt.Println(a) // 15 (inner block)
    }
    fmt.Println(a) // 10 (package-level again - shadow ends)
}
~~~

## Key Points

- **Inner blocks** can access outer block variables
- **Same-name variables** in inner blocks shadow outer ones
- **Shadowing ends** when inner block completes
- **Package variables** are shared across files in same package
