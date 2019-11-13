package evmop

import (
	"io"
)

// Add represents the "ADD" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Add{}.WriteTo(w)
type Add struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Add) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeAdd)
}
