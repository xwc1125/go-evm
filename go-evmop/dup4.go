package evmop

import (
	"io"
)

// Dup4 represents the "DUP4" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Dup4{}.WriteTo(w)
type Dup4 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Dup4) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeDup4)
}
