package evmop

import (
	"io"
)

// Swap10 represents the "SWAP10" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Swap10{}.WriteTo(w)
type Swap10 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Swap10) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSwap10)
}
