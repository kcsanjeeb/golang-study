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