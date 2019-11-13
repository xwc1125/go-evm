package evmop

import (
	"io"
)

// Not represents the "NOT" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Not{}.WriteTo(w)
type Not struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Not) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeNot)
}
