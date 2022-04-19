package robust_test

import (
	"fmt"
	"testing"

	robust "neilpa.me/cgo-shewchuk-robust"
)

func Test_Orient2d(t *testing.T) {
	tests := []struct {
		ax, ay, bx, by, cx, cy float64
	}{
		{0, 0, 0, 1, 1, 0},
		{0, 0, 1, 0, 0, 1},

		{0, 1e-66, 0, 1e-66, 1e-65, 0}, // golang is wrong, there are two coincident points
		{0, 1e-66, 0, 2e-66, 1e-65, 0},
		{0, 1e-30, 0, 1e-30, 1e-65, 0}, // golang is wrong, there are two coincident points
		{1e-30, 0, 1e-30, 0, 0, 1e-65}, // golang is wrong, there are two coincident points
		{1e-30, 0, 0, 1e-65, 1e-30, 0},

		// https://blog.devgenius.io/floating-point-round-off-errors-in-geometric-algorithms-a8779662904b
		//
		// Failure 1 - seems to "work" in all cases, but golang differs slightly
		{24.00000000000005, 24.000000000000053, 7.3000000000000194, 7.3000000000000167, 0.50000000000001621, 0.5000000000},
		// Failure 2 - This one produces 0 in faster and a slightly different version in golang
		{27.643564356435643, -21.881188118811881, 83.366336633663366, 15.544554455445542, 73.415841584158414, 8.8613861386138595},
		// Failure 3 - All work, but faster gets a slightly different answer
		{-233.33333333333334, 50.93333333333333, 200.0, 49.200000000000003, 166.66666666666669, 49.333333333333336},
		// Failure 4 - Both faster and golang are broken in this test
		{0.50000000000001243, 0.50000000000000189, 24.000000000000068, 24.000000000000071, 17.300000000000001, 17.300000000000001},

		// TODO:
		// http://wwwisg.cs.uni-magdeburg.de/ag/ClassroomExample/another_classroom_example_slides.pdf
		// http://wwwisg.cs.uni-magdeburg.de/ag/ClassroomExample/another_classroom_example.pdf
	}
	for i, tt := range tests {
		a := [2]float64{tt.ax, tt.ay}
		b := [2]float64{tt.bx, tt.by}
		c := [2]float64{tt.cx, tt.cy}
		fmt.Println(i, "robust", robust.Orient2d(a, b, c))
		fmt.Println(i, "faster", robust.Orient2dFast(a, b, c))

		acx := a[0] - c[0]
		bcx := b[0] - c[0]
		acy := a[1] - c[1]
		bcy := b[1] - c[1]
		fmt.Println(i, "golang", acx*bcy-acy*bcx)
		fmt.Println()
	}
}

func Test_Orient3d(t *testing.T) {
	tests := []struct {
		ax, ay, az, bx, by, bz, cx, cy, cz, dx, dy, dz float64
	}{
		{0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 1, 1},
	}
	for i, tt := range tests {
		a := [3]float64{tt.ax, tt.ay, tt.az}
		b := [3]float64{tt.bx, tt.by, tt.bz}
		c := [3]float64{tt.cx, tt.cy, tt.cz}
		d := [3]float64{tt.dx, tt.dy, tt.dz}

		fmt.Println(i, "robust", robust.Orient3d(a, b, c, d))
	}
}

func Test_InCircle(t *testing.T) {
	tests := []struct {
		ax, ay, bx, by, cx, cy, dx, dy float64
	}{
		{0, 0, 0, 1, 1, 0, 0.5, 0.5},
	}
	for i, tt := range tests {
		a := [2]float64{tt.ax, tt.ay}
		b := [2]float64{tt.bx, tt.by}
		c := [2]float64{tt.cx, tt.cy}
		d := [2]float64{tt.dx, tt.dy}

		fmt.Println(i, "robust", robust.InCircle(a, b, c, d))
	}
}

func Test_InSphere(t *testing.T) {
	tests := []struct {
		ax, ay, az, bx, by, bz, cx, cy, cz, dx, dy, dz, ex, ey, ez float64
	}{
		{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0.5, 0.5, 0.5},
	}
	for i, tt := range tests {
		a := [3]float64{tt.ax, tt.ay, tt.az}
		b := [3]float64{tt.bx, tt.by, tt.bz}
		c := [3]float64{tt.cx, tt.cy, tt.cz}
		d := [3]float64{tt.dx, tt.dy, tt.dz}
		e := [3]float64{tt.ex, tt.ey, tt.ez}

		fmt.Println(i, "robust", robust.InSphere(a, b, c, d, e))
	}
}
