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

## Error Handling

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
template:green

# Interlude
## New Syntax: `defer`

```
	func main() {
	  fileHandle := os.Open("some file")
	  defer fileHandle.Close()
		fmt.Println("hello, world!")
	}
```

* `defer` causes a function call to run when the function returns
* Very useful for cleaning up resources, like closing file handles or network connections.

* Try it: Update your "hello world" so that `main` defers a print of `goodbye`.

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
	* Add some context with `fmt.Errorf`
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
1. Add a `validate` function that accepts one argument- the user's name as a `string`, and that returns one value, an `error`.
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
