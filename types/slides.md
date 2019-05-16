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

## Type System

Simple Types

User-defined Types

---
template:blue

# Simple Types

* Boolean (`bool`, `true` or `false`)
* Numeric (`int`, `uint`, `float`)
  * `int8`, `int16`, `int32` (aka `rune`), `int64`
  * `uint8` (aka `byte`), …
  * `float32`, `float64`
  * `complex64`, `complex128`
* String (`string`)

---
template:blue

# Conversions

## Implicit

```go
  // This does not compile!
  var example int8 = 1
  var exampleB int16 = example
  // cannot use example (type int8) as type int16 in assignment
```

## Explicit conversions look like function calls

```go
  var example int8 = 1
  var exampleB int16 = int16(example)
```

### We'll see more in User-Defined Types

---
template:blue

# Strings

## Constants

```go
  "string" // standard form
  `string` // *every byte* from first ` to until the next `
```

## Operations

```go
  s[0] = 'H' // NO! Immutable! Does not work.

  "a" + "b" == "ab" // Concatenation
  len("ohai") == 4  // Length
  "ohai"[0] == 'o'  // Subscripting; 'o' is a `rune`
```

* Other more complex operations are in the `strings` package

---
template:blue

---
template:blue

# Do it: Embrace and Extend

* Take the calculator from before, and add in the ideas of maps and function types.

* Currently, you probably have a list of operations (`result = opA + opB`)
* Replace that list with a map, which maps the operand, a string `"+"`, to a function with type `func(int,int)(int)`
* Your result should have no `switch` or nested `if-else` for dispatching to operators!

* Hint: `map[string]func(opA, opB int)(int)`

---
template:title

# Thanks!

## Onward, to Chapter 2

![tada](tada.png)
