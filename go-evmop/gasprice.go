package evmop

import (
	"io"
)

// GasPrice represents the "GASPRICE" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.GasPrice{}.WriteTo(w)
type GasPrice struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (GasPrice) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeGasPrice)
}
