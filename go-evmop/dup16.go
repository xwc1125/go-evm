package evmop

import (
	"io"
)

// Dup16 represents the "DUP16" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Dup16{}.WriteTo(w)
type Dup16 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Dup16) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeDup16)
}
