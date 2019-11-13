package evmop

import (
	"io"
)

// Swap5 represents the "SWAP5" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Swap5{}.WriteTo(w)
type Swap5 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Swap5) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSwap5)
}
