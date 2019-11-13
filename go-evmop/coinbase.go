package evmop

import (
	"io"
)

// Coinbase represents the "COINBASE" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Coinbase{}.WriteTo(w)
type Coinbase struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Coinbase) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeCoinbase)
}
