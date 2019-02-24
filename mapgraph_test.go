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
	for i := 0; i < 5; i++ {
		g.Add(Node(i))
	}
	for i := 0; i < 2; i++ {
		for j := 2; j < 10; j++ {
			g.AddEdge(Node(i), Node(j))
		}
	}
	g.AddEdge(Node(-1), Node(-2))
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
		a Node
		b Node
	}
	tests := []struct {
		name string
		g    Graph
		args args
		want bool
	}{
		{name: "simple: remove 0 from 0-2", g: constructGraph1(),
			args: args{a: Node(0), b: Node(2)}, want: false},
		{name: "simple: remove 1 from 1-9", g: constructGraph1(),
			args: args{a: Node(1), b: Node(9)}, want: false},
		{name: "simple: remove -2 from (-1)-(-2)", g: constructGraph1(),
			args: args{a: Node(-2), b: Node(-1)}, want: false},

		{name: "edge: both added, not connected 9-10", g: constructGraph1(),
			args: args{a: Node(9), b: Node(10)}, want: false},
		{name: "edge: remove unadded 20", g: constructGraph1(),
			args: args{a: Node(0), b: Node(20)}, want: false},
		{name: "edge: one not added, one added 20-0", g: constructGraph1(),
			args: args{a: Node(20), b: Node(0)}, want: false},
		{name: "edge: no self-reference 0-0", g: constructGraph1(),
			args: args{a: Node(0), b: Node(0)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.Remove(tt.args.a)
			if tt.g.Has(tt.args.a) {
				t.Log(tt.g)
				t.Errorf("mapgraph.Remove: did not remove (%v)", tt.args.a)
			}
			if got := tt.g.HasEdge(tt.args.b, tt.args.a); got != tt.want {
				t.Log(tt.g)
				t.Errorf("mapgraph.RemoveEdge: g.RemoveEdge((%v)-(%v))==%v. want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}

}
