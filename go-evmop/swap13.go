package evmop

import (
	"io"
)

// Swap13 represents the "SWAP13" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Swap13{}.WriteTo(w)
type Swap13 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Swap13) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSwap13)
}
