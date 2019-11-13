package evmop

import (
	"io"
)

// Stop represents the "STOP" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Stop{}.WriteTo(w)
type Stop struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Stop) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeStop)
}
