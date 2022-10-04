package robust_test

import (
	"fmt"
	"testing"

	robust "neilpa.me/cgo-shewchuk-robust"
)

func Test_Orient2d(t *testing.T) {
	tests := []struct {
		ax, ay, bx, by, cx, cy float64
		want                   int
	}{
		{0, 0, 0, 1, 1, 0, -1},
		{0, 0, 1, 0, 0, 1, 1},

		{0, 1e-66, 0, 1e-66, 1e-65, 0, 0},
		{0, 1e-66, 0, 2e-66, 1e-65, 0, -1},
		{0, 1e-30, 0, 1e-30, 1e-65, 0, 0},
		{1e-30, 0, 1e-30, 0, 0, 1e-65, 0},
		{1e-30, 0, 0, 1e-65, 1e-30, 0, 0},

		// https://blog.devgenius.io/floating-point-round-off-errors-in-geometric-algorithms-a8779662904b
		//
		// Failure 1 - seems to "work" in all cases, but golang differs slightly
		{24.00000000000005, 24.000000000000053, 7.3000000000000194, 7.3000000000000167, 0.50000000000001621, 0.5000000000, 1},
		// Failure 2 - This one produces 0 in faster and a slightly different version in golang
		{27.643564356435643, -21.881188118811881, 83.366336633663366, 15.544554455445542, 73.415841584158414, 8.8613861386138595, 1},
		// Failure 3 - All work, but faster gets a slightly different answer
		{-233.33333333333334, 50.93333333333333, 200.0, 49.200000000000003, 166.66666666666669, 49.333333333333336, 1},
		// Failure 4 - Both faster and golang are broken in this test
		{0.50000000000001243, 0.50000000000000189, 24.000000000000068, 24.000000000000071, 17.300000000000001, 17.300000000000001, 1},

		// TODO:
		// http://wwwisg.cs.uni-magdeburg.de/ag/ClassroomExample/another_classroom_example_slides.pdf
		// http://wwwisg.cs.uni-magdeburg.de/ag/ClassroomExample/another_classroom_example.pdf
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("basic: %d", i), func(t *testing.T) {
			a := []float64{tt.ax, tt.ay}
			b := []float64{tt.bx, tt.by}
			c := []float64{tt.cx, tt.cy}
			assert(t, tt.want, robust.Orient2Ds(a, b, c))
			assert(t, tt.want, robust.Orient2D(&a[0], &b[0], &c[0]))
		})
	}

	fixtures := load(t, "orient2d.txt", 6)
	for i, tt := range fixtures {
		t.Run(fmt.Sprintf("data: %d", i+1), func(t *testing.T) {
			a := []float64{tt.args[0], tt.args[1]}
			b := []float64{tt.args[2], tt.args[3]}
			c := []float64{tt.args[4], tt.args[5]}
			res := robust.Orient2Ds(a, b, c)
			assert(t, tt.sign, res)

			va := Vec2{tt.args[0], tt.args[1]}
			vb := Vec2{tt.args[2], tt.args[3]}
			vc := Vec2{tt.args[4], tt.args[5]}
			res = robust.Orient2D(&va.X, &vb.X, &vc.X)
			assert(t, tt.sign, res)
		})
	}
}

func Benchmark_Orient2D_Vec2(b *testing.B) {
	fixtures := load(b, "orient2d.txt", 6)

	tests := make([][3]Vec2, len(fixtures))
	for i, tt := range fixtures {
		tests[i] = [3]Vec2{
			{tt.args[0], tt.args[1]},
			{tt.args[2], tt.args[3]},
			{tt.args[4], tt.args[5]},
		}
	}

	b.ResetTimer()
	var res float64
	for n := 0; n < b.N; n++ {
		for _, vecs := range tests {
			va, vb, vc := vecs[0], vecs[1], vecs[2]
			res = robust.Orient2D(&va.X, &vb.X, &vc.X)
		}
	}
	result = res
}

func Benchmark_Orient2D_Ptr(b *testing.B) {
	fixtures := load(b, "orient2d.txt", 6)

	tests := make([][3]*float64, len(fixtures))
	for i, tt := range fixtures {
		va := Vec2{tt.args[0], tt.args[1]}
		vb := Vec2{tt.args[2], tt.args[3]}
		vc := Vec2{tt.args[4], tt.args[5]}
		tests[i] = [3]*float64{&va.X, &vb.X, &vc.X}
	}

	b.ResetTimer()
	var res float64
	for n := 0; n < b.N; n++ {
		for _, ptrs := range tests {
			res = robust.Orient2D(ptrs[0], ptrs[1], ptrs[2])
		}
	}
	result = res
}

func Benchmark_Orient2D_Slice(b *testing.B) {
	fixtures := load(b, "orient2d.txt", 6)
	tests := make([][3][]float64, len(fixtures))
	for i, tt := range fixtures {
		tests[i] = [3][]float64{
			{tt.args[0], tt.args[1]},
			{tt.args[2], tt.args[3]},
			{tt.args[4], tt.args[5]},
		}
	}

	b.ResetTimer()
	var res float64
	for n := 0; n < b.N; n++ {
		for _, arr := range tests {
			res = robust.Orient2Ds(arr[0], arr[1], arr[2])
		}
	}
	result = res
}
