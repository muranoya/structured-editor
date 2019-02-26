package tui

import "github.com/muranoya/structured-editor/src/format"

type containerType uint

const (
	// None is not container data type, this type is used in function call
	None containerType = 0
	// String represents the node which has string value
	String containerType = 1 << iota
	// Boolean represents the node which has bool value
	Boolean
	// Integer represents the node which has integer value
	Integer
	// Float represents the node which has float value
	Float
	// Null represents the node which has null value
	Null
	// Array represents the root node of array
	Array
	// ArrayValue represents the node which has array value
	ArrayValue
	// Map represents the root node of map
	Map
	// MapKey represents the node of key
	MapKey
	// MapValue represents the node of value
	MapValue
)

type container struct {
	data    format.DataObject
	conType containerType
}

// newContainer creates new instance of container
func newContainer(data format.DataObject, ct containerType) *container {
	return &container{
		data:    data,
		conType: ct,
	}
}
