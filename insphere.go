package robust

// double insphere(double *pa, double *pb, double *pc, double *pd, double *pe);
// double insphereadapt(double *pa, double *pb, double *pc, double *pd, double *pe, double permanent);
import "C"
import "math"

// InSphere returns a positive value if the point e lies inside the
// sphere passing through a, b, c, and d; a negative value if it lies
// outside; and zero if the five points are cospherical. The points a,
// b, c, and d must be ordered so that they have a positive orientation
// (as defined by `Orient3`), or the sign of the result will be reversed.
//
// Each slice parameter must contain at least 3 values.
func InSphere(a, b, c, d, e []float64) float64 {
	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	pe := (*C.double)(&e[0])

	return inSphere(pa, pb, pc, pd, pe,
		a[0]-e[0], b[0]-e[0], c[0]-e[0], d[0]-e[0],
		a[1]-e[1], b[1]-e[1], c[1]-e[1], d[1]-e[1],
		a[2]-e[2], b[2]-e[2], c[2]-e[2], d[2]-e[2],
	)
}

// InSphereVec is similiar to `InSphere` but takes a point-like struct
// pointer rather than a slice.
func InSphereVec(a, b, c, d, e *XYZ) float64 {
	pa := (*C.double)(&a.X)
	pb := (*C.double)(&b.X)
	pc := (*C.double)(&c.X)
	pd := (*C.double)(&d.X)
	pe := (*C.double)(&e.X)

	return inSphere(pa, pb, pc, pd, pe,
		a.X-e.X, b.X-e.X, c.X-e.X, d.X-e.X,
		a.Y-e.Y, b.Y-e.Y, c.Y-e.Y, d.Y-e.Y,
		a.Z-e.Z, b.Z-e.Z, c.Z-e.Z, d.Z-e.Z,
	)
}

// InSpherePtr is the direct wrapper of `insphere` from `predicates.c`.
// See `InSphere` for additional details.
func InSpherePtr(a, b, c, d, e *float64) float64 {
	pa := (*C.double)(a)
	pb := (*C.double)(b)
	pc := (*C.double)(c)
	pd := (*C.double)(d)
	pe := (*C.double)(e)
	return float64(C.insphere(pa, pb, pc, pd, pe))
}

// inCircle implements the basic error bound checks to minimize
// CGO calls to the adaptive implementation.
func inSphere(pa, pb, pc, pd, pe *C.double,
	aex, bex, cex, dex float64,
	aey, bey, cey, dey float64,
	aez, bez, cez, dez float64,
) float64 {

	var aexbey, bexaey, bexcey, cexbey, cexdey, dexcey, dexaey, aexdey float64
	var aexcey, cexaey, bexdey, dexbey float64
	var alift, blift, clift, dlift float64
	var ab, bc, cd, da, ac, bd float64
	var abc, bcd, cda, dab float64
	var aezplus, bezplus, cezplus, dezplus float64
	var aexbeyplus, bexaeyplus, bexceyplus, cexbeyplus float64
	var cexdeyplus, dexceyplus, dexaeyplus, aexdeyplus float64
	var aexceyplus, cexaeyplus, bexdeyplus, dexbeyplus float64
	var det float64
	var permanent, errbound float64

	aexbey = aex * bey
	bexaey = bex * aey
	ab = aexbey - bexaey
	bexcey = bex * cey
	cexbey = cex * bey
	bc = bexcey - cexbey
	cexdey = cex * dey
	dexcey = dex * cey
	cd = cexdey - dexcey
	dexaey = dex * aey
	aexdey = aex * dey
	da = dexaey - aexdey

	aexcey = aex * cey
	cexaey = cex * aey
	ac = aexcey - cexaey
	bexdey = bex * dey
	dexbey = dex * bey
	bd = bexdey - dexbey

	abc = aez*bc - bez*ac + cez*ab
	bcd = bez*cd - cez*bd + dez*bc
	cda = cez*da + dez*ac + aez*cd
	dab = dez*ab + aez*bd + bez*da

	alift = aex*aex + aey*aey + aez*aez
	blift = bex*bex + bey*bey + bez*bez
	clift = cex*cex + cey*cey + cez*cez
	dlift = dex*dex + dey*dey + dez*dez

	det = (dlift*abc - clift*dab) + (blift*cda - alift*bcd)

	aezplus = math.Abs(aez)
	bezplus = math.Abs(bez)
	cezplus = math.Abs(cez)
	dezplus = math.Abs(dez)
	aexbeyplus = math.Abs(aexbey)
	bexaeyplus = math.Abs(bexaey)
	bexceyplus = math.Abs(bexcey)
	cexbeyplus = math.Abs(cexbey)
	cexdeyplus = math.Abs(cexdey)
	dexceyplus = math.Abs(dexcey)
	dexaeyplus = math.Abs(dexaey)
	aexdeyplus = math.Abs(aexdey)
	aexceyplus = math.Abs(aexcey)
	cexaeyplus = math.Abs(cexaey)
	bexdeyplus = math.Abs(bexdey)
	dexbeyplus = math.Abs(dexbey)
	permanent =
		((cexdeyplus+dexceyplus)*bezplus+
			(dexbeyplus+bexdeyplus)*cezplus+
			(bexceyplus+cexbeyplus)*dezplus)*
			alift +
			((dexaeyplus+aexdeyplus)*cezplus+
				(aexceyplus+cexaeyplus)*dezplus+
				(cexdeyplus+dexceyplus)*aezplus)*
				blift +
			((aexbeyplus+bexaeyplus)*dezplus+
				(bexdeyplus+dexbeyplus)*aezplus+
				(dexaeyplus+aexdeyplus)*bezplus)*
				clift +
			((bexceyplus+cexbeyplus)*aezplus+
				(cexaeyplus+aexceyplus)*bezplus+
				(aexbeyplus+bexaeyplus)*cezplus)*
				dlift
	errbound = isperrboundA * permanent
	if (det > errbound) || (-det > errbound) {
		return det
	}

	return float64(C.insphereadapt(pa, pb, pc, pd, pe, C.double(permanent)))
}
