---
title: Control Flow 
description: Break, Continue, Labels, Goto, and Switch statements in Go
date: 2024-01-30
tags: ['go', 'control-flow', 'switch', 'break', 'continue']
---

# Control Flow

## Loop Control Statements

### Break Statement
Exits the innermost loop immediately:

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // Loop stops when i reaches 5
    }
    fmt.Println(i)
}
```

### Continue Statement
Skips the current iteration and continues with the next:

```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue // Skip even numbers
    }
    fmt.Println(i) // Prints only odd numbers
}
```

### Labels
Labels allow breaking or continuing specific outer loops:

```go
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break outer // Breaks both loops
        }
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}
```

### Goto Statement
Jumps to a labeled statement (use sparingly):

```go
func main() {
    i := 0
start:
    if i < 5 {
        fmt.Println(i)
        i++
        goto start
    }
}
```

## Switch Statement

Basic switch with cases:

```go
day := "Monday"

switch day {
case "Monday":
    fmt.Println("It's Monday")
case "Tuesday":
    fmt.Println("It's Tuesday")
default:
    fmt.Println("Other day")
}
```

### Switch Features
- **No break needed** - cases don't fall through by default
- **Multiple values** per case: `case "Mon", "Monday"`
- **Expression switches** - compare against expressions
- **Type switches** - switch on variable type

