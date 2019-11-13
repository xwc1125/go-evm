package evmop

import (
	"io"
)

// Address represents the "ADDRESS" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Address{}.WriteTo(w)
type Address struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Address) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeAddress)
}
