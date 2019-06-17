package structtag

import "reflect"

const tagName = "mytag"

// Struct is an example struct used to marshal data into
type Struct struct {
	Field1 string `mytag:"field1"`
	Field2 string `mytag:"field2"`
	Field3 string `mytag:"field3"`
	Field4 string `mytag:"field4"`
}

// Marshal inserts data in the struct
func Marshal(dst *Struct, key, val string) *Struct {
	ret := dst
	v := reflect.ValueOf(ret).Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() == reflect.Struct {
			continue // No non string field is supported yet
		}

		sf := v.Type().Field(i)
		if tag, ok := sf.Tag.Lookup(tagName); ok {
			if tag == key && f.CanSet() {
				switch f.Kind() {
				case reflect.String:
					f.SetString(val)
				}

				break
			}
		}
	}

	return ret
}
