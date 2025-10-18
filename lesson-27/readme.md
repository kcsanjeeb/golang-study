# Dependancy Management
* go mopdules are used for manaing your dependencies in go project. 
* go modules use go.mod and go.sum files to track dependencies of your go project. 
* one go module per repository is preferred. 
* the go.mod also defines your go modules. 

## Initializing Go Module
* The go mod init command is usually used to initialize a module. 
* Name of your module is usually the whole path of your repository that hosts the module.

```go
go mod init github.com/kcsanjeeb/first-module
go: creating new go.mod: module github.com/kcsanjeeb/first-module
go: to add module requirements and sums:
        go mod tidy

```
#### go.mod
```go
module github.com/kcsanjeeb/first-module

go 1.25.1
```

### Adding Module
* The go get command is used to add module dependencies. 
* By default it pulls the default version. 
* The go mod tidy command can be used to clean up the modules not used in your project. 

```go 
go get github.com/fatih/color

go: downloading github.com/fatih/color v1.18.0
go: downloading github.com/mattn/go-isatty v0.0.20
go: downloading github.com/mattn/go-colorable v0.1.13
go: downloading golang.org/x/sys v0.25.0
go: added github.com/fatih/color v1.18.0
go: added github.com/mattn/go-colorable v0.1.13
go: added github.com/mattn/go-isatty v0.0.20
go: added golang.org/x/sys v0.25.0
```
#### main.go
```go 
package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	str := color.CyanString("Hello World !")
	fmt.Println(str)
}

```
#### go.mod
```go
module github.com/kcsanjeeb/first-module

go 1.25.1

require (
	github.com/fatih/color v1.18.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.25.0 // indirect
)
```
#### go.sum
* The go.sum serves as a lock file. 
* It stores cryptographic hashes for specific module version.
* The go.mod and go.sum files together ensure that same versions are downloaded everytime.
```go
github.com/fatih/color v1.18.0 h1:S8gINlzdQ840/4pfAwic/ZE0djQEH3wM94VfqLTZcOM=
github.com/fatih/color v1.18.0/go.mod h1:4FelSpRwEGDpQ12mAdzqdOukCy4u8WUtOY6lkT/6HfU=
github.com/mattn/go-colorable v0.1.13 h1:fFA4WZxdEF4tXPZVKMLwD8oUnCTTo08duU7wxecdEvA=
github.com/mattn/go-colorable v0.1.13/go.mod h1:7S9/ev0klgBDR4GtXTXX8a3vIGJpMovkB8vQcUbaXHg=
github.com/mattn/go-isatty v0.0.16/go.mod h1:kYGgaQfpe5nmfYZH+SKPsOc2e4SrIfOl2e/yFXSvRLM=
github.com/mattn/go-isatty v0.0.20 h1:xfD0iDuEKnDkl03q4limB+vH+GxLEtL/jb4xVJSWWEY=
github.com/mattn/go-isatty v0.0.20/go.mod h1:W+V8PltTTMOvKvAeJH7IuucS94S2C6jfK/D7dTCTo3Y=
golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.6.0/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.25.0 h1:r+8e+loiHxRqhXVl6ML1nO3l1+oFoWbnlu2Ehimmi34=
golang.org/x/sys v0.25.0/go.mod h1:/VUhepiaJMQUp4+oa/7Zr1D23ma6VTLIYjOOTFZPUcA=
```

### The go list
* It is used to list all the imported packages in your module.
* You can supply the -m flag to only list the modules imported.
```go
go list all
bytes
cmp
errors
fmt
github.com/fatih/color
github.com/kcsanjeeb/first-module
github.com/mattn/go-colorable
github.com/mattn/go-isatty
golang.org/x/sys/unix
internal/abi
internal/asan
internal/bisect
internal/bytealg
internal/byteorder
internal/chacha8rand
internal/coverage/rtcov
internal/cpu
internal/filepathlite
internal/fmtsort
internal/goarch
internal/godebug
internal/godebugs
internal/goexperiment
internal/goos
internal/itoa
internal/msan
internal/oserror
internal/poll
internal/profilerecord
internal/race
internal/reflectlite
internal/runtime/atomic
internal/runtime/exithook
internal/runtime/gc
internal/runtime/maps
internal/runtime/math
internal/runtime/strconv
internal/runtime/sys
internal/stringslite
internal/sync
internal/synctest
internal/syscall/execenv
internal/syscall/unix
internal/testlog
internal/trace/tracev2
internal/unsafeheader
io
io/fs
iter
math
math/bits
os
path
reflect
runtime
slices
sort
strconv
strings
sync
sync/atomic
syscall
time
unicode
unicode/utf8
unsafe
```

### The go mod tidy
*  The go mod tidy command as the name suggest is used for tidying up dependencies of your Go project.
* Running go mod tidy remove any unused modules from your project.
* Running go mod tidy will also add missing modules that is imported but missing from go.mod file.

```go 
package main

func main() {
	// str := color.CyanString("Hello World !")
	// fmt.Println(str)
}
```
```go 
go mod tidy
```

