package maze

import (
	"reflect"
	"testing"
)

func TestNodeSlice_Append(t *testing.T) {
	type args struct {
		n []Node
	}
	tests := []struct {
		name  string
		slice NodeSlice
		args  args
		want  NodeSlice
	}{
		{name: "1", slice: NodeSlice{}, args: args{[]Node{Node(0)}}, want: NodeSlice{Node(0)}},
		{name: "2", slice: NodeSlice{}, args: args{[]Node{Node(0), Node(1)}}, want: NodeSlice{Node(0), Node(1)}},
		{name: "3", slice: NodeSlice{Node(0)}, args: args{[]Node{Node(1)}}, want: NodeSlice{Node(0), Node(1)}},
		{name: "4", slice: NodeSlice{Node(0)}, args: args{[]Node{Node(0)}}, want: NodeSlice{Node(0), Node(0)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.slice.Append(tt.args.n...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NodeSlice.Append() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeSlice_AppendUnique(t *testing.T) {
	type args struct {
		n Node
	}
	tests := []struct {
		name  string
		slice NodeSlice
		args  args
		want  NodeSlice
	}{
		{name: "standard", slice: NodeSlice{}, args: args{Node(0)}, want: NodeSlice{Node(0)}},
		{name: "add new", slice: NodeSlice{Node(0)}, args: args{Node(1)}, want: NodeSlice{Node(0), Node(1)}},
		{name: "add existing", slice: NodeSlice{Node(0)}, args: args{Node(0)}, want: NodeSlice{Node(0)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.slice.AppendUnique(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NodeSlice.AppendUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeSlice_index(t *testing.T) {
	type args struct {
		n Node
	}
	tests := []struct {
		name  string
		slice NodeSlice
		args  args
		want  int
	}{
		{name: "simple", slice: NodeSlice{Node(0)}, args: args{Node(0)}, want: 0},
		{name: "simple, multiple elems", slice: NodeSlice{Node(10), Node(-1), Node(0)}, args: args{Node(0)}, want: 2},
		{name: "empty list, not found", slice: NodeSlice{}, args: args{Node(0)}, want: notFound},
		{name: "mutiple elems, not found", slice: NodeSlice{Node(0), Node(-1), Node(10)}, args: args{Node(20)}, want: notFound},
		{name: "duplicate elems, finds first", slice: NodeSlice{Node(1), Node(5), Node(0), Node(4), Node(0)}, args: args{Node(0)}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.slice.index(tt.args.n); got != tt.want {
				t.Errorf("NodeSlice.index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeSlice_Has(t *testing.T) {
	type args struct {
		n Node
	}
	tests := []struct {
		name  string
		slice NodeSlice
		args  args
		want  bool
	}{
		{name: "simple", slice: NodeSlice{Node(0)}, args: args{Node(0)}, want: true},
		{name: "duplicates", slice: NodeSlice{Node(1), Node(0), Node(3), Node(0)}, args: args{Node(0)}, want: true},
		{name: "not in list", slice: NodeSlice{Node(1), Node(2)}, args: args{Node(0)}, want: false},
		{name: "empty list", slice: NodeSlice{}, args: args{Node(0)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.slice.Has(tt.args.n); got != tt.want {
				t.Errorf("NodeSlice.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeSlice_removeAt(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name  string
		slice NodeSlice
		args  args
		want  NodeSlice
	}{
		{name: "empty", slice: NodeSlice{}, args: args{0}, want: NodeSlice{}},
		{name: "bad index: negative", slice: NodeSlice{Node(0), Node(1)}, args: args{-1}, want: NodeSlice{Node(0), Node(1)}},
		{name: "bad index: too high", slice: NodeSlice{Node(0), Node(1)}, args: args{5}, want: NodeSlice{Node(0), Node(1)}},

		{name: "len(1), remove [0]", slice: NodeSlice{Node(0)}, args: args{0}, want: NodeSlice{}},
		{name: "len(2), remove [0]", slice: NodeSlice{Node(0), Node(1)}, args: args{0}, want: NodeSlice{Node(1)}},
		{name: "len(3), remove [1]", slice: NodeSlice{Node(0), Node(1), Node(2)}, args: args{1}, want: NodeSlice{Node(0), Node(2)}},
		{name: "len(3), remove [2]", slice: NodeSlice{Node(0), Node(1), Node(2)}, args: args{2}, want: NodeSlice{Node(0), Node(1)}},
		{name: "len(4), remove [0]", slice: NodeSlice{Node(0), Node(1), Node(2), Node(4)}, args: args{0}, want: NodeSlice{Node(4), Node(1), Node(2)}},

		{name: "duplicates", slice: NodeSlice{}, args: args{0}, want: NodeSlice{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.slice.removeAt(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NodeSlice.removeAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeSlice_Remove(t *testing.T) {
	type args struct {
		n Node
	}
	tests := []struct {
		name  string
		slice NodeSlice
		args  args
		want  NodeSlice
	}{
		{name: "empty", slice: NodeSlice{}, args: args{Node(0)}, want: NodeSlice{}},
		{name: "node not in slice", slice: NodeSlice{Node(0), Node(1)}, args: args{Node(5)}, want: NodeSlice{Node(0), Node(1)}},

		{name: "len(4), remove Node(0)", slice: NodeSlice{Node(0), Node(1), Node(2), Node(3)}, args: args{Node(0)}, want: NodeSlice{Node(3), Node(1), Node(2)}},
		{name: "len(4), remove Node(1)", slice: NodeSlice{Node(0), Node(1), Node(2), Node(3)}, args: args{Node(1)}, want: NodeSlice{Node(0), Node(3), Node(2)}},
		{name: "len(4), remove Node(2)", slice: NodeSlice{Node(0), Node(1), Node(2), Node(3)}, args: args{Node(2)}, want: NodeSlice{Node(0), Node(1), Node(3)}},
		{name: "len(4), remove Node(3)", slice: NodeSlice{Node(0), Node(1), Node(2), Node(3)}, args: args{Node(3)}, want: NodeSlice{Node(0), Node(1), Node(2)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.slice.Remove(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NodeSlice.Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeSlice_pop(t *testing.T) {
	tests := []struct {
		name  string
		slice NodeSlice
		want  NodeSlice
		want1 Node
	}{
		{name: "empty", slice: NodeSlice{}, want: NodeSlice{}, want1: nil},
		{name: "len(1)", slice: NodeSlice{Node(0)}, want: NodeSlice{}, want1: Node(0)},
		{name: "len(2)", slice: NodeSlice{Node(0), Node(1)}, want: NodeSlice{Node(0)}, want1: Node(1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.slice.pop()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NodeSlice.pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NodeSlice.pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

// func TestToNodeSlice(t *testing.T) {
// 	type point struct{ x, y int }

// 	type args struct {
// 		slice interface{}
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want NodeSlice
// 	}{
// 		{name: "[]float32", args: args{slice: []float32{0.1, -10, 3.14159, 100021.1}},
// 			want: NodeSlice{Node(float32(0.1)), Node(float32(-10)), Node(float32(3.14159)), Node(float32(100021.1))}},

// 		{name: "[]uint8", args: args{slice: []uint8{1}},
// 			want: NodeSlice{Node(uint8(1))}},

// 		{name: "[]struct", args: args{slice: []point{{0, 0}, {1, 5}}},
// 			want: NodeSlice{Node(point{0, 0}), Node(point{1, 5})}},

// 		{name: "empty", args: args{slice: []int{}}, want: NodeSlice{}},
// 		{name: "identity", args: args{slice: NodeSlice{Node(0)}}, want: NodeSlice{Node(0)}},
// 		{name: "bad: non-slice arg", args: args{slice: point{}}, want: nil},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := ToNodeSlice(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("ToNodeSlice() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
