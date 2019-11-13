package evmop

import (
	"io"
)

// GasLimit represents the "GASLIMIT" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.GasLimit{}.WriteTo(w)
type GasLimit struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (GasLimit) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeGasLimit)
}
