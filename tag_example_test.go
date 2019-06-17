package structtag

import (
	"fmt"
)

func ExampleMarshal() {

	source := map[string]string{
		"field1": "foo",
		"field3": "baz",
	}

	result := new(Struct)
	for k, v := range source {
		result = Marshal(result, k, v)
	}

	fmt.Printf("result: %#v\n", result)
	// Output: result: &structtag.Struct{Field1:"foo", Field2:"", Field3:"baz", Field4:""}
}
