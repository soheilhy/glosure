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

  fromNode.Dependencies.PushBack(toNode)
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
  return g.getDependencies(nodes, make([]*Node, 0))
}

func (g *DependencyGraph) getDependencies(nodes []*Node, deps []*Node) []*Node {
  for _, node := range nodes {
    if contains(deps, node) {
      continue
    }

    for e := node.Dependencies.Front(); e != nil; e = e.Next() {
      eDeps := g.getDependencies([]*Node{e.Value.(*Node)}, deps)
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

func containsPkg(nodes []*Node, pkg string) bool {
  for _, node := range nodes {
    if node.Pkg == pkg {
      return true
    }
  }

  return false
}


func (n *Node) isRecursivelyDependentOn(pkg string) bool {
  return n.isRecursivelyDependentOnWithCache(pkg, make([]*Node, 0))
}

func (n *Node) isRecursivelyDependentOnWithCache(pkg string,
                                                 checked []*Node) bool {
  if pkg == n.Pkg {
    return true;
  }

  if containsPkg(checked, pkg) {
    return false;
  }

  for e := n.Dependencies.Front(); e != nil; e = e.Next() {
    if e.Value.(*Node).isRecursivelyDependentOnWithCache(pkg, checked) {
      return true
    }
  }

  checked = append(checked, n)
  return false
}

