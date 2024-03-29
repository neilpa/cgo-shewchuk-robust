package robust_test

import (
	"testing"

	robust "neilpa.me/cgo-shewchuk-robust"
)

func Test_Orient3(t *testing.T) {
	fixtures := loadCases(t, "orient3.txt", 12)
	for _, tt := range fixtures {
		t.Run(tt.label, func(t *testing.T) {
			a := []float64{tt.args[0], tt.args[1], tt.args[2]}
			b := []float64{tt.args[3], tt.args[4], tt.args[5]}
			c := []float64{tt.args[6], tt.args[7], tt.args[8]}
			d := []float64{tt.args[9], tt.args[10], tt.args[11]}
			res := robust.Orient3(a, b, c, d)
			assert(t, tt.sign, res)

			va := Vec3{tt.args[0], tt.args[1], tt.args[2]}
			vb := Vec3{tt.args[3], tt.args[4], tt.args[5]}
			vc := Vec3{tt.args[6], tt.args[7], tt.args[8]}
			vd := Vec3{tt.args[9], tt.args[10], tt.args[11]}
			res = robust.Orient3Ptr(&va.X, &vb.X, &vc.X, &vd.X)
			assert(t, tt.sign, res)

			res = robust.Orient3Vec((*robust.XYZ)(&va), (*robust.XYZ)(&vb), (*robust.XYZ)(&vc), (*robust.XYZ)(&vd))
			assert(t, tt.sign, res)
		})
	}
}

func Benchmark_Orient3(b *testing.B) {
	fixtures := loadCases(b, "orient3.txt", 12)
	tests := make([][4][]float64, len(fixtures))
	for i, tt := range fixtures {
		tests[i] = [4][]float64{
			{tt.args[0], tt.args[1], tt.args[2]},
			{tt.args[3], tt.args[4], tt.args[5]},
			{tt.args[6], tt.args[7], tt.args[8]},
			{tt.args[9], tt.args[10], tt.args[11]},
		}
	}

	b.ResetTimer()
	var res float64
	for n := 0; n < b.N; n++ {
		for _, arr := range tests {
			res = robust.Orient3(arr[0], arr[1], arr[2], arr[3])
		}
	}
	result = res
}

func Benchmark_Orient3Ptr(b *testing.B) {
	fixtures := loadCases(b, "orient3.txt", 12)

	tests := make([][4]*float64, len(fixtures))
	for i, tt := range fixtures {
		va := Vec3{tt.args[0], tt.args[1], tt.args[2]}
		vb := Vec3{tt.args[3], tt.args[4], tt.args[5]}
		vc := Vec3{tt.args[6], tt.args[7], tt.args[8]}
		vd := Vec3{tt.args[9], tt.args[10], tt.args[11]}
		tests[i] = [4]*float64{&va.X, &vb.X, &vc.X, &vd.X}
	}

	b.ResetTimer()
	var res float64
	for n := 0; n < b.N; n++ {
		for _, ptrs := range tests {
			res = robust.Orient3Ptr(ptrs[0], ptrs[1], ptrs[2], ptrs[3])
		}
	}
	result = res
}

func Benchmark_Orient3Vec(b *testing.B) {
	fixtures := loadCases(b, "orient3.txt", 12)

	tests := make([][4]*Vec3, len(fixtures))
	for i, tt := range fixtures {
		va := Vec3{tt.args[0], tt.args[1], tt.args[2]}
		vb := Vec3{tt.args[3], tt.args[4], tt.args[5]}
		vc := Vec3{tt.args[6], tt.args[7], tt.args[8]}
		vd := Vec3{tt.args[9], tt.args[10], tt.args[11]}
		tests[i] = [4]*Vec3{&va, &vb, &vc, &vd}
	}

	b.ResetTimer()
	var res float64
	for n := 0; n < b.N; n++ {
		for _, vecs := range tests {
			res = robust.Orient3Vec((*robust.XYZ)(vecs[0]), (*robust.XYZ)(vecs[1]), (*robust.XYZ)(vecs[2]), (*robust.XYZ)(vecs[3]))
		}
	}
	result = res
}
