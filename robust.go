// Package robust provides Go bindings to the `predicates.c` library from
// Jonathan Shewchuk. Only the primary adaptive functions are exported from
// this package.
package robust

// void exactinit();
// double orient2d(double *pa, double *pb, double *pc);
// double orient2dfast(double *pa, double *pb, double *pc);
// double orient3d(double *pa, double *pb, double *pc, double *pd);
// double incircle(double *pa, double *pb, double *pc, double *pd);
// double insphere(double *pa, double *pb, double *pc, double *pd, double *pe);
import "C"

func init() {
	C.exactinit()
}

// Orient2d TODO
func Orient2d(a, b, c [2]float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	return float64(C.orient2d(pa, pb, pc))
}

func Orient2dFast(a, b, c [2]float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	return float64(C.orient2dfast(pa, pb, pc))
}

// Orient3d TODO
func Orient3d(a, b, c, d [3]float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	return float64(C.orient3d(pa, pb, pc, pd))
}

// InCircle TODO
func InCircle(a, b, c, d [2]float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	return float64(C.incircle(pa, pb, pc, pd))
}

// InSphere TODO
func InSphere(a, b, c, d, e [3]float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	pe := (*C.double)(&e[0])
	return float64(C.insphere(pa, pb, pc, pd, pe))
}
