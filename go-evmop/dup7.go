package evmop

import (
	"io"
)

// Dup7 represents the "DUP7" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Dup7{}.WriteTo(w)
type Dup7 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Dup7) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeDup7)
}
