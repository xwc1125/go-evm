package evmop

import (
	"io"
)

// Log2 represents the "LOG2" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Log2{}.WriteTo(w)
type Log2 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Log2) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeLog2)
}
