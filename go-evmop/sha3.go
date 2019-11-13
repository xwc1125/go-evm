package evmop

import (
	"io"
)

// SHA3 represents the "SHA3" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.SHA3{}.WriteTo(w)
type SHA3 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (SHA3) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSHA3)
}
