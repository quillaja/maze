package maze

import (
	"reflect"
	"testing"
)

func TestMakeGrid(t *testing.T) {
	type args struct {
		dx int
		dy int
		dz int
	}
	tests := []struct {
		name string
		args args
		want Graph
	}{
		{name: "3x3x3", args: args{dx: 3, dy: 3, dz: 3},
			want: Graph(mapgraph{
				// z=0
				0: NodeSlice{1, 3, 9},
				1: NodeSlice{0, 2, 4, 10},
				2: NodeSlice{1, 5, 11},

				3: NodeSlice{0, 4, 6, 12},
				4: NodeSlice{1, 3, 5, 7, 13},
				5: NodeSlice{2, 4, 8, 14},

				6: NodeSlice{3, 7, 15},
				7: NodeSlice{4, 6, 8, 16},
				8: NodeSlice{5, 7, 17},

				// z=1
				9:  NodeSlice{0, 10, 12, 18},
				10: NodeSlice{1, 9, 11, 13, 19},
				11: NodeSlice{2, 10, 14, 20},

				12: NodeSlice{3, 9, 13, 15, 21},
				13: NodeSlice{4, 10, 12, 14, 16, 22},
				14: NodeSlice{5, 11, 13, 17, 23},

				15: NodeSlice{6, 12, 16, 24},
				16: NodeSlice{7, 13, 15, 17, 25},
				17: NodeSlice{8, 14, 16, 26},

				// z=2
				18: NodeSlice{9, 19, 21},
				19: NodeSlice{10, 18, 20, 22},
				20: NodeSlice{11, 19, 23},

				21: NodeSlice{12, 18, 22, 24},
				22: NodeSlice{13, 19, 21, 23, 25},
				23: NodeSlice{14, 20, 22, 26},

				24: NodeSlice{15, 21, 25},
				25: NodeSlice{16, 22, 24, 26},
				26: NodeSlice{17, 23, 25},
			})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeGrid(tt.args.dx, tt.args.dy, tt.args.dz); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
