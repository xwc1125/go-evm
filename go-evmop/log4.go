package evmop

import (
	"io"
)

// Log4 represents the "LOG4" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Log4{}.WriteTo(w)
type Log4 struct {}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (Log4) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodeLog4)
}
