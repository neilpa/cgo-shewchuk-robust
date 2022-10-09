package robust_test

import (
	"testing"

	robust "neilpa.me/cgo-shewchuk-robust"
)

func Test_InSphere(t *testing.T) {
	fixtures := loadCases(t, "insphere.txt", 15)
	for _, tt := range fixtures {
		t.Run(tt.label, func(t *testing.T) {
			a := []float64{tt.args[0], tt.args[1], tt.args[2]}
			b := []float64{tt.args[3], tt.args[4], tt.args[5]}
			c := []float64{tt.args[6], tt.args[7], tt.args[8]}
			d := []float64{tt.args[9], tt.args[10], tt.args[11]}
			e := []float64{tt.args[12], tt.args[13], tt.args[14]}
			res := robust.InSphere(a, b, c, d, e)
			assert(t, tt.sign, res)

			va := Vec3{tt.args[0], tt.args[1], tt.args[2]}
			vb := Vec3{tt.args[3], tt.args[4], tt.args[5]}
			vc := Vec3{tt.args[6], tt.args[7], tt.args[8]}
			vd := Vec3{tt.args[9], tt.args[10], tt.args[11]}
			ve := Vec3{tt.args[12], tt.args[13], tt.args[14]}
			res = robust.InSpherePtr(&va.X, &vb.X, &vc.X, &vd.X, &ve.X)
			assert(t, tt.sign, res)

			res = robust.InSphereVec((*robust.XYZ)(&va), (*robust.XYZ)(&vb), (*robust.XYZ)(&vc), (*robust.XYZ)(&vd), (*robust.XYZ)(&ve))
			assert(t, tt.sign, res)
		})
	}
}

func Benchmark_InSphere(b *testing.B) {
	fixtures := loadCases(b, "insphere.txt", 15)

	tests := make([][5][]float64, len(fixtures))
	for i, tt := range fixtures {
		a := []float64{tt.args[0], tt.args[1], tt.args[2]}
		b := []float64{tt.args[3], tt.args[4], tt.args[5]}
		c := []float64{tt.args[6], tt.args[7], tt.args[8]}
		d := []float64{tt.args[9], tt.args[10], tt.args[11]}
		e := []float64{tt.args[12], tt.args[13], tt.args[14]}
		tests[i] = [5][]float64{a, b, c, d, e}
	}

	b.ResetTimer()
	var res float64
	for n := 0; n < b.N; n++ {
		for _, arrs := range tests {
			res = robust.InSphere(arrs[0], arrs[1], arrs[2], arrs[3], arrs[4])
		}
	}
	result = res
}

func Benchmark_InSpherePtr(b *testing.B) {
	fixtures := loadCases(b, "insphere.txt", 15)

	tests := make([][5]*float64, len(fixtures))
	for i, tt := range fixtures {
		va := Vec3{tt.args[0], tt.args[1], tt.args[2]}
		vb := Vec3{tt.args[3], tt.args[4], tt.args[5]}
		vc := Vec3{tt.args[6], tt.args[7], tt.args[8]}
		vd := Vec3{tt.args[9], tt.args[10], tt.args[11]}
		ve := Vec3{tt.args[12], tt.args[13], tt.args[14]}
		tests[i] = [5]*float64{&va.X, &vb.X, &vc.X, &vd.X, &ve.X}
	}

	b.ResetTimer()
	var res float64
	for n := 0; n < b.N; n++ {
		for _, ptrs := range tests {
			res = robust.InSpherePtr(ptrs[0], ptrs[1], ptrs[2], ptrs[3], ptrs[4])
		}
	}
	result = res
}

func Benchmark_InSphereVec(b *testing.B) {
	fixtures := loadCases(b, "insphere.txt", 15)

	tests := make([][5]*Vec3, len(fixtures))
	for i, tt := range fixtures {
		va := Vec3{tt.args[0], tt.args[1], tt.args[2]}
		vb := Vec3{tt.args[3], tt.args[4], tt.args[5]}
		vc := Vec3{tt.args[6], tt.args[7], tt.args[8]}
		vd := Vec3{tt.args[9], tt.args[10], tt.args[11]}
		ve := Vec3{tt.args[12], tt.args[13], tt.args[14]}
		tests[i] = [5]*Vec3{&va, &vb, &vc, &vd, &ve}
	}

	b.ResetTimer()
	var res float64
	for n := 0; n < b.N; n++ {
		for _, vecs := range tests {
			res = robust.InSphereVec((*robust.XYZ)(vecs[0]), (*robust.XYZ)(vecs[1]), (*robust.XYZ)(vecs[2]), (*robust.XYZ)(vecs[3]), (*robust.XYZ)(vecs[4]))
		}
	}
	result = res
}
