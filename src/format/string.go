package format

import "fmt"

// DataString represents the string
type DataString struct {
	value  string
	parent DataObject
}

// NewDataString creates new instance DataString
func NewDataString(val string, parent DataObject) *DataString {
	return &DataString{
		value:  val,
		parent: parent,
	}
}

// Type returns the data type
func (p DataString) Type() DataType {
	return STRING
}

// GetValue returns object value
func (p *DataString) GetValue() string {
	return p.value
}

// SetValue sets object value
func (p *DataString) SetValue(val string) {
	p.value = val
}

// ParentNode returns the parent node
func (p DataString) ParentNode() DataObject {
	return p.parent
}

func (p DataString) String() string {
	return fmt.Sprintf("\"%s\"", p.value)
}
