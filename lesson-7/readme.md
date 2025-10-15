---
title: Arrays 
description: Fixed-size sequences and their usage in Go
date: 2024-01-23
tags: ['go', 'arrays', 'data-structures', 'fixed-size']
---

# Arrays

## What are Arrays?

Arrays are data structures that store elements of the same type in **contiguous memory blocks**.

- **Elements**: Each piece of data in the array
- **Same Type**: All elements must be identical type
- **Zero-Indexed**: First element at index 0
- **Fixed Size**: Length determined at declaration

## Benefits of Arrays

1. **Memory Efficiency** - Contiguous memory allocation
2. **Fast Access** - O(1) random access by index
3. **Type Safety** - Compile-time type checking

## Array Declaration

### Basic Declaration
~~~go
var a [5]int          // Zero-valued array
b := [5]int{1,2,3,4,5} // With initial values
~~~

### Sparse Array
~~~go
c := [5]int{1, 3:44, 5} // Index-based initialization
// Result: [1, 0, 0, 44, 5]
~~~

### Implicit Length
~~~go
d := [...]int{1,2,3,4,5,6,7,8} // Compiler determines length
~~~

## Array Operations

### Length and Access
~~~go
length := len(d)      // Get array length
element := d[3]       // Access by index
~~~

### Multi-dimensional
~~~go
f := [2][2]int{{1,2}, {3,4}} // 2D array
~~~

## Limitations

- **Fixed Size** - Cannot resize after declaration
- **Type Includes Size** - `[5]int` and `[10]int` are different types
- **No Variable Size** - Cannot use variables for array size
- **Slices Preferred** - More flexible alternative for most use cases
