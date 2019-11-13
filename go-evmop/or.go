package evmop

import (
	"io"
)

// Or represents the "OR" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Or{}.WriteTo(w)
type Or struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Or) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeOr)
}
