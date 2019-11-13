package evmop

import (
	"io"
)

// CallDataCopy represents the "CALLDATACOPY" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.CallDataCopy{}.WriteTo(w)
type CallDataCopy struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (CallDataCopy) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeCallDataCopy)
}
