package evmop

import (
	"io"
)

// Dup14 represents the "DUP14" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Dup14{}.WriteTo(w)
type Dup14 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Dup14) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeDup14)
}
