package evmop

import (
	"io"
)

// Push2 represents the "PUSH2" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Push2{ 0x01, 0x02 }.WriteTo(w)
type Push2 struct {
	Byte1 byte
	Byte2 byte
}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (receiver Push2) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodePush2,
		receiver.Byte1,
		receiver.Byte2,
	)
}
