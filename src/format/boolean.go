package format

import "fmt"

// DataBoolean represents the boolean
type DataBoolean struct {
	value  bool
	parent DataObject
}

// NewDataBoolean creates new instance DataBoolean
func NewDataBoolean(val bool, parent DataObject) *DataBoolean {
	return &DataBoolean{
		value:  val,
		parent: parent,
	}
}

// Type returns the data type
func (p DataBoolean) Type() DataType {
	return BOOLEAN
}

// GetValue returns object value
func (p *DataBoolean) GetValue() bool {
	return p.value
}

// SetValue sets object value
func (p *DataBoolean) SetValue(val bool) {
	p.value = val
}

// ParentNode returns the parent node
func (p DataBoolean) ParentNode() DataObject {
	return p.parent
}

func (p DataBoolean) String() string {
	return fmt.Sprintf("%t", p.value)
}
