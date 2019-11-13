package evmop

import (
	"io"
)

// DelegateCall represents the "DELEGATECALL" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.DelegateCall{}.WriteTo(w)
type DelegateCall struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (DelegateCall) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeDelegateCall)
}
