package evmop

import (
	"io"
)

// Swap15 represents the "SWAP15" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Swap15{}.WriteTo(w)
type Swap15 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Swap15) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSwap15)
}
