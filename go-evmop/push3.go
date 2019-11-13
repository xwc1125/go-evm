package evmop

import (
	"io"
)

// Push3 represents the "PUSH3" OpCode in the Ethereum Virtual Machine (EVM).
//
// Example Usage:
//
//	var w io.Writer
//	
//	// ...
//	
//	n, err := evmop.Push3{ 0x01, 0x02, 0x03 }.WriteTo(w)
type Push3 struct {
	Byte1 byte
	Byte2 byte
	Byte3 byte
}

// WriteTo writers the bytecodes for this Ethereum Virtual Machine OpCode to a io.Writer.
//
// WriteTo makes this struct fit the io.WriterTo interface.
func (receiver Push3) WriteTo(writer io.Writer) (int64, error) {

	return writeTo(writer, CodePush3,
		receiver.Byte1,
		receiver.Byte2,
		receiver.Byte3,
	)
}
