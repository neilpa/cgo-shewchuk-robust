package robust_test

import (
	"testing"

	robust "neilpa.me/cgo-shewchuk-robust"
)

func Test_Orient2(t *testing.T) {
	fixtures := loadCases(t, "orient2.txt", 6)
	for _, tt := range fixtures {
		t.Run(tt.label, func(t *testing.T) {
			a := []float64{tt.args[0], tt.args[1]}
			b := []float64{tt.args[2], tt.args[3]}
			c := []float64{tt.args[4], tt.args[5]}
			res := robust.Orient2(a, b, c)
			assert(t, tt.sign, res)

			va := Vec2{tt.args[0], tt.args[1]}
			vb := Vec2{tt.args[2], tt.args[3]}
			vc := Vec2{tt.args[4], tt.args[5]}
			res = robust.Orient2Ptr(&va.X, &vb.X, &vc.X)
			assert(t, tt.sign, res)

			res = robust.Orient2Vec((*robust.XY)(&va), (*robust.XY)(&vb), (*robust.XY)(&vc))
			assert(t, tt.sign, res)
		})
	}
}

func Benchmark_Orient2(b *testing.B) {
	fixtures := loadCases(b, "orient2.txt", 6)
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
			res = robust.Orient2(arr[0], arr[1], arr[2])
		}
	}
	result = res
}

func Benchmark_Orient2Ptr(b *testing.B) {
	fixtures := loadCases(b, "orient2.txt", 6)

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
			res = robust.Orient2Ptr(ptrs[0], ptrs[1], ptrs[2])
		}
	}
	result = res
}

func Benchmark_Orient2Vec(b *testing.B) {
	fixtures := loadCases(b, "orient2.txt", 6)

	tests := make([][3]*Vec2, len(fixtures))
	for i, tt := range fixtures {
		tests[i] = [3]*Vec2{
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
			res = robust.Orient2Vec((*robust.XY)(va), (*robust.XY)(vb), (*robust.XY)(vc))
		}
	}
	result = res
}
