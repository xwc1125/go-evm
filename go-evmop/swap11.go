package evmop

import (
	"io"
)

// Swap11 represents the "SWAP11" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Swap11{}.WriteTo(w)
type Swap11 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Swap11) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSwap11)
}
