[Advent of code 2019](adventofcode.com/2019) written in Go. 

My first attempt at writing Go. Here are the resources used for each challenge.

### Day 1
1. [How to write go code](https://golang.org/doc/code.html)
2. [Effective go](https://golang.org/doc/effective_go.html)

### Day 2
1. [Function definitions in Go](https://go101.org/article/function-declarations-and-calls.html)
2. [Functions in Go](https://go101.org/article/function.html)
3. [Pointers](https://www.golang-book.com/books/intro/8)

## Notes

### Things I like
1. Static type checking. Code runs correctly the first time more often than a dyanmic language. I am used to writing dynamic code (python).
2. Easy to test. The testing framework is very lightweight. It's easy to add testfiles and run them. 
3. Less boilerplate. 

### Things I don't like
1. No obvious way to do dependency management for separate projects. Everything is in one workspace. Packages are saved in the workspace. A single go path. 
2. Some types are automatically dereferenced, some are not. See [selectors](https://golang.org/ref/spec#Selectors). I was expecting to use pointers/references to pass arrays by value, but this is not needed. Because this is a C-like language and pointers are emphasized, this was surprising to me. Ultimately, it makes me use case easier but it seems inconsistent.
