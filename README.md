# Go kata code

## A brief history of GO
Go was designed at Google in 2007 to improve programming productivity in an era of multicore, networked machines and large codebases.
The designers wanted to address criticism of other languages in use at Google (shared dislike of C++), but keep their useful characteristics.
* 2007 - Started and built by Robert Griesemer, Rob Pike and Ken Thompson as a part-time project.
* 2008 - A lot of other people help to bring go from prototype to reality.
* 2009 - Go was publicly announced in November and became an Open Source project.
* 2010 - Starts to be adopted by other programmers.

## What is GO
Go is a compiled, statically typed, imperative, concurrent and structured language.
* compiled - gc compiler, memory managed by garbage collector.
* statically typed with some dinamically typed capabilities
* imperative - tells computer how to do e.g ruby
* declarative - tells computer what to do e.g scala
* concurrent - multithreading as well as CPU parallelism 

## Why GO
No major programming language has been developed in over a decade while the computing landscape has changed a lot.
Those trends are:
* Computers are enormously quicker but software development is not faster.
* Dependency management is a big part of software development today but the “header ﬁles” of languages in the C tradition are antithetical to clean dependency analysis—and fast compilation.
* There is a growing rebellion against cumbersome type systems like those of Java and C++, pushing people towards dynamically typed languages such as Python and JavaScript.
* Some fundamental concepts such as garbage collection and parallel computation are not well supported by popular systems languages.
* The emergence of multicore computers has generated worry and confusion.
 We believe it's worth trying again with a new language, a concurrent, garbage-collected language with fast compilation. Regarding the points above:
 * It is possible to compile a large Go program in a few seconds on a single computer.
 * Go provides a model for software construction that makes dependency analysis easy and avoids much of the overhead of C-style include ﬁles and libraries.
 * Go's type system has no hierarchy, so no time is spent deﬁning the relationships between types. Also, although Go has static types the language attempts to make types feel lighter weight than in typical OO languages.
 * Go is fully garbage-collected and provides fundamental support for concurrent execution and communication.
 * By its design, Go proposes an approach for the construction of system software on multicore machines.

## Guiding principles in Design
* Reduce the amount of typing in both sense of the word 
* Reduce clutter and complexity 
* No forward declarations and no header files; everything is declared exactly once 
* Most radically, there is no type hierarchy: types just are, they don't have to announce their relationships 
For more info GO read [FAQs](https://golang.org/doc/faq)

## Install Go
Follow the [installation](https://golang.org/doc/install) procedure specific to each OS.
### Test the installation by running:
```go run main.go```
### Install [ginkgo](https://onsi.github.io/ginkgo/) testing framework:
* `$ go get github.com/onsi/ginkgo/ginkgo`
* `$ go get github.com/onsi/gomega/...`

