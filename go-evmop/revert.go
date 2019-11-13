package evmop

import (
	"io"
)

// Revert represents the "REVERT" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Revert{}.WriteTo(w)
type Revert struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Revert) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeRevert)
}
