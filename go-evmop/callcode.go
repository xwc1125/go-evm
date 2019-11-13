package evmop

import (
	"io"
)

// CallCode represents the "CALLCODE" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.CallCode{}.WriteTo(w)
type CallCode struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (CallCode) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeCallCode)
}
