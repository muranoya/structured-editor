package format

import "fmt"

// DataString represents the string
type DataString struct {
	value string
}

// NewDataString creates new instance DataString
func NewDataString(val string) *DataString {
	dat := DataString{}
	dat.value = val
	return &dat
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

func (p DataString) String() string {
	return fmt.Sprintf("\"%s\"", p.value)
}
