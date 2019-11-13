package evmop

import (
	"io"
)

// SGT represents the "SGT" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.SGT{}.WriteTo(w)
type SGT struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (SGT) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSGT)
}
