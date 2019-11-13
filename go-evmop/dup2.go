package evmop

import (
	"io"
)

// Dup2 represents the "DUP2" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Dup2{}.WriteTo(w)
type Dup2 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Dup2) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeDup2)
}
