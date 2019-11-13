package evmop

import (
	"io"
)

// CallDataLoad represents the "CALLDATALOAD" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.CallDataLoad{}.WriteTo(w)
type CallDataLoad struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (CallDataLoad) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeCallDataLoad)
}
