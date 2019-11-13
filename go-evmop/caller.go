package evmop

import (
	"io"
)

// Caller represents the "CALLER" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Caller{}.WriteTo(w)
type Caller struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Caller) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeCaller)
}
