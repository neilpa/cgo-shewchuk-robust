# cgo-shewchuk-robust

[![CI][ci-badge]][ci] [![GoDoc][docs-badge]][docs]

Go bindings for the [robust geometric predicates][robust] from [Jonathan Shewchuk][shewchuk].

## Usage

There are four base predicates, each taking slice arguments.

* [`Orient2`][docs-orient2] - orientation of a point relative to a directed line
* [`Orient3`][docs-orient3] - orientation of a point relative to a directed plane
* [`InCircle`][docs-incircle] - containment of a point in a directed circle
* [`InSphere`][docs-insphere] - containment of a point in a directed sphere

Each predicate has two other flavors taking struct (`*Vec`) and C-array style pointers (`*Ptr`). See the [docs][] for more details.

## Tests

Large set of test cases from [here][tests] with props to [mourner/robust-predicates][tests-mourner] for the pointer.

## Licence

Like the original [`predicates.c`][predicates.c], this is released into the public domain.


[ci]: https://github.com/neilpa/cgo-shewchuk-robust/actions
[ci-badge]: https://github.com/neilpa/cgo-shewchuk-robust/workflows/Test/badge.svg
[docs]: https://godoc.org/neilpa.me/cgo-shewchuk-robust#section-documentation
[docs-badge]: https://godoc.org/neilpa.me/cgo-shewchuk-robust?status.svg
[docs-incircle]: https://pkg.go.dev/neilpa.me/cgo-shewchuk-robust#InCircle
[docs-insphere]: https://pkg.go.dev/neilpa.me/cgo-shewchuk-robust#InSphere
[docs-orient2]: https://pkg.go.dev/neilpa.me/cgo-shewchuk-robust#Orient2
[docs-orient3]: https://pkg.go.dev/neilpa.me/cgo-shewchuk-robust#Orient3
[predicates.c]: http://www.cs.cmu.edu/afs/cs/project/quake/public/code/predicates.c
[robust]: https://www.cs.cmu.edu/~quake/robust.html
[shewchuk]: https://people.eecs.berkeley.edu/~jrs/
[tests]: https://www.cs.cmu.edu/afs/cs/project/pscico/pscico/src/arithmetic/compiler1/test/
[tests-mourner]: https://github.com/mourner/robust-predicates/tree/master/test/fixtures