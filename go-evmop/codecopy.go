package evmop

import (
	"io"
)

// CodeCopy represents the "CODECOPY" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.CodeCopy{}.WriteTo(w)
type CodeCopy struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (CodeCopy) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeCodeCopy)
}
