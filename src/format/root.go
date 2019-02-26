package format

import "fmt"

// DataRoot represents the root special object
type DataRoot struct {
	data DataObject
}

// NewDataRoot creates new instance DataRoot
func NewDataRoot(d DataObject) *DataRoot {
	dat := DataRoot{
		data: d,
	}
	return &dat
}

// Type returns the data type
func (p DataRoot) Type() DataType {
	return ROOT
}

// GetValue returns object value
func (p *DataRoot) GetValue() DataObject {
	return p.data
}

// SetValue sets object value
func (p *DataRoot) SetValue(d DataObject) {
	p.data = d
}

func (p DataRoot) String() string {
	return fmt.Sprint(p.data)
}
