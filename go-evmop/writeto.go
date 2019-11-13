package evmop

import (
	"io"
)

func writeTo(writer io.Writer, opcode byte, moreBytes ...byte) (int64, error) {

	if nil == writer {
		return 0, errNilWriter
	}

	var n64 int64

	{
		var b [1]byte = [1]byte{opcode}
		p := b[:]

		n, err := writer.Write(p)
		if nil != err {
			return n64, err
		}
		if 0 > n {
			return n64, errNegativeNumberOfBytesWritten
		}
		n64 += int64(n)
	}

	if 0 < len(moreBytes) {
		p := moreBytes

		n, err := writer.Write(p)
		if nil != err {
			return n64, err
		}
		if 0 > n {
			return n64, errNegativeNumberOfBytesWritten
		}
		n64 += int64(n)
	}

	return n64, nil
}
