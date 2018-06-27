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

# HTTP Servers in Go

## Basic Servers
## Frameworks

---
template:blue

# Hello, World Wide Web

```go
func main() {
  http.HandleFunc(
    "/foo",
    func(w http.ResponseWriter, r *http.Request) {
      rw.Write([]byte("Hello, world."))
    },
  )
  http.ListenAndServe(":8080", nil);
}
```

* The "Hello World" of Go web servers

---
template:blue

# Pattern Matching

```go
func main() {
  http.HandleFunc(
    "/foo",                                                 ←  A "pattern"
    func(rw http.ResponseWriter, r *http.Request) {
      rw.Write([]byte("Hello, world."))
    },
  )
  http.ListenAndServe(":8080", nil);
}
```

* The default `ServeMux` (`DefaultServeMux`) matches URL "patterns":

> Patterns name fixed, rooted paths, like "/favicon.ico", or rooted subtrees, like "/images/" (note the trailing slash). Longer patterns take precedence over shorter ones

---
template:blue

# Handler Anatomy

```go
func main() {
  http.HandleFunc(
    "/foo",
    func(rw http.ResponseWriter, r *http.Request) {         ←  A "handler"
      rw.Write([]byte("Hello, world."))
    },
  )
  http.ListenAndServe(":8080", nil);
}
```

* `HandleFunc` registers a function as a handler (clever, huh?)
* `Handle` registers an object that implements the interface `Handler { ServeHttp(w http.ResponseWriter, req *http.Request) }`

---
template:blue

# Writing

```go
func main() {
  http.HandleFunc(
    "/foo",
    func(rw http.ResponseWriter, r *http.Request) {
      rw.Write([]byte("Hello, world."))                     ←  Begin writing a resopnse body
    },
  )
  http.ListenAndServe(":8080", nil);
}
```

* Calling `Write` for the first time implicitly:
  * Stop accepting Request Body data from the client (HTTP 1.x; HTTP/2 with a client that supports it can do both at once)
  * Writes any headers that have been queued up via `ResponseWriter.Header().Add(…)`
  * Writes an HTTP status `200 OK` (for other statuses, we need to call `.WriteHeader(int)` ourselves)
* We can then continue to call `Write` to send more data, until we return from the handler.

---
template:blue

# Try it out!

* Wire up an HTTP server to dump the message history from one of your micro-slack channels
* Either:
    * Register a handler specific to the channel; or
    * Register a handler that parses `http.Request.URL` to determine the channel

Basic Structure:

```go
func main() {
  http.HandleFunc(
    "/foo",
    func(rw http.ResponseWriter, r *http.Request) {
      rw.Write([]byte("Hello, world."))
    },
  )
  http.ListenAndServe(":8080", nil)
}
```
