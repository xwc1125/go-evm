package evmop

import (
	"io"
)

// Push1 represents the "PUSH1" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Push1{ 0x01 }.WriteTo(w)
type Push1 struct {
	Byte1 byte
}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (receiver Push1) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodePush1,
		receiver.Byte1,
	)
}
