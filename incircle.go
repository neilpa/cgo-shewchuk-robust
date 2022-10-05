package robust

// double incircle(double *pa, double *pb, double *pc, double *pd);
// double incircleadapt(double *pa, double *pb, double *pc, double *pd, double permanent);
import "C"
import "math"

// InCircle returns a positive value if the point d lies inside the
// circle passing through a, b, and c; a negative value if it lies
// outside; and zero if the four points are cocircular. The points
// a, b, and c must be in counterclockwise order, or the sign of the
// result will be reversed.
//
// Each slice parameter must contain at least 2 values.
func InCircle(a, b, c, d []float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	return inCircle(pa, pb, pc, pd,
		a[0]-d[0], b[0]-d[0], c[0]-d[0],
		a[1]-d[1], b[1]-d[1], c[1]-d[1])
}

// InCircleVec is similiar to `InCircle` but takes a point-like struct
// pointer rather than a slice.
func InCircleVec(a, b, c, d *XY) float64 {
	pa := (*C.double)(&a.X)
	pb := (*C.double)(&b.X)
	pc := (*C.double)(&c.X)
	pd := (*C.double)(&d.X)
	return inCircle(pa, pb, pc, pd,
		a.X-d.X, b.X-d.X, c.X-d.X,
		a.Y-d.Y, b.Y-d.Y, c.Y-d.Y)
}

// InCirclePtr is the direct wrapper of `incircle` from `predicates.c`.
// See `InCircle` for additional details.
func InCirclePtr(a, b, c, d *float64) float64 {
	pa := (*C.double)(a)
	pb := (*C.double)(b)
	pc := (*C.double)(c)
	pd := (*C.double)(d)
	return float64(C.incircle(pa, pb, pc, pd))
}

// inCircle implements the basic error bound checks to minimize
// CGO calls to the adaptive implementation.
func inCircle(pa, pb, pc, pd *C.double,
	adx, bdx, cdx, ady, bdy, cdy float64) float64 {

	var bdxcdy, cdxbdy, cdxady, adxcdy, adxbdy, bdxady float64
	var alift, blift, clift float64
	var det float64
	var permanent, errbound float64

	bdxcdy = bdx * cdy
	cdxbdy = cdx * bdy
	alift = adx*adx + ady*ady

	cdxady = cdx * ady
	adxcdy = adx * cdy
	blift = bdx*bdx + bdy*bdy

	adxbdy = adx * bdy
	bdxady = bdx * ady
	clift = cdx*cdx + cdy*cdy

	det =
		alift*(bdxcdy-cdxbdy) +
			blift*(cdxady-adxcdy) +
			clift*(adxbdy-bdxady)

	permanent =
		(math.Abs(bdxcdy)+math.Abs(cdxbdy))*alift +
			(math.Abs(cdxady)+math.Abs(adxcdy))*blift +
			(math.Abs(adxbdy)+math.Abs(bdxady))*clift

	errbound = iccerrboundA * permanent
	if (det > errbound) || (-det > errbound) {
		return det
	}

	return float64(C.incircleadapt(pa, pb, pc, pd, C.double(permanent)))
}
