package evmop

import (
	"io"
)

// Dup9 represents the "DUP9" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Dup9{}.WriteTo(w)
type Dup9 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Dup9) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeDup9)
}