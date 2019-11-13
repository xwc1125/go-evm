# go-evmop

Package **evmop** provides tools for turning Ethereum Virtual Machine (EVM) OpCodes into bytecodes, for the Go programming language.

This might be useful to someone writing an AOT or JIT compiler targeting the Ethereum Virtual Machine (EVM).


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-evmop

[![GoDoc](https://godoc.org/github.com/reiver/go-evmop?status.svg)](https://godoc.org/github.com/reiver/go-evmop)


## Example

The following stack-oriented style of assembly for the Ethereum Virtual Machine (EVM):
```
push1 3 push1 2 add push1 1 mul
```

Would be the equivalent to the following Go code using the evmop package: 
```go
var writer io.Writer

// ...

evmop.Push1{3}.WriteTo(writer)

evmop.Push1{2}.WriteTo(writer)

evmop.Add{}.WriteTo(writer)

evmop.Push1{1}.WriteTo(writer)

evmop.Mul{}.WriteTo(writer)
```

(Note that the code above didn't bother checking for errors. That was to make that example easier to understand. **BUT YOUR CODE SHOULD HANDLE ERRORS!**) 

 At this point, writer would have had written to it, the equivalent of:
```go
[]byte{
	0x60, 0x03, // push1 3
	0x60, 0x02, // push1 2
	0x01,       // add
	0x60, 0x01, // push1 1
	0x02,       // mul
}
```
