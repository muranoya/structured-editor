package format

import "fmt"

// DataMap represents the map
type DataMap struct {
	value map[DataString]DataObject
}

// NewDataMap creates new instance DataMap
func NewDataMap() *DataMap {
	dat := DataMap{}
	dat.value = map[DataString]DataObject{}
	return &dat
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

func (p DataMap) String() string {
	str := "{"
	for k, v := range p.value {
		str += fmt.Sprintf("%s:%s, ", k, v)
	}
	str += "}"
	return str
}
