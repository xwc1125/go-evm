package evmop

import (
	"io"
)

// MStore8 represents the "MSTORE8" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.MStore8{}.WriteTo(w)
type MStore8 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (MStore8) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeMStore8)
}
