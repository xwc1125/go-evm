package evmop

import (
	"io"
)

// Swap14 represents the "SWAP14" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Swap14{}.WriteTo(w)
type Swap14 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Swap14) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSwap14)
}
