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
cc := glosure.NewCompiler("./example/js/")
// Set compiler options.
http.Handle("/", glosure.GlosureServer(cc))
http.ListenAndServe(":8080", nil);
```

Or even simpler if you do not need to customize the compiler:
```go
http.Handle("/", glosure.GlosureServerWithRoot("./example/js/"))
http.ListenAndServe(":8080", nil);
```

```GlosureServer``` serves only the requests for compiled
JavaScript (by default ```*.min.js```) and returns error otherwise.
For example, ```http://localhost:8080/sample.min.js``` returns the 
compiled version of ```./example/js/sample.js```, but
dialing ```http://localhost:8080/sample.js``` results in a 404.
You can change this behavior by setting a customized
handler in ```cc.ErrorHandler```.

For a more comprehensive example, take a look at
```example/server.go```. You can run the example by:

    # go run server.go --logtostderr -v=1

