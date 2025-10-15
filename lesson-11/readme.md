---
title: Conditional Statements 
description: if, if-else, and scoped variables in Go
date: 2024-01-27
tags: ['go', 'conditionals', 'if-else', 'control-flow']
---

# Conditional Statements

## The if Statement

The `if` statement executes code blocks based on boolean conditions.

```go
if age >= 18 {
    fmt.Println("You are an adult!")
}
```

## The if-else Statement

`if-else` provides alternative execution paths when the condition is false.

```go
if age >= 18 {
    fmt.Println("You are an adult!")
} else {
    fmt.Println("You are not an adult")
}
```

## The if-else if Ladder

Chain multiple conditions using `else if`:

```go
if age >= 18 {
    fmt.Println("You are an adult!")
} else if age >= 13 {
    fmt.Println("You are a teenager!")
} else {
    fmt.Println("You are a child!")
}
```

## Scoped Variables in if

Go allows variable initialization within the `if` statement. The variable is scoped exclusively to the if-else block.

```go
if even := isEven(age); even {
    fmt.Println("Age is even!")
}
// 'even' variable is not accessible here
```

**Benefits:**
- Cleaner code by limiting variable scope
- Variables exist only where needed
- Prevents namespace pollution
