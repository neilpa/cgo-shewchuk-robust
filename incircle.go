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
	var adx, bdx, cdx, ady, bdy, cdy float64
	var bdxcdy, cdxbdy, cdxady, adxcdy, adxbdy, bdxady float64
	var alift, blift, clift float64
	var det float64
	var permanent, errbound float64

	adx = a[0] - d[0]
	bdx = b[0] - d[0]
	cdx = c[0] - d[0]
	ady = a[1] - d[1]
	bdy = b[1] - d[1]
	cdy = c[1] - d[1]

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

	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	return float64(C.incircleadapt(pa, pb, pc, pd, C.double(permanent)))
}

// InCircleVec is similiar to `InCircle` but takes a point-like struct
// pointer rather than a slice.
func InCircleVec(a, b, c, d *XY) float64 {
	return InCirclePtr(&a.X, &b.X, &c.X, &d.X) // TODO
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
