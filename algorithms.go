package maze

import (
	"math/rand"
)

func getUnvisited(g, maze Graph) Node {
	if g.NodeCount() == maze.NodeCount() {
		return nil
	}

	for {
		n := g.RandomNode()
		if !maze.Has(n) {
			return n
		}
	}
}

// Wilson implements Wilson's algorithm.
// see: http://weblog.jamisbuck.org/2011/1/20/maze-generation-wilson-s-algorithm.html
func Wilson(g Graph) (maze Graph) {
	maze = NewMapGraph()
	if n := getUnvisited(g, maze); n != nil {
		maze.Add(n) // add random initial node to maze
	}

	// while there are unvisited nodes, create random acyclic walks
	// through g and add those paths to maze.
	for n := getUnvisited(g, maze); n != nil; n = getUnvisited(g, maze) {

		// create a random walk through unvisited graph
		path := NodeSlice{n}
		for pathCreated := false; !pathCreated; {
			neighbors := g.Neighbors(n)
			n = neighbors[rand.Intn(len(neighbors))] // get random neighbor
			path = path.Append(n)

			// check if next is already in the path
			if prev := path.index(n); prev != notFound && prev != len(path)-1 {
				// it's already in the path, we have a loop.
				// must cut loop out
				path = path[:prev+1]
			}

			pathCreated = maze.Has(n) // check if next is currently in the 'maze' (visted nodes)
		}

		// add the path to the maze
		for i := 0; i < len(path)-1; i++ {
			maze.AddEdge(path[i], path[i+1])
		}
	}

	return
}
