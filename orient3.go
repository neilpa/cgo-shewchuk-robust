package robust

// double orient3d(double *pa, double *pb, double *pc, double *pd);
// double orient3dadapt(double *pa, double *pb, double *pc, double *pd, double permanent);
import "C"
import "math"

// Orient3D returns a positive value if the point pd lies below the
// plane passing through a, b, and c; "below" is defined so that a, b,
// and c appear in counterclockwise order when viewed from above the
// plane. Returns a negative value if d lies above the plane. Returns
// zero if the points are coplanar. The result is also a rough
// approximation of six times the signed volume of the tetrahedron
// defined by the four points.
//
// Each slice parameter must contain at least 3 values.
func Orient3(a, b, c, d []float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	return orient3(pa, pb, pc, pd,
		a[0]-d[0], b[0]-d[0], c[0]-d[0],
		a[1]-d[1], b[1]-d[1], c[1]-d[1],
		a[2]-d[2], b[2]-d[2], c[2]-d[2])
}

// Orient3Vec is similiar to `Orient3` but takes a point-like struct
// pointer rather than a slice.
func Orient3Vec(a, b, c, d *XYZ) float64 {
	pa := (*C.double)(&a.X)
	pb := (*C.double)(&b.X)
	pc := (*C.double)(&c.X)
	pd := (*C.double)(&d.X)
	return orient3(pa, pb, pc, pd,
		a.X-d.X, b.X-d.X, c.X-d.X,
		a.Y-d.Y, b.Y-d.Y, c.Y-d.Y,
		a.Z-d.Z, b.Z-d.Z, c.Z-d.Z)
}

// Orient3Ptr is the direct wrapper of `orient3d` from `predicates.c`.
// See `Orient3` for additional details.
func Orient3Ptr(a, b, c, d *float64) float64 {
	pa := (*C.double)(a)
	pb := (*C.double)(b)
	pc := (*C.double)(c)
	pd := (*C.double)(d)
	return float64(C.orient3d(pa, pb, pc, pd))
}

// orient3 implements the basic error bound checks to minimize
// CGO calls to the adaptive implementation.
func orient3(pa, pb, pc, pd *C.double,
	adx, bdx, cdx, ady, bdy, cdy, adz, bdz, cdz float64) float64 {

	var bdxcdy, cdxbdy, cdxady, adxcdy, adxbdy, bdxady float64
	var det float64
	var permanent, errbound float64

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

	return float64(C.orient3dadapt(pa, pb, pc, pd, C.double(permanent)))
}
