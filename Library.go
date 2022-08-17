package gogenerics

import (
	"reflect"
)

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

 @param `toClone` type{any}
 @returns type{any}
*/
func CloneInterface(toClone any) any {
	return reflect.New(reflect.ValueOf(toClone).Elem().Type()).Interface()
}
