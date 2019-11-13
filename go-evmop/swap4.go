package evmop

import (
	"io"
)

// Swap4 represents the "SWAP4" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Swap4{}.WriteTo(w)
type Swap4 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Swap4) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSwap4)
}
