package robust_test

import (
	"fmt"
	"testing"

	robust "neilpa.me/cgo-shewchuk-robust"
)

func Test_InCircle(t *testing.T) {
	tests := []struct {
		ax, ay, bx, by, cx, cy, dx, dy float64
		want                           int
	}{
		{0, 0, 1, 0, 0, 1, 0.5, 0.5, 1},
		{0, 0, 0, 1, 1, 0, 0.5, 0.5, -1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("basic: %d", i), func(t *testing.T) {
			a := []float64{tt.ax, tt.ay}
			b := []float64{tt.bx, tt.by}
			c := []float64{tt.cx, tt.cy}
			d := []float64{tt.dx, tt.dy}
			assert(t, tt.want, robust.InCircle2Ds(a, b, c, d))
		})
	}

	fixtures := load(t, "incircle2d.txt", 8)
	for i, tt := range fixtures {
		t.Run(fmt.Sprintf("data: %d", i+1), func(t *testing.T) {
			a := []float64{tt.args[0], tt.args[1]}
			b := []float64{tt.args[2], tt.args[3]}
			c := []float64{tt.args[4], tt.args[5]}
			d := []float64{tt.args[6], tt.args[7]}
			res := robust.InCircle2Ds(a, b, c, d)
			assert(t, tt.sign, res)

			va := Vec2{tt.args[0], tt.args[1]}
			vb := Vec2{tt.args[2], tt.args[3]}
			vc := Vec2{tt.args[4], tt.args[5]}
			vd := Vec2{tt.args[6], tt.args[7]}
			res = robust.InCircle2D(&va.X, &vb.X, &vc.X, &vd.X)
			assert(t, tt.sign, res)
		})
	}
}

func Benchmark_InCircle_Ptr(b *testing.B) {
	fixtures := load(b, "incircle2d.txt", 8)

	tests := make([][4]*float64, len(fixtures))
	for i, tt := range fixtures {
		va := Vec2{tt.args[0], tt.args[1]}
		vb := Vec2{tt.args[2], tt.args[3]}
		vc := Vec2{tt.args[4], tt.args[5]}
		vd := Vec2{tt.args[6], tt.args[7]}
		tests[i] = [4]*float64{&va.X, &vb.X, &vc.X, &vd.X}
	}

	b.ResetTimer()
	var res float64
	for n := 0; n < b.N; n++ {
		for _, ptrs := range tests {
			res = robust.InCircle2D(ptrs[0], ptrs[1], ptrs[2], ptrs[3])
		}
	}
	result = res
}

func Benchmark_InCircle_Slice(b *testing.B) {
	fixtures := load(b, "incircle2d.txt", 8)

	tests := make([][4][]float64, len(fixtures))
	for i, tt := range fixtures {
		va := []float64{tt.args[0], tt.args[1]}
		vb := []float64{tt.args[2], tt.args[3]}
		vc := []float64{tt.args[4], tt.args[5]}
		vd := []float64{tt.args[6], tt.args[7]}
		tests[i] = [4][]float64{va, vb, vc, vd}
	}

	b.ResetTimer()
	var res float64
	for n := 0; n < b.N; n++ {
		for _, ptrs := range tests {
			res = robust.InCircle2Ds(ptrs[0], ptrs[1], ptrs[2], ptrs[3])
		}
	}
	result = res
}
