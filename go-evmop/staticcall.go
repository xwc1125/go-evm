package evmop

import (
	"io"
)

// StaticCall represents the "STATICCALL" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.StaticCall{}.WriteTo(w)
type StaticCall struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (StaticCall) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeStaticCall)
}
