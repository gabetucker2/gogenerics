package sample

import (
	"github.com/gabetucker2/govariadics"
)

/**
 The following is the recommended template for how you should structure functions using this library:
*/
func myFunction(variadic ...any) {

	// unpack variadic into optional parameters
	var name, age, height any
	govariadics.UnpackVariadic(variadic, &name, &age, &height)

	// sample body stuff below:
	if name == nil {
		// no name was entered
	}
	if age == 34 {
		// age 34 was entered
	}
	if height == nil {
		// no height was entered
	}

}

func sample() {

	// call A
	myFunction("Gerry", 20, 5.9)

	// call B (all are the same)
	myFunction("Gerry", 20)
	myFunction("Gerry", 20, nil)

	// call C (all are the same)
	myFunction(nil, 20)
	myFunction(nil, 20, nil)

	// call D (all are the same)
	myFunction()
	myFunction(nil)
	myFunction(nil, nil)
	myFunction(nil, nil, nil)

}
