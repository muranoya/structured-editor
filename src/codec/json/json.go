package json

import (
	"encoding/json"
	"io"
	"regexp"

	"github.com/muranoya/structured-editor/src/format"
	"github.com/pkg/errors"
)

// CodecJSON represents the codec for json
type CodecJSON struct {
}

var rexp *regexp.Regexp

func init() {
	rexp = regexp.MustCompile("^[0-9-][0-9]*$")
}

// NewCodecJSON create new instance of CodecJSON
func NewCodecJSON() *CodecJSON {
	codecJSON := CodecJSON{}
	return &codecJSON
}

func isIntegerString(s string) bool {
	return rexp.MatchString(s)
}

func decode(obj interface{}) (format.DataObject, error) {
	if obj == nil {
		return format.NewDataNull(), nil
	} else if val, ok := obj.(bool); ok {
		return format.NewDataBoolean(val), nil
	} else if val, ok := obj.(json.Number); ok {
		if isIntegerString(val.String()) {
			i64, err := val.Int64()
			if err != nil {
				return nil, errors.Wrapf(err, "Cannot convert number to int64 from \"%v\"", val.String())
			}
			return format.NewDataInteger(i64), nil
		}

		f64, err := val.Float64()
		if err != nil {
			return nil, errors.Wrapf(err, "Cannot convert number to float64 from \"%v\"", val.String())
		}
		return format.NewDataFloat(f64), nil
	} else if val, ok := obj.(string); ok {
		return format.NewDataString(val), nil
	} else if val, ok := obj.([]interface{}); ok {
		arrayObj := format.NewDataArray()
		for _, v := range val {
			if valObj, err := decode(v); err == nil {
				arrayObj.AppendValue(valObj)
			} else {
				return nil, errors.Cause(err)
			}
		}
		return arrayObj, nil
	} else if val, ok := obj.(map[string]interface{}); ok {
		mapObj := format.NewDataMap()
		for k, v := range val {
			if valObj, err := decode(v); err == nil {
				mapObj.GetValue()[*format.NewDataString(k)] = valObj
			} else {
				return nil, errors.Cause(err)
			}
		}
		return mapObj, nil
	}

	return nil, errors.New("Unknown type")
}

// Decode a text to a DataObject
func (p *CodecJSON) Decode(reader io.Reader) (*format.DataRoot, error) {
	var obj interface{}
	decoder := json.NewDecoder(reader)
	decoder.UseNumber()
	if err := decoder.Decode(&obj); err != nil {
		return nil, errors.Cause(err)
	}

	data, err := decode(obj)
	if err != nil {
		return nil, err
	}
	return format.NewDataRoot(data), nil
}

// Encode a DataObject to a text
func (p *CodecJSON) Encode(obj format.DataObject) (string, error) {
	return "", nil
}

// SupportedTypes returns supported data types
func (p *CodecJSON) SupportedTypes() uint {
	return uint(format.ARRAY | format.BOOLEAN | format.FLOAT | format.INTEGER | format.MAP | format.NULL | format.STRING)
}
