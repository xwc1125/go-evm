package evmop

import (
	"io"
)

// Xor represents the "XOR" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Xor{}.WriteTo(w)
type Xor struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Xor) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeXor)
}
