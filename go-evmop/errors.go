package evmop

import (
	"errors"
)

var (
	errNegativeNumberOfBytesWritten = errors.New("Negative Number Of Bytes Written")
	errNilWriter                    = errors.New("Nil Writer")
)
