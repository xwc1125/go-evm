package evmop

import (
	"io"
)

// Log0 represents the "LOG0" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Log0{}.WriteTo(w)
type Log0 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Log0) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeLog0)
}
