---
title: Maps
description: Key-value data structures and their operations in Go
date: 2024-01-25
tags: ['go', 'maps', 'key-value', 'data-structures']
---

# Maps

## What are Maps?

Maps are **key-value data structures** implemented as hash tables. They store unordered pairs where each key maps to a value.

- **Key-Value Pairs**: Store data as `key: value` combinations
- **Hash Implementation**: Uses hash tables behind the scenes
- **Length**: Use `len()` to get number of key-value pairs

## Map Declaration

### Literal Declaration
~~~go
nameAge := map[string]int{
    "Foo": 25,
    "Bar": 20,
}
~~~

### Empty Map Declaration
~~~go
var nameAge map[string]int        // Nil map
nameAge := map[string]int{}       // Empty map (shorthand)
nameAge := make(map[string]int)   // Empty map (make function)
var nameAge = map[string]int{}    // Empty map (explicit)
~~~

## Map Operations

### Adding and Updating
~~~go
nameAge := map[string]int{}
nameAge["foo"] = 23
nameAge["bar"] = 24
nameAge["foo"] = 25  // Update existing
~~~

### Reading Values
~~~go
fmt.Println(nameAge["foo"])  // Access value
fmt.Println(len(nameAge))    // Get count of pairs
~~~

### Comma OK Idiom
~~~go
grade, exists := nameGrade["buka"]
fmt.Println(grade, exists)  // 0 true (even if value is 0)
~~~

## Nil vs Empty Maps

| Aspect | Nil Map | Empty Map |
|--------|---------|-----------|
| Declaration | `var m map[string]int` | `m := map[string]int{}` |
| Underlying | No storage allocated | Storage allocated |
| Read Operations | Returns zero value | Returns stored/zero value |
| Write Operations | Panics (runtime error) | Works normally |
| Length | 0 | 0 |

## Key Characteristics

- **Zero Value**: `nil` for maps
- **Key Types**: Must be comparable (strings, integers, etc.)
- **Value Types**: Any valid Go type
- **Order**: Iteration order is not guaranteed
- **Performance**: O(1) average case for operations
