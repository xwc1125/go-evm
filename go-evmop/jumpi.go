package evmop

import (
	"io"
)

// JumpI represents the "JUMPI" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.JumpI{}.WriteTo(w)
type JumpI struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (JumpI) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeJumpI)
}
