package evmop

import (
	"io"
)

// Dup11 represents the "DUP11" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Dup11{}.WriteTo(w)
type Dup11 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Dup11) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeDup11)
}
