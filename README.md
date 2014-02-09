Glosure: Closure Compiler for the Go programming language.
==========================================================
Glosure is a simple http.Handler for compiling JavaScript using the
Closure Compiler.

### Prerequisites:

    * Java: The Java Runtime Environment is required to run the
      closure compiler. You need to have java in your $PATH.

### Installation:

    # go get github.com/soheilhy/glosure

### Example:
Take a look at ```example/server.go``` directory. You can run the
example by:

    # go run server.go --logtostderr -v=1

