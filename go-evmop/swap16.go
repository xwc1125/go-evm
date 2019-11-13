package evmop

import (
	"io"
)

// Swap16 represents the "SWAP16" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Swap16{}.WriteTo(w)
type Swap16 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Swap16) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSwap16)
}
