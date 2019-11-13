package evmop

import (
	"io"
)

// IsZero represents the "ISZERO" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.IsZero{}.WriteTo(w)
type IsZero struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (IsZero) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeIsZero)
}
