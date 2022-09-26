package gogenerics

import (
	"reflect"
)

/** Removes the "not used" compiler error from variables and functions
	  OR
	Removes the library imported but not used error by inputting a function in that library to this function

 @param `variadic` type{...[]any}
*/
 func RemoveUnusedError(variadic ...any) { }

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

/** Returns, from any array type, a version of that array which is converted to type []any

 @param `arr` type{any}
 @return type{[]any}
 @requires `arr` is an array
 */
 func UnpackArray(arr any) []any {
    valType := reflect.ValueOf(arr)
    new := make([]any, valType.Len())
    for i := 0; i < valType.Len(); i++ {
        new[i] = valType.Index(i).Interface()
    }
    return new
}

/** Returns, from any map type, a version of that map which is converted to type map[any]any

 @param `arr` type{any}
 @return type{[]any}
 @requires `arr` is an array
 */
 func UnpackMap(s any) map[any]any {
	v := reflect.ValueOf(s)
    m := make(map[any]any, v.Len())
    for _, k := range v.MapKeys() {
		m[k.Interface()] = v.MapIndex(k).Interface()
    }
    return m
}

/** Returns a clone of this interface
 NOTE: Pass memory address

 @param `toClone` type{any}
 @returns type{any}
 */
 func CloneInterface(toClone any) any {
	return reflect.New(reflect.ValueOf(toClone).Elem().Type()).Interface()
}

/** Returns a version of the object as an interface

 @param `obj` type{any}
 @return type{any} interface version of `obj`*/
 func MakeInterface(obj any) any {
	return reflect.ValueOf(obj).Interface()
 } 

/** Returns out1 if test is true; else return out2
 
 @param `test` type{bool}
 @param `out1` type{any}
 @param `out2` type{any}
 @returns any `out1` or `out2`
 @requires neither param yields a syntax error
 */
 func IfElse(test bool, out1, out2 any) any {
	if test { return out1 } else { return out2 }
}

/** Returns whether an argument is the address of another object

 @param `arg` type{any}
 @return type{bool}
 */
 func IsPointer(arg any) bool {
	return reflect.ValueOf(arg).Kind() == reflect.Ptr
}

/** If an argument is an address, then return the value stored at that address; else return nil

 @param `ptr` type{any}
 @param optional `returnIfNotPointer` type{bool} default true
 @return type{any}
 */
 func GetPointer(ptr any, variadic ...any) any {

	// unpack variadic into optional parameters
	var returnIfNotPointer any
	UnpackVariadic(variadic, &returnIfNotPointer)
	// set default
	if returnIfNotPointer == nil {
		returnIfNotPointer = false
	}

	if IsPointer(ptr) {
		return *ptr.(*any)
	} else {
		if returnIfNotPointer.(bool) {
			return ptr
		} else {
			return nil
		}
	}
}

/** If an argument is an address, then update the value stored at that address to `to` and return the value; else return nil

 @param `ptr` type{any}
	Must be an interface type upon being passed into the function
 @param `to` type{any}
 @return type{bool}
 */
 func SetPointer(arg, to any) any {
	if IsPointer(arg) {
		reflect.ValueOf(arg).Elem().Set(reflect.ValueOf(&to).Elem())
		return arg
	} else {
		return nil
	}
}
