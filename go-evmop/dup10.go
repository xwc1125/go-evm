package evmop

import (
	"io"
)

// Dup10 represents the "DUP10" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Dup10{}.WriteTo(w)
type Dup10 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Dup10) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeDup10)
}
