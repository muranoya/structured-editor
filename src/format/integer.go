package format

import "fmt"

// DataInteger represents the 64bit signed integer
type DataInteger struct {
	value  int64
	parent DataObject
}

// NewDataInteger creates new instance DataInteger
func NewDataInteger(val int64, parent DataObject) *DataInteger {
	return &DataInteger{
		value:  val,
		parent: parent,
	}
}

// Type returns the data type
func (p DataInteger) Type() DataType {
	return INTEGER
}

// GetValue returns object value
func (p *DataInteger) GetValue() int64 {
	return p.value
}

// SetValue sets object value
func (p *DataInteger) SetValue(val int64) {
	p.value = val
}

// ParentNode returns the parent node
func (p DataInteger) ParentNode() DataObject {
	return p.parent
}

func (p DataInteger) String() string {
	return fmt.Sprintf("%d", p.value)
}
