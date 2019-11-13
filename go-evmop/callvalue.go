package evmop

import (
	"io"
)

// CallValue represents the "CALLVALUE" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.CallValue{}.WriteTo(w)
type CallValue struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (CallValue) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeCallValue)
}
