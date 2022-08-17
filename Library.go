package govariadics

/** Sets a set of variables to the variable set passed into a variadic parameter

 @param `variadic` type{...[]any}
 @param `var1, var2, ..., varN` type{any}
 @updates `var1, var2, ..., varN` are set to each of the values in the variadic array, or nil if undefined, respectively
 */
 func UnpackVariadic(variadic []any, into ...*any) {
	vLen := len(variadic)
	for i, v := range into {
		if i < vLen {
			*v = variadic[i]
		} else {
			*v = nil
		}
	}
}
