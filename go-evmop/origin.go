package evmop

import (
	"io"
)

// Origin represents the "ORIGIN" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Origin{}.WriteTo(w)
type Origin struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Origin) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeOrigin)
}
