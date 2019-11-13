package evmop

import (
	"io"
)

// MSize represents the "MSIZE" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.MSize{}.WriteTo(w)
type MSize struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (MSize) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeMSize)
}
