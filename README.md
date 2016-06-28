# Learning Golang

My small repository to track my golang learning process

Understanding the go language to apply it to my research of blockchains

## Notes

### Variables

* Can be defined multiple ways

```go

var name type

name := ____

var a, b type = tuple1, tuple2

```

###Arrays and Slices

* Setting array ``var name [size]type``
* Can also do ``name := [size]type{element, element,...}``

Slices are different, but very similar to the way python does it

* Simple range: ``array[1:5]``
* Or ``array[:2] , array[2:]``
* They are ***immutable***

### Variadic Functions

Use the `...` after the argument name to have multiple arguments into one

```go
func sum(nums ...int) {}

foo := []int{1,2,3}
sum(foo)
```

### Interfaces
* Very similar to every programming language
* Define an interface, make sure that the 'object' you pass it has the function 

### Implementing Errors
* Import the error class ``import "errors"``
* then to throw a new error ``errors.New("message here")``

### Goroutines 
* Similar to the threads 
* All you have to do is call ``go`` with the function 
* Runs *async*

### Channels
* Pipes that connect concurrent goroutines
* Send values from one, receive from the other
* create a channel: ``make(chan val-type)``
* Send: ``<-``
* Receive: ``<-channel``
* By default, sends and receives block until sender AND receiver are ready
* Channels by default unbuffered (accept sends if there is a receive ready to receive
* Buffer allows for the receiving of many
* How: ``make(chan type, limit)``


