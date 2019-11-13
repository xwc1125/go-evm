package evmop

import (
	"io"
)

// Log1 represents the "LOG1" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Log1{}.WriteTo(w)
type Log1 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Log1) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeLog1)
}
