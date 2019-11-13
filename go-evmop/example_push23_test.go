package evmop_test

import (
	"github.com/reiver/go-evmop"

	"bytes"
	"fmt"
	"io"
	"os"
)

func ExamplePush23() {
	var buffer bytes.Buffer

	var writer io.Writer = &buffer

	n, err := evmop.Push23{0x05, 0x07, 0x09, 0x0b, 0x0d, 0x0f, 0x11, 0x13, 0x15, 0x17, 0x19, 0x1b, 0x1d, 0x1f, 0x2b, 0x2d, 0x2f, 0x31, 0x33, 0x35, 0x37, 0x39, 0x3b}.WriteTo(writer)
	if nil != err {
		fmt.Fprintf(os.Stderr, "Problem writing “PUSH23 0x05 0x07 0x09 0x0b 0x0d 0x0f 0x11 0x13 0x15 0x17 0x19 0x1b 0x1d 0x1f 0x2b 0x2d 0x2f 0x31 0x33 0x35 0x37 0x39 0x3b” to buffer.\n")
		return
	}

	fmt.Printf("Wrote %d bytes to buffer.\n", n)
	fmt.Printf("Buffer: %#v\n", buffer.Bytes())

	// Output:
	// Wrote 24 bytes to buffer.
	// Buffer: []byte{0x76, 0x5, 0x7, 0x9, 0xb, 0xd, 0xf, 0x11, 0x13, 0x15, 0x17, 0x19, 0x1b, 0x1d, 0x1f, 0x2b, 0x2d, 0x2f, 0x31, 0x33, 0x35, 0x37, 0x39, 0x3b}
}
