package format

// DataNull represents the null
type DataNull struct {
}

// NewDataNull creates new instance DataNull
func NewDataNull() *DataNull {
	dat := DataNull{}
	return &dat
}

// Type returns the data type
func (p DataNull) Type() DataType {
	return NULL
}

func (p DataNull) String() string {
	return "null"
}
