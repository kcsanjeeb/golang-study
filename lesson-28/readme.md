---
title: Packages, Exports and Imports 
description: Comprehensive guide to understanding Go's package system and visibility rules
date: 2024-02-12
tags: ['go', 'packages', 'exports', 'imports', 'beginners']
---

# Packages, Exports and Imports in Go

## 📦 What are Packages?

### Definition
Packages are like **folders for your code** that help you organize related files together. Think of them as containers that hold code for specific purposes.

### Key Rules:
1. **Every Go file must start with a package declaration**
2. **One package per folder** (directory)
3. **Files in the same folder must belong to the same package**
4. **Exception**: Test files can have `_test` suffix

### Example:
```go
// This file is part of the 'math' package
package math

// Your code goes here...
```

## 🔒 Exports: Controlling What Others Can See

### The Capital Letter Rule
Go uses a simple but powerful rule: **Capital letters mean "public", lowercase means "private"**

### Export Rules:
| Case | Visibility | Who can access? |
|------|------------|-----------------|
| **Uppercase** | Exported | Anyone who imports your package |
| **Lowercase** | Unexported | Only code within the same package |

### Practical Example:
```go
package math

// ✅ EXPORTED - Can be used by other packages
const PI = 3.14159
func Add(x, y int) int {
    return x + y
}
type Calculator struct {
    Brand string  // Brand is exported (uppercase)
}

// ❌ NOT EXPORTED - Only usable within this package
const secretFormula = "x + y"
func subtract(x, y int) int {
    return x - y
}
type internalConfig struct {  // internalConfig is private
    debug bool
}
```

## 📥 Imports: Using Other People's Code

### How Imports Work
To use code from another package, you need to **import** it first.

### Import Rules:
1. **Use the full path**: `module-name/package-folder`
2. **Must use what you import** (Go compiler is strict!)
3. **Each file imports separately**

### Complete Working Example:

#### Project Structure:
```
my-math-project/
├── go.mod ← Module definition file
├── main.go ← Main program
└── math/ ← Math package folder
└── math.go ← Math package code
```


#### Step 1: Define the Module (go.mod)
```go
module my-math-project  // Your project's name
go 1.25.1              // Go version
```

#### Step 2: Create the Math Package (math/math.go)
```go
package math

// These functions can be used by other packages
func Add(x, y int) int {
    return x + y
}

func Multiply(x, y int) int {
    return x * y
}

// This function is private to the math package
func secretCalculation(x int) int {
    return x * 2
}
```

#### Step 3: Use the Package (main.go)
```go
package main

import (
    "fmt"
    "my-math-project/math"  // Import our math package
)

func main() {
    // Use the exported functions
    sum := math.Add(3, 4)
    product := math.Multiply(3, 4)
    
    fmt.Printf("Sum: %d, Product: %d\n", sum, product)
    // Output: Sum: 7, Product: 12
    
    // This would cause an ERROR:
    // math.secretCalculation(5) ← secretCalculation is not exported!
}
```

## 🏷️ Package Naming Best Practices

### Good Names:
- **Short and clear**: `math`, `http`, `database`
- **Lowercase only**: `utils` not `Utils`
- **Descriptive**: `userauth` not `ua`

### Folder Structure Examples:
```
my-app/
├── calculator/ ← package calculator
│ └── calc.go
├── database/ ← package database
│ └── db.go
└── utils/ ← package utils
└── helpers.go
```

### Exceptions:
- **Versioned packages**: `mypackage/v2`, `mypackage/v3`
- **Internal packages**: `internal/` (special restricted access)

## 🚀 Common Patterns & Tips

### 1. Group Related Imports
```go
import (
    "fmt"
    "strings"
    "myapp/database"
    "myapp/utils"
)
```

### 2. Use Package Aliases for Clarity
```go
import (
    "database/sql"
    "fmt"
    m "myapp/math"  // 'm' is now an alias for math package
)

func main() {
    result := m.Add(5, 3)  // Clear and concise
    fmt.Println(result)
}
```

### 3. Blank Import for Side Effects
```go
import (
    _ "github.com/lib/pq"  // Only runs init() functions
)
// Used when a package needs to register itself but you don't call its functions directly
```

## ❌ Common Mistakes to Avoid

### Mistake 1: Wrong Package Declaration
```go
// ❌ WRONG - File in math folder but different package
package calculator  // Should be package math

// ✅ CORRECT
package math
```

### Mistake 2: Forgetting to Export
```go
package math

// ❌ WRONG - Other packages can't use this
func add(x, y int) int {
    return x + y
}

// ✅ CORRECT - Other packages can use this
func Add(x, y int) int {
    return x + y
}
```

### Mistake 3: Unused Imports
```go
import (
    "fmt"
    "strings"  // ❌ ERROR: Imported but not used
    "myapp/math"
)

func main() {
    result := math.Add(1, 2)
    fmt.Println(result)
    // strings package is imported but never used!
}
```

## 🎯 Quick Summary

| Concept | What to Remember |
|---------|------------------|
| **Packages** | Organize code into logical groups |
| **Exports** | Uppercase = public, Lowercase = private |
| **Imports** | Use full path, must use what you import |
| **Naming** | Folder name usually matches package name |

## 💡 Real-World Analogy

Think of packages like **departments in a company**:
- **Package** = Department (Sales, Engineering, HR)
- **Exported functions** = Public phone numbers anyone can call
- **Unexported functions** = Internal extensions only department members can use
- **Importing** = Getting the phone number to call another department
