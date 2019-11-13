package evmop

import (
	"io"
)

// Log3 represents the "LOG3" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Log3{}.WriteTo(w)
type Log3 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Log3) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeLog3)
}
