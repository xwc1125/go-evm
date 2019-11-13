package evmop_test

import (
	"github.com/reiver/go-evmop"

	"bytes"
	"fmt"
	"io"
	"os"
)

func ExamplePush9() {
	var buffer bytes.Buffer

	var writer io.Writer = &buffer

	n, err := evmop.Push9{0x05, 0x07, 0x09, 0x0b, 0x0d, 0x0f, 0x11, 0x13, 0x15}.WriteTo(writer)
	if nil != err {
		fmt.Fprintf(os.Stderr, "Problem writing “PUSH9 0x05 0x07 0x09 0x0b 0x0d 0x0f 0x11 0x13 0x15” to buffer.\n")
		return
	}

	fmt.Printf("Wrote %d bytes to buffer.\n", n)
	fmt.Printf("Buffer: %#v\n", buffer.Bytes())

	// Output:
	// Wrote 10 bytes to buffer.
	// Buffer: []byte{0x68, 0x5, 0x7, 0x9, 0xb, 0xd, 0xf, 0x11, 0x13, 0x15}
}
