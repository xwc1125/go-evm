package evmop

import (
	"io"
)

// CallBlackBox represents the "CALLBLACKBOX" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.CallBlackBox{}.WriteTo(w)
type CallBlackBox struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (CallBlackBox) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeCallBlackBox)
}
