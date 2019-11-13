package evmop

import (
	"io"
)

// AddMod represents the "ADDMOD" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.AddMod{}.WriteTo(w)
type AddMod struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (AddMod) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeAddMod)
}
