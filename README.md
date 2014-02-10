Glosure: Closure Compiler for Go
================================
Glosure is a simple http.Handler for compiling JavaScript using the
Closure Compiler.

### Prerequisites:

    * Java: The Java Runtime Environment is required to run the
      closure compiler. You need to have java in your $PATH.

### Installation:

    # go get github.com/soheilhy/glosure

### Example:
First create a closure compiler, and then serve HTTP requests: 
```go
cc := NewCompiler("./example/js/")
// Set compiler options.
http.Handle("/", GlosureServer(cc))
http.ListenAndServe(":8080", nil);
```

Or even simpler if you do not need to customize the compiler:
```go
http.Handle("/", GlosureServerWithRoot("./example/js/"))
http.ListenAndServe(":8080", nil);
```

For a more comprehensive example, take a look at
```example/server.go```. You can run the example by:

    # go run server.go --logtostderr -v=1

