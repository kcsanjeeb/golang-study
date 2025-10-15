---
title: Operators
description: Comprehensive guide to operators, operands, and precedence in Go
date: 2024-01-22
tags: ['go', 'operators', 'precedence', 'expressions']
---

# Operators

## Overview

**Operators** are symbols that perform arithmetic or logical operations on **operands**, combining them into expressions.

- **Binary operators**: Work with two operands
- **Unary operators**: Work with one operand
- **Postfix/Suffix**: Some operators are used after the operand

## Operator Types

### Unary Operators
| Operator | Operation | Description |
|----------|-----------|-------------|
| `+` | Positive | Unary plus |
| `-` | Negative | Unary minus |
| `!` | Logical NOT | Boolean negation |
| `^` | Bitwise complement | Bitwise NOT |
| `*` | Pointer dereference | Access value through pointer |
| `&` | Address | Get memory address |
| `<-` | Channel receive | Receive from channel |

### Arithmetic Operators
| Operator | Operation | Description |
|----------|-----------|-------------|
| `+` | Addition | Sum of operands |
| `-` | Subtraction | Difference of operands |
| `*` | Multiplication | Product of operands |
| `/` | Division | Quotient of operands |
| `%` | Modulus | Remainder of division |

### Bitwise, Logical and Shift Operators
| Operator | Operation | Description |
|----------|-----------|-------------|
| `&` | Bitwise AND | Bits that are set in both operands |
| `\|` | Bitwise OR | Bits that are set in either operand |
| `^` | Bitwise XOR | Bits that are set in one but not both |
| `&^` | Bit clear (AND NOT) | Clear bits from first operand |
| `<<` | Left shift | Shift bits left |
| `>>` | Right shift | Shift bits right |

### Comparison Operators
| Operator | Operation | Description |
|----------|-----------|-------------|
| `==` | Equal | Returns true if values are equal |
| `!=` | Not equal | Returns true if values are different |
| `<` | Less than | Returns true if left is less than right |
| `<=` | Less than or equal | Returns true if left is less than or equal to right |
| `>` | Greater than | Returns true if left is greater than right |
| `>=` | Greater than or equal | Returns true if left is greater than or equal to right |

### Logical Operators
| Operator | Operation | Description |
|----------|-----------|-------------|
| `&&` | Logical AND | True if both operands are true |
| `\|\|` | Logical OR | True if either operand is true |
| `!` | Logical NOT | Reverse the boolean value |

### Other Operators
| Operator | Operation | Description |
|----------|-----------|-------------|
| `&` | Address of | Returns pointer to variable |
| `*` | Pointer dereference | Accesses value through pointer |
| `<-` | Channel operations | Send/receive from channel |

## Operator Precedence

**Precedence** determines the order of operations (PEMDAS concept).

| Precedence | Operators | Evaluation |
|------------|-----------|------------|
| 5 | `* / % << >> & &^` | Multiplication, division, modulus, shifts, bitwise AND, bit clear |
| 4 | `+ - \| ^` | Addition, subtraction, bitwise OR, bitwise XOR |
| 3 | `== != < <= > >=` | Comparison operators |
| 2 | `&&` | Logical AND |
| 1 | `\|\|` | Logical OR |

**Note**: Operators with higher precedence are evaluated first. Use parentheses `()` to override precedence.

---
