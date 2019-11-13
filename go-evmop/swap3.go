package evmop

import (
	"io"
)

// Swap3 represents the "SWAP3" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Swap3{}.WriteTo(w)
type Swap3 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Swap3) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSwap3)
}
