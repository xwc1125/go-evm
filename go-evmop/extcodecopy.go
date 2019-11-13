package evmop

import (
	"io"
)

// ExtCodeCopy represents the "EXTCODECOPY" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.ExtCodeCopy{}.WriteTo(w)
type ExtCodeCopy struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (ExtCodeCopy) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeExtCodeCopy)
}
