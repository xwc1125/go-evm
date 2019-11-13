package evmop

import (
	"io"
)

// GT represents the "GT" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.GT{}.WriteTo(w)
type GT struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (GT) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeGT)
}
