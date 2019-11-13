package test

import (
	"fmt"
	"github.com/xwc1125/go-evm/go-evm/core"
	"github.com/xwc1125/go-evm/go-evm/state"
	"github.com/xwc1125/go-evm/go-evm/vm"
	"github.com/xwc1125/go-evm/pkg/types"
	"github.com/xwc1125/go-evm/pkg/util/hexutil"
	"io/ioutil"
	"math/big"
	"os"
	"testing"
)

var (
	testHash    = types.BytesToHash([]byte("xujingshi"))
	fromAddress = types.BytesToAddress([]byte("xujingshi"))
	toAddress   = types.BytesToAddress([]byte("andone"))
	amount      = big.NewInt(0)
	nonce       = uint64(0)
	gasLimit    = uint64(100000)
	coinbase    = fromAddress
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}
func loadBin(filename string) []byte {
	code, err := ioutil.ReadFile(filename)
	must(err)
	return hexutil.MustDecode("0x" + string(code))
	//return []byte("0x" + string(code))
}
//func loadAbi(filename string) abi.ABI {
//	abiFile, err := os.Open(filename)
//	must(err)
//	defer abiFile.Close()
//	abiObj, err := abi.JSON(abiFile)
//	must(err)
//	return abiObj
//}

func TestEVM(t *testing.T) {
	//abiFileName := "./coin_sol_Coin.abi"
	binFileName := "./coin_sol_Coin.bin"
	data := loadBin(binFileName)

	// init db
	msg := core.NewMessage(fromAddress, &toAddress, nonce, amount, gasLimit, big.NewInt(0), data, false)
	cc := ChainContext{}
	ctx := core.NewEVMContext(msg,  cc, &fromAddress)
	dataPath := "/Users/xwc1125/Workspaces/Git/JoinChain/joinchain/go-evm/test/data"
	os.Remove(dataPath)
	//mdb, err := ethdb.NewLDBDatabase(dataPath, 100, 100)
	//must(err)
	//
	//// FIXME: whats the diff from db and statedb?
	db := state.NewDatabase()
	root := types.Hash{}
	statedb, err := state.New(root, db)
	must(err)
	//
	//set balance
	statedb.GetOrNewStateObject(fromAddress)
	statedb.GetOrNewStateObject(toAddress)
	statedb.AddBalance(fromAddress, big.NewInt(1e18))
	fmt.Println("init testBalance =", statedb.GetBalance(fromAddress))
	must(err)

	// log config
	////config := params.MainnetChainConfig
	logConfig := vm.LogConfig{}
	//// types.Address => Storage
	structLogger := vm.NewStructLogger(&logConfig)
	vmConfig := vm.Config{Debug: true, Tracer: structLogger /*, JumpTable: vm.NewByzantiumInstructionSet()*/}

	// load evm
	evm := vm.NewEVM(ctx, statedb, vmConfig)
	// caller
	contractRef := vm.AccountRef(fromAddress)
	contractCode, contractAddr, gasLeftover, vmerr := evm.Create(contractRef, data, statedb.GetBalance(fromAddress).Uint64(), big.NewInt(0))
	must(vmerr)
	fmt.Println(gasLeftover)
	fmt.Printf("getcode:%x\n%x\n", contractCode, statedb.GetCode(contractAddr))

	statedb.SetBalance(fromAddress, big.NewInt(0).SetUint64(gasLeftover))
	fmt.Println("after create contract, testBalance =", statedb.GetBalance(fromAddress))

	//abiObj := loadAbi(abiFileName)
	//
	//// method_id(4B) + args0(32B) + args1(32B) + ...
	//input, err := abiObj.Pack("minter")
	//must(err)
	//// 调用minter()获取矿工地址
	//outputs, gasLeftover, vmerr := evm.Call(contractRef, contractAddr, input, statedb.GetBalance(fromAddress).Uint64(), big.NewInt(0))
	//must(vmerr)
	//
	//// fmt.Printf("minter is %x\n", types.BytesToAddress(outputs))
	//// fmt.Printf("call address %x\n", contractRef)
	//
	//sender := types.BytesToAddress(outputs)
	//
	//if !bytes.Equal(sender.Bytes(), fromAddress.Bytes()) {
	//	fmt.Println("caller are not equal to minter!!")
	//	os.Exit(-1)
	//}
	//
	////senderAcc := vm.AccountRef(sender)
	//
	////// mint
	////input, err = abiObj.Pack("mint", sender, big.NewInt(1000000))
	////must(err)
	////outputs, gasLeftover, vmerr = evm.Call(senderAcc, contractAddr, input, statedb.GetBalance(fromAddress).Uint64(), big.NewInt(0))
	////must(vmerr)
	////statedb.SetBalance(fromAddress, big.NewInt(0).SetUint64(gasLeftover))
	////fmt.Println("after mint, testBalance =", gasLeftover)
	////
	//////send
	////input, err = abiObj.Pack("send", toAddress, big.NewInt(11))
	////outputs, gasLeftover, vmerr = evm.Call(senderAcc, contractAddr, input, statedb.GetBalance(fromAddress).Uint64(), big.NewInt(0))
	////must(vmerr)
	////fmt.Println("after send 11, testBalance =", gasLeftover)
	////
	//////send
	////input, err = abiObj.Pack("send", toAddress, big.NewInt(19))
	////must(err)
	////outputs, gasLeftover, vmerr = evm.Call(senderAcc, contractAddr, input, statedb.GetBalance(fromAddress).Uint64(), big.NewInt(0))
	////must(vmerr)
	////fmt.Println("after send 19, testBalance =", gasLeftover)
	////
	////// get receiver balance
	////input, err = abiObj.Pack("balances", toAddress)
	////must(err)
	////outputs, gasLeftover, vmerr = evm.Call(contractRef, contractAddr, input, statedb.GetBalance(fromAddress).Uint64(), big.NewInt(0))
	////must(vmerr)
	////Print(outputs, "balances")
	////fmt.Println("after get receiver balance, testBalance =", gasLeftover)
	////
	////// get sender balance
	////input, err = abiObj.Pack("balances", sender)
	////must(err)
	////outputs, gasLeftover, vmerr = evm.Call(contractRef, contractAddr, input, statedb.GetBalance(fromAddress).Uint64(), big.NewInt(0))
	////must(vmerr)
	////Print(outputs, "balances")
	////fmt.Println("after get sender balance, testBalance =", gasLeftover)
	//
	//// get event
	////logs := statedb.Logs()
	////
	////for _, log := range logs {
	////	fmt.Printf("%#v\n", log)
	////	for _, topic := range log.Topics {
	////		fmt.Printf("topic: %#v\n", topic)
	////	}
	////	fmt.Printf("data: %#v\n", log.Data)
	////}
	//
	////root, err = statedb.Commit(true)
	////must(err)
	////err = db.TrieDB().Commit(root, true)
	////must(err)
	//
	//fmt.Println("Root Hash", root.Hex())
	////mdb.Close()
	////
	////mdb2, err := ethdb.NewLDBDatabase(dataPath, 100, 100)
	////defer mdb2.Close()
	////must(err)
	//db2 := state.NewDatabase()
	//statedb2, err := state.New(root, db2)
	//must(err)
	//fmt.Println("get testBalance =", statedb2.GetBalance(fromAddress))
	//if !bytes.Equal(contractCode, statedb2.GetCode(contractAddr)) {
	//	fmt.Println("BUG!,the code was changed!")
	//	os.Exit(-1)
	//}
	//getVariables(statedb2, contractAddr)
}

func getVariables(statedb *state.StateDB, hash types.Address) {
	//cb := func(key, value types.Hash) bool {
	//	fmt.Printf("key=%x,value=%x\n", key, value)
	//	return true
	//}

	//statedb.ForEachStorage(hash, cb)

}

func Print(outputs []byte, name string) {
	fmt.Printf("method=%s, output=%x\n", name, outputs)
}

type ChainContext struct{}

// get block header
//func (cc ChainContext) GetHeader(hash types.Hash, number uint64) *types.Header {
//
//	//return &types.Header{
//	//	// ParentHash: types.Hash{},
//	//	// UncleHash:  types.Hash{},
//	//	Coinbase: fromAddress,
//	//	//	Root:        types.Hash{},
//	//	//	TxHash:      types.Hash{},
//	//	//	ReceiptHash: types.Hash{},
//	//	//	Bloom:      types.BytesToBloom([]byte("xujingshi")),
//	//	Difficulty: big.NewInt(1),
//	//	Number:     big.NewInt(1).SetUint64(number),
//	//	GasLimit:   1000000,
//	//	GasUsed:    0,
//	//	Time:       big.NewInt(time.Now().Unix()),
//	//	Extra:      nil,
//	//	//MixDigest:  testHash,
//	//	//Nonce:      types.EncodeNonce(1),
//	//}
//}

//func (cc ChainContext) Engine() consensus.Engine {
//	return nil
//}
