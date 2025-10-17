---
title: Interfaces 
description: Abstract types, method contracts, and type assertions in Go
date: 2024-02-08
tags: ['go', 'interfaces', 'type-assertion', 'empty-interface']
---

# Interfaces

## What are Interfaces?

Interfaces are **abstract types** that define method signatures without implementation. They specify **what a type should do** rather than how it's implemented.

### Interface Naming Convention
- Standard library interfaces often end with `...er`
- Examples: `io.Writer`, `io.Reader`, `fmt.Stringer`

## Implementing Interfaces

Go uses **implicit interface implementation** - a type automatically satisfies an interface if it implements all its methods.

### Basic Interface Example
~~~go
type Writer interface {
    Write([]byte) (int, error)
}

type FileWriter struct{}

func (f FileWriter) Write(data []byte) (int, error) {
    // Implementation
    return len(data), nil
}
// FileWriter implicitly satisfies Writer interface
~~~

## Benefits of Interfaces

- **Abstraction**: Code works with functionality, not implementation
- **Flexibility**: Easy to swap implementations (MySQL ↔ PostgreSQL)
- **Testability**: Mock implementations for testing
- **Maintainability**: Clear contracts between components

## Empty Interface

The empty interface `interface{}` (or `any`) represents any type:

### Using Empty Interface
~~~go
func printValue(value any) {
    fmt.Println(value)
}

func main() {
    mixedSlice := []any{42, "hello", 3.14, true}
    for _, value := range mixedSlice {
        printValue(value)
    }
}
~~~

## Type Assertion

Type assertion extracts the concrete value from an empty interface:

### Safe Type Assertion
~~~go
var emptyInterface any = "Hello World"

if str, ok := emptyInterface.(string); ok {
    fmt.Println("The value is a string:", str)
} else {
    fmt.Println("Not a string")
}
~~~

### Unsafe Type Assertion (May Panic)
~~~go
str := emptyInterface.(string) // Panics if not string
~~~

## Type Switch

Type switch checks against multiple types in a clean syntax:

### Type Switch Example
~~~go
func describe(value any) {
    switch v := value.(type) {
    case string:
        fmt.Printf("String: %s\n", v)
    case int:
        fmt.Printf("Integer: %d\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
~~~

## Practical Examples

### Database Interface Example
~~~go
type Database interface {
    Connect() error
    Query(string) ([]string, error)
    Close() error
}

type MySQL struct{}
func (m MySQL) Connect() error { return nil }
func (m MySQL) Query(q string) ([]string, error) { return []string{}, nil }
func (m MySQL) Close() error { return nil }

type Postgres struct{}
func (p Postgres) Connect() error { return nil }
func (p Postgres) Query(q string) ([]string, error) { return []string{}, nil }
func (p Postgres) Close() error { return nil }

func useDatabase(db Database) {
    db.Connect()
    results, _ := db.Query("SELECT * FROM users")
    fmt.Println(results)
    db.Close()
}
~~~

### Complete Type Handling Example
~~~go
package main

import "fmt"

func processValue(val any) {
    // Type switch for comprehensive type handling
    switch v := val.(type) {
    case int:
        fmt.Printf("Processing integer: %d (doubled: %d)\n", v, v*2)
    case string:
        fmt.Printf("Processing string: %s (length: %d)\n", v, len(v))
    case []int:
        fmt.Printf("Processing slice of integers: %v\n", v)
    case map[string]int:
        fmt.Printf("Processing map: %v\n", v)
    default:
        fmt.Printf("Unhandled type: %T\n", v)
    }
}

func main() {
    processValue(42)
    processValue("hello")
    processValue([]int{1, 2, 3})
    processValue(map[string]int{"a": 1, "b": 2})
    processValue(3.14)
}
~~~

## Key Points

### Interface Implementation
- **Implicit**: No explicit `implements` keyword
- **Duck Typing**: "If it walks like a duck and quacks like a duck..."
- **Multiple Interfaces**: One type can satisfy multiple interfaces

### Empty Interface Usage
- **Flexible containers**: Store any type
- **Function parameters**: Accept any argument type
- **JSON handling**: Common in marshaling/unmarshaling

### Type Assertion vs Type Conversion
- **Type Assertion**: `interfaceValue.(ConcreteType)` - for interfaces
- **Type Conversion**: `Type(value)` - for compatible types

## Best Practices

1. **Keep interfaces small** (Single Responsibility Principle)
2. **Use meaningful names** that describe behavior
3. **Prefer concrete types** when type is known
4. **Use empty interfaces sparingly** - they bypass type safety
5. **Always use safe type assertions** with the `ok` idiom
