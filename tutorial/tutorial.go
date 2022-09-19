package tutorial

import (
	"github.com/gabetucker2/gogenerics"
)

//lint:ignore U1000 Ignore unused function warning
func tutorial() {


	// UnpackVariadics tutorial: -----------------------------------------------------------------------------------------

	/**
	 Assume you declared a function/variable and that you are not using them in your script, i.e., for debugging purposes:*/

	myVar := "Hey"
	myFunc := func() {}

	gogenerics.RemoveUnusedError(myVar, myFunc) // notice we don't do myFunc(), as that would call the function

	/**
	 If you had an unused library import, you could also just call RemoveUnusedError(myFuncInTheLibrary) to remove that library
	 	import error.*/
	
	/**
	 Say you have several parameters that you will only invoke in some function calls and not others.  This method makes
	 	it easy and concise to create "optional" parameters from a variadic in golang.  It is HIGHLY encouraged to
		document your parameter guidelines, as any parameter type can be inputted to the `...any` variadic.  For
		good practice, you should also ensure the parameter types match your expectations as not to raise any errors.
	*/

	myFunction := func(variadic ...any) {

		// unpack variadic into optional parameters (this line, and the following two lines, are essential to make UnpackVariadic work)
		var name, age, height any
		gogenerics.UnpackVariadic(variadic, &name, &age, &height)

		// sample body stuff below:
		//lint:ignore SA9003 Ignore empty if statement warning
		if name == nil {
			// no name was entered
		}
		//lint:ignore SA9003 Ignore empty if statement warning
		if age == 34 {
			// age 34 was entered
		}
		//lint:ignore SA9003 Ignore empty if statement warning
		if height == nil {
			// no height was entered
		}

		// If you need to convert height to a float, rather than an interface, do height.(float).
		
	}

	/**
	 You now have a go equivalent of "optional" parameters in the function.  In order to call 
	*/

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


	// UnpackArray tutorial: ---------------------------------------------------------------------------------------------

	/**
	 If you have an array of a specific value type which you would like to convert to an array of interfaces, it can be
	 	irritating having to case-by-case convert the array to an interface array.  For instance, if you have an array
		of arrays of different type ([]any {[]string, []int, []float} ) or you have a function parameter accepting
		an array whose contents you would like to convert into interfaces, this conversion process can be redundant and
		irritating.  Our UnpackArray function automates this process.*/

	myStringArray := []string {"Hey", "Hello", "Hi"}
	myIntArray := []int {4, 2}

	myArraysAsOriginalTypes := []any {myStringArray, myIntArray}
	myArraysAsGenericTypes := []any {}

	for _, thisArray := range myArraysAsOriginalTypes {
		myArraysAsGenericTypes = append(myArraysAsGenericTypes, gogenerics.UnpackArray(thisArray))//lint:ignore SA4010 Ignore unused warning
	}

	// myArraysAsGenericTypes := []any {[]any, []any}


	// UnpackMap tutorial: -----------------------------------------------------------------------------------------------

	/**
	 Say you have a map of a specific type which you would like to convert into a map[any]any map.  This could be
	 	because you have a map parameter that takes any map type or because you want to access the interface
		properties.  UnpackMap automates this process:*/

	mapOriginalType := map[string]int {"Bob" : 5, "Joe" : 2}
	//mapGenericType :=
	gogenerics.UnpackMap(mapOriginalType)


	// CloneInterface tutorial: -----------------------------------------------------------------------------------------------

	/**
	 Say you have an interface, and you would like to make a copy of this interface.  This process is very unintuitive in go,
	 	especially since you cannot clone the object on pass like for mutable types (e.g., strings or ints).  So
		CloneInterface automates this process:*/

	var myInterface any
	//myInterface2 := 
	gogenerics.CloneInterface(myInterface)

	// IfElse tutorial: -------------------------------------------------------------------------------------------------------

	/**
	 If something is true, you want to set a variable to one value.  Otherwise, you want to set it to another value.  IfElse
	 	simply expedites this process:*/
		
	//lemonadeBalance :=
	gogenerics.IfElse(3 > 5, 3, -5) // returns -5 since condition is false

	// Pointer tutorial: ------------------------------------------------------------------------------------------------------
	myStr := "Hey"
	myStrAdr := &myStr

	gogenerics.IsPointer("Hi") // false
	gogenerics.IsPointer(myStr) // false
	gogenerics.IsPointer(&myStr) // true
	gogenerics.IsPointer(myStrAdr) // true

	gogenerics.GetPointer(myStr) // nil
	gogenerics.GetPointer(&myStr) // "Hey"
	gogenerics.GetPointer(myStrAdr) // "Hey"

	gogenerics.SetPointer(myStr, "Hi") // nil (and doesn't update)
	gogenerics.SetPointer(&myStr, "Hi") // "Hi"
	gogenerics.SetPointer(myStrAdr, "Hi") // "Hi"

}
