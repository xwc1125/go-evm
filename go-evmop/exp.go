package evmop

import (
	"io"
)

// Exp represents the "EXP" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Exp{}.WriteTo(w)
type Exp struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Exp) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeExp)
}
