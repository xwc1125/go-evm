package evmop

import (
	"io"
)

// SMod represents the "SMOD" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.SMod{}.WriteTo(w)
type SMod struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (SMod) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeSMod)
}
