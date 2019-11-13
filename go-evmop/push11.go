package evmop

import (
	"io"
)

// Push11 represents the "PUSH11" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Push11{ 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b }.WriteTo(w)
type Push11 struct {
	Byte1 byte
	Byte2 byte
	Byte3 byte
	Byte4 byte
	Byte5 byte
	Byte6 byte
	Byte7 byte
	Byte8 byte
	Byte9 byte
	Byte10 byte
	Byte11 byte
}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (receiver Push11) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodePush11,
		receiver.Byte1,
		receiver.Byte2,
		receiver.Byte3,
		receiver.Byte4,
		receiver.Byte5,
		receiver.Byte6,
		receiver.Byte7,
		receiver.Byte8,
		receiver.Byte9,
		receiver.Byte10,
		receiver.Byte11,
	)
}
