package reflection

import (
	"errors"
	"reflect"
	"strings"
)

// ErrNotFunction is returned when the type is not a reflect.Func.
var ErrNotFunction = errors.New("Not A Function Type")

// IsFuncType returns true/false if the interface provided is a func type.
func IsFuncType(elem interface{}) bool {
	_, err := FuncType(elem)
	if err != nil {
		return false
	}
	return true
}

// FuncValue return the Function reflect.Value of the interface provided else
// returns a non-nil error.
func FuncValue(elem interface{}) (reflect.Value, error) {
	tl := reflect.ValueOf(elem)

	if tl.Kind() == reflect.Ptr {
		tl = tl.Elem()
	}

	if tl.Kind() != reflect.Func {
		return tl, ErrNotFunction
	}

	return tl, nil
}

// FuncType return the Function reflect.Type of the interface provided else
// returns a non-nil error.
func FuncType(elem interface{}) (reflect.Type, error) {
	tl := reflect.TypeOf(elem)

	if tl.Kind() == reflect.Ptr {
		tl = tl.Elem()
	}

	if tl.Kind() != reflect.Func {
		return nil, ErrNotFunction
	}

	return tl, nil
}

// HasArgumentSize return true/false to indicate if the function type has the
// size of arguments. It will return false if the interface is not a function
// type.
func HasArgumentSize(elem interface{}, len int) bool {
	tl := reflect.TypeOf(elem)

	if tl.Kind() == reflect.Ptr {
		tl = tl.Elem()
	}

	if tl.Kind() != reflect.Func {
		return false
	}

	if tl.NumIn() != len {
		return false
	}

	return true
}

// GetFuncArgumentsType returns the arguments type of function which should be
// a function type,else returns a non-nil error.
func GetFuncArgumentsType(elem interface{}) ([]reflect.Type, error) {
	tl := reflect.TypeOf(elem)

	if tl.Kind() == reflect.Ptr {
		tl = tl.Elem()
	}

	if tl.Kind() != reflect.Func {
		return nil, ErrNotFunction
	}

	totalFields := tl.NumIn()

	var input []reflect.Type

	for i := 0; i < totalFields; i++ {
		indElem := tl.In(i)

		// if indElem.Kind() == reflect.Ptr {
		// 	indElem = indElem.Elem()
		// }

		input = append(input, indElem)
	}

	return input, nil
}

// MakeValueFor makes a new reflect.Value for the reflect.Type.
func MakeValueFor(t reflect.Type) reflect.Value {
	var input reflect.Value

	mtl := reflect.New(t)

	if mtl.Kind() == reflect.Ptr {
		mtl = mtl.Elem()
	}

	return input
}

// MakeArgumentsValues takes a list of reflect.Types and returns a new version of
// those types, ensuring to dereference if it receives a pointer reflect.Type.
func MakeArgumentsValues(args []reflect.Type) []reflect.Value {
	var inputs []reflect.Value

	for _, tl := range args {
		inputs = append(inputs, MakeValueFor(tl))
	}

	return inputs
}

// InterfaceFromValues returns a list of interfaces representing the concrete
// values within the lists of reflect.Value types.
func InterfaceFromValues(vals []reflect.Value) []interface{} {
	var data []interface{}

	for _, val := range vals {
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		data = append(data, val.Interface())
	}

	return data
}

//==============================================================================

// ErrNotStruct is returned when the reflect type is not a struct.
var ErrNotStruct = errors.New("Not a struct type")

// Field defines a specific tag field with its details from a giving struct.
type Field struct {
	Name  string
	Tag   string
	Type  reflect.Type
	Index int
}

// Fields defines a lists of Field instances.
type Fields []Field

// GetTagFields retrieves all fields of the giving elements with the giving tag
// type.
func GetTagFields(elem interface{}, tag string, allowNaturalNames bool) (Fields, error) {
	if !IsStruct(elem) {
		return nil, ErrNotStruct
	}

	tl := reflect.TypeOf(elem)

	if tl.Kind() == reflect.Ptr {
		tl = tl.Elem()
	}

	var fields Fields

	for i := 0; i < tl.NumField(); i++ {
		field := tl.Field(i)

		// Get the specified tag from this field if it exists.
		tagVal := strings.TrimSpace(field.Tag.Get(tag))

		// If its a - item in the tag then skip or if its an empty string.
		if tagVal == "-" {
			continue
		}

		if !allowNaturalNames && tagVal == "" {
			continue
		}

		if tagVal == "" {
			tagVal = strings.ToLower(field.Name)
		}

		fields = append(fields, Field{
			Name:  field.Name,
			Type:  field.Type,
			Index: i,
			Tag:   tagVal,
		})
	}

	return fields, nil
}

// ToMap returns a map of the giving values from a struct using a provided
// tag to capture the needed values, it extracts those tags values out into
// a map. It returns an error if the element is not a struct.
func ToMap(tag string, elem interface{}, allowNaturalNames bool) (map[string]interface{}, error) {
	// Collect the fields that match the giving tag.
	fields, err := GetTagFields(elem, tag, allowNaturalNames)
	if err != nil {
		return nil, err
	}

	// If there exists no field matching the tag skip.
	if len(fields) == 0 {
		return nil, errors.New("No Tag Matches")
	}

	tl := reflect.ValueOf(elem)

	if tl.Kind() == reflect.Ptr {
		tl = tl.Elem()
	}

	data := make(map[string]interface{})

	// Loop through  the fields and set the appropriate value as needed.
	for _, field := range fields {
		fl := tl.Field(field.Index)
		data[field.Tag] = fl.Interface()
	}

	return data, nil
}

// MergeMap merges the key names of the provided map into the appropriate field
// place where the element has the provided tag.
func MergeMap(tag string, elem interface{}, values map[string]interface{}, allowAll bool) error {

	// Collect the fields that match the giving tag.
	fields, err := GetTagFields(elem, tag, allowAll)
	if err != nil {
		return err
	}

	// If there exists no field matching the tag skip.
	if len(fields) == 0 {
		return nil
	}

	tl := reflect.ValueOf(elem)

	if tl.Kind() == reflect.Ptr {
		tl = tl.Elem()
	}

	// Loop through  the fields and set the appropriate value as needed.
	for _, field := range fields {

		item := values[field.Tag]

		if item == nil {
			continue
		}

		fl := tl.Field(field.Index)

		// If we can't set this field, then skip.
		if !fl.CanSet() {
			continue
		}

		fl.Set(reflect.ValueOf(item))
	}

	return nil
}

// IsStruct returns true/false if the elem provided is a type of struct.
func IsStruct(elem interface{}) bool {
	mc := reflect.TypeOf(elem)

	if mc.Kind() == reflect.Ptr {
		mc = mc.Elem()
	}

	if mc.Kind() != reflect.Struct {
		return false
	}

	return true
}

// MakeNew returns a new version of the giving type, returning a nonpointer type.
// If the type is not a struct then an error is returned.
func MakeNew(elem interface{}) (interface{}, error) {
	mc := reflect.TypeOf(elem)

	if mc.Kind() != reflect.Struct {
		return nil, ErrNotStruct
	}

	return reflect.New(mc).Interface(), nil
}
