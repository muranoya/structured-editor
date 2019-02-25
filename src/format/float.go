package format

import "fmt"

// DataFloat represents the float
type DataFloat struct {
	value float64
}

// NewDataFloat creates new instance DataFloat
func NewDataFloat(val float64) *DataFloat {
	dat := DataFloat{}
	dat.value = val
	return &dat
}

// Type returns the data type
func (p DataFloat) Type() DataType {
	return FLOAT
}

// GetValue returns object value
func (p *DataFloat) GetValue() float64 {
	return p.value
}

// SetValue sets object value
func (p *DataFloat) SetValue(val float64) {
	p.value = val
}

func (p DataFloat) String() string {
	return fmt.Sprintf("%f", p.value)
}
