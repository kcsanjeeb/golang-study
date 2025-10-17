---
title: Derived Types and Iota in Go
description: Type derivation, composition, and iota constants in Go
date: 2024-02-06
tags: ['go', 'derived-types', 'iota', 'constants', 'composition']
---

# Derived Types and Iota in Go

## Derived Types

Derived types create new types from existing ones, enabling type safety and clear intent.

### From Built-in Types
```go
type Age int
type Score float64
type Name string
```

### From User-Defined Types
```go
type Human struct {
    Age int
}

type Student Human  // Derived from Human type
```

## Type Derivation vs Inheritance

**Important**: Go uses **composition**, not inheritance.

### Key Differences:
- **Different Types**: Derived and base types are distinct
- **No Method Sharing**: Methods don't automatically transfer
- **Composition Over Inheritance**: Go favors embedding over classical inheritance

### Example Demonstration:
```go
type Human struct {
    Age int
}

func (h Human) printAge() {
    fmt.Println(h.Age)
}

type Student Human  // Derived type

func main() {
    var s = Student{Age: 10}
    fmt.Println(s.Age)     // Works: 10
    // s.printAge()        // ERROR: Student doesn't have printAge method
}
```

## Iota: Sequential Constants

Iota generates successive integer constants, ideal for enumerated values.

### Basic Iota Usage
```go
const (
    New int = iota  // 0
    Pending         // 1
    Completed       // 2
)
```

### Advanced Iota Patterns
```go
const (
    ExtraSmall Size = iota * 2  // 0
    Small                       // 2
    Medium                      // 4
    Large                       // 6
    ExtraLarge                  // 8
)

const (
    Read = 1 << iota  // 1 (binary: 001)
    Write             // 2 (binary: 010)
    Execute           // 4 (binary: 100)
)
```

### Iota with Expressions
```go
const (
    _ = iota                    // Skip 0
    KB Size = 1 << (10 * iota)  // 1 << 10 = 1024
    MB                          // 1 << 20 = 1048576
    GB                          // 1 << 30 = 1073741824
)
```

## Practical Examples

### Size Enumeration with Iota
```go
type Size int

const (
    ExtraSmall Size = iota * 2
    Small
    Medium
    Large
    ExtraLarge
)

func printSize(s Size) {
    switch s {
    case ExtraSmall:
        fmt.Println("Extra Small")
    case Small:
        fmt.Println("Small")
    case Medium:
        fmt.Println("Medium")
    case Large:
        fmt.Println("Large")
    case ExtraLarge:
        fmt.Println("Extra Large")
    }
}
```

### File Permissions with Bit Shifting
```go
const (
    ReadPermission = 1 << iota  // 1
    WritePermission             // 2
    ExecutePermission           // 4
)

// Combined permissions
userPermissions := ReadPermission | WritePermission  // 3 (binary: 011)
```

## Key Characteristics

### Derived Types
- Create type-safe aliases
- Prevent accidental mixing of types
- Don't inherit methods automatically
- Require explicit conversion between types

### Iota
- Resets to 0 in each const block
- Increments for each constant declaration
- Can be used in expressions
- Ideal for enumerated constants

## Best Practices

1. **Use derived types** for domain-specific concepts
2. **Prefer composition** over type derivation for behavior reuse
3. **Use iota** for related constant values
4. **Document iota patterns** for complex expressions
5. **Consider string representations** for iota-based enums

