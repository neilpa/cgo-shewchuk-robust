package robust

// double orient2d(double *pa, double *pb, double *pc);
// double orient2dadapt(double *pa, double *pb, double *pc, double detsum);
import "C"

// TODO: Updated API Surface
//
// Orient3
// Orient3Vec
// Orient3Ptr
// InCircle
// InCircleVec
// InCirclePtr
// InSphere
// InCircleVec
// InCirclePtr
//

// Orient2 returns a positive value if the points a, b, and c occur in
// counterclockwise order; a negative value if they occur in clockwise
// order; and zero if they are collinear. The result is also a rough
// approximation of twice the signed area of the triangle defined by the
// three points.
//
// Each slice parameter must contain at least 2 values.
func Orient2(a, b, c []float64) float64 {
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

// Orient2Vec is similiar to `Orient2` but takes a point-like struct
// pointer rather than a slice.
func Orient2Vec(a, b, c *XY) float64 {
	return Orient2Ptr(&a.X, &b.X, &c.X) // TODO: go error bounds check
}

// Orient2Ptr is the direct wrapper of `orient2d` from `predicates.c`.
// See `Orient2` for additional details.
func Orient2Ptr(a, b, c *float64) float64 {
	pa := (*C.double)(a)
	pb := (*C.double)(b)
	pc := (*C.double)(c)
	return float64(C.orient2d(pa, pb, pc))
}
