package codec

import (
	"io"

	"github.com/muranoya/structured-editor/src/format"
)

// Codec represents the codec
type Codec interface {
	Decode(io.Reader) (*format.DataRoot, error)
	Encode(format.DataObject) (string, error)
	SupportedTypes() uint
}
