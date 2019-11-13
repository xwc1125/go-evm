package evmop

import (
	"io"
)

// Dup5 represents the "DUP5" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Dup5{}.WriteTo(w)
type Dup5 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Dup5) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeDup5)
}
