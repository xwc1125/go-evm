package evmop

import (
	"io"
)

// Byte represents the "BYTE" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Byte{}.WriteTo(w)
type Byte struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Byte) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeByte)
}
