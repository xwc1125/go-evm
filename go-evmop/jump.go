package evmop

import (
	"io"
)

// Jump represents the "JUMP" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Jump{}.WriteTo(w)
type Jump struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Jump) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeJump)
}
