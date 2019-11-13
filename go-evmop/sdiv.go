package evmop

import (
	"io"
)

// SDiv represents the "SDIV" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.SDiv{}.WriteTo(w)
type SDiv struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (SDiv) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSDiv)
}
