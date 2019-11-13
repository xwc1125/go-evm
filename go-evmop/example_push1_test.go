package evmop_test

import (
	"github.com/reiver/go-evmop"

	"bytes"
	"fmt"
	"io"
	"os"
)

func ExamplePush1() {
	var buffer bytes.Buffer

	var writer io.Writer = &buffer

	n, err := evmop.Push1{0x05}.WriteTo(writer)
	if nil != err {
		fmt.Fprintf(os.Stderr, "Problem writing “PUSH1 0x05” to buffer.\n")
		return
	}

	fmt.Printf("Wrote %d bytes to buffer.\n", n)
	fmt.Printf("Buffer: %#v\n", buffer.Bytes())

	// Output:
	// Wrote 2 bytes to buffer.
	// Buffer: []byte{0x60, 0x5}
}
