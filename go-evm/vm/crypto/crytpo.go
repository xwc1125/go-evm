// description: joinchain 
// 
// @author: xwc1125
// @date: 2019/11/12
package crypto

import (
	"github.com/xwc1125/go-evm/go-evm/vm/crypto/prime256v1"
	"github.com/xwc1125/go-evm/go-evm/vm/params"
	"github.com/xwc1125/go-evm/models/rlp"
	"github.com/xwc1125/go-evm/pkg/crypto/keccak"
	"github.com/xwc1125/go-evm/pkg/types"
	"math/big"
)

var (
	secp256k1N, _  = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	secp256k1halfN = new(big.Int).Div(secp256k1N, big.NewInt(2))
)

func ValidateSignatureValues(v byte, r, s *big.Int, homestead bool) bool {
	if r.Cmp(params.Big1) < 0 || s.Cmp(params.Big1) < 0 {
		return false
	}
	// reject upper range of s values (ECDSA malleability)
	// see discussion in secp256k1/libsecp256k1/include/secp256k1.h
	if homestead && s.Cmp(secp256k1halfN) > 0 {
		return false
	}
	// Frontier: allow s to be in full N range
	return r.Cmp(secp256k1N) < 0 && s.Cmp(secp256k1N) < 0 && (v == 0 || v == 1)
}

func Ecrecover(hash, sig []byte) ([]byte, error) {
	return prime256v1.RecoverPubkey(hash, sig)
}


// CreateAddress creates an ethereum address given the bytes and the nonce
func CreateAddress(b types.Address, nonce uint64) types.Address {
	data, _ := rlp.EncodeToBytes([]interface{}{b, nonce})
	return types.BytesToAddress(keccak.Keccak256(data)[12:])
}

// CreateAddress2 creates an ethereum address given the address bytes, initial
// contract code and a salt.
func CreateAddress2(b types.Address, salt [32]byte, code []byte) types.Address {
	return types.BytesToAddress(keccak.Keccak256([]byte{0xff}, b.Bytes(), salt[:], keccak.Keccak256(code))[12:])
}
