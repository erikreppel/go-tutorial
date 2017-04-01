<!-- $theme: default -->

# Golang, a violent introduction

---

# Intro: Project structure, gopath, package management, tooling

- Imperative, static, concurrent language
- Written by Robert Griesemer, Rob Pike, and Ken Thompson @ Google
- Primary goals
	- Pragmatism
	- Fast compiler
	- Easy & safe concurrency
	- Low garbage collection jitter
	- Uniformity
- Primary flaws
	- Package management
	- Gopath
- Pro/Con: completely and utterly unmagical


---

# 0: Project structure

Read [Five suggestions for setting up a Go project](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project), its way better than anything I’m going to tell you.

Tl;dr: 
- Binaries have `package main` in each file and have a main function
- Packages have no main function and have `package main` in each file

---

# 0: GOPATH

- $GOPATH is where go stuff lives
- Kind of a pain to set up (until Go1.8, intros default path)
- Python has a PYTHONPATH, its just better abstracted
- Typically gopath looks like this ->

---

```
erik@Zeus-ubuntu: ~/dev/go
 $ tree -L 2                                                     
.
├── bin
│   ├── dlv
│   ├── glide
│   ├── gocode
│   ├── golint
│   ├── go-outline
│   ├── gopkgs
│   ├── gorename
│   ├── goreturns
│   ├── go-symbols
│   ├── gotests
│   ├── kube-lego
│   └── launchaco_prod
├── pkg
│   └── linux_amd64
└── src
    ├── cloud.google.com
    ├── github.com
    ├── golang.org
    ├── google.golang.org
    ├── gopkg.in
    └── sourcegraph.com

```

---

# 0: GOPATH
See `0_intro/gopath.sh` for what my go config looks like in my `.profile` file,

---

# 0: Package management

- Currently an official tool in development: [dep](https://github.com/golang/dep)
- For now, [glide](https://glide.sh/) is best imo

---

# 0: Package management
Glide usage:
```
# start glide in repo, read the options it gives you
$ glide init
...
# add a dependancy
$ glide get github.com/pkg/errors
# resolve existing dependancies
$ glide update
# install from existing glide lock file
$ glide install
```

---

# 0: Package management

Best solution I've found: 
- Binaries: `glide` + check in `vendor` and squash commits where you update vendor or add dependancies.
- Never checkin `vendor` for packages, it will innevitably break things

---

# 0: Tooling

I use VSCode with the go plugin. It installs all the tools you need for you and the debugger works really well for go. I've heard Jetbrains Gogland is good but its also still in beta.

---

# 0: Tooling

## Tips
- Have your editor auto format your code on save
- Use godoc
- Take advantage of golang standard tools, they are great
- VSCode's debugger is great and works with web servers

---

# 1: First program: Hello, world

## Three ways to run a program

---

# 1: First program: Hello, world

1. Build and run binary

```
../1_first_program/hello_world master!
$ go build
 
../1_first_program/hello_world master!
$ ./hello_world
Hello, world
```

---

# 1: First program: Hello, world

2. Run the program (compiles and executes binary in one step)
```
../1_first_program/hello_world master!
$ go run main.go
Hello, world!
```

---

# 1: First program: Hello, world

3. Install the program (adds to `$GOPATH/bin`)
```
../1_first_program/hello_world master!
$ go install .
 
../1_first_program/hello_world master!
 $ hello_world
Hello, world!

 $ tree $GOPATH/bin
/home/erik/dev/go/bin
├── dlv
├── glide
├── gocode
...
├── gotests
├── guru
├── hello_world
├── kube-cert-manager
├── kube-lego
└── launchaco_prod
```

---

# 1: First program: Hello, world

## Packages
- Public/exported functions/structs/variables start with a capital letter `Greeter`
- Private/local functions/structs start with lowercase letter `greeter`
- Public/exported things should have a comment above them describing what they do, this makes it really easy to autogen docs
- 

---

# 1: First program: Hello, world

## Imports

Are relative to `$GOPATH/src`
Ex:
```
import (
	"github.com/erikreppel/go-tutorial/1_first_program/hellopackage"
)


```

---

# 2: Testing and Benchmarking

## Its super easy and has great tooling built into stdlib, see example.

---

# 3: Error handling

---

# 3: Error handling

- Go forces you to think of all possible things that could go wrong and gracefully handle situations of failure
- This makes run time panics (what go calls exceptions) very rare

---

# 4: Interfaces and Structs

---

# 4: Interfaces and Structs

## Interfaces
- a criteria that something must fufil
- In go `error` is an interface

---

This is the entire definition of `error` in go:

```
type error interface {
    Error() string
}
```

[https://godoc.org/builtin#error](https://godoc.org/builtin#error)

---

# 4: Interfaces and Structs

This means anything can be used as an `error` as long as it has a function of `Error` that takes to args and returns a string.

You can define functions on any type, and you can define a wrapper type on almost anything, so you and shape almost any type of data into a interface if needed (really, powerful, but can add unnecasarry complexity to code very quickly)

See `4_interfaces_and_structs/error_interface` for examples.

See `4_interfaces_and_structs/storage_interface` for an example of an interface that allows to easilly swap storage backend for an application.


---

Interfaces are all nice because they allow you to test logic seperately from I/O. In the storage example we can implement the `Storage` interface in memory to make it easy to test our logic without having to deal with an FS or setting up an S3 bucket, yet still be able to test our implentation of an interface seperately.

--- 

## A fun example

---

golang's built in `http` library has a function `ListenAndServe`

![Imgur](http://i.imgur.com/deoTABe.png)

---

![Imgur](http://i.imgur.com/WSfpTNM.png)

---

Since its an interface we can make any interface fufil it, see `server_interface/`

---

