package sample

import (
	. "github.com/gabetucker2/govariadics"
)

/**
 The following is the recommended template for how you should structure functions using this library:
*/
func sampleFunction(variadic ...any) {

	// unpack variadic into optional parameters
	var name, age, height any
	UnpackVariadic(variadic, &name, &age, &height)

	// sample body stuff below:
	if name == "Gerry" {
		// do stuff
	}
	if age == 34 {
		// do stuff
	}
	if height == 2.6 {
		// do stuff
	}

}

func main() {
	sampleFunction("Gerry", 20, 5.9)
}
