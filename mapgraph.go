package maze

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
			//g.assignUnusedID(n)
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
			g[neighbor].Remove(n)
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
	neighbors, in := g[a]
	if !in {
		return false
	}
	return neighbors.Has(b)
}

// AddEdge adds an undirected edge connecting a and b. Nodes a and b are added
// to the graph if not already present.
func (g mapgraph) AddEdge(a, b Node) {
	if a == b {
		return
	}
	g.Add(a, b)
	g[a] = g[a].AppendUnique(b)
	g[b] = g[b].AppendUnique(a)
}

// RemoveEdge removes the edge connecting a and b. It does nothing if a, b, or
// the edge are not in the graph.
func (g mapgraph) RemoveEdge(a, b Node) {
	if g.Has(a) {
		g[a] = g[a].Remove(b)
	}

	if g.Has(b) {
		g[b] = g[b].Remove(a)
	}
}

// assignUnusedID assigns an id to n.
// func (g mapgraph) assignUnusedID(n Node) {
// 	var id ID
// 	used := true
// 	for used {
// 		id := randID(maxID)
// 		_, used = g[id]
// 	}
// 	n.SetID(id)
// }
