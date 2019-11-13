package evmop

import (
	"io"
)

// JumpDest represents the "JUMPDEST" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.JumpDest{}.WriteTo(w)
type JumpDest struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (JumpDest) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeJumpDest)
}
