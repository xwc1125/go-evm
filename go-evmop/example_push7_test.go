package evmop_test

import (
	"github.com/reiver/go-evmop"

	"bytes"
	"fmt"
	"io"
	"os"
)

func ExamplePush7() {
	var buffer bytes.Buffer

	var writer io.Writer = &buffer

	n, err := evmop.Push7{0x05, 0x07, 0x09, 0x0b, 0x0d, 0x0f, 0x11}.WriteTo(writer)
	if nil != err {
		fmt.Fprintf(os.Stderr, "Problem writing “PUSH7 0x05 0x07 0x09 0x0b 0x0d 0x0f 0x11” to buffer.\n")
		return
	}

	fmt.Printf("Wrote %d bytes to buffer.\n", n)
	fmt.Printf("Buffer: %#v\n", buffer.Bytes())

	// Output:
	// Wrote 8 bytes to buffer.
	// Buffer: []byte{0x66, 0x5, 0x7, 0x9, 0xb, 0xd, 0xf, 0x11}
}
