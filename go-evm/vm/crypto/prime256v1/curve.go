package prime256v1

import (
	"crypto/elliptic"
	"math/big"
)

func fromHex(s string) *big.Int {
	r, ok := new(big.Int).SetString(s, 16)
	if !ok {
		panic("invalid hex in source file: " + s)
	}
	return r
}

func P256() elliptic.Curve {
	return elliptic.P256()
}
