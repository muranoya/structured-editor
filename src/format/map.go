package format

import "fmt"

// DataMap represents the map
type DataMap struct {
	value  map[DataString]DataObject
	parent DataObject
}

// NewDataMap creates new instance DataMap
func NewDataMap(parent DataObject) *DataMap {
	return &DataMap{
		value:  map[DataString]DataObject{},
		parent: parent,
	}
}

// Type returns the data type
func (p DataMap) Type() DataType {
	return MAP
}

// GetValue returns object value
func (p *DataMap) GetValue() map[DataString]DataObject {
	return p.value
}

// SetValue sets object value
func (p *DataMap) SetValue(val map[DataString]DataObject) {
	p.value = val
}

// ParentNode returns the parent node
func (p DataMap) ParentNode() DataObject {
	return p.parent
}

func (p DataMap) String() string {
	str := "{"
	for k, v := range p.value {
		str += fmt.Sprintf("%s:%s, ", k, v)
	}
	str += "}"
	return str
}
