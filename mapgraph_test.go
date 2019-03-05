package maze

import (
	"testing"
)

func Test_mapgraph_Has(t *testing.T) {
	nodein := Node(0)
	nodeout := Node(1)
	g := NewMapGraph()
	g.Add(nodein)

	type args struct {
		n Node
	}
	tests := []struct {
		name string
		g    Graph
		args args
		want bool
	}{
		{name: "positive", g: g, args: args{nodein}, want: true},
		{name: "negative", g: g, args: args{nodeout}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Has(tt.args.n); got != tt.want {
				t.Errorf("mapgraph.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapgraph_Add(t *testing.T) {
	type args struct {
		nodes []Node
	}
	tests := []struct {
		name string
		g    Graph
		args args
		len  int
	}{
		{name: "add one", g: NewMapGraph(), args: args{nodes: []Node{Node(0)}}, len: 1},
		{name: "add two", g: NewMapGraph(), args: args{nodes: []Node{Node(0), Node(10)}}, len: 2},
		{name: "add dups", g: NewMapGraph(), args: args{nodes: []Node{Node(0), Node(0)}}, len: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.Add(tt.args.nodes...)
			if l := len(tt.g.(mapgraph)); tt.len != l {
				t.Errorf("mapgraph.Add(): len(g)=%v was not the expected %v", l, tt.len)
			}
			for _, n := range tt.args.nodes {
				if !tt.g.Has(n) {
					t.Errorf("mapgraph.Add(): g does not include (%v)", n)
				}
			}
		})
	}
}

func constructGraph1() Graph {
	g := NewMapGraph()
	// add nodes 0-4
	for i := 0; i < 5; i++ {
		g.Add(Node(i))
	}
	// make edges between 0 & 1 and 2-9, adding nodes 5-9
	for i := 0; i < 2; i++ {
		for j := 2; j < 10; j++ {
			g.AddEdge(Node(i), Node(j))
		}
	}
	// creates nodes -1 and -2 and connects them
	g.AddEdge(Node(-1), Node(-2))
	// attempts to add a "loop" to node 0, which is expected silently fail
	g.AddEdge(Node(0), Node(0))
	return g
}

func Test_mapgraph_AddEdge(t *testing.T) {
	g := constructGraph1()
	t.Log("finished constructing g")
	// t.Log(g)
	type args struct {
		a Node
		b Node
	}
	tests := []struct {
		name string
		g    Graph
		args args
		want bool
	}{
		{name: "check nil-0", g: g, args: args{a: nil, b: Node(2)}, want: false},
		{name: "both pre-added 0-2", g: g, args: args{a: Node(0), b: Node(2)}, want: true},
		{name: "one added in AddEdge 1-9", g: g, args: args{a: Node(1), b: Node(9)}, want: true},
		{name: "both added in AddEdge (-1)-(-2)", g: g, args: args{a: Node(-1), b: Node(-2)}, want: true},
		{name: "both added, not connected 9-10", g: g, args: args{a: Node(9), b: Node(10)}, want: false},
		{name: "one added, one not 0-20", g: g, args: args{a: Node(0), b: Node(20)}, want: false},
		{name: "one not added, one added 20-0", g: g, args: args{a: Node(20), b: Node(0)}, want: false},
		{name: "no self-reference 0-0", g: g, args: args{a: Node(0), b: Node(0)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.HasEdge(tt.args.a, tt.args.b); got != tt.want {
				t.Log(tt.g)
				t.Errorf("mapgraph.AddEdge: g.AddEdge((%v)-(%v))==%v. want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}
}

func Test_mapgraph_HasEdge(t *testing.T) {
	Test_mapgraph_AddEdge(t)
}

func Test_mapgraph_RemoveEdge(t *testing.T) {
	type args struct {
		a Node
		b Node
	}
	tests := []struct {
		name string
		g    Graph
		args args
		want bool
	}{
		{name: "both pre-added 0-2", g: constructGraph1(), args: args{a: Node(0), b: Node(2)}, want: false},
		{name: "one added in AddEdge 1-9", g: constructGraph1(), args: args{a: Node(1), b: Node(9)}, want: false},
		{name: "both added in AddEdge (-1)-(-2)", g: constructGraph1(), args: args{a: Node(-1), b: Node(-2)}, want: false},
		{name: "both added, not connected 9-10", g: constructGraph1(), args: args{a: Node(9), b: Node(10)}, want: false},
		// {name: "one added, one not 0-20", g: constructGraph1(), args: args{a: Node(0), b: Node(20)}, want: false},
		// {name: "one not added, one added 20-0", g: constructGraph1(), args: args{a: Node(20), b: Node(0)}, want: false},
		// {name: "no self-reference 0-0", g: constructGraph1(), args: args{a: Node(0), b: Node(0)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.RemoveEdge(tt.args.a, tt.args.b)
			if got := tt.g.HasEdge(tt.args.a, tt.args.b); got != tt.want {
				t.Log(tt.g)
				t.Errorf("mapgraph.RemoveEdge: g.RemoveEdge((%v)-(%v))==%v. want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}
}

func Test_mapgraph_Remove(t *testing.T) {
	type args struct {
		r NodeSlice
		a Node
		b Node
	}
	tests := []struct {
		name string
		g    Graph
		args args
		want bool
	}{
		{name: "remove 9, check 0-2", g: constructGraph1(),
			args: args{r: NodeSlice{Node(9)}, a: Node(0), b: Node(2)}, want: true},
		{name: "remove 0 check 0-9", g: constructGraph1(),
			args: args{r: NodeSlice{Node(0)}, a: Node(0), b: Node(9)}, want: false},
		{name: "remove 0 check 0-2", g: constructGraph1(),
			args: args{r: NodeSlice{Node(0)}, a: Node(0), b: Node(2)}, want: false},
		{name: "remove 0 and 1 check 1-9", g: constructGraph1(),
			args: args{r: NodeSlice{Node(0), Node(1)}, a: Node(1), b: Node(9)}, want: false},
		{name: "remove -2 from (-1)-(-2)", g: constructGraph1(),
			args: args{r: NodeSlice{Node(-2)}, a: Node(-1), b: Node(-2)}, want: false},
		{name: "remove 0 and 1 check 1-9", g: constructGraph1(),
			args: args{r: NodeSlice{Node(0), Node(1)}, a: Node(1), b: Node(9)}, want: false},
		{name: "remove 3 and 6; check affected 1-9", g: constructGraph1(),
			args: args{r: NodeSlice{Node(3), Node(6)}, a: Node(1), b: Node(9)}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.Remove(tt.args.r...)
			// t.Log(tt.g)
			for _, r := range tt.args.r {
				if tt.g.Has(r) {
					t.Log(tt.g)
					t.Errorf("mapgraph.Remove: did not remove (%v)", r)
				}
			}
			if got := tt.g.HasEdge(tt.args.a, tt.args.b) || tt.g.HasEdge(tt.args.b, tt.args.a); got != tt.want {
				t.Log(tt.g)
				t.Errorf("mapgraph.RemoveEdge: g.RemoveEdge(%v, %v)==%v. want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}

}

func Test_mapgraph_Neighbors(t *testing.T) {
	g := constructGraph1()
	// t.Log(g)

	type args struct {
		n Node
	}
	tests := []struct {
		name string
		g    Graph
		args args
		want NodeSlice
	}{
		{name: "Node(0)", g: g, args: args{Node(0)}, want: NodeSlice{Node(2), Node(3), Node(4), Node(5), Node(6), Node(7), Node(8), Node(9)}},
		{name: "Node(2)", g: g, args: args{Node(2)}, want: NodeSlice{Node(1), Node(0)}},
		{name: "Node(-1)", g: g, args: args{Node(-1)}, want: NodeSlice{Node(-2)}},
		{name: "Node(-2)", g: g, args: args{Node(-2)}, want: NodeSlice{Node(-1)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// order in slice deep equal matters, but doesnt for the graph,
			// so have to test each Node in tt.want against the Neighbors()
			got := tt.g.Neighbors(tt.args.n)
			gotAll := true
			for _, n := range tt.want {
				gotAll = gotAll && got.Has(n) // one false will make gotAll false
			}
			if !gotAll {
				t.Errorf("mapgraph.Neighbors() = %v, want %v", got, tt.want)
				t.Log(g)
			}
		})
	}
}

func Test_mapgraph_RandomNode(t *testing.T) {
	tests := []struct {
		name string
		g    Graph
	}{
		{name: "", g: constructGraph1()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.RandomNode(); !tt.g.Has(got) {
				t.Errorf("mapgraph.RandomNode() = %v,", got)
			}
		})
	}
}

func Test_mapgraph_NodeCount(t *testing.T) {
	tests := []struct {
		name string
		g    Graph
		want int
	}{
		{name: "empty", g: NewMapGraph(), want: 0},
		{name: "normal", g: constructGraph1(), want: 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.NodeCount(); got != tt.want {
				t.Errorf("mapgraph.NodeCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
