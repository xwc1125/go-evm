package evmop

import (
	"io"
)

// Balance represents the "BALANCE" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Balance{}.WriteTo(w)
type Balance struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Balance) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeBalance)
}
