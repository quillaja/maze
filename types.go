package maze

// TODO: for future consideration.
// type void struct{}
// type nodeSet map[Node]void

// Node allows types to become "nodes" or "cells" in a maze, which is
// essentially just an undirected graph.
type Node interface {
	IsNeighbor(other Node) bool
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
func (slice NodeSlice) Append(n ...Node) NodeSlice {
	return append(slice, n...)
}

// AppendUnique adds the node to the slice only if it is not already
// in the slice.
func (slice NodeSlice) AppendUnique(n Node) NodeSlice {
	if !slice.Has(n) {
		slice = slice.Append(n)
	}
	return slice
}

// Index returns the first index where 'n' is found, or 'len(slice)'
// if the node is not found.
func (slice NodeSlice) index(n Node) int {
	// sort.Search() doesn't work as a general purpose search.
	// return sort.Search(len(slice), func(j int) bool {
	// 	return slice[j] == n
	// })
	for i := 0; i < len(slice); i++ {
		if slice[i] == n {
			return i
		}
	}
	return len(slice)
}

// Has returns true if the slice contains 'n'.
func (slice NodeSlice) Has(n Node) bool {
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
func (slice NodeSlice) Remove(n Node) NodeSlice {
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

// Graph provides the basic interface needed by maze making algorithms, since
// mazes are essentially just undirected graphs.
type Graph interface {
	// Has returns true if the graph contains Node.
	Has(Node) bool
	// Adds the node(s) to the graph.
	Add(...Node)
	// Removes the node(s) from the graph and any edges associated with the node(s).
	Remove(...Node)
	// Neighbors provides a list of nodes connected to Node.
	// It is expected to return an empty NodeSlice if the Node has no associated
	// edges, and nil if the Node is not in the graph.
	Neighbors(Node) NodeSlice
	// HasEdge returns true if a and b are connected by an edge.
	HasEdge(a, b Node) bool
	// AddEdge adds an undirected edge between a and b. It may add nodes a
	// and/or b if they are not already in the graph.
	AddEdge(a, b Node)
	// RemoveEdge removes the edge between a and b if a, b, and the edge are
	// in the graph.
	RemoveEdge(a, b Node)
}

// mapgraph maintains an undirected graph of Nodes and edges using a map.
type mapgraph map[Node]NodeSlice

// NewMapGraph makes a Graph using a map. Therefore it is suitable for graphs
// with an arbitrary topology.
func NewMapGraph() Graph {
	return make(mapgraph)
}

// Has returns true if the node 'n' is in the graph.
func (g mapgraph) Has(n Node) bool {
	_, in := g[n]
	return in
}

// Adds the node(s) to the graph.
func (g mapgraph) Add(nodes ...Node) {
	for _, n := range nodes {
		if !g.Has(n) {
			g[n] = make(NodeSlice, 0)
		}
	}
}

// Removes the node(s) from the graph. It also removes edges between the
// deleted nodes and their former neighbors.
func (g mapgraph) Remove(nodes ...Node) {
	for _, n := range nodes {
		// remove edges so that current neighbors do not maintain
		// "dangling" directed edges to n
		for _, neighbor := range g.Neighbors(n) {
			g.RemoveEdge(n, neighbor)
		}
		// remove the node
		delete(g, n)
	}
}

// Neighbors returns the nodes with an edge to n. If n is isolated, an empty
// NodeSlice is returned. If n is not in the graph, nil is returned.
func (g mapgraph) Neighbors(n Node) NodeSlice {
	return g[n]
}

// HadEdge returns true if a and b and connected with an edge. False is returned
// if a and b are not connected or if a is not in the graph.
func (g mapgraph) HasEdge(a, b Node) bool {
	n, in := g[a]
	if !in {
		return false
	}
	return n.Has(b)
}

// AddEdge adds an undirected edge connecting a and b. Nodes a and b are added
// to the graph if not already present.
func (g mapgraph) AddEdge(a, b Node) {
	g.Add(a, b)
	g[a] = g[a].AppendUnique(b)
	g[b] = g[b].AppendUnique(a)
}

// RemoveEdge removes the edge connecting a and b. It does nothing if a, b, or
// the edge are not in the graph.
func (g mapgraph) RemoveEdge(a, b Node) {
	if !g.Has(a) || !g.Has(b) || !g.HasEdge(a, b) {
		return
	}

	g[a] = g[a].Remove(b)
	g[b] = g[b].Remove(a)
}
