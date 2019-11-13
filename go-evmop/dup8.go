package evmop

import (
	"io"
)

// Dup8 represents the "DUP8" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Dup8{}.WriteTo(w)
type Dup8 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Dup8) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeDup8)
}
