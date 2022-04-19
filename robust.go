// Package robust provides Go bindings to the `predicates.c` library from
// Jonathan Shewchuk. Only the primary adaptive functions are exported from
// this package.
package robust

// void exactinit();
// double orient2d(double *pa, double *pb, double *pc);
// double orient2dfast(double *pa, double *pb, double *pc);
import "C"

func init() {
	C.exactinit()
}

func Orient2d(a, b, c [2]float64) float64 {
	return float64(C.orient2d((*C.double)(&a[0]), (*C.double)(&b[0]), (*C.double)(&c[0])))
}

func Orient2dFast(a, b, c [2]float64) float64 {
	return float64(C.orient2dfast((*C.double)(&a[0]), (*C.double)(&b[0]), (*C.double)(&c[0])))
}

// func Orient3d() float64 { panic("todo") }
// func InCircle() float64 { panic("todo") }
// func InSphere() float64 { panic("todo") }
