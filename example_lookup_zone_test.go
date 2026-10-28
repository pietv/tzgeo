package tzgeo_test

import (
	"fmt"

	"github.com/pietv/tzgeo"
)

func ExampleLookupZone() {
	z, ok := tzgeo.LookupZone("Europe/Berlin")
	if !ok {
		return
	}
	fmt.Println(z.Name)
	fmt.Println(z.Countries)
	// Output:
	// Europe/Berlin
	// [DE DK NO SE SJ]
}
