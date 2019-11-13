package evmop

import (
	"io"
)

// Blockhash represents the "BLOCKHASH" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Blockhash{}.WriteTo(w)
type Blockhash struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Blockhash) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeBlockhash)
}
