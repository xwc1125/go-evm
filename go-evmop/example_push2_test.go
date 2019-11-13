package evmop_test

import (
	"github.com/reiver/go-evmop"

	"bytes"
	"fmt"
	"io"
	"os"
)

func ExamplePush2() {
	var buffer bytes.Buffer

	var writer io.Writer = &buffer

	n, err := evmop.Push2{0x05, 0x07}.WriteTo(writer)
	if nil != err {
		fmt.Fprintf(os.Stderr, "Problem writing “PUSH2 0x05 0x07” to buffer.\n")
		return
	}

	fmt.Printf("Wrote %d bytes to buffer.\n", n)
	fmt.Printf("Buffer: %#v\n", buffer.Bytes())

	// Output:
	// Wrote 3 bytes to buffer.
	// Buffer: []byte{0x61, 0x5, 0x7}
}
