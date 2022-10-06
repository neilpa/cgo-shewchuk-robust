// Package robust provides Go bindings to the `predicates.c` library from
// Jonathan Shewchuk. Only the primary adaptive functions are exported from
// this package.
//
// <https://www.cs.cmu.edu/~quake/robust.html>
//
// Each of the core functions takes `[]float64` arguments that should be
// at least 2 or 3 values in length for the respective dimensionality. This
// makes it easy to work with large flat buffers of many points that can be
// sliced efficiently.
//
// There are also `*Vec` suffixed functions that take point-like values of
// the shape `*struct{ X,Y float64 }` or `*struct{ X,Y,Z float64 }`. Use
// requires casting to the target type of this package, e.g.
//
//	var a, b, c MyVec2
//	res := robust.Orient2Vec((*robust.XY)(&a), (*robust.XY)(&b), (*robust.XY)(&c))
//
// Both the slice and struct variants do initial error bounds check in go
// which avoids uncessary CGO calls in the simple cases. Only if those
// fail are the corresponding `*adapt` C functions called. This provides
// the most notable performance impact in the `Orient*` methods.
//
// Finally, there are `*Ptr` suffixed functions that take C-like arrays of
// at least 2 or 3 `*float64` values. These directly wrap the equivalent
// C functions and don't do any error bounds checking in go. As such, they
// always incur the CGO overhead but allow directly passing pointers, e.g.
//
//	res := robust.Orient2Ptr(&p0.x, &p1.x, &p2.x)
package robust

// void exactinit();
// extern double ccwerrboundA, o3derrboundA, iccerrboundA, isperrboundA;
import "C"

// Cache values from CGO init
var (
	ccwerrboundA, o3derrboundA, iccerrboundA, isperrboundA float64
)

// XY is a "template" for 2D vector types. It's not intended for use
// directly but as a pointer cast target. See Orient2Vec and InCircleVec.
type XY struct {
	X, Y float64
}

// XYZ is a "template" for 3D vector types. It's not intended for use
// directly but as a pointer cast target. See Orient3Vec and InSphereVec.
type XYZ struct {
	X, Y, Z float64
}

func init() {
	C.exactinit()
	ccwerrboundA = float64(C.ccwerrboundA)
	o3derrboundA = float64(C.o3derrboundA)
	iccerrboundA = float64(C.iccerrboundA)
	isperrboundA = float64(C.isperrboundA)
}
