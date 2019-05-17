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

## Structs
## User-Defined Types

---
template:blue

# Type Declarations

* We can declare new types made from fundamental types, using the `type` keyword:
    ```go
      type UserId int
      type TeamId int
    ```
--

* The declaration `type UserId int` binds the **type name** `UserId` to a **new type** with **underlying type** `int`.
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

# Compound Types: Structs

* `struct`s are **new types** composed of zero or more existing types

    ```go
    struct {
      UserName string                     or           struct { UserName string; UserId int }
      UserId   int
    }
    ```

--
* We can use this anywhere we might use `int`:
    ```go
      s := struct{ UserName string }{ "bob" }
      func f( user struct { UserName string } ) {…}
    ```

--
* **More often,** we declare a new type, with a struct as the underlying type:
    ```go
      type User struct {
        UserName string
        UserId   int
      }
    ```

---
template:blue

# Struct Types

```go
  type User struct {
    UserName string
    UserId   int
  }
```

```go
  // Struct constant with named fields
  bUser := User{ UserName: "bob", UserId: 2 }

  // With named fields, all fields are optional
  var cUser = User{ UserName: "charlie" }

  // …and order doesn't mater
  var dUser = User{ UserId: 3, UserName: "diane" }

  // Struct constant with ordered fields: all fields required, in exact order
  aUser := User{ "alice", 1 }
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
template:title

![tada](tada.png)
