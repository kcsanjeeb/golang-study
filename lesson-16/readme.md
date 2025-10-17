---
title: Variadic Parameters
description: Functions that accept any number of arguments
date: 2024-02-01
tags: ['go', 'functions', 'variadic', 'parameters']
---

# Variadic Parameters

## What are Variadic Parameters?

Variadic parameters allow functions to accept **any number of arguments** of the same type. They provide flexibility when you don't know the exact number of parameters in advance.

## Rules for Variadic Parameters

- Use **three dots** `...` before the type
- Must be the **final parameter** in the function
- Accessed as a **slice** within the function
- Can be called with zero or more arguments

## Syntax

```go
func functionName(regularParams ...type) returnType {
    // parameters are available as a slice
}
```
