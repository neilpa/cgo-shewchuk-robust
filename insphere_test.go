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
			assert(t, tt.want, robust.InSphere3Ds(a, b, c, d, e))
		})
	}

	fixtures := load(t, "insphere.3d", 15)
	for i, tt := range fixtures {
		t.Run(fmt.Sprintf("data: %d", i+1), func(t *testing.T) {
			a := []float64{tt.args[0], tt.args[1], tt.args[2]}
			b := []float64{tt.args[3], tt.args[4], tt.args[5]}
			c := []float64{tt.args[6], tt.args[7], tt.args[8]}
			d := []float64{tt.args[9], tt.args[10], tt.args[11]}
			e := []float64{tt.args[12], tt.args[13], tt.args[14]}
			res := robust.InSphere3Ds(a, b, c, d, e)
			assert(t, tt.sign, res)

			va := Vec3{tt.args[0], tt.args[1], tt.args[2]}
			vb := Vec3{tt.args[3], tt.args[4], tt.args[5]}
			vc := Vec3{tt.args[6], tt.args[7], tt.args[8]}
			vd := Vec3{tt.args[9], tt.args[10], tt.args[11]}
			ve := Vec3{tt.args[12], tt.args[13], tt.args[14]}
			res = robust.InSphere3D(&va.X, &vb.X, &vc.X, &vd.X, &ve.X)
			assert(t, tt.sign, res)
		})
	}
}

func Benchmark_InSphere_Ptr(b *testing.B) {
	fixtures := load(b, "insphere.3d", 15)

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
			res = robust.InSphere3D(ptrs[0], ptrs[1], ptrs[2], ptrs[3], ptrs[4])
		}
	}
	result = res
}

func Benchmark_InSphere_Slice(b *testing.B) {
	fixtures := load(b, "insphere.3d", 15)

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
			res = robust.InSphere3Ds(arrs[0], arrs[1], arrs[2], arrs[3], arrs[4])
		}
	}
	result = res
}
