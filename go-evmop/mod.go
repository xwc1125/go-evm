package evmop

import (
	"io"
)

// Mod represents the "MOD" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Mod{}.WriteTo(w)
type Mod struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Mod) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeMod)
}
