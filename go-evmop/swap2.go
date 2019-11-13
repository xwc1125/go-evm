package evmop

import (
	"io"
)

// Swap2 represents the "SWAP2" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Swap2{}.WriteTo(w)
type Swap2 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Swap2) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSwap2)
}
