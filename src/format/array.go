package format

import (
	"fmt"

	"github.com/pkg/errors"
)

// DataArray represents the array
type DataArray struct {
	value  []DataObject
	parent DataObject
}

// NewDataArray creates new instance DataArray
func NewDataArray(parent DataObject) *DataArray {
	return &DataArray{
		value:  make([]DataObject, 0, 10),
		parent: parent,
	}
}

// Type returns the data type
func (p DataArray) Type() DataType {
	return ARRAY
}

// GetValue returns object value
func (p DataArray) GetValue() []DataObject {
	return p.value
}

// SetValue sets object value
func (p *DataArray) SetValue(val []DataObject) {
	p.value = val
}

// ParentNode returns the parent node
// The parent of array value node is the array node.
func (p DataArray) ParentNode() DataObject {
	return p.parent
}

// RemoveValue removes data object has the index of idx
func (p *DataArray) RemoveValue(idx int) {
	p.SetValue(append(p.value[:idx], p.value[idx+1:]...))
}

// FindValue finds data object, and returns index
func (p *DataArray) FindValue(val DataObject) (int, error) {
	for i, v := range p.value {
		if v == val {
			return i, nil
		}
	}
	return -1, errors.New("DataObject is not found")
}

// AppendValue appends data object
func (p *DataArray) AppendValue(obj DataObject) {
	p.value = append(p.value, obj)
}

func (p DataArray) String() string {
	str := "["
	for _, v := range p.value {
		str += fmt.Sprintf("%s, ", v)
	}
	str += "]"
	return str
}
