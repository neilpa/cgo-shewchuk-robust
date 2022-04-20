// Package robust provides Go bindings to the `predicates.c` library from
// Jonathan Shewchuk. Only the primary adaptive functions are exported from
// this package.
//
// All arguments are []float64 values presenting either 2D or 3D points
// depending on the function and must have at least that many elements. This
// makes it convenient to directly operate on flat buffers of polygon or
// polyhedra vertices.
package robust

// void exactinit();
// double orient2d(double *pa, double *pb, double *pc);
// double orient2dfast(double *pa, double *pb, double *pc);
// double orient3d(double *pa, double *pb, double *pc, double *pd);
// double orient3dfast(double *pa, double *pb, double *pc, double *pd);
// double incircle(double *pa, double *pb, double *pc, double *pd);
// double incirclefast(double *pa, double *pb, double *pc, double *pd);
// double insphere(double *pa, double *pb, double *pc, double *pd, double *pe);
// double inspherefast(double *pa, double *pb, double *pc, double *pd, double *pe);
import "C"

func init() {
	C.exactinit()
}

// Orient2d returns a positive value if the points a, b, and c occur in
// counterclockwise order; a negative value if they occur in clockwise
// order; and zero if they are collinear. The result is also a rough
// approximation of twice the signed area of the triangle defined by the
// three points.
//
// Each point slice must be at least 2 elements long
func Orient2d(a, b, c []float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	return float64(C.orient2d(pa, pb, pc))
}

// Orient3d returns a positive value if the point pd lies below the
// plane passing through a, b, and c; "below" is defined so that a, b,
// and c appear in counterclockwise order when viewed from above the
// plane. Returns a negative value if d lies above the plane. Returns
// zero if the points are coplanar. The result is also a rough
// approximation of six times the signed volume of the tetrahedron
// defined by the four points.
func Orient3d(a, b, c, d []float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	return float64(C.orient3d(pa, pb, pc, pd))
}

// InCircle returns a positive value if the point d lies inside the
// circle passing through a, b, and c; a negative value if it lies
// outside; and zero if the four points are cocircular. The points
// a, b, and c must be in counterclockwise order, or the sign of the
// result will be reversed.
func InCircle(a, b, c, d []float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	return float64(C.incircle(pa, pb, pc, pd))
}

// InSphere eturns a positive value if the point e lies inside the
// sphere passing through a, b, c, and d; a negative value if it lies
// outside; and zero if the five points are cospherical. The points a,
// b, c, and d must be ordered so that they have a positive orientation
// (as defined by orient3d()), or the sign of the result will be reversed.
func InSphere(a, b, c, d, e []float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	pe := (*C.double)(&e[0])
	return float64(C.insphere(pa, pb, pc, pd, pe))
}

// The following "Fast" algorithms are non-robust. They've been included
// to check against the robust variants but should be avoided otherwise.

// Orient2dFast is the naive, non-robust orient2d check.
func Orient2dFast(a, b, c []float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	return float64(C.orient2dfast(pa, pb, pc))
}

// Orient3dFast is the naive, non-robust orient3d check.
func Orient3dFast(a, b, c, d []float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	return float64(C.orient3dfast(pa, pb, pc, pd))
}

// InCircleFast is the naive, non-robust incircle check.
func InCircleFast(a, b, c, d []float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	return float64(C.incirclefast(pa, pb, pc, pd))
}

// InSphereFast is the naive, non-robust insphere check.
func InSphereFast(a, b, c, d, e []float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	pe := (*C.double)(&e[0])
	return float64(C.inspherefast(pa, pb, pc, pd, pe))
}
