package evmop

import (
	"io"
)

// Dup15 represents the "DUP15" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Dup15{}.WriteTo(w)
type Dup15 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Dup15) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeDup15)
}
