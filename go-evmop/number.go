package evmop

import (
	"io"
)

// Number represents the "NUMBER" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Number{}.WriteTo(w)
type Number struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Number) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeNumber)
}
