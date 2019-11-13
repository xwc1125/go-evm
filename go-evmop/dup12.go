package evmop

import (
	"io"
)

// Dup12 represents the "DUP12" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Dup12{}.WriteTo(w)
type Dup12 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Dup12) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeDup12)
}
