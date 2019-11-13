package evmop

import (
	"io"
)

// MLoad represents the "MLOAD" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.MLoad{}.WriteTo(w)
type MLoad struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (MLoad) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeMLoad)
}
