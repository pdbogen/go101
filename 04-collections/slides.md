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

## Collections

---
template:blue

# Collections: Arrays

* For any type T and integer length N, "N-length array of T" is spelled `[N]T`
* `N` is _part of the type_
  * `[5]int` is a **different type** from `[3]int`: neither comparable nor assignable
* Arrays cannot grow or shrink

## Constants

```go
  var ints [3]int = [3]int{1,2,3}
  var strings [5]string = [5]string{"one", "two", "three", "four", "five"}
```

## Operations

```go
  ints[0] = 2     // sub-element assignment
  len(ints) == 3  // length
```

---
template:blue

# Collections: Arrays

## Assignment

* Assigning an array **copies** the data (this can be expensive!)
* To pass by reference, use a Slice (stay tuned!)

## Comparison

* Comparing two arrays compares their contents

```go
  a := [5]int{1,2,3}
  b := [5]int{2,2,3}
  a == b // false!
  b[0] = 1
  a == b // true!
```

---
template:blue

# Collections: Arrays

## Zero Value

* The zero value of an array is the array with all elements set to their zero values

```go
  var a [5]int // uninitialized! value is [5]int{0,0,0,0,0}
```

---
template:blue

# Collections: Slices

* For any type T, "slice Of Ts" is spelled `[]T`
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
  len(strings) == 3      // Length
  strings[0] == "one"    // Subscripting 
```

---
template:blue

# Collections: Slices

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

# Collections: Slices: Internals

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

# Collections: Slices: Append

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

# Collections: Slice expressions

* The slice expression `s[low:high]` creates slices from arrays or from other slices
    * `low` ranges from `0` to `len(s)` inclusive; `0` if blank
    * `high` ranges from `low` to `len(s)` inclusive; `len(s)` if blank
    * The new length is `high - low`. _(tip: work back from the length you want)_
    * The new capacity is `cap - low`.

* Recipes:
  * `s[:]` creates a new slice object from an existing array or slice- shares data!
      * same as `s[0:]` or `s[:len(s)]` or `s[0:len(s)]`
  * `s[1:]` creates a slice with the first item removed
  * `s[:len(s)-1]` creates a slice with the last item removed

---
template:blue

# Collections: Slice expressions

* `s[1:8]` has `low` higher than `0`, so length _and_ capacity decrease

```
         .---------.----------.-----.
  Slice: | len = 8 | cap = 12 | ptr |
         '---------'----------'--+--'
           .--------------------/
           v
  Array: | 0 | 1 | 1 | 3 | 1 | 2 | 4 | 5 |   |   |   |   |

  s[1:8]
         .---------.----------.-----.
  Slice: | len = 7 | cap = 11 | ptr |
         '---------'----------'--+--'
               .----------------/
               v
  Array: | ? | 1 | 1 | 3 | 1 | 2 | 4 | 5 |   |   |   |   |
```

???
notice that the capacity has decreased. we can't get capacity back.
mention we can re-slice an array to get a slice pointing to it.

---
template:blue

# Collections: Slice expressions

* `s[0:3]` has `high` less than `len(s)` reduces length but not capacity, but the data is gone

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

# Collections: Iterating: `range`

* By key or index
```go
// i is each index (or key, for maps)
for i := range s {
}
```

* By key or index _and_ value
```go
// each pair of index/value (or key/value, for maps)
for i, v := range s {
}
```

* By value
```go
// discard index/key, iterate over each value
for _, v := range s {
}
```

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
    * Ex: `calculate([]string{"10", "2", "/"})` should return `5`

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
