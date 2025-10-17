---
title: Embedding and Composition 
description: Code reuse through embedding and method promotion
date: 2024-02-07
tags: ['go', 'embedding', 'composition', 'methods', 'structs']
---

# Embedding and Composition

## What is Embedding?

Embedding is Go's approach to **code reuse through composition** rather than inheritance. It allows one struct to include another struct, promoting the embedded type's fields and methods to the containing struct.

### Basic Syntax
```go
type Engine struct {
    HorsePower int
    Model      string
}

type Car struct {
    Model string
    Engine  // Embedded type (no field name)
}
```

## Embedding vs Inheritance

### Key Differences:
- **Composition over Inheritance**: Go favors "has-a" relationships over "is-a"
- **No Type Hierarchy**: Embedded types remain distinct
- **Explicit Overrides**: Same-name fields/methods require explicit access
- **No Polymorphism**: No virtual methods or dynamic dispatch

## Method and Field Promotion

### Automatic Promotion
Embedded type's methods and fields become directly accessible:

```go
type Engine struct {
    HorsePower int
}

func (e *Engine) Start() {
    fmt.Println("Engine Started")
}

type Car struct {
    Engine  // Embedded
}

func main() {
    myCar := Car{Engine: Engine{HorsePower: 200}}
    myCar.Start()           // Direct method access
    fmt.Println(myCar.HorsePower)  // Direct field access
}
```

### Name Conflicts Resolution
When fields/methods have same names, use explicit embedding:

```go
type Car struct {
    Model string
    Engine
    GPS
}

myCar := Car{
    Model:  "Nissan",
    Engine: Engine{Model: "Nissan HR engine"},
    GPS:    GPS{Model: "Garmin"},
}

fmt.Println(myCar.Model)           // Car's Model: "Nissan"
fmt.Println(myCar.Engine.Model)    // Engine's Model: "Nissan HR engine"
fmt.Println(myCar.GPS.Model)       // GPS's Model: "Garmin"
```

## Practical Example

```go
package main

import "fmt"

type Engine struct {
    HorsePower int
    Model      string
}

func (e *Engine) Start() {
    fmt.Println("Engine Started")
}

type GPS struct {
    Model string
}

type Car struct {
    Model string
    Engine  // Embedded
    GPS     // Embedded
}

func (c *Car) Drive() {
    fmt.Printf("Driving my %s...\n", c.Model)
}

func main() {
    myCar := Car{
        Model:  "Nissan",
        Engine: Engine{HorsePower: 200, Model: "Nissan HR engine"},
        GPS:    GPS{Model: "Garmin"},
    }
    
    // Direct access to promoted fields/methods
    fmt.Println("Car Model:", myCar.Model)
    fmt.Println("Car HorsePower:", myCar.HorsePower) // Promoted from Engine
    
    // Method calls
    myCar.Start()  // Promoted from Engine
    myCar.Drive()  // Car's own method
    
    // Explicit access for conflicting names
    fmt.Println("Car Model:", myCar.Model)
    fmt.Println("Engine Model:", myCar.Engine.Model)
    fmt.Println("GPS Model:", myCar.GPS.Model)
}
```

## Benefits of Embedding

1. **Code Reuse**: Share functionality without inheritance
2. **Type Safety**: Maintain clear type relationships
3. **Flexibility**: Mix and match behaviors through composition
4. **Explicit Design**: Clear about what's being embedded
5. **No Diamond Problem**: Avoids multiple inheritance issues

## Limitations

- **No Method Overriding**: Can't override embedded methods
- **Name Conflicts**: Require explicit resolution
- **No Polymorphism**: Embedded types don't satisfy interface requirements of their parent

## Best Practices

1. **Use meaningful names** to avoid conflicts
2. **Document embedded relationships** for clarity
3. **Prefer shallow embedding** over deep nesting
4. **Use interfaces** for polymorphic behavior
5. **Keep embeddings purposeful** - don't embed unnecessarily
