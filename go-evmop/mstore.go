package evmop

import (
	"io"
)

// MStore represents the "MSTORE" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.MStore{}.WriteTo(w)
type MStore struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (MStore) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeMStore)
}
