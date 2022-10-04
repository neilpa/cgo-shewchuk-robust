package robust

// double insphere(double *pa, double *pb, double *pc, double *pd, double *pe);
// double insphereadapt(double *pa, double *pb, double *pc, double *pd, double *pe, double permanent);
import "C"
import "math"

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
	var aex, bex, cex, dex float64
	var aey, bey, cey, dey float64
	var aez, bez, cez, dez float64
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

	aex = a[0] - e[0]
	bex = b[0] - e[0]
	cex = c[0] - e[0]
	dex = d[0] - e[0]
	aey = a[1] - e[1]
	bey = b[1] - e[1]
	cey = c[1] - e[1]
	dey = d[1] - e[1]
	aez = a[2] - e[2]
	bez = b[2] - e[2]
	cez = c[2] - e[2]
	dez = d[2] - e[2]

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

	pa := (*C.double)(&a[0])
	pb := (*C.double)(&b[0])
	pc := (*C.double)(&c[0])
	pd := (*C.double)(&d[0])
	pe := (*C.double)(&e[0])
	return float64(C.insphereadapt(pa, pb, pc, pd, pe, C.double(permanent)))
}
