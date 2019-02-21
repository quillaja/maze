package maze

import (
	"sort"
)

// Node allows types to become "nodes" or "cells" in a maze, which is
// essentially just an undirected graph.
type Node interface {
	Neighbors() NodeSlice
}

// NodeSlice is a slice of Nodes.
type NodeSlice []Node

// ToNodeSlice is a convenience method to convert a slice of any type that
// implements Node to a NodeSlice.
func ToNodeSlice(list []interface{}) NodeSlice {
	s := make(NodeSlice, len(list))
	for i := range list {
		s[i] = list[i].(Node)
	}
	return s
}

// Append adds the nodes to the slice.
func (slice NodeSlice) append(n ...Node) NodeSlice {
	return append(slice, n...)
}

// AppendUnique adds the node to the slice only if it is not already
// in the slice.
func (slice NodeSlice) appendUnique(n Node) NodeSlice {
	if !slice.has(n) {
		slice = slice.append(n)
	}
	return slice
}

// Index returns the first index where 'n' is found, or 'len(slice)'
// if the node is not found.
func (slice NodeSlice) index(n Node) int {
	return sort.Search(len(slice), func(j int) bool {
		return slice[j] == n
	})
}

// Has returns true if the slice contains 'n'.
func (slice NodeSlice) has(n Node) bool {
	return slice.index(n) != len(slice)
}

// RemoveAt removes the node at i. The original order of the slice
// is not preserved.
func (slice NodeSlice) removeAt(i int) NodeSlice {
	l := len(slice)
	if l == 0 {
		return slice
	}

	slice[i] = slice[l-1] // overwrite i with end
	slice[l-1] = nil      // prevent memory leak
	slice = slice[:l-1]   // slice off end

	return slice
}

// Remove removes the first occurence of 'n' from the slice. The
// original order is not preserved.
func (slice NodeSlice) remove(n Node) NodeSlice {
	i := slice.index(n)
	if i < len(slice) {
		return slice.removeAt(i)
	}

	return slice
}

// Pop removes and returns the last node on the slice.
func (slice NodeSlice) pop() (NodeSlice, Node) {
	n := slice[len(slice)-1]
	return slice.removeAt(len(slice) - 1), n
}

type nodeSet map[Node]struct{}

// Graph maintains an undirected graph of Nodes and edges between Nodes.
type graph map[Node]NodeSlice

// Has returns true if the node 'n' is in the graph.
func (g graph) has(n Node) bool {
	_, in := g[n]
	return in
}

// AddNode adds the node to the graph.
func (g graph) addNode(nodes ...Node) {
	for _, n := range nodes {
		if !g.has(n) {
			g[n] = n.Neighbors() // NOTE: do addEdge() instead of this?
		}
	}
}

// AddEdge adds an undirected edge connecting a and b.
func (g graph) addEdge(a, b Node) {
	g.addNode(a, b)
	g[a] = g[a].appendUnique(b)
	g[b] = g[b].appendUnique(a)
}

// RemoveEdge removes the edge connecting a and b.
func (g graph) removeEdge(a, b Node) {
	if !g.has(a) || !g.has(b) {
		return
	}

	g[a] = g[a].remove(b)
	g[b] = g[b].remove(a)
}

// ThreeToOne calculates the 1D index from a 3D index given the indices
// 'x, y, z' and the x-row width 'dx' and the y-column height 'dy'. Both
// dx and dy should be in the range [1, N].
func ThreeToOne(x, y, z, dx, dy int) int {
	return x + y*dx + z*dx*dy
}

// OneToThree calculates the 3D indices from a 1D index given
// the x-row width 'dx' and the y-column height 'dy'. Both
// dx and dy should be in the range [1, N].
func OneToThree(i, dx, dy int) (x, y, z int) {
	z = i / (dx * dy)
	i = i % (dx * dy)
	y = i / dx
	x = i % dx
	return
}
