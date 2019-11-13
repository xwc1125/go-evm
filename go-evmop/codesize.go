package evmop

import (
	"io"
)

// CodeSize represents the "CODESIZE" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.CodeSize{}.WriteTo(w)
type CodeSize struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (CodeSize) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeCodeSize)
}
