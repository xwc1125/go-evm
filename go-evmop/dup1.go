package evmop

import (
	"io"
)

// Dup1 represents the "DUP1" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Dup1{}.WriteTo(w)
type Dup1 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Dup1) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeDup1)
}
