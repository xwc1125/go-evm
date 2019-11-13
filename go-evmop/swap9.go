package evmop

import (
	"io"
)

// Swap9 represents the "SWAP9" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Swap9{}.WriteTo(w)
type Swap9 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Swap9) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSwap9)
}
