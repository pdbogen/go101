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

Collections (maps, slices, arrays)

Functions as Objects

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

# Collections: The Slice

* For any type T, "slice Of Ts" is spelled `[]T`
    * "N-length Array of Ts" is spelled `[N]T`
* `[][]string` is pronounced "slice of slices of strings"

## Constants

```go
  var ints []int = []int{1, 2, 3}
  strings := []string{"one", "two", "three"}
```

## Operations

```go
  ints[0] = 2            // Sub-element assignment
  ints = append(ints, 4) // Appending items to the end. This _might_ create a copy of the data.
  len(strings) == 3       // Length
  strings[0] == "one"    // Subscripting 
```

???
...but we rarely use arrays directly. explain they're fixed length, can't be appended.

---
template:blue

# Collections: The Slice

* Slices are like pointers- the zero value of a slice variable is nil:
  ```go
    var zero []int
    zero == nil // true!
  ```

* Assigning a slice creates a reference to the same data:
  ```go
  a := []int{1,2,3} ← a is a new slice
  b := a            ← b is a reference to a
  a[0] = 2          ← we change a[0]
  b[0] == 2         ← b[0] is changed as well
  ```

---
template:blue

# Collections: The Slice: Internals

* Slices have three fields, internally:
    * A value indicating the length -- how much of the array we're using -- `len(s)`
    * A pointer to the underlying array -- implying the first element
    * A value indicating the capacity -- the total length of the array -- `cap(s)`

```
         .---------.----------.-----.
  Slice: | len = 8 | cap = 12 | ptr |
         '---------'----------'-----'
             ___________________/        .- len - 1      . cap - 1
            /                           /               /
           v                           v               v
         .---.---.---.---.---.---.---.---.---.---.---.---.
  Array: | 0 | 1 | 1 | 3 | 1 | 2 | 4 | 5 |   |   |   |   |
         '---'---'---'---'---'---'---'---'---'---'---'---'
           0   1   2   3   4   5   6   7   8   9   10  11
```

---
template:blue

# Collections: The Slice: Append

* `append` increases the length, as long as the length is less than the capacity
    * Actually, `append` makes a new _slice_, with a pointer to the same data, if there's room. That's why we assign the result of `append` to the original slice.
    * If there isn't room, `append` copies the entire array.

```
  Before:
         .---------.----------.-----.
  Slice: | len = 8 | cap = 12 | ptr |
         '---------'----------'-----'
         .---.---.---.---.---.---.---.---.---.---.---.---.
  Array: | 0 | 1 | 1 | 3 | 1 | 2 | 4 | 5 |   |   |   |   |
         '---'---'---'---'---'---'---'---'---'---'---'---'

  After s = append(s,9)
         .---------.----------.-----.
  Slice: | len = 9 | cap = 12 | ptr |
         '---------'----------'-----'
         .---.---.---.---.---.---.---.---.---.---.---.---.
  Array: | 0 | 1 | 1 | 3 | 1 | 2 | 4 | 5 | 9 |   |   |   |
         '---'---'---'---'---'---'---'---'---'---'---'---'
```

???
discuss performance implications of naive append

---
template:blue

# Collections: The Slice: Re-slice

* We can re-slice, to move the slice around the underlying array: `s = s[low:high]`
    * `0 <= low < len(s)`; default is `0`
    * `low < high < len(s)`; default is `len(s)`
    * The new length is `high-low`.
    * The new capacity is `cap - low`.

```
         .---------.----------.-----.
  Slice: | len = 8 | cap = 12 | ptr |
         '---------'----------'-----'
         .---.---.---.---.---.---.---.---.---.---.---.---.
  Array: | 0 | 1 | 1 | 3 | 1 | 2 | 4 | 5 |   |   |   |   |
         '---'---'---'---'---'---'---'---'---'---'---'---'

  s[1:8]
         .---------.----------.-----.
  Slice: | len = 7 | cap = 11 | ptr |
         '---------'----------'-----'
         .---.---.---.---.---.---.---.---.---.---.---.
  Array: | 1 | 1 | 3 | 1 | 2 | 4 | 5 |   |   |   |   |
         '---'---'---'---'---'---'---'---'---'---'---'
```

???
notice that the capacity has decreased. we can't get capacity back.
mention we can re-slice an array to get a slice pointing to it.

---
template:blue

# Collections: The Slice: Re-slice

```
         .---------.----------.-----.
  Slice: | len = 8 | cap = 12 | ptr |
         '---------'----------'-----'
         .---.---.---.---.---.---.---.---.---.---.---.---.
  Array: | 0 | 1 | 1 | 3 | 1 | 2 | 4 | 5 |   |   |   |   |
         '---'---'---'---'---'---'---'---'---'---'---'---'

  s[0:3]
         .---------.----------.-----.
  Slice: | len = 3 | cap = 12 | ptr |
         '---------'----------'-----'
         .---.---.---.---.---.---.---.---.---.---.---.---.
  Array: | 0 | 1 | 1 |   |   |   |   |   |   |   |   |   |
         '---'---'---'---'---'---'---'---'---'---'---'---'
```

???
since we don't change `low`, capacity doesn't decrease, but our data is "gone".

---
template:blue

# Collections: The Slice: Iterating

* C-style, `for i := 0; i < len(s); i++` is alright, but is a lot of boilerplate.
* Go gives us the `range` operator.

```go
for i := range s {
  // do something with s[i]
}

for i, v := range s {
  // s[i] == v, even cleaner
}

for _, v := range s {
  // if we don't use the index, we can skip saving it to a variable
}
```

* Bonus: syntax is identical for maps (coming soon to a slide near you!)

---
template:blue

# Do It: Implement Common Operations on Slices

* `func Push(in []int, item int) (out []int) {…}`
    * Push treats a slice as a stack, and pushes an item on to it. It returns the new stack as a slice:
* `func Top(in []int) (top int) {…}`
    * Top treats a slice as a stack, and returns the top item (the most recently pushed). It panics if there's no top item.
* `func Pop(in []int) (out []int) {…}`
    * Pop treats a slice as a stack, and removes the top item. It panics if there's no top item.
* Bonus: Implement PopAndTop, which does both of the previous two: `func PopAndTop(in []int) (top int, out[]int) {…}`
* Bonus bonus: Rewrite the previous three to return an error instead of panicing

---
template:blue

# Project: Bring it all together

* No cheat sheet this time; use:
    * the ref-spec:  `https://golang.org/ref/spec`
    * and pkg docs:  `https://golang.org/pkg/strconv`
* Let’s create a basic postfix calculator (plus, minus, multiply, divide):
    * A function that accepts a slice of strings (`[]string`), where each element is either an operand (a number) or an operator
    * The function should iterate over the slice:
        * use strconv.Atoi to determine if the element is a number or not
        * Numbers are pushed onto a stack (we built one, remember?)
        * Operators pop two numbers off of the stack, perform their operation, and push the result onto the stack
        * When there are no more elements in the input, the function should return the top item on the stack- the result.
    * Ex: `calculate([]string{"1", "2", "+", "4", "*"})` should return `12`.

---
template:title

# Break

Relax, maybe meditate.

Have some coffee.

Both might be counterproductive, though.

---
template:blue

# Maps

* Easier than slices somehow
* Also pointer-like, but none of these re-slice shenanigans
* For two types K and V, `mapping of Ks to Vs` is spelled `map[K]V`

```go
  numerals := map[int]string{ 0: "zero", 1: "one", 2: "two" }
  numerals_l10n := map[int]map[string]string{
    0: map[string]string{ "english": "zero", "latin": "nulla" },
  }
  numerals[0] == "zero"
  numerals[99] == ""

  numerals_l10n[0]["english"] == "zero"
  numerals_l10n[0]["spanish"] == ""
  numerals_l10n[99] == nil
  numerals_l10n[99]["english"] == "" ← _accessing_ a `nil` map is OK
  numerals_l10n[99]["english"] = "ninety-nine" ← panic! _writing_ to a `nil` map is _not_ OK
```

---
template:blue

# Maps -- "comma-ok"

* Sometimes we want to differentiate between an actually-zero value and a missing key
* When we assign variables from a map, we can optionally assign _two_ variables.

```go
  numerals := map[int]string{ 0: "zero", 1: "one", 2: "two" }
  n, ok := numerals[0] // ok is a bool
  ok == true // numerals[0] _does_ exist

  numerals[99] // returns ""
  n, ok := numerals[99]
  ok == false // numerals[99] _does not_ exist
```

---
template:blue

# Functions as Objects

* Almost exactily what you expect
* Old hat: `func foo(a int, b string) (float32, error) {…}`
* Foo is just a constant declaration of a function! Its type is `func(a int, b string) (float32, error)`:

```go
  var fooObj func(a int, b string) (float32, error) = foo
  fooObj(1, "bar") == foo(1, "bar")

  var undefined func(int i)(error)
  undefined == nil // the zero value of a function is `nil`
  undefined(0) ← panic! cannot call a nil function

  anonymous := func(int a)(error){
    if a == 0 { return errors.New("no zeroes allowed!") }
    return nil
  }
  fmt.Println(anonymous(1)) // prints `nil`
  fmt.Println(anonymous(0)) // prints "no zeroes allowed!"
```

---
template:blue

# Side note: Closures

* Simply: Functions can access variables from the scope in which they're defined

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
template:blue

# Do it: Embrace and Extend

* Take the calculator from before, and add in the ideas of maps and function types.

* Instead of a static set of operators, have a mapping of operator strings to functions that execute them.

* Hint: `map[string]func(opA, opB int)(int)`

---
template:title

# Thanks!

## Onward, to Chapter 2

![tada](tada.png)
