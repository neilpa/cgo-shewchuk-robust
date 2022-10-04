// Package robust provides Go bindings to the `predicates.c` library from
// Jonathan Shewchuk. Only the primary adaptive functions are exported from
// this package.
//
// Each of the core functions takes `*float64` arguments that should be
// C-like arrays of at least 2 or 3 values for the respective dimensionality.
// For use with slice types or flat buffers of point values, take the address
// of the x-coordinate of the target index, e.g.
//
//	res := robust.Orient2D(&buf[i], &buf[i+2], &buf[i+4])
//
// Alternatively, if you have point-like struct types with X, Y or X, Y, Z
// coordinates, use the address of the X field, e.g.
//
//	res := robust.Orient2D(&p0.X, &p1.X, &p2.X)
//
// There are also `s` suffixed convenience functions taking `[]float64`
// arguments that use the first element as the X coordinate.
package robust

// void exactinit();
// extern double ccwerrboundA, o3derrboundA;
// double orient2d(double *pa, double *pb, double *pc);
// double orient2dadapt(double *pa, double *pb, double *pc, double detsum);
// double orient3d(double *pa, double *pb, double *pc, double *pd);
// double orient3dadapt(double *pa, double *pb, double *pc, double *pd, double permanent);
// double incircle(double *pa, double *pb, double *pc, double *pd);
// double insphere(double *pa, double *pb, double *pc, double *pd, double *pe);
import "C"
import "math"

// Cache values from CGO init
var (
	ccwerrboundA, o3derrboundA float64
)

func init() {
	C.exactinit()
	ccwerrboundA = float64(C.ccwerrboundA)
	o3derrboundA = float64(C.o3derrboundA)
}

// Orient2D returns a positive value if the points a, b, and c occur in
// counterclockwise order; a negative value if they occur in clockwise
// order; and zero if they are collinear. The result is also a rough
// approximation of twice the signed area of the triangle defined by the
// three points.
//
// Each pointer must at least 2 contiguous values.
func Orient2D(a, b, c *float64) float64 {
	pa := (*C.double)(a)
	pb := (*C.double)(b)
	pc := (*C.double)(c)
	return float64(C.orient2d(pa, pb, pc))
}

// Orient2Ds is a convenience wrapper for Orient2D. Each slice must
// be at least 2 elements long, additional elements are ignored.
func Orient2Ds(a, b, c []float64) float64 {
	var detleft, detright, det float64
	var detsum float64

	detleft = (a[0] - c[0]) * (b[1] - c[1])
	detright = (a[1] - c[1]) * (b[0] - c[0])
	det = detleft - detright

	if detleft > 0.0 {
		if detright <= 0.0 {
			return det
		} else {
			detsum = detleft + detright
		}
	} else if detleft < 0.0 {
		if detright >= 0.0 {
			return det
		} else {
			detsum = -detleft - detright
		}
	} else {
		return det
	}

	errbound := ccwerrboundA * detsum
	if (det >= errbound) || (-det >= errbound) {
		return det
	}

	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	return float64(C.orient2dadapt(pa, pb, pc, C.double(detsum)))
}

// Orient3D returns a positive value if the point pd lies below the
// plane passing through a, b, and c; "below" is defined so that a, b,
// and c appear in counterclockwise order when viewed from above the
// plane. Returns a negative value if d lies above the plane. Returns
// zero if the points are coplanar. The result is also a rough
// approximation of six times the signed volume of the tetrahedron
// defined by the four points.
func Orient3D(a, b, c, d *float64) float64 {
	pa := (*C.double)(a)
	pb := (*C.double)(b)
	pc := (*C.double)(c)
	pd := (*C.double)(d)
	return float64(C.orient3d(pa, pb, pc, pd))
}

// Orient3Ds is a convenience wrapper for Orient3D. Each slice must
// be at least 3 elements long, additional elements are ignored.
func Orient3Ds(a, b, c, d []float64) float64 {
	var adx, bdx, cdx, ady, bdy, cdy, adz, bdz, cdz float64
	var bdxcdy, cdxbdy, cdxady, adxcdy, adxbdy, bdxady float64
	var det float64
	var permanent, errbound float64

	adx = a[0] - d[0]
	bdx = b[0] - d[0]
	cdx = c[0] - d[0]
	ady = a[1] - d[1]
	bdy = b[1] - d[1]
	cdy = c[1] - d[1]
	adz = a[2] - d[2]
	bdz = b[2] - d[2]
	cdz = c[2] - d[2]

	bdxcdy = bdx * cdy
	cdxbdy = cdx * bdy

	cdxady = cdx * ady
	adxcdy = adx * cdy

	adxbdy = adx * bdy
	bdxady = bdx * ady

	det =
		adz*(bdxcdy-cdxbdy) +
			bdz*(cdxady-adxcdy) +
			cdz*(adxbdy-bdxady)

	permanent =
		(math.Abs(bdxcdy)+math.Abs(cdxbdy))*math.Abs(adz) +
			(math.Abs(cdxady)+math.Abs(adxcdy))*math.Abs(bdz) +
			(math.Abs(adxbdy)+math.Abs(bdxady))*math.Abs(cdz)

	errbound = o3derrboundA * permanent
	if (det > errbound) || (-det > errbound) {
		return det
	}

	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	return float64(C.orient3dadapt(pa, pb, pc, pd, C.double(permanent)))
}

// InCircle2D returns a positive value if the point d lies inside the
// circle passing through a, b, and c; a negative value if it lies
// outside; and zero if the four points are cocircular. The points
// a, b, and c must be in counterclockwise order, or the sign of the
// result will be reversed.
func InCircle2D(a, b, c, d *float64) float64 {
	pa := (*C.double)(a)
	pb := (*C.double)(b)
	pc := (*C.double)(c)
	pd := (*C.double)(d)
	return float64(C.incircle(pa, pb, pc, pd))
}

// InCircle2Ds is a convenience wrapper for InCircle2D. Each slice must
// be at least 2 elements long, additional elements are ignored.
func InCircle2Ds(a, b, c, d []float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	return float64(C.incircle(pa, pb, pc, pd))
}

// InSphere3D returns a positive value if the point e lies inside the
// sphere passing through a, b, c, and d; a negative value if it lies
// outside; and zero if the five points are cospherical. The points a,
// b, c, and d must be ordered so that they have a positive orientation
// (as defined by orient3d()), or the sign of the result will be reversed.
func InSphere3D(a, b, c, d, e *float64) float64 {
	pa := (*C.double)(a)
	pb := (*C.double)(b)
	pc := (*C.double)(c)
	pd := (*C.double)(d)
	pe := (*C.double)(e)
	return float64(C.insphere(pa, pb, pc, pd, pe))
}

// InSphere3Ds is a convenience wrapper for InSphere3D. Each slice must
// be at least 3 elements long, additional elements are ignored.
func InSphere3Ds(a, b, c, d, e []float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	pe := (*C.double)(&e[0])
	return float64(C.insphere(pa, pb, pc, pd, pe))
}
