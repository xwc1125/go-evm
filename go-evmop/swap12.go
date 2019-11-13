package evmop

import (
	"io"
)

// Swap12 represents the "SWAP12" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Swap12{}.WriteTo(w)
type Swap12 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Swap12) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSwap12)
}
