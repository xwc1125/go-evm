package evmop

import (
	"io"
)

// CallDataSize represents the "CALLDATASIZE" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.CallDataSize{}.WriteTo(w)
type CallDataSize struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (CallDataSize) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeCallDataSize)
}
