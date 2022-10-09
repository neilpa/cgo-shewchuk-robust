package robust_test

import (
	"testing"

	robust "neilpa.me/cgo-shewchuk-robust"
)

func Test_InCircle(t *testing.T) {
	fixtures := loadCases(t, "incircle.txt", 8)
	for _, tt := range fixtures {
		t.Run(tt.label, func(t *testing.T) {
			a := []float64{tt.args[0], tt.args[1]}
			b := []float64{tt.args[2], tt.args[3]}
			c := []float64{tt.args[4], tt.args[5]}
			d := []float64{tt.args[6], tt.args[7]}
			res := robust.InCircle(a, b, c, d)
			assert(t, tt.sign, res)

			va := Vec2{tt.args[0], tt.args[1]}
			vb := Vec2{tt.args[2], tt.args[3]}
			vc := Vec2{tt.args[4], tt.args[5]}
			vd := Vec2{tt.args[6], tt.args[7]}
			res = robust.InCirclePtr(&va.X, &vb.X, &vc.X, &vd.X)
			assert(t, tt.sign, res)

			res = robust.InCircleVec((*robust.XY)(&va), (*robust.XY)(&vb), (*robust.XY)(&vc), (*robust.XY)(&vd))
			assert(t, tt.sign, res)
		})
	}
}

func Benchmark_InCircle(b *testing.B) {
	fixtures := loadCases(b, "incircle.txt", 8)

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
			res = robust.InCircle(ptrs[0], ptrs[1], ptrs[2], ptrs[3])
		}
	}
	result = res
}

func Benchmark_InCirclePtr(b *testing.B) {
	fixtures := loadCases(b, "incircle.txt", 8)

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
			res = robust.InCirclePtr(ptrs[0], ptrs[1], ptrs[2], ptrs[3])
		}
	}
	result = res
}

func Benchmark_InCircleVec(b *testing.B) {
	fixtures := loadCases(b, "incircle.txt", 8)

	tests := make([][4]*Vec2, len(fixtures))
	for i, tt := range fixtures {
		va := Vec2{tt.args[0], tt.args[1]}
		vb := Vec2{tt.args[2], tt.args[3]}
		vc := Vec2{tt.args[4], tt.args[5]}
		vd := Vec2{tt.args[6], tt.args[7]}
		tests[i] = [4]*Vec2{&va, &vb, &vc, &vd}
	}

	b.ResetTimer()
	var res float64
	for n := 0; n < b.N; n++ {
		for _, vecs := range tests {
			res = robust.InCircleVec((*robust.XY)(vecs[0]), (*robust.XY)(vecs[1]), (*robust.XY)(vecs[2]), (*robust.XY)(vecs[3]))
		}
	}
	result = res
}
