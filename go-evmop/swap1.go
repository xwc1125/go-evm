package evmop

import (
	"io"
)

// Swap1 represents the "SWAP1" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Swap1{}.WriteTo(w)
type Swap1 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Swap1) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSwap1)
}
