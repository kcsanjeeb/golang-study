---
title: Closures and Anonymous Functions 
description: Functions that capture state and inline function definitions
date: 2024-02-03
tags: ['go', 'closures', 'anonymous-functions', 'functional']
---

# Closures and Anonymous Functions 

## What are Closures?

Closures are functions that **close over values** from their surrounding scope. They capture variables from the enclosing function block and maintain access to them even after the outer function has finished executing.

### Key Characteristics
- **Capture State**: Access variables from surrounding scope
- **Persistent State**: Maintain access to captured variables
- **Stateful Functions**: Functions with associated state

## Benefits of Closures
- Access variables not directly passed as arguments
- Operate on captured variables independently
- Enable patterns like memoization and callbacks
- Create factory functions with encapsulated state

## Anonymous Functions

Anonymous functions are **nameless functions** defined inline. They're useful for small, one-off functionality without needing formal naming.

### Use Cases
- **Inline operations**: Quick function definitions
- **Callbacks**: Passing behavior as parameters
- **Immediate execution**: Define and run simultaneously

