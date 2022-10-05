package robust_test

import (
	"fmt"
	"testing"

	robust "neilpa.me/cgo-shewchuk-robust"
)

func Test_InSphere(t *testing.T) {
	tests := []struct {
		ax, ay, az, bx, by, bz, cx, cy, cz, dx, dy, dz, ex, ey, ez float64
		want                                                       int
	}{
		{0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0.5, 0.5, 0.5, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0.5, 0.5, 0.5, -1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("basic: %d", i), func(t *testing.T) {
			a := []float64{tt.ax, tt.ay, tt.az}
			b := []float64{tt.bx, tt.by, tt.bz}
			c := []float64{tt.cx, tt.cy, tt.cz}
			d := []float64{tt.dx, tt.dy, tt.dz}
			e := []float64{tt.ex, tt.ey, tt.ez}
			assert(t, tt.want, robust.InSpherePtr(&a[0], &b[0], &c[0], &d[0], &e[0]))
			assert(t, tt.want, robust.InSphere(a, b, c, d, e))

			va := Vec3{tt.ax, tt.ay, tt.az}
			vb := Vec3{tt.bx, tt.by, tt.bz}
			vc := Vec3{tt.cx, tt.cy, tt.cz}
			vd := Vec3{tt.dx, tt.dy, tt.dz}
			ve := Vec3{tt.ex, tt.ey, tt.ez}
			res := robust.InSphereVec((*robust.XYZ)(&va), (*robust.XYZ)(&vb), (*robust.XYZ)(&vc), (*robust.XYZ)(&vd), (*robust.XYZ)(&ve))
			assert(t, tt.want, res)
		})
	}

	fixtures := load(t, "insphere3d.txt", 15)
	for i, tt := range fixtures {
		t.Run(fmt.Sprintf("data: %d", i+1), func(t *testing.T) {
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
	fixtures := load(b, "insphere3d.txt", 15)

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
	fixtures := load(b, "insphere3d.txt", 15)

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
	fixtures := load(b, "insphere3d.txt", 15)

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
