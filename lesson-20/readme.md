---
title: Pointers, Memory Management, and Mutability in Go
description: Comprehensive guide to pointers, memory allocation, and value vs reference semantics
date: 2024-02-05
tags: ['go', 'pointers', 'memory', 'mutability', 'slices', 'maps']
---

# Pointers, Memory Management, and Mutability in Go

## Memory Allocation

Programs are allocated memory by the OS for execution, divided into two key segments:

- **Stack**: Manages function call execution
- **Heap**: Handles dynamic storage of variables
- **Garbage Collector**: Go runtime automatically manages memory allocation and deallocation

## Pointers Fundamentals

A pointer is a variable that holds the **memory address** of another variable.

### Memory Illustration
* Variable Name: name isEven age section
* Variable Value: "a" true 10 "b"
* Memory Address: 0x00 0x08 0x16 0x24


Pointer variables hold these memory address values (0x00, 0x08, 0x16, 0x24).

### Basic Pointer Operations
```go
var intPtr *int        // Pointer declaration (nil by default)
age := 10
intPtr = &age          // Address-of operator
fmt.Println(*intPtr)   // Dereferencing: prints 10
```

## Mutability

### Mutable vs Immutable
- **Mutable**: Values can be modified after assignment (`var`)
- **Immutable**: Values cannot be changed (`const`)

```go
var a int = 10        // Mutable
a = 20                // Allowed

const b = 30          // Immutable
// b = 25             // Compilation error
```

## Function Parameter Passing

### Call by Value (Default in Go)
```go
func increment(x int) {
    x++ // Operates on copy, original unchanged
}

c := 10
increment(c)
fmt.Println(c) // Still 10
```

### Call by Reference (Using Pointers)
```go
func incrementActual(x *int) {
    *x++ // Dereferences and modifies original
}

d := 10
incrementActual(&d)
fmt.Println(d) // Now 11
```

## Slices: The Special Case

### Slice Internal Structure
A slice is a data structure containing:
- **Pointer** to underlying array
- **Length** (number of elements)
- **Capacity** (total available space)

### Slice Behavior Demonstration
```go
func main() {
    s := []int{1, 2}
    fmt.Println(s) // [1 2]
    somethingSlice(s)
    fmt.Println(s) // Behavior depends on operation
}

func somethingSlice(s []int) {
    s[0] = 100     // MODIFIES original (same underlying array)
    // OR
    s = append(s, 1000) // Creates new slice, original UNCHANGED
}
```

### Why Append Behaves Differently
- **Direct modification**: Uses same underlying array
- **Append**: May create new array if capacity exceeded
- **Length changes**: Only affect the copy, not original

## Maps: Reference Type Behavior

Maps are reference types - modifications in functions affect the original:

```go
func modifyMap(m map[string]int) {
    m["key"] = 100 // Affects original map
}
```

## Practical Examples

### Struct Modification with Pointers
```go
type student struct {
    firstname string
    lastname  string
}

func updateLastname(s *student, newlastname string) {
    s.lastname = newlastname // Direct field access works
}

func updateLastnameWithErr(s *student, newlastname string) (*string, error) {
    if newlastname == "" {
        return nil, errors.New("empty new last name")
    }
    previous := s.lastname
    s.lastname = newlastname
    return &previous, nil // Return pointer to previous value
}
```

## Key Concepts Summary

### Pointer Rules
- Use `*type` for pointer declaration
- Use `&variable` to get address
- Use `*pointer` to dereference
- Nil pointers cause runtime panic when dereferenced

### Go's Strict Call-by-Value
- Even pointers are passed by value (copy of pointer)
- Maps and slices contain pointers internally
- Direct modifications affect original, but structural changes may not

### Performance Considerations
- Pass large structs by pointer for performance
- Be mindful of garbage collector overhead
- Use value semantics unless performance impact is substantial

## Best Practices

1. **Use pointers** when functions need to modify arguments
2. **Return pointers** to distinguish between zero value and no value
3. **Document mutability** in function behavior
4. **Prefer value semantics** for small data structures
5. **Use pointers judiciously** for large structs to avoid copying
