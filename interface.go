package maze

// TODO: for future consideration.
// type void struct{}
// type nodeSet map[Node]void

// Node allows types to become "nodes" or "cells" in a maze, which is
// essentially just an undirected graph.
type Node interface {
	IsNeighbor(other Node) bool
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
