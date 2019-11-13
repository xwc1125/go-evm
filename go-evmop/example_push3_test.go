package evmop_test

import (
	"github.com/reiver/go-evmop"

	"bytes"
	"fmt"
	"io"
	"os"
)

func ExamplePush3() {
	var buffer bytes.Buffer

	var writer io.Writer = &buffer

	n, err := evmop.Push3{0x05, 0x07, 0x09}.WriteTo(writer)
	if nil != err {
		fmt.Fprintf(os.Stderr, "Problem writing “PUSH3 0x05 0x07 0x09” to buffer.\n")
		return
	}

	fmt.Printf("Wrote %d bytes to buffer.\n", n)
	fmt.Printf("Buffer: %#v\n", buffer.Bytes())

	// Output:
	// Wrote 4 bytes to buffer.
	// Buffer: []byte{0x62, 0x5, 0x7, 0x9}
}
