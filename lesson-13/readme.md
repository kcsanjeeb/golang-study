---
title: Loops in Go
description: Different types of for loops and their usage in Go
date: 2024-01-29
tags: ['go', 'loops', 'for', 'range', 'iteration']
---

# Loops in Go

## Types of For Loops

Go has only one looping construct - the `for` loop, but it comes in several forms:

### For with Condition
Similar to `while` loops in other languages:

```go
a := 10
for a > 0 {
    fmt.Println(a)
    a--
}
```

### Infinite Loop with Break
Loop indefinitely and break when condition met:

```go
b := 10
for {
    fmt.Println(b)
    b--
    if b == 0 {
        break
    }
}
```

### For with Clause
Traditional three-component for loop:

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

### For with Range Clause
Iterate over arrays, slices, strings, maps, and channels:

```go
// Over slices
c := []int{1, 2, 3, 4, 5}
for index, value := range c {
    fmt.Printf("index: %d, value: %d\n", index, value)
}

// Over strings (iterates runes)
d := "你好"
for index, runeValue := range d {
    fmt.Printf("index: %d, rune value: %d\n", index, runeValue)
}
```

## Key Points

- **Only `for` loop** - no while or do-while
- **Range loops** handle Unicode correctly in strings
- **Break statement** exits loops early
- **Continue statement** skips to next iteration
