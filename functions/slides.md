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

## Functions

---
template:green

### Functions, Formally

```go
//.`func`, a literal keyword that begins every functions
//
func (m MyType) Method(a int, b int) string { … }
```

* All functions, or function types, begin with the `func` keyword

---
template:green

### Functions, Formally

```go
//    . "receiver" -- optional, a type declared in the same package
//   /
func (m MyType) Method(a int, b int) string { … }

var m MyType
m.Method(…)
```

* A function with a receiver is in the Method Set of the received type
    * Pointer receivers `m *MyType` allow modification of the object.
    * Non-pointer receivers `m MyType` make a _copy_ of `m`, but a shallow copy- pointers still point to the same things.

```go
func Method(a int, b int) string { … }
```

* A function without a receiver is package-level.

---
template:green

### Functions, Formally

```go
//               . the method name
//              /
func (m MyType) Method(a int, b int) string { … }
func (m MyType) privateMethod() { … }
func PackageFunc() { … }
func packagePrivate() { … }
```

* Capitalized function names (`Method`) are exported by the package
* Non-capitalized function names (`method`) are private to the package

---
template:green

### Functions, Formally

```go
//                     . the argument list, two named arguments
//                    /            
func (m MyType) Method(a int, b int) string { … }
//                       . same as above
//                      /
func (m MyType) Compact(a, b int) string { … }

//                       . no arguments
//                      /
func (m MyType) Niladic() string { … }

//                          . un-named arguments
//                         /  (must be the last argument! one only!)
func (m MyType) Anonymized(int,string) string { … }
func (m MyType) Anonymized(_ int, _ string) string { … }

//                        . zero or more integers
//                       /
func (m MyType) Variadic(numbers ... int) { … }

//                     . named argument
//                    /       . un-named argument
//                    |      /          . variadic argument
//                    |      |         /
func (m MyType) Combo(a int, _ string, others ... byte) { … }
```

---
template:green

### Functions, Formally

```go
//                  the return type .
//                                   \            
func (m MyType) Method(a int, b int) string { … }

//           named and pre-declared .
//                                   \
func (m MyType) Method(a int, b int) (s string) { … }

//                 two return types .
//                                   \
func (m MyType) Method(a int, b int) (string, error) { … }

//     two return types, both named .
//                                   \
func (m MyType) Method(a int, b int) (s string, err error) { … }
```

---
template:green

### Functions, Formally

```go
//                       the function body .
//                                          \
func (m MyType) Method(a int, b int) string { … }

```

* The simplest and most complicated part of a function.

---
template:green

# Functions as Objects

* Almost exactily what you expect
* Old hat: `func foo(a int, b string) (float32, error) {…}`
* `foo` is just a constant declaration of a function! Its type is `func(a int, b string) (float32, error)`:

```go
  var fooObj func(a int, b string) (float32, error) = foo
  fooObj(1, "bar") == foo(1, "bar")

  var undefined func(int i)(error)
  undefined == nil // the zero value of a function is `nil`
  undefined(0)     // panic! cannot call a nil function

  anonymous := func(int a)(error){
    if a == 0 { return errors.New("no zeroes allowed!") }
    return nil
  }
  fmt.Println(anonymous(1)) // prints `nil`
  fmt.Println(anonymous(0)) // prints "no zeroes allowed!"
```

---
template:green

# Side note: Closures

* Simply: Functions can access (read **and** change) variables from the scope in which they're defined

```go
  func main() {
    a := 1
    b := "two"

    // Create _and execute_ an anonymous function
    func() {
      fmt.Println(a) // prints 1
      fmt.Println(b) // prints "two"
    }()
  }
```

---
template:green

# Do It: Middleware

* Functions can accept functions as arguments
* Functions can return other functions
* And both at the same time

* We can implement cross-cutting concerns (like logging or authentication) as middleware, when our functions have common signatures

```go
  func Plus(a, b int) (int) { … }
  func Minus(a, b int) (int) { … }
  func Logged(func(int,int)(int)) func(int,int)(int) { … }

  func main() {
    plus := Logged(Plus)
    fmt.Println("1+1 is", plus(1,1))
    fmt.Println("5-4 is", Logged(Minus)(5,4))
  }
```

* Write the implementation of Plus, Minus, and Logged!
