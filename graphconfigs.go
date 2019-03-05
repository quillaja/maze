package maze

// MakeGrid generates a Graph representing a regular "square" grid where
// each node has at most 6 neighbors. dx, dy, and dz should be in [1,1024].
func MakeGrid(dx, dy, dz int) Graph {
	max := 1024
	dx = clamp(dx, 1, max)
	dy = clamp(dy, 1, max)
	dz = clamp(dz, 1, max)

	g := NewMapGraph()

	for z := 0; z < dz; z++ {
		for y := 0; y < dy; y++ {
			for x := 0; x < dx; x++ {
				n := Node(ThreeToOne(x, y, z, dx, dy))
				if x > 0 {
					g.AddEdge(n, Node(ThreeToOne(x-1, y, z, dx, dy)))
				}
				if x < dx-1 {
					g.AddEdge(n, Node(ThreeToOne(x+1, y, z, dx, dy)))
				}
				if y > 0 {
					g.AddEdge(n, Node(ThreeToOne(x, y-1, z, dx, dy)))
				}
				if y < dy-1 {
					g.AddEdge(n, Node(ThreeToOne(x, y+1, z, dx, dy)))
				}
				if z > 0 {
					g.AddEdge(n, Node(ThreeToOne(x, y, z-1, dx, dy)))
				}
				if z < dz-1 {
					g.AddEdge(n, Node(ThreeToOne(x, y, z+1, dx, dy)))
				}
			}
		}
	}

	return g
}
