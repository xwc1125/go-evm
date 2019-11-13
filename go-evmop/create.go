package evmop

import (
	"io"
)

// Create represents the "CREATE" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Create{}.WriteTo(w)
type Create struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Create) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeCreate)
}
