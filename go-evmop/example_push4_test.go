package evmop_test

import (
	"github.com/reiver/go-evmop"

	"bytes"
	"fmt"
	"io"
	"os"
)

func ExamplePush4() {
	var buffer bytes.Buffer

	var writer io.Writer = &buffer

	n, err := evmop.Push4{0x05, 0x07, 0x09, 0x0b}.WriteTo(writer)
	if nil != err {
		fmt.Fprintf(os.Stderr, "Problem writing “PUSH4 0x05 0x07 0x09 0x0b” to buffer.\n")
		return
	}

	fmt.Printf("Wrote %d bytes to buffer.\n", n)
	fmt.Printf("Buffer: %#v\n", buffer.Bytes())

	// Output:
	// Wrote 5 bytes to buffer.
	// Buffer: []byte{0x63, 0x5, 0x7, 0x9, 0xb}
}
