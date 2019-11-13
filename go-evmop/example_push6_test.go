package evmop_test

import (
	"github.com/reiver/go-evmop"

	"bytes"
	"fmt"
	"io"
	"os"
)

func ExamplePush6() {
	var buffer bytes.Buffer

	var writer io.Writer = &buffer

	n, err := evmop.Push6{0x05, 0x07, 0x09, 0x0b, 0x0d, 0x0f}.WriteTo(writer)
	if nil != err {
		fmt.Fprintf(os.Stderr, "Problem writing “PUSH6 0x05 0x07 0x09 0x0b 0x0d 0x0f” to buffer.\n")
		return
	}

	fmt.Printf("Wrote %d bytes to buffer.\n", n)
	fmt.Printf("Buffer: %#v\n", buffer.Bytes())

	// Output:
	// Wrote 7 bytes to buffer.
	// Buffer: []byte{0x65, 0x5, 0x7, 0x9, 0xb, 0xd, 0xf}
}
