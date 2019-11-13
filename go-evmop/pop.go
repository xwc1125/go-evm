package evmop

import (
	"io"
)

// Pop represents the "POP" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Pop{}.WriteTo(w)
type Pop struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Pop) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodePop)
}
