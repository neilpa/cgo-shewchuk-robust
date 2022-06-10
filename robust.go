// Package robust provides Go bindings to the `predicates.c` library from
// Jonathan Shewchuk. Only the primary adaptive functions are exported from
// this package.
//
// There are two variants of each core function. One with `[]float64` arguments
// that represent 2D or 3D points and require at least that many elements. The
// other use a "template" pointer to struct point-like argument. This package
// contains `XY` and `XYZ` definitions of these that can be used as cast targets
// for similarly defined point structs.
//
// TODO: A third variant taking `*float` arguments which could cover both above
// cases when numbers are layed out in memory as the C code expects.
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

// XY is a "template" for 2D vector types. It's not intended for use
// directly but as a pointer cast target. See OrientXY and InCircleXY.
type XY struct {
	X, Y float64
}

// XYZ is a "template" for 3D vector types. It's not intended for use
// directly but as a pointer cast target. See OrientXYZ and InSphereXYZ.
type XYZ struct {
	X, Y, Z float64
}

// Orient2D returns a positive value if the points a, b, and c occur in
// counterclockwise order; a negative value if they occur in clockwise
// order; and zero if they are collinear. The result is also a rough
// approximation of twice the signed area of the triangle defined by the
// three points.
//
// Each point slice must be at least 2 elements long
func Orient2D(a, b, c []float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	return float64(C.orient2d(pa, pb, pc))
}

// OrientXY is the same as Orient2D but takes a pointer to a vector-2
// like struct with fields X and Y. This exploits struct layout to avoid
// copying values or allocating new slices.
func OrientXY(a, b, c *XY) float64 {
	pa := (*C.double)(&a.X)
	pb := (*C.double)(&b.X)
	pc := (*C.double)(&c.X)
	return float64(C.orient2d(pa, pb, pc))
}

// Orient3D returns a positive value if the point pd lies below the
// plane passing through a, b, and c; "below" is defined so that a, b,
// and c appear in counterclockwise order when viewed from above the
// plane. Returns a negative value if d lies above the plane. Returns
// zero if the points are coplanar. The result is also a rough
// approximation of six times the signed volume of the tetrahedron
// defined by the four points.
func Orient3D(a, b, c, d []float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	return float64(C.orient3d(pa, pb, pc, pd))
}

// OrientXYZ is the same as Orient3D but takes a pointer to a vector-3
// like struct with fields X, Y, and Z. This exploits struct layout to
// avoid copying values or allocating new slices.
func OrientXYZ(a, b, c, d *XYZ) float64 {
	pa := (*C.double)(&a.X)
	pb := (*C.double)(&b.X)
	pc := (*C.double)(&c.X)
	pd := (*C.double)(&d.X)
	return float64(C.orient3d(pa, pb, pc, pd))
}

// InCircle2D returns a positive value if the point d lies inside the
// circle passing through a, b, and c; a negative value if it lies
// outside; and zero if the four points are cocircular. The points
// a, b, and c must be in counterclockwise order, or the sign of the
// result will be reversed.
func InCircle2D(a, b, c, d []float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	return float64(C.incircle(pa, pb, pc, pd))
}

// InCircleXY is the same as InCircle but takes a pointer to a vector-2
// like struct with fields X and Y. This exploits struct layout to avoid
// copying values or allocating new slices.
func InCircleXY(a, b, c, d *XY) float64 {
	pa := (*C.double)(&a.X)
	pb := (*C.double)(&b.X)
	pc := (*C.double)(&c.X)
	pd := (*C.double)(&d.X)
	return float64(C.incircle(pa, pb, pc, pd))
}

// InSphere3D returns a positive value if the point e lies inside the
// sphere passing through a, b, c, and d; a negative value if it lies
// outside; and zero if the five points are cospherical. The points a,
// b, c, and d must be ordered so that they have a positive orientation
// (as defined by orient3d()), or the sign of the result will be reversed.
func InSphere3D(a, b, c, d, e []float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	pe := (*C.double)(&e[0])
	return float64(C.insphere(pa, pb, pc, pd, pe))
}

// InSphereXYZ is the same as InSphere but takes a pointer to a vector-3
// like struct with fields X, Y, and Z. This exploits struct layout to
// avoid copying values or allocating new slices.
func InSphereXYZ(a, b, c, d, e *XYZ) float64 {
	pa := (*C.double)(&a.X)
	pb := (*C.double)(&b.X)
	pc := (*C.double)(&c.X)
	pd := (*C.double)(&d.X)
	pe := (*C.double)(&e.X)
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

// InCircle2DFast is the naive, non-robust incircle check.
func InCircle2DFast(a, b, c, d []float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	return float64(C.incirclefast(pa, pb, pc, pd))
}

// InSphere3DFast is the naive, non-robust insphere check.
func InSphere3DFast(a, b, c, d, e []float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	pe := (*C.double)(&e[0])
	return float64(C.inspherefast(pa, pb, pc, pd, pe))
}
