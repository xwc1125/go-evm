package evmop

import (
	"io"
)

// SStore represents the "SSTORE" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.SStore{}.WriteTo(w)
type SStore struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (SStore) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSStore)
}
