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
* Errors (`error`)

---
template:blue

# "Different Types" and Conversions

* `int8` and `int16` are "different types"
* As different as `string` and `float32`
* We can convert:
  * amongst numeric types
  * integer types to strings (special case, converts to unicode; doesn't render integer)
  * string to/from byte slices or rune slices

```go
  var i8 int8 = 65
  var i16 int16 = 16384

  i8 == i16       // compiler error!
  i8 == int8(i16) // compiles, but integer overflow 

  s := string(i8) // string is `"A"`, because integer 65 is Unicode A
```

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

# Errors

* Errors are `interfaces`, which we'll cover later.

```go
  import (
    "errors"
    "fmt
  )

  // Many functions return an error
  func GetStringCouldFail() (string,error) {
    // Return a new "constant" error
    return errors.New("this function is not implemented!")
  }

  output, err := GetStringCouldFail()
  if err != nil { // errors are `nil` when no error occurred
    fmt.Printf("getting string failed: %s", err)
  }
```

---
template:title

![tada](tada.png)
