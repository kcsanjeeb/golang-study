---
title: Identifiers 
description: Naming rules and predeclared identifiers in Go
date: 2024-01-21
tags: ['go', 'identifiers', 'naming', 'syntax']
---

# Identifiers

## What are Identifiers?

Identifiers are names given to program entities like variables, constants, types, and functions. They are fundamental building blocks for naming elements in your code.

## Identifier Rules

- **Composition**: Formed by one or more Unicode letters or digits
- **First Character**: Must be a letter (Unicode letter)
- **Underscore**: Considered as a letter (`_` is valid)
- **Case Sensitivity**: `myVar` and `myvar` are different identifiers
- **Unicode Support**: Can use characters like `α`, `β`, `字`

## Valid Identifier Examples

- `myInt`
- `_myString` 
- `MyName`
- `αβ`
- `user123`
- `MAX_SIZE`

## Predeclared Identifiers

Go has built-in identifiers that are implicitly declared:

### Types
- Basic types: `bool`, `string`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`, `byte`, `rune`, `float32`, `float64`, `complex64`, `complex128`
- Error type: `error`

### Constants
- `true`, `false`
- `iota`

### Zero Value
- `nil`

### Functions
- `append`, `cap`, `close`, `complex`, `copy`, `delete`, `imag`, `len`, `make`, `new`, `panic`, `print`, `println`, `real`, `recover`

## Key Points

- Predeclared identifiers are available in every Go file
- You can shadow predeclared identifiers (but not recommended)
- Follow naming conventions for readable and maintainable code

