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

package glosure

import (
  "flag"
  "fmt"
  "io/ioutil"
  "net/http"
  "testing"
)

func TestDialClosureApi(t *testing.T) {
  cc := NewCompiler(".")
  res, err := cc.dialClosureApi("var i = 0; i += 1; window.alert(1)", "")
  if err != nil {
    t.Error("Error(s) in closure api call: ", res.Errors)
  }

  if len(res.CompiledCode) == 0 {
    t.Error("Got empty results from the closure compiler.")
  }
}

func TestGetClosureDependecies(t *testing.T) {
  deps, err := getClosureDependecies("./test_resources/pkg1.js")
  if err != nil {
    t.Error(err)
    return
  }
  if deps == nil || deps[0] != "pkg2" || deps[1] != "pkg3" {
    t.Error("Invalid dependecies loaded from javascript file: ", deps)
  }
}

func TestGetClosurePackage(t *testing.T) {
  pkgs, err := getClosurePackage("./test_resources/pkg1.js")
  if err != nil {
    t.Error(err)
    return
  }

  if len(pkgs) != 1 || pkgs[0] != "pkg1" {
    t.Error("Invalid package loaded from javascript file: ", pkgs)
  }
}

func TestCompilerJar(t *testing.T) {
  cc := NewCompiler("./test_resources")
  err := cc.Compile("pkg1.min.js")
  if err != nil {
    t.Error(err)
    return
  }

  _, err = ioutil.ReadFile("./test_resources/pkg1.min.js")
  if err != nil {
    t.Error(err)
  }
}

func TestCompilerApi(t *testing.T) {
  cc := NewCompiler("./test_resources")
  cc.UseClosureApi = true
  err := cc.Compile("pkg1.min.js")
  if err != nil {
    t.Error(err)
    return
  }

  _, err = ioutil.ReadFile("./test_resources/pkg1.min.js")
  if err != nil {
    t.Error(err)
  }
}

func Example() {
  // Parse the flags if you want to use glog.
  flag.Parse()

  // Creat a new compiler assuming javascript files are in "example/js".
  cc := NewCompiler("./example/js/")
  // Use strict mode for the closure compiler. All warnings are treated as
  // error.
  cc.Strict()
  // Or use debug mode.
  // cc.Debug()

  // Use advanced optimizations.
  cc.CompilationLevel = AdvancedOptimizations

  http.Handle("/", GlosureServer(cc))
  fmt.Println("Checkout http://localhost:8080/sample.min.js")
  http.ListenAndServe(":8080", nil);
}

