package evmop

import (
	"io"
)

// Mul represents the "MUL" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Mul{}.WriteTo(w)
type Mul struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Mul) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeMul)
}
