package evmop

import (
	"io"
)

// Swap6 represents the "SWAP6" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Swap6{}.WriteTo(w)
type Swap6 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Swap6) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSwap6)
}
