package evmop

import (
	"io"
)

// Dup6 represents the "DUP6" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Dup6{}.WriteTo(w)
type Dup6 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Dup6) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeDup6)
}
