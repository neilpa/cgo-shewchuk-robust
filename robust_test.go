package robust_test

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	robust "neilpa.me/cgo-shewchuk-robust"
)

type Vec2 struct{ X, Y float64 }
type Vec3 struct{ X, Y, Z float64 }

var result float64 // benchmark results

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

	fixtures := load(t, "orient.2d", 6)
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
	fixtures := load(b, "orient.2d", 6)

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
	fixtures := load(b, "orient.2d", 6)

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
	fixtures := load(b, "orient.2d", 6)
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

func Test_Orient3d(t *testing.T) {
	tests := []struct {
		ax, ay, az, bx, by, bz, cx, cy, cz, dx, dy, dz float64
		want                                           int
	}{
		{0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 1, 1, -1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("basic: %d", i), func(t *testing.T) {
			a := []float64{tt.ax, tt.ay, tt.az}
			b := []float64{tt.bx, tt.by, tt.bz}
			c := []float64{tt.cx, tt.cy, tt.cz}
			d := []float64{tt.dx, tt.dy, tt.dz}
			assert(t, tt.want, robust.Orient3Ds(a, b, c, d))
		})
	}

	fixtures := load(t, "orient.3d", 12)
	for i, tt := range fixtures {
		t.Run(fmt.Sprintf("data: %d", i+1), func(t *testing.T) {
			a := []float64{tt.args[0], tt.args[1], tt.args[2]}
			b := []float64{tt.args[3], tt.args[4], tt.args[5]}
			c := []float64{tt.args[6], tt.args[7], tt.args[8]}
			d := []float64{tt.args[9], tt.args[10], tt.args[11]}
			res := robust.Orient3Ds(a, b, c, d)
			assert(t, tt.sign, res)

			va := Vec3{tt.args[0], tt.args[1], tt.args[2]}
			vb := Vec3{tt.args[3], tt.args[4], tt.args[5]}
			vc := Vec3{tt.args[6], tt.args[7], tt.args[8]}
			vd := Vec3{tt.args[9], tt.args[10], tt.args[11]}
			res = robust.Orient3D(&va.X, &vb.X, &vc.X, &vd.X)
			assert(t, tt.sign, res)
		})
	}
}

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

	fixtures := load(t, "insphere.2d", 8)
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

func assert(t *testing.T, want int, got float64) {
	t.Helper()
	if sign(got) != want {
		t.Errorf("want: %d; got: sign(%g)", want, got)
	}
}

func sign(n float64) int {
	if n < 0 {
		return -1
	}
	if n > 0 {
		return 1
	}
	return 0
}

type testcase struct {
	args []float64
	sign int
}

func load(t testing.TB, path string, coords int) []testcase {
	f, err := os.Open("test_data/" + path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var tests []testcase
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var tt testcase
		parts := strings.Split(scanner.Text(), " ")
		if len(parts) != coords+2 {
			t.Fatalf("Coord count doens't match, got: %d want: %d", len(parts)-2, coords)
		}

		for _, field := range parts[1 : len(parts)-1] {
			n, err := strconv.ParseFloat(field, 64)
			if err != nil {
				t.Fatal(err)
			}
			tt.args = append(tt.args, n)
		}
		tt.sign, err = strconv.Atoi(parts[len(parts)-1])
		if err != nil {
			t.Fatal(err)
		}
		tests = append(tests, tt)
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
	return tests
}
