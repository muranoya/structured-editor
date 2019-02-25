package format

import "fmt"

// DataArray represents the array
type DataArray struct {
	value []DataObject
}

// NewDataArray creates new instance DataArray
func NewDataArray() *DataArray {
	dat := DataArray{}
	dat.value = make([]DataObject, 0, 10)
	return &dat
}

// Type returns the data type
func (p DataArray) Type() DataType {
	return ARRAY
}

// GetValue returns object value
func (p *DataArray) GetValue() *[]DataObject {
	return &p.value
}

// SetValue sets object value
func (p *DataArray) SetValue(val []DataObject) {
	p.value = val
}

func (p DataArray) String() string {
	str := "["
	for _, v := range p.value {
		str += fmt.Sprintf("%s, ", v)
	}
	str += "]"
	return str
}
