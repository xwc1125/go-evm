package evmop

import (
	"io"
)

// Return represents the "RETURN" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Return{}.WriteTo(w)
type Return struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Return) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeReturn)
}
