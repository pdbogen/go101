name:title
background-image:url(bg.png)
background-color:white
background-size:50% auto
background-position:100% 0%
layout:true

class:middle, center

.grey[
{{content}}
]

---
name:blue
layout:true
background-image:none
background-color:#4b6bc6

---
name:white
layout:true
background-image:none
background-color:white

.black[
{{content}}
]
---
name:green
layout:true
background-image:none
background-color:#63d297

.darkgrey[
{{content}}
]
---
layout:false
template: title

![:scale 50%](slack.png)

# Go 101

## Basic Syntax

## Workflow

---
template:blue

# Before We Start

## Do this: obtain `Go`

```
$ brew install go
$ go version
go version go1.8.1 linux/amd64
```

#### Other versions are OK as long as they're 1.8 or newer.

## Do this: exercise your `GOPATH`: `$HOME/go`

```
$ go get -v github.com/golang/example/hello
github.com/golang/example (download)
github.com/golang/example/stringutil
github.com/golang/example/hello
```

---
template:blue

# Go Workspace

## Your workspace holds all of your Go source code and go-managed binaries and libraries.

* Default: `$HOME/go`
* Set by `GOPATH` environment variable


* `$GOPATH/bin` — Go-managed binaries
* `$GOPATH/pkg` — Go-managed compiled libraries
* `$GOPATH/src` — Go Source Code
    * `$GOPATH/src/github.com/golang/example` — The package we downloaded on the previous slide

---
template:green
background-image:url(bg.png)
background-size:50% auto
background-position:100% 0%

# Agenda

### — Go's "Hello, World"

### — Creating a Package

### — Personalizing "Hello, World"

### — Error Handling

---
template:title

## Part 1
# Hello, World

---
template:blue

# The `Package` Statement

```go
package main ⇐

import (
	"fmt"
)

func main() {
	fmt.Println("hello, playground")
}
```

## **Package Name** is `main`

### Package Name is used in code to reference _exported_ functions, variables, and types.

### Runnable packages **must** be named `main`

---
template:blue

# The `import` Statement

```go
package main

import ( ⇐
	"fmt"
)

func main() {
	fmt.Println("hello, playground")
}
```

## Two forms:

```go
        import (
          "pkgA"                         or                             import "pkgA"
          "pkgB"                                                        import "pkgB"
        )
```

### Import accepts an **Import Path**, distinct from Package Name.

---
template:blue

# The `import` Statement

```go
package main

import (
	"fmt" ⇐
)

func main() {
	fmt.Println("hello, playground")
}
```

### **Import Path** is `fmt`
### Non-URL Import path means built-in Package
### Package name is ALSO `fmt`
#### But it doesn't have to be.

---
template:blue

# Function Definition

```go
package main

import (
	"fmt"
)

func main() { ⇐
	fmt.Println("hello, playground")
}
```

### `main` function must accept no arguments and have no return

```go
func NAME() { BODY }
func NAME(Arg1Name Arg1Type, Arg2Name Arg2Type, …) { BODY }
func NAME(Arg1Name Arg1Type, Arg2Name Arg2Type, …) ReturnType { BODY }
func NAME(Arg1Name Arg1Type, Arg2Name Arg2Type, …) (Return1Name Return1Type,
  Return2Name Return2Type, …) { BODY }
```

---
template:blue

# Function Call

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello, playground") ⇐
}
```

### Only capitalized members are exported from packages

### `Println` is capitalized, so is exported by `fmt`

---
template:blue

# Return

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello, playground")
} ⇐
```

### Implicit `return`, only because `main` has no return type

---
template:blue

# Return

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello, playground")
	return ⇐
}
```

### Explicit `return` is OK too

---
template:title

## Part 2
# Creating a Package

---
template:blue
# Creating a new Package

### Import Path:
```
github.com/USERNAME/go101/hello
```

### Package Name:
```
main
```
#### (it's a runnable package)

---
template:blue

## Do this: make the directory
```
$ mkdir -p $HOME/go/src/github.com/USERNAME/go101/hello
```

## Do this: add some source code

#### Create `$HOME/go/src/github.com/USERNAME/go101/hello/hello.go` with your favorite editor.

#### Fill it with:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
}
```

#### _Tip: Copy this from `https://play.golang.org`_

---
template:blue

## Do this: build a binary

```
$ go build -v -o $HOME/hello github.com/USERNAME/go101/hello
github.com/pdbogen/go101/hello
```

#### Without `-o $HOME/hello`, the binary will be built in your current directory.

## Do this: run your binary
```
$ $HOME/hello
Hello, playground
```

#### `$HOME/hello` is a statically-linked binary; if you build on Linux, you can run it on Linux, without worrying about library versions.

---
template:blue

## Do this: Create a GitHub repo for your new package

### 1. Visit `https://github.com` and login
### 2. Click on the `+` next to your image in the top-right
### 3. Select "New repository" and name it `go101`. (Do not initialize it with a README)

---
template:blue

## Do this: Initialize and Push your new package

```
cd $HOME/go/src/github.com/USERNAME/go101
git init
git add hello
git commit -m "hello, world"
git remote add upstream git@github.com:USERNAME/go101.git
git push --set-upstream upstream master
```

![tada](tada.png)

---
template:title

## Part 3
# Personalizing Your Hello

---
template:green
## Our "Hello, World" is kind of impersonal.

_After all,_ my _name isn't "playground."_

## Let's make it better:

* Prompt the user to enter their name.
* Read the name in from the command-line and save it in a variable.
* Greet the user personally by name.

---
template:green
# Declaring Variables

### We declare variables with the `var` keyword:

```go
  var aNumber int       // default: 0
  var someString string // default: "" (empty string)
```

#### _The default value is called the "zero value."_

### We can initialize them at the same time:

```go
  var aNumber int = 10
  var someString string = "foo"
```

---
template:green
# Declaring Variables: Short Forms

### If we initialize it to an unambiguous type, we can leave off the type:

```go
  var aNumber = 10       // Still an `int`
  var someString = "foo" // Still a `string`
```

### In fact, we do this so often, there's an even shorter syntax:

```go
  aNumber := 10
  someString := "foo"
```

---
template:green
# Reading Input: A Magic Spell

### First, we need a place to hold the string we'll read.

```go
  var someString string
```

---
template:green
# Reading Input: A Magic Spell

### We'll also need to import the `fmt` package

```go
  import "fmt"

  var someString string
```

---
template:green
# Reading Input: A Magic Spell

### Now, we can call `fmt.Scanln` and pass it the _address_ of our string

```go
  import "fmt"

  var someString string

  fmt.Scanln(&someString)
```

#### — `&` is the address-of operator. It converts a variable of type `T` to a pointer type `*T`.

#### — If `someString` is a `string`, then `&someString` is a `*string`.

#### — If `somePtr` is a `*string`, then `*somePtr` is a `string`. (`&somePtr` would become a `**string`)

---
template:green
# Let's do it

### Outline
1. Print a prompt asking the user for their name.
2. Read the name and save it to a variable.
3. Print out a message greeting the user by name.

### Reminders

```
Declare Variables          Print a String                 Read a string
---------------------      -----------------              -----------------------
var someString string      fmt.Println("hi")              fmt.Scanln(&someString)
someString := "init"       fmt.Println("hi,", someStr)


Run your Code
----------------------------------------------------------------------
go run $HOME/go/src/github.com/USERNAME/go101/hello/hello.go
```

---
template:green
# Solution

```go
package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print(“What is your name? “)
	fmt.Scanln(&name)
	fmt.Println("Hello,", name, “!”)
}
```

---
template:title

## Part 4
# Error Handling

---
template:blue

# We've got both kinds of errors

## `panic`

* Severe errors & programming errors
* Functions like an exception
* Really more like a C `assert()`

## `error`

* Recoverable errors
* Returned value of `error` type
* `nil` (keyword for unset pointer) means "no error"

---
template:blue

# When to `panic`

--
template:blue

* Never

--
* Except for _very simple_ programs.

--
* And `Must…` functions
	* `Must…` functions are wrappers around functions that return `error` that instead panic when the wrapped function would error.
	* i.e., `regexp.Compile` and `regexp.MustCompile`
	* If a _static_ regular exprssion string is compiled and it has a syntax error, program probably can't continue; so `regexp.MustCompile` can be used to simply panic.
		* You can do this at global scope, so the program discloses this error and exits immediately on startup.

---
template:blue

# Recovering a `panic`

```
	func fearful() {
		panic("you're doin me a frighten")
		fmt.Printf("this doesn't run")
	}
```

## If the function `fearful` calls `panic`:

1. That function ends immediately
2. Functions it `defer`ed are called
3. Repeat #1 and #2 for the function that called `fearful`

---
template:blue

# Recovering a `panic`

```
	func fearful() {
		panic("you're doin me a frighten")
		fmt.Printf("this doesn't run")
	}
	func something() {
		fearful()
		fmt.Printf("this also doesn't run")
	}
```

--

* While `defer`ed functions are being called, they can call `recover`, which returns the object passed to `panic`
* If that function does not subsequently call `panic`, the `panic` ends.

---
template:blue

# Recovering a `panic`

```
	func fearful() {
		panic("you're doin me a frighten")
		fmt.Printf("this doesn't run")
	}
	func something() {
		defer func() { recover(); fmt.Printf("something we did called panic()") }()
		fearful()
		fmt.Printf("this *still* doesn't run")
	}
```

* While `defer`ed functions are being called, they can call `recover`, which returns the object passed to `panic`
* If that function does not subsequently call `panic`, the `panic` ends.

---
template:blue

# Recovering a `panic`

```
	func fearful() {
		defer func() { recover(); fmt.Printf("something we did called panic()") }()
		panic("you're doin me a frighten")
		fmt.Printf("this doesn't run")
	}
	func something() {
		fearful()
		fmt.Printf("now this runs")
	}
```

* While `defer`ed functions are being called, they can call `recover`, which returns the object passed to `panic`
* If that function does not subsequently call `panic`, the `panic` ends.

---
template:blue

# When to `error`

* _Any_ time something can fail
	* Mathematical operations that may be undefined (square root of negatives, etc.)
	* Network calls
	* Database lookups

--
* Any time you call something that returns an `error`
	* Add some context with `fmt.Errorf`:
		```
			func convert(foo string) (string, error) {
				var bar, err = doSomething(foo)
				if err != nil {
					return "", fmt.Errorf("converting %s: %s", foo, err)
				}
				return bar, nil
			}
			func main() {
				conversion, err := convert(os.Args[1])
				if err != nil { panic(err) }
				fmt.Print(conversion)
			}
		```

---
template:blue

# Let's do it

### Outline
1. Add a `validate` function that accepts one argument- the user's name as a `string`, and that returns one value, an error.
2. Check if the name exactly matches yours; if it _doesn't_, return an error (created using `errors.New`) describing the problem in English. (If the name _does_ match, return `nil` to indicate "no error")
3. In your `main` function, call validate after the user provides their name, and print the error message instead of a greeting, if the error is not nil.

### New Syntax

For an `if` statement, you can declare a variable and check its value at the
same time by separating these two statement with a semi-colon, like this:

```go
if intValue := someFunction(); intValue != 1 {
   // `intValue` can be used here
}

// `intValue` does not exist outside the `if`
```

---
template:blue

# Solution

```go
package main

import (
	"fmt"
)

func validate(name string) error {
  if name != "patrick" { return errors.New("you are not Patrick!"); }
  return nil
}

func main() {
	var name string
	fmt.Print(“What is your name? “)
	fmt.Scanln(&name)
	if err := validate(name); err != nil {
	  fmt.Println("Uh, oh:", err)
  } else {
    fmt.Println("Hello,", name, “!”)
	}
}
```
