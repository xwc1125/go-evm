package evmop

import (
	"io"
)

// Push17 represents the "PUSH17" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Push17{ 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11 }.WriteTo(w)
type Push17 struct {
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
	Byte12 byte
	Byte13 byte
	Byte14 byte
	Byte15 byte
	Byte16 byte
	Byte17 byte
}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (receiver Push17) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodePush17,
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
		receiver.Byte12,
		receiver.Byte13,
		receiver.Byte14,
		receiver.Byte15,
		receiver.Byte16,
		receiver.Byte17,
	)
}
