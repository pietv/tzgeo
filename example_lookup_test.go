package tzgeo_test

import (
	"fmt"

	"github.com/pietv/tzgeo"
)

func ExampleLookup() {
	lat, lon, ok := tzgeo.Lookup("Europe/Berlin")
	fmt.Printf("%.4f %.4f %v\n", lat, lon, ok)

	_, _, ok = tzgeo.Lookup("Mars/Olympus")
	fmt.Println(ok)
	// Output:
	// 52.5000 13.3667 true
	// false
}
