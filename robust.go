// Package robust provides Go bindings to the `predicates.c`
// library from Jonathan Shewchuk
package robust

// extern double epsilon;
// void exactinit();
import "C"
import "fmt"

func init() {
	C.exactinit()
	fmt.Println(C.epsilon) // TODO: quick hack for testing
}
