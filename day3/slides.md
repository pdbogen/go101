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

## User-Defined Types
## Ecosystem

---
template:title

## Part 1
# User-Defined Types

---
template:blue

# Anatomy of a Type

* We've seen: `int`, `float`, `string`, `bool`, `[]T`, `map[T1]T2`, `func(/* parameter list */) (/* return list */)`
    * Where `T`, `T1`, etc. name other types
* These are recursive: Just like we can have `[]int` (slice of ints), we can have:
    * `[][]int` (slice of slices of ints) or
    * `[]map[string]int` (slice of maps mapping strings to ints) or 
    * `[][][]map[[]string][][]int` (slice of slices of slices of maps mapping slices of strings to slices of slices of ints)
    * `map[string]func(int,int) int` (map of strings to functions that take two integers and return an integer)

---
template:blue

# Type Declarations

* We can declare new types made from fundamental types, using the `type` keyword:
    ```go
      type UserId int
      type TeamId int
      type AttributeMap map[string]string
    ```
--

* The declaration `type UserId int` binds the **type name** `UserId` to a new type with **underlying type** `int`.
--

* `UserId`, `TeamId`, and `int` are all **different**
    * Not Assignable: `var u UserId; var i int = u` does not compile:
        * `cannot use u (type UserId) as type int in assignment`
    * We _can_ convert, because the _underlying type_ is the same:
        * **new type** to **underlying type**: `integer := int(u)`
        * **underlying type** to **new type**: `user_id := UserId(i)`
        * between two types with the same underlying type: `team_id := TeamId(u)`

---
template:blue

# Compound Types

* There are two additional building blocks we haven't properly met yet:
    * `struct { /* field declarations */ }`
    * `interface { /* method spec */ }`

* Structs first.
* Interfaces later.

---
template: blue

# Structs

```go
struct {
  UserName string                     or           struct { UserName string; UserId int }
  UserId   int
}
```

* We can use this anywhere we might use `int`:
    ```go
      int_map    := map[string]int                                  { "a": 1        }
      struct_map := map[string]struct{ UserName string; UserId int }{ "a": {"b", 1} }
    ```
* **More often,** we declare a new type, with a struct as the underlying type:
    ```go
      type User struct {
        UserName string
        UserId   int
      }
    ```
---
template: blue
# Exercise: Struct Basics

* Imagine we make a tool that lets users exchange messages in channels.
* Define three new structs:
  * `User`, which contains a user's name and an `enabled` bool.
  * `Message`, which contains a pointer to the author, a timestamp (hint: `time` package, `time.Time`), and the text of a message.
  * `Channel`, which contains a slice of pointers to Users and another slice of pointers to messages
* Write some simple code to work with these:
  * Create a user
  * Add a user to a channel
  * Add a message to a channel
---
template: blue
# Structs: Method Set

* Functions have a property we have not yet used, called the "receiver":
    ```go
      type Something struct {}
      func (s Something) Print() { fmt.Println(s) }
      func (s *Something) Combobulate() {…}
      …
      func main() {
        athing := Something{}
        athing.Print()
      }
    ```
* `Print` is part of the "method set" of `Something`
* `Print` is _also_ part of the "method set" of `*Something` (pointer-to-`Something`); because _"The method set of the pointer type `*T` is the set of all methods declared with receiver `*T` or `T`."_
* `Combobulate` is in the method set of `*Something` _but not_ `Something`.
---
template: blue
# Exercise: Method Set
* Define a method with a receiver of `*Channel` (_pointer to Channel_) that retrieves **a slice of message pointers** for messages in the channel, where the user is enabled.
* Make sure you have some enabled users and some disabled users.
* Make sure you have some messages from both type of user.
* Be sure your method returns `[]*Message`
---
template: blue
# Interfaces

* An `interface` describes a method set, like this:
    ```go
      // Writer implements `Write(s)`, which accepts a string argument and returns an error if
      // it cannot be written.
      type Writer interface {
        Write(string) error
      }
    ```
* An `interface` is a (pointer-flavored) type:
    ```go
      var w Writer // default/zero value is `nil`
      w.Write("something") // but this will crash, because w is nil
    ```
---
template: blue
# Interfaces

* A struct that "implements" an interface can be assigned to variables of that interface type:
    ```go
      type AnActualWriter struct { }
      func (a *AnActualWriter) Write(string) error { return nil }

      var w Writer = &AnActualWriter{}
      w.Write("something") // this doesn't crash!
    ```
--

* Remember how pointer-type method sets differ from non-pointer method sets?
    ```go
      var w Writer = &AnActualWriter{} // `&AnActualWriter{}` has type `*AnActualWriter`
      var w2 Writer = AnActualWriter{} // this one is type non-pointer `AnAntualWriter`
    ```
    * Because the non-pointer type `AnActualWriter` does not have methods that are for the pointer-type, `AnActualWriter` does not have a `Write` method; so it doesn't implement the `Writer` interface, and cannot be assigned to a variable of that type.
---
template: blue
# Exercise: Interfaces

* Define an interface, `Sender`, that describes the sending of a message to a destination.
* Add receiver functions as necessary so that your `Channel` struct type implements this interface
* Create a new `OneToOneChat` struct type that _also_ implements this interface.
* Hint: You can hint to the compiler that you expect a struct to satisfy an interface by assigning an instance of the struct type to a variable of the interface type:
    ```go
      type Demonstrator interface {
        Demonstrate()
      }
      type Demonstration struct {}
      func (d *Demonstration) Demonstrate() {}

      var evidence Demonstrator = (*Demonstration)(&Demonstration{})
      // This does the same thing, but saves a little bit of RAM:
      var _        Demonstrator = (*Demonstration)(nil)
    ```
