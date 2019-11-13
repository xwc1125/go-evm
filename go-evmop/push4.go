package evmop

import (
	"io"
)

// Push4 represents the "PUSH4" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Push4{ 0x01, 0x02, 0x03, 0x04 }.WriteTo(w)
type Push4 struct {
	Byte1 byte
	Byte2 byte
	Byte3 byte
	Byte4 byte
}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (receiver Push4) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodePush4,
		receiver.Byte1,
		receiver.Byte2,
		receiver.Byte3,
		receiver.Byte4,
	)
}
