---
title: Structs in Go
description: Custom data types and struct operations in Go
date: 2024-01-26
tags: ['go', 'structs', 'custom-types', 'data-structures']
---

# Structs in Go

## What are Structs?

Structs are **custom data types** that group related data together. They allow you to create complex data structures beyond primitive types.

- **Custom Types**: Create your own data structures
- **Group Data**: Combine related fields into one type
- **Flexible**: Can contain any data type including slices, maps, other structs

## Struct Declaration

### Basic Struct Definition
~~~go
type student struct {
    firstName string
    lastName  string
    age       int
    classes   []string
}
~~~

## Struct Initialization

### Implicit Initialization
~~~go
foo := student{
    "John",
    "Doe",
    25,
    []string{"maths", "english"},
}
~~~

### Explicit Initialization
~~~go
foo := student{
    firstName: "John",
    lastName:  "Doe",
    age:       25,
    classes:   []string{"maths", "english"},
}
~~~

### Zero Value Struct
~~~go
var student1 student
// Fields get zero values: {firstName: lastName: age:0 classes:[]}
~~~

## Dot Notation

Use dot notation to **access and modify** struct fields:

~~~go
student4 := student{firstName: "Jade", lastName: "Doe"}
fmt.Println(student4.firstName) // Read field
student4.classes = append(student4.classes, "arts") // Modify field
~~~

## Anonymous Structs

Create structs without defining a type:

~~~go
guardian := struct {
    firstName string
    lastName  string
}{
    firstName: "Alex",
    lastName:  "Moon",
}
~~~

## Key Points

- **Implicit Initialization**: Must provide all values in order
- **Explicit Initialization**: Can omit fields (get zero values)
- **Zero Values**: Uninitialized fields get appropriate zero values
- **Mutable**: Struct fields can be modified after creation
