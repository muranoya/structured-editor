package format

// DataType represents the type of data object
type DataType uint

const (
	// ROOT is a root object
	ROOT DataType = iota
	// MAP is a map object
	MAP
	// ARRAY is a array object
	ARRAY
	// STRING is a string object
	STRING
	// INTEGER is a 64bit signed integer object
	INTEGER
	// BOOLEAN is a boolean object
	BOOLEAN
	// FLOAT is a 64bit float object
	FLOAT
	// NULL is a null object
	NULL
)

// DataObject represents the structured-data object
type DataObject interface {
	Type() DataType
	ParentNode() DataObject
}
