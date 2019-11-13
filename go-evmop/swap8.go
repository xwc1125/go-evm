package evmop

import (
	"io"
)

// Swap8 represents the "SWAP8" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Swap8{}.WriteTo(w)
type Swap8 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Swap8) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSwap8)
}
