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

# Interlude (Bear with me…)
## New Syntax: `defer`

```go
func main() {
  fileHandle := os.Open("some file")
  defer fileHandle.Close()
  fmt.Println("hello, world!")
}
```

* `defer` causes a function call to run when the function returns
* Very useful for cleaning up resources, like closing file handles or network connections.

* **Try it**: Update your "hello world" so that `main` defers a print of `Goodbye!`.

---
template:blue

# Recovering a `panic`

```go
  func main() {
    fearful()
    fmt.Println("this doesn't run")
  }
  func fearful() {
    panic("you're doin me a frighten")
    fmt.Println("this doesn't run")
  }
```

## If the function `fearful` calls `panic`:

1. That function ends immediately
2. Functions it `defer`ed are called
3. Repeat #1 and #2 for the function that called `fearful`

---
template:blue

# Recovering a `panic`

```go
  func main() {
    fearful()
    fmt.Println("this doesn't run")
  }
  func fearful() {
    panic("you're doin me a frighten")
    fmt.Println("this doesn't run")
  }
```

--

* While `defer`ed functions are being called, they can call `recover`, which returns the object passed to `panic`
* If that function does not subsequently call `panic`, the `panic` ends.

---
template:blue

# Recovering a `panic`

```go
  func dontpanic() { fmt.Printf("Swallowing panic (maybe): %v", recover()) }
  func main() {
    defer dontpanic()
    fearful()
    fmt.Println("NOW, this runs")
  }
  func fearful() {
    panic("you're doin me a frighten")
    fmt.Println("this doesn't run")
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

```go
  func convert(input string) (string, error) {
    var result, err = doSomething(input)
    if err != nil {
      return "", fmt.Errorf("converting %s: %s", input, err)
    }
    return result, nil
  }
  func main() {
    conversion, err := convert(os.Args[1])
    if err != nil { panic(err) }
    fmt.Println(conversion)
  }
```

---
template: blue

# New Syntax

For an `if` or `switch` statement, you can declare a variable and check its
value at the same time by separating these two statement with a semi-colon,
like this:

```go
if intValue := someFunction(); intValue != 1 {
 // `intValue` can be used here
}

if err := maybeFunction(); err != nil {
  // Handle the non-nil error `err`
}

// neither `intValue` nor `err` exist here, outside the `if`

select option := someOption; option {
  case "a": …
  case "b": …
  default: …
}
```

---
template:blue

## Let's do it

### Outline
1. Add a `validate` function that accepts one argument, the user's name (a `string`), and that returns one value, an `error`.
2. If the name exactly matches yours, return `nil`. if it _doesn't match_, return an error (created using `errors.New`) describing the problem in English.
3. Use `validate` in your main function to validate the name, or print an error.

### Hints

* Read about basic language features at the "ref-spec": `https://golang.org/ref/spec`
* Read about standard libraries like `errors` via `golang.org/pkg/…`: `https://golang.org/pkg/errors`

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
    return
  }
  fmt.Println("Hello,", name, “!”)
}
```
