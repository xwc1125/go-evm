package evmop

import (
	"io"
)

// Sub represents the "SUB" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Sub{}.WriteTo(w)
type Sub struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Sub) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSub)
}
