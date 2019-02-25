package format

import "fmt"

// DataBoolean represents the boolean
type DataBoolean struct {
	value bool
}

// NewDataBoolean creates new instance DataBoolean
func NewDataBoolean(val bool) *DataBoolean {
	dat := DataBoolean{}
	dat.value = val
	return &dat
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

func (p DataBoolean) String() string {
	return fmt.Sprintf("%t", p.value)
}
