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

## Concurrency
## Type System, Pt2
## Ecosystem

---
template:title

## Part 1
# Concurrency

---
template:blue

# Goroutines: `go` keyword

* Less like `fork`, more like `bash`/`sh` `&`
* Any function can be run as a Goroutine:
  ```go
    package main

    import "fmt"

    func foo() { fmt.Println("My name is foo.") }

    func main() {
      fmt.Println("Before foo")
      go foo()
      fmt.Println("After foo?")
    }
  ```

---
template:blue

# Semi-Cooperative

* But wait, that didn't work.
  * Output:
    ```
    Before foo
    After foo?
    ```
  * Main never "yields" to let the goroutines actually run
  * We need some kind of synchronization or at least a delay
* Goroutines are semi-cooperative
  * Cooperative: Another routine/thread cannot run unless the current routine/thread "yields"
  * Preemptive: Routines/threads can be stopped mid-statement
* Goroutines need to yield, but lots of things implicitly yield

---
template:blue

# Fixing Foo: Easy & Bad Way

* Add a short delay to the end of `main`:
  ```go
    package main

    import "fmt"
    import "time"

    func foo() { fmt.Println("My name is foo.") }

    func main() {
      fmt.Println("Before foo")
      go foo()
      fmt.Println("After foo?")
      time.Sleep(time.Millisecond)
    }
  ```
* Output:
  ```
  Before foo
  After foo?
  My name is foo.
  ```

---
template:blue
# Fixing Foo: (Slightly) Harder Way

* Use a synchronization object
* `sync.WaitGroup` from `sync` package is appropriate:
  ```go
    package main

    import "sync"
    
    func foo(wg *sync.WaitGroup) {
      fmt.Println("My name is foo.")
      wg.Done() // decrements the waitgroup counter
    }

    func main() {
      wg := &sync.WaitGroup{} // Create a new zero-valued WaitGroup and get a pointer to it
      wg.Add(1) // We're going to be starting one goroutine
      fmt.Print("definitely before")
      go foo(wg)
      wg.Wait() // Yield & block until the waitgroup counter reaches zero
      fmt.Print("definitely after")
    }
  ```

???
The sync documentation instructs is not to copy a WaitGroup, which is how we know to use pointers for it.

---
template:blue

# Try it: Five Goroutines went to market

* Start five goroutines, each of which prints a different message.
* Use a `sync.WaitGroup` to ensure each goroutine gets to run before `main` exits
* Bonus: Try a mix of regular function definitions and anonymous functions

```go
  wg := &sync.WaitGroup{} // Initialize our WaitGroup
  wg.Add(1) // Indicate we're adding one to the counter
  go firstPrint(wg) // firstPrint needs to call `wg.Done()`
  wg.Wait() // Wait until the counter reaches zero
```

---
template:blue
# Implicit & Explicit Yields

* Explicit yield via `runtime` package:
  ```
    import "runtime"
    
    …
    runtime.Gosched()
    …
  ```
* Implicit Yield
  * Definitely anything that waits on a syscall or sleeps via `time.Sleep`
  * Sometimes just by calling a function
  * Usually don't worry, unless you're writing a tight loop:
  ```
    var updated bool
    go someFunction(&updated)
    for !updated { /* do nothing until updated is true */ }
  ```
    * This never yields. It's also very bad Go.
---
template:blue
# Off to the Races

* Consider a program with two goroutines:

.column3[
```go
func One() {
  fmt.Println("A")
  fmt.Println("B")
  fmt.Println("C")
}
```
]

.column3[
```go
func Two() {
  fmt.Println("X")
  fmt.Println("Y")
  fmt.Println("Z")
}
```
]

.column3[
```go
  func main() {
    go One(); 
    go Two(); 
  }
```
]

* Go **only** guarantees that:
  * `A` is printed before `B` and `B` is printed before `C`
  * `X` is printed before `Y` and `Y` is printed before `Z`
* Is `A` printed before `X`?
  * I don't know.
  * Nobody knows.
---
template:blue
# Race Conditions, Formally

* The order of `A` and `X` is “unspecified”
  * `A` and `X` can happen in either order, or even simultaneously.
  * Go says that these two operations are “concurrent.”
* If `A` and `X` are computations, and `A` uses the output of `X` (perhaps saved to shared memory):
  * We have a race condition
  * If `A` happens before `X`, our program will misbehave.
---
template:blue
# Race Detection

.column2[
`test.go`:
```go
package main
import ( "fmt"
         "time" )
func manipulate(shared *string) {
  *shared = "changed!"
}
func main() {
  var shared string = "pristine"
  go manipulate(&shared)
  for shared == "pristine" {
    fmt.Println(shared)
    time.Sleep(time.Millisecond)
  }
  fmt.Println(shared)
}
```
]

.column2[
Run with race detection:
```
$ go run -race test.go
Pristine
==================
WARNING: DATA RACE
Read at 0x00c420076190 by main goroutine:
  main.main()
        /home/pdbogen/test.go:10 +0xb2
        
        Previous write at 0x00c420076190 by
goroutine 6:
          main.manipulate()
                /home/pdbogen/test.go:5 +0x3b
```
]
---
template:blue
# Three tools to help

* `sync` package, likes we saw with `sync.WaitGroup`.
  * `sync.Cond` lets multiple goroutines wait on a shared signal
  * `sync.Mutex` implements a simple shared mutual-exclusion lock
  * `sync.RWMutex` implements a mutex lock that can have multiple readers
  * `sync.Once` implements a wrapper around a function ensuring the function only executes once
  * `sync.Pool` is designed to manage a group of temporary items silently shared among and potentially reused by concurrent independent clients of a package, i.e., a connection pool.
--

* `sync/atomic` contains lower-level primitives that are fast, but easy to use incorrectly
--

* `channels` are safe, powerful, and atomic.

---
template:title

## Part 2
# Type System, Continued

---
template:green
# Channels
## Communication between Goroutines

* Channels are safe
  * Sends & receives from channels have a strict ordering
  * Channel operations apparently block
* Channels are like pipes
  * You put objects into a channel and the same object comes out the other end
  * Even with multiple receivers, a sent object is received exactly once.
* Channels are typed
  * Like a map, array, or slice, channels send and receive objects of a specific type.

---
template:green
# Channel Creation

```go
  var biDir        chan int   = make(chan int) // can send to or receive from

  var sendOnly     chan<- int = make(chan int) // Sends here block forever
  var sendSide     chan<- int = biDir          // Sends here can be received on `biDir`

  var receiveOnly  <-chan int = make(chan int) // Receives here block forever
  var receiveSide  <-chan int = biDir          // Receives from `sendSide` or `biDir`
```

* _Mnemonic: `<-` always points left, and is always `chan`-adjacent._
  * Send-only points into `chan`
  * Receive-only points out of `chan`.
* **Cannot** get `send` version from `receive` version or vice-versa
* `make` can also create a buffer: `make(chan int, 100)`
  * Without a buffer, all operations are synchronous / blocking

---
template:green
# Channel Usage: Send _Statement_

```go
  var someChan chan int = make(chan int)
  someChan <- 1
```

* Arrow still points left.
  * We're putting a value into the channel, so the arrow points into the channel.
* This is a _statement_.
  * Not an operator or expression. It has no return value.
* Sends block until there is room in the buffer or a receiver waiting to receive.
* Sending on a `nil` channel blocks forever.
* Sending on a `close`d channel panics.


* _Idiom: Only the goroutines that send on a channel should close that channel._

---
template:green
# Channel Usage: Receive _Operator_

```go
  var someChan chan int = make(chan int)
  i := <-someChan
  i, ok := <-someChan
  if <-someChan > 0 { … }
```

* Arrow still points left.
  * We're taking a value out of the channel, so the arrow points out of the channel.
* This is an _operator_.
  * Returns one or two results
  * First is always an object of the channel's type
  * Second (optional) is a `bool` indicating whether the receive was real or a `close`d receive
* Receive blocks until there is something to receive (or the channel is `close`d).
* Receiving on a `nil` channel blocks forever.
* Receiving on a `close`d channel happens immediately, and returns the zero-value typed and `false` (for the two-variable return).

---
template:green
# Try it: Silly Queue

* We can mis-use channels as a queue. Never do this! But we will.
* In `main`, create a *buffered* channel of ints.
* Send two numbers to it.
* Read two numbers from it and print their sum.

```go
ch := make(chan string, 1) // a bi-directional channel of strings
ch <- "hi"
fmt.Println(<-ch)          // should print "hi"
```

* Note, this is unsafe, because of the buffer fills up on accident, we deadlock, because the 'send' will block.

---
template:green
# Channel Recipes

### Blocking to Non-blocking with `select`

```go
select {
  case input := <- inChan: fmt.Println("received input!")
  case outChan <- output:  fmt.Println("was able to send output!")
  default: fmt.Println("could not send OR receive")
}
```

* Leave off `default`, and the `select` will block until we can do anything.

### Timeouts

```go
select {
  …
  case <- time.After(time.Second): fmt.Println("waited 1s, got tired")
}
```

---
template:green
# Channel Recipes: Workers

```go
type Request struct {…}
type Response struct {…}

func worker(reqIn <-chan Request, resOut chan<- Response) {
  for req := range reqIn { // loop as long as reqIn is not closed
    /* do some kind of work to turn `req Request` into `res Response` */
    resOut <- res
  }
}

func main() {
  var requests = make(chan Request)
  var responses = make(chan Response)
  go worker(requests, responses)
  requests <- Request{…} // create and send some request
  response, ok := <-Responses // Wait for a response
  if !ok { panic("response channel closed!") }
  fmt.Println(response)
}
```
---
template:title

# Whew.
### Let's take a short break.
---
template:blue

# Warm-up Exercise

* Write a program that communicates with a single goroutine via a single channel.
  * Goroutine should receive a `string` from a channel and print it out via `fmt.Println`
  * Main should:
      * Create the channel (hint: `make`)
      * Start the goroutine, with the channel as an argument (hint: `go`)
      * Send some strings via the channel (hint: "send statement")
      * Sleep a little bit so the goroutine can finish (hint: `time.Sleep`)
  * Goroutine should:
      * Receive from a channel (hint: "receive operator")
      * Return if the channel was closed (hint: "not ok")
      * Print the string otherwise
  * Bonus points:
      * Use a `sync.WaitGroup` to make sure the goroutine finishes before `main` exits.
* For hints:
  * Use the ref-spec (`https://golang.org/ref/spec`) and 
  * Package Docs (`https://golang.org/pkg`)

---
template:blue

# Solution

```go
  import "fmt"
  import "time"

  func Printer(stringIn <-chan string) {
    for {
      str, ok := <-stringIn
      if !ok {
        return
      }
      fmt.Println(str)
    }
  }
  
  func main() {
    stringCh := make(chan string)
    go Printer(stringCh)
    stringCh <- "hello"
    stringCh <- "world"
    close(stringCh)
    // Sleep to get Printer time to finish its work
    time.Sleep(time.Second)
  }
```

---
template:blue

# Bonus Solution

```go
  import "fmt"
  import "sync"

  func Printer(stringIn <-chan string, wg *sync.WaitGroup) {
    defer wg.Done()               // On return, decrement the WaitGroup counter
    for str := range stringIn {   // Get strings from the chan until it's closed
      fmt.Println(str)            // Print the strings
    }
  }
  
  func main() {
    stringCh := make(chan string) // Create a bi-directional channel
    wg := &sync.WaitGroup{}       // Create a (pointer to) a WaitGroup
    wg.Add(1)                     // Increment the counter by one for our one goroutine
    go Printer(stringCh, wg)      // Start the goroutine
    stringCh <- "hello"           // Send "hello" to the printer
    stringCh <- "world"           // Send "world" to the printer
    close(stringCh)               // Signal to the printer that it can terminate
    wg.Wait()                     // Wait for the printer to finish
  }
```

---
template:blue

# MEGACALC

--

* Our calculator from last time was fine, but you know what would be better?

--

# MORE GOROUTINES

---
template:blue

# MEGACALC

* Write a goroutine to read tokens from stdin (hint: `fmt.Scan`) and emit the tokens, one at a time, on a channel.
* Write a core goroutine that receives tokens from a channel and decides what to do with them (i.e., operands go onto the stack, operators get executed)
    * Implement a special operator `end` that will print the top of the stack and then exit.
* Write one goroutine per operation that receives two numbers (operands) from a channel and emits the result on another channel.
    * Each of these goroutines should loop forever waiting for work
    * Bonus points: Instead of using a big `if`-`else if` chain, use a `map[string]chan int` to make it easy to add new operations.
* Bonus points: Use the `testing` package to write unit tests for each of your operator goroutines
* _Caveat: Yes, this is a bit silly/dumb. It's 1000% about learning to use channels & goroutines and not about learning to write an efficient calculator._
