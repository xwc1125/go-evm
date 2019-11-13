package evmop

import (
	"io"
)

// Eq represents the "EQ" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Eq{}.WriteTo(w)
type Eq struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Eq) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeEq)
}
