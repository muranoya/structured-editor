package format

// DataNull represents the null
type DataNull struct {
	parent DataObject
}

// NewDataNull creates new instance DataNull
func NewDataNull(parent DataObject) *DataNull {
	return &DataNull{
		parent: parent,
	}
}

// Type returns the data type
func (p DataNull) Type() DataType {
	return NULL
}

// ParentNode returns the parent node
func (p DataNull) ParentNode() DataObject {
	return p.parent
}

func (p DataNull) String() string {
	return "null"
}
