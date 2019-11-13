package evmop

import (
	"io"
)

// LT represents the "LT" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.LT{}.WriteTo(w)
type LT struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (LT) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeLT)
}
