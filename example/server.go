// Copyright (c) 2014 The Glosure Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
  "flag"
  "fmt"
  "net/http"

  "github.com/soheilhy/glosure"
)

// To see all logs from the server, run this file as:
//
//    go run server.go --logtostderr -v=1
//
func main() {
  debug := flag.Bool("debug", false, "run the compiler in debug mode.")
  advanced := flag.Bool("advanced", false, "use advanced optimizations.")
  noJava := flag.Bool("nojava", false, "use closure rest api instead of java.")

  // Parse the flags if you want to use glog.
  flag.Parse()

  // Creat a new compiler.
  cc := glosure.NewCompiler("./js/")
  if *debug {
    cc.Debug()
  } else {
    // Use strict mode for the closure compiler. All warnings are treated as
    // error.
    cc.Strict()
  }

  if *advanced {
    // Use advanced optimizations.
    cc.CompilationLevel = glosure.AdvancedOptimizations
  }

  cc.UseClosureApi = *noJava

  http.Handle("/", glosure.GlosureServer(cc))
  fmt.Println("Checkout http://localhost:8080/sample.min.js?force=1")
  http.ListenAndServe(":8080", nil);
}


