package evmop

import (
	"io"
)

// SignExtend represents the "SIGNEXTEND" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.SignExtend{}.WriteTo(w)
type SignExtend struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (SignExtend) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSignExtend)
}
