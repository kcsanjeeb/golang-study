---
title: Error Handling 
description: Go's approach to errors with practical examples
date: 2024-02-09
tags: ['go', 'errors', 'error-handling', 'custom-errors']
---

# Error Handling

## Go's Error Philosophy

Go handles errors explicitly without `try-catch` blocks. The convention is to return errors as the **last return value** and check them with `if` statements.

### Error String Conventions
- **No capitalization** (start with lowercase)
- **No punctuation** at the end
- **Descriptive and clear** messages

## The Error Interface

The `error` type is a built-in interface with a single method:

```go
type error interface {
    Error() string
}
```

## Creating Errors

### 1. Using `errors.New()` - Simple Errors
```go
import "errors"

func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
```

### 2. Using `fmt.Errorf()` - Formatted Errors
```go
import "fmt"

func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero %f", b)
    }
    return a / b, nil
}
```

### 3. Custom Error Types - Structured Errors
```go
type CustomError struct {
    Code    int
    Message string
}

func (c CustomError) Error() string {
    return fmt.Sprintf("Error %d: %s", c.Code, c.Message)
}

func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, CustomError{Code: 1, Message: "cannot divide by zero"}
    }
    return a / b, nil
}
```

## Error Handling Pattern

### Standard Error Checking
```go
result, err := Divide(10, 0)
if err != nil {
    fmt.Println("error:", err)
    return // or handle appropriately
}
fmt.Println("result:", result)
```

### Multiple Error Checks
```go
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    data, err := io.ReadAll(file)
    if err != nil {
        return err
    }
    
    // Process data...
    return nil
}
```

---

# Sentinal Errrors
* Indicate occurance of a unique events. 
* Exported package level variables 
* Example from database/sql package : 
    var ErrNoRows = errors.New("sql: no rows in result set")
* To check for specific error , you can do 
    if err == sql.ErrNoRows

Sentinel errors are **predefined, exported error variables** that indicate specific error conditions. They allow callers to check for particular error types.

### Characteristics
- **Package-level variables**: Exported for external use
- **Unique events**: Represent specific error conditions
- **Constant comparison**: Check with `==` operator

### Example from Standard Library
```go
// database/sql package
var ErrNoRows = errors.New("sql: no rows in result set")

// Usage
if err == sql.ErrNoRows {
    // Handle no rows case
}
```

### Creating Sentinel Errors
```go
package mypkg

import "errors"

var (
    ErrInvalidInput = errors.New("mypkg: invalid input")
    ErrNotFound     = errors.New("mypkg: resource not found")
    ErrTimeout      = errors.New("mypkg: operation timed out")
)
```


## Error Wrapping
* wrap errors to add additional information 
* an error e wraps another error if e's type has one of the methods: 
    * Unwrap() error
    * Unwrap() []error
* an error chain refers to series of errors that are wrapped together
* you can use fmt.Errof() with %w to wrap errors. 


Error wrapping adds **contextual information** to errors while preserving the original error.

### Wrapping with `fmt.Errorf(%w)`
```go
func firstFunction() error {
    return fmt.Errorf("original error: something went wrong in first function")
}

func secondFunction() error {
    firstErr := firstFunction()
    if firstErr != nil {
        return fmt.Errorf("failed in second function: %w", firstErr)
    }
    return nil
}

func main() {
    err := secondFunction()
    fmt.Println(err)
    // Output: failed in second function: original error: something went wrong in first function
}
```

### Wrapping Multiple Errors with `errors.Join()`
```go
func processData() error {
    var errs []error
    
    if err := validateInput(); err != nil {
        errs = append(errs, err)
    }
    if err := processFile(); err != nil {
        errs = append(errs, err)
    }
    
    return errors.Join(errs...)
}

// Or directly
func secondFunction() error {
    firstErr := firstFunction()
    if firstErr != nil {
        secondErr := errors.New("failed in second function")
        return errors.Join(firstErr, secondErr)
    }
    return nil
}
```


## errors.Join()
* to wrap multiple errors
* join wraps provided errros, discarding nil errors.
* if all errors are nil, it returns nil 
* A non-nil error returned by join implements the Unwrap() []error method. 


## Unwrapping Errors

The `errors` package provides functions to inspect error chains.

* standard lib errors package has unwrap function to unwrap errors 
* unwrap calls the unwrap method implemented on the error or returns nil 
* fmt.Errof() create error implements the Unwrap method. 
* please note unwrap does not unwrap errors returned by errors.Join()
* if e.Unwrap() returns a non nil error w or. a slice containing w, then we say that e wraps w 
* it is invalid for a unwrap method to return an [] error containing a nil error value 
* a nil error returned from e.Unwrap() indicates that e does not wrap any error. 

### Basic Unwrapping
```go
import "errors"

func main() {
    err := secondFunction()
    
    // Unwrap single error
    original := errors.Unwrap(err)
    fmt.Println("Unwrapped:", original)
    
    // Check if error chain contains specific error
    if errors.Is(err, sql.ErrNoRows) {
        fmt.Println("Contains ErrNoRows")
    }
    
    // Extract specific error type from chain
    var customErr *CustomError
    if errors.As(err, &customErr) {
        fmt.Println("Found CustomError:", customErr)
    }
}
```

### Custom Error with Unwrap
```go
type CustomError struct {
    Message string
    Wrapped error
}

func (e CustomError) Error() string {
    return fmt.Sprintf("%s: %v", e.Message, e.Wrapped)
}

func (e CustomError) Unwrap() error {
    return e.Wrapped
}

func SomeFunction() error {
    return CustomError{
        Message: "original error: something went wrong",
        Wrapped: errors.New("wrapped error"),
    }
}
```

## Key Functions for Error Inspection

### errors.Unwrap()
- Extracts the next error in the chain
- Returns nil if no wrapped error

### errors.Is()
- Checks if error chain contains specific error
- Uses `==` comparison for sentinel errors

### errors.As()
- Extracts specific error type from chain
- Useful for custom error types

## Best Practices

1. **Use sentinel errors** for predictable, checkable error conditions
2. **Wrap errors** to add context while preserving original
3. **Use `errors.Join()`** for multiple independent errors
4. **Implement `Unwrap()`** for custom error types
5. **Use `errors.Is()` and `errors.As()`** for error inspection

---

# error.Is() & error.As()

## error.Is()
* The error.Is() function accepts two parameters, the first is the error and second is the instance of error we are checking against. 
* Function Signature: `func Is(err, target error) bool `
* Iterate through the errors in the error chain and returns true if matching error instance found. 
* Own Implementations: `func Is(error) bool`

```go 
package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Message string
}

var ErrFirstError = CustomError{Message: "first error"}

func (c CustomError) Error() string {
	return c.Message
}

func SomeFunction() error {
	return ErrFirstError
}

func main() {
	err := SomeFunction()
	fmt.Println(err)

	if err != nil {
		if errors.Is(err, ErrFirstError) {
			fmt.Println("sentinal error found")
		}
	}
}
```
```go
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	if _, err := os.Open("non-existing"); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("file does not exist")
		} else {
			fmt.Println(err)
		}
	}
}
```

## error.As()
* The errors.As function accepts two parameters the first is the error and the second a non-nil pointer to either a type that implements error, or to any interface type.
* Function signature: `func As(err error, target any) bool`
* Iterate through the errors in the error chain and matching target type, if found returns true and set the target to that error value.
* Own implementation: `func As(interface{}) bool`

```go
package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Message string
	Status  int
}

func (e CustomError) Error() string {
	return e.Message
}

func SomeFunction() error {
	return CustomError{Message: "original error: something went wrong", Status: 400}
}

func main() {
	err := SomeFunction()
	var customErr CustomError
	if errors.As(err, &customErr) {
		fmt.Println("Extracted CustomError: ", customErr)
	}
}
```