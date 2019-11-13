package evmop

import (
	"io"
)

// ExtCodeSize represents the "EXTCODESIZE" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.ExtCodeSize{}.WriteTo(w)
type ExtCodeSize struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (ExtCodeSize) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeExtCodeSize)
}
