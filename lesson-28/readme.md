# Packages
* Packages are collection of source code files that logically grouped.
* First line of code is the package declaration in a Go file.
* Only one package can live inside a directory with exception of test packages.

```go 
package math

func main() {

}
```

# Exports
* Go uses the case of constants ,variables, constants , types etc to declare if they are exported or not.
* constants ,variables, constants , types etc that start with uppercase letters are exported.
```go
package math

// PI Exported (Uppercase letter)
const PI = 3.14159

// Add Exported (Uppercase letter)
func Add(x, y int) int {
	return x + y
}

// substract not Exported (Lowercase letter)
func substract(x, y int) int {
	return x - y
}
```

# Imports
* To import package you need the package path which is the module name plus the relative directory path of the package.
* The Go compiler does not allow packages to be imported but not use them.
* Each code file needs to have its own imports.

```go
package main

import (
	"example-1/math"
	"fmt"
)

func main() {
	sum := math.Add(3, 4)
	fmt.Println(sum) // Output: 7
}

```
```go
module example-1

go 1.25.1
```
```go
package math

func Add(x, y int) int {
	return x + y
}
```
```go
├── go.mod
├── main.go
├── math
    └── math.go
```

# Package Naming
* Idiomatically a package directory should have the same name as the package.
* Their are exceptions to this rule, example when you have different version of the same package. 
