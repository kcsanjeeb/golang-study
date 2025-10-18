# Panic and Recovery

## Panic
* Panic occurs when a go program has reached a state of no return 
* Example: trying to access an element that is out of bounds of an array causes runtime panic.
* You can raise your own panic with panic() function
* The panic function accepts an argument of type any 
* When panic is called code execution stops and panic prints the error stack of where the panic occured.
* Panic has a relationship with defer, it can be better understood with the help of illustration

```go
package main

import "fmt"

func function1() {
	// Registers a deferred func that will print: “Function 1 deffered function called.”
	defer func() {
		fmt.Println("Function 1 deffered function called.")
	}()
	// Calls function2.
	function2()
}

func function2() {
	// Registers a deferred func that will print: “Function 2 deffered function called.”
	defer func() {
		fmt.Println("Function 2 deffered function called.")
	}()
	// Calls panic("Function 2 panicked !").
	panic("Function 2 panicked !")
}

func main() {
	// Registers a deferred func that will print: “Function Main deffered function called.”
	defer func() {
		fmt.Println("Function Main deffered function called.")
	}()

	//Calls function1.
	function1()
}
```
```
Output : 
term@mac lesson-26 $ go run main.go
Function 2 deffered function called.
Function 1 deffered function called.
Function Main deffered function called.
panic: Function 2 panicked !
...
```

Here’s exactly what happens, step by step:
1. main starts
    * Registers a deferred func that will print: “Function Main deffered function called.”
Calls function1.

2. function1 runs
    * Registers a deferred func that will print: “Function 1 deffered function called.”
    * Calls function2.

3. function2 runs
    * Registers a deferred func that will print: “Function 2 deffered function called.”
    * Calls panic("Function 2 panicked !").

4. Panic unwinds the stack (no recover anywhere)
    * Go starts popping frames and running each frame’s deferred calls before unwinding it.
    * In function2, its deferred func runs → prints “Function 2 deffered function called.”
    * Unwind to function1, run its deferred → prints “Function 1 deffered function called.”
    * Unwind to main, run its deferred → prints “Function Main deffered function called.”

5. Program terminates abnormally
    * After all defers run, the panic is still unhandled, so the runtime prints the panic message and a stack trace to stderr, then exits with a non-zero status.


---

## Recover
* The recover() function can be used to recover from panics . 
* Function recover() need to be called in a deferred function. 
* The return value of calling recover() is the argument provided to panic().

```go
package main

import "fmt"

func panicExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic: ", r)
		}
	}()
	panic("Something went Wrong !")
}

func main() {
	fmt.Println("Program has started")
	panicExample()
	fmt.Println("Exited !!")

}
```
```
Output:
Program has started
Recovered from panic:  Something went Wrong !
Exited !!
```

### Use of Recover
* Purpose of recover is to gracefully shutdown.
* Not designed to act like exception handling.
* If you recover, code can and most probably will panic again.
* Used in packages to handle panic and translate to a meaningful error