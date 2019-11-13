package evmop

import (
	"io"
)

// Push6 represents the "PUSH6" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Push6{ 0x01, 0x02, 0x03, 0x04, 0x05, 0x06 }.WriteTo(w)
type Push6 struct {
	Byte1 byte
	Byte2 byte
	Byte3 byte
	Byte4 byte
	Byte5 byte
	Byte6 byte
}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (receiver Push6) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodePush6,
		receiver.Byte1,
		receiver.Byte2,
		receiver.Byte3,
		receiver.Byte4,
		receiver.Byte5,
		receiver.Byte6,
	)
}
