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

package depgraph

import (
  "container/list"
  "errors"
)

type DependencyGraph struct {
  Nodes map[string]*Node
}

func New() DependencyGraph {
  return DependencyGraph{make(map[string]*Node)}
}

func (g *DependencyGraph) AddFile(pkg string, path string) {
  g.Nodes[pkg] = &Node{pkg, path, list.New()}
}

func (g *DependencyGraph) AddDependency(from string, to string) error {
  fromNode, ok := g.Nodes[from]
  if !ok {
    return errors.New("Package not found: " + from)
  }

  toNode, ok := g.Nodes[to]
  if !ok {
    return errors.New("Package not found: " + to)
  }

  if toNode.isRecursivelyDependentOn(from) {
    return errors.New("Circular dependency between " + from + " and " + to)
  }

  fromNode.Dependencies.PushFront(toNode)
  return nil
}

func (g *DependencyGraph) GetDependenciesOfPackage(pkg string) []*Node {
  node, ok := g.Nodes[pkg]
  if !ok {
    return nil
  }

  return g.GetDependencies([]*Node{node})
}

func (g *DependencyGraph) GetDependencies(nodes []*Node) []*Node {
  deps := []*Node{}
  for _, node := range nodes {
    for e := node.Dependencies.Front(); e != nil; e = e.Next() {
      eDeps := g.GetDependencies([]*Node{e.Value.(*Node)})
      for _, dependencyNode := range eDeps {
        if contains(deps, dependencyNode) {
          continue
        }
        deps = append(deps, dependencyNode)
      }
    }
    deps = append(deps, node)
  }
  return deps
}

type Node struct {
  Pkg string
  Path string
  Dependencies *list.List
}

func contains(nodes []*Node, newNode *Node) bool {
  for _, node := range nodes {
    if node.Path == newNode.Path {
      return true
    }
  }

  return false
}

func (n *Node) isRecursivelyDependentOn(pkg string) bool {
  if pkg == n.Pkg {
    return true
  }

  for e := n.Dependencies.Front(); e != nil; e = e.Next() {
    if e.Value.(*Node).isRecursivelyDependentOn(pkg) {
      return true
    }
  }

  return false
}

