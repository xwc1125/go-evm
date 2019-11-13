package prime256v1

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

type signatureTest struct {
	name    string
	sig     []byte
	der     bool
	isValid bool
}

// decodeHex decodes the passed hex string and returns the resulting bytes.  It
// panics if an error occurs.  This is only used in the tests as a helper since
// the only way it can fail is if there is an error in the test source code.
func decodeHex(hexStr string) []byte {
	b, err := hex.DecodeString(hexStr)
	if err != nil {
		panic("invalid hex string in test source: err " + err.Error() +
			", hex: " + hexStr)
	}

	return b
}

func testSignCompact(priv *PrivateKey, tag string, data []byte, isCompressed bool) {
	//priv, _ := GeneratePrivateKey()

	hashed := []byte("testing")
	sig, err := SignCompact(priv, hashed, isCompressed)
	if err != nil {
		fmt.Println(" error signing: ", tag, err)
		return
	}
	fmt.Printf("签名长度%x,%d \n", sig, len(sig))

	// 校验签名内容
	key := priv.PublicKey
	publicKey := NewPublicKey(key.X, key.Y)
	verify := VerifySignature(publicKey.Serialize(), hashed, sig[:64])

	fmt.Println("签名校验", verify)

	//fmt.Println("签名内容", hexutil.Encode(sig))

	pk, wasCompressed, err := RecoverCompact(sig, hashed)

	fmt.Println("公钥", pk)
	if err != nil {
		fmt.Println(" error recovering: ", tag, err)
		return
	}
	if pk.X.Cmp(priv.X) != 0 || pk.Y.Cmp(priv.Y) != 0 {
		fmt.Println("%s: recovered pubkey doesn't match original "+
			"(%v,%v) vs (%v,%v) ", tag, pk.X, pk.Y, priv.X, priv.Y)
		return
	}
	if wasCompressed != isCompressed {
		fmt.Println("%s: recovered pubkey doesn't match compressed state "+
			"(%v vs %v)", tag, isCompressed, wasCompressed)
		return
	}

	// If we change the compressed bit we should get the same key back,
	// but the compressed flag should be reversed.
	if isCompressed {
		sig[len(sig)-1] -= 4
	} else {
		sig[len(sig)-1] += 4
	}

	pk, wasCompressed, err = RecoverCompact(sig, hashed)
	fmt.Println("恢复的公钥", pk)
	if err != nil {
		fmt.Println(" error recovering (2):", tag, err)
		return
	}
	if pk.X.Cmp(priv.X) != 0 || pk.Y.Cmp(priv.Y) != 0 {
		fmt.Println("%s: recovered pubkey (2) doesn't match original "+
			"(%v,%v) vs (%v,%v) ", tag, pk.X, pk.Y, priv.X, priv.Y)
		return
	}
	if wasCompressed == isCompressed {
		fmt.Println("%s: recovered pubkey doesn't match reversed "+
			"compressed state (%v vs %v)", tag, isCompressed,
			wasCompressed)
		return
	}
}

func TestSignCompact(t *testing.T) {
	//priv, _ := GeneratePrivateKey()

	key := "cca9fbcc1b41e5a95d369eaa6ddcff73b61a4efaa279cfc6567e8daa39cbaf50"
	priv, _ := PrivKeyFromBytes(decodeHex(key))

	s := hex.EncodeToString(priv.Serialize())
	fmt.Println("私钥String", s)

	fmt.Println("私钥", priv.D)
	fmt.Println("公钥", priv.PublicKey)
	//fmt.Println("地址", crypto.PubkeyToAddress(priv.PublicKey))

	//for i := 0; i < 256; i++ {
	for i := 0; i < 1; i++ {
		name := fmt.Sprintf("test %d", i)
		data := make([]byte, 32)
		_, err := rand.Read(data)
		if err != nil {
			t.Errorf("failed to read random data for %s", name)
			continue
		}
		compressed := i%2 != 0
		testSignCompact(priv, name, data, compressed)
	}
}

func TestRFC6979(t *testing.T) {
	// Test vectors matching Trezor and CoreBitcoin implementations.
	// - https://github.com/trezor/trezor-crypto/blob/9fea8f8ab377dc514e40c6fd1f7c89a74c1d8dc6/tests.c#L432-L453
	// - https://github.com/oleganza/CoreBitcoin/blob/e93dd71207861b5bf044415db5fa72405e7d8fbc/CoreBitcoin/BTCKey%2BTests.m#L23-L49
	tests := []struct {
		key       string
		msg       string
		nonce     string
		signature string
	}{
		{
			"cca9fbcc1b41e5a95d369eaa6ddcff73b61a4efaa279cfc6567e8daa39cbaf50",
			"sample",
			"2df40ca70e639d89528a6b670d9d48d9165fdc0febc0974056bdce192b8e16a3",
			"3045022100af340daf02cc15c8d5d08d7735dfe6b98a474ed373bdb5fbecf7571be52b384202205009fb27f37034a9b24b707b7c6b79ca23ddef9e25f7282e8a797efe53a8f124",
		},
		{
			// This signature hits the case when S is higher than halforder.
			// If S is not canonicalized (lowered by halforder), this test will fail.
			"0000000000000000000000000000000000000000000000000000000000000001",
			"Satoshi Nakamoto",
			"8f8a276c19f4149656b280621e358cce24f5f52542772691ee69063b74f15d15",
			"3045022100934b1ea10a4b3c1757e2b0c017d0b6143ce3c9a7e6a4a49860d7a6ab210ee3d802202442ce9d2b916064108014783e923ec36b49743e2ffa1c4496f01a512aafd9e5",
		},
		{
			"fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364140",
			"Satoshi Nakamoto",
			"33a19b60e25fb6f4435af53a3d42d493644827367e6453928554f43e49aa6f90",
			"3045022100fd567d121db66e382991534ada77a6bd3106f0a1098c231e47993447cd6af2d002206b39cd0eb1bc8603e159ef5c20a5c8ad685a45b06ce9bebed3f153d10d93bed5",
		},
		{
			"f8b8af8ce3c7cca5e300d33939540c10d45ce001b8f252bfbc57ba0342904181",
			"Alan Turing",
			"525a82b70e67874398067543fd84c83d30c175fdc45fdeee082fe13b1d7cfdf1",
			"304402207063ae83e7f62bbb171798131b4a0564b956930092b33b07b395615d9ec7e15c022058dfcc1e00a35e1572f366ffe34ba0fc47db1e7189759b9fb233c5b05ab388ea",
		},
		{
			"0000000000000000000000000000000000000000000000000000000000000001",
			"All those moments will be lost in time, like tears in rain. Time to die...",
			"38aa22d72376b4dbc472e06c3ba403ee0a394da63fc58d88686c611aba98d6b3",
			"30450221008600dbd41e348fe5c9465ab92d23e3db8b98b873beecd930736488696438cb6b0220547fe64427496db33bf66019dacbf0039c04199abb0122918601db38a72cfc21",
		},
		{
			"e91671c46231f833a6406ccbea0e3e392c76c167bac1cb013f6f1013980455c2",
			"There is a computer disease that anybody who works with computers knows about. It's a very serious disease and it interferes completely with the work. The trouble with computers is that you 'play' with them!",
			"1f4b84c23a86a221d233f2521be018d9318639d5b8bbd6374a8a59232d16ad3d",
			"3045022100b552edd27580141f3b2a5463048cb7cd3e047b97c9f98076c32dbdf85a68718b0220279fa72dd19bfae05577e06c7c0c1900c371fcd5893f7e1d56a37d30174671f6",
		},
	}

	for i, test := range tests {
		privKey, _ := PrivKeyFromBytes(decodeHex(test.key))
		hash := sha256.Sum256([]byte(test.msg))

		// Ensure deterministically generated nonce is the expected value.
		gotNonce := NonceRFC6979(privKey.D, hash[:], nil, nil).Bytes()
		wantNonce := decodeHex(test.nonce)
		if !bytes.Equal(gotNonce, wantNonce) {
			t.Errorf("NonceRFC6979 #%d (%s): Nonce is incorrect: "+
				"%x (expected %x)", i, test.msg, gotNonce,
				wantNonce)
			continue
		}

		// Ensure deterministically generated signature is the expected value.
		gotSig, err := privKey.Sign(hash[:])
		if err != nil {
			t.Errorf("Sign #%d (%s): unexpected error: %v", i,
				test.msg, err)
			continue
		}
		gotSigBytes := gotSig.Serialize()
		wantSigBytes := decodeHex(test.signature)
		if !bytes.Equal(gotSigBytes, wantSigBytes) {
			t.Errorf("Sign #%d (%s): mismatched signature: %x "+
				"(expected %x)", i, test.msg, gotSigBytes,
				wantSigBytes)
			continue
		}
	}
}

func TestSignatureIsEqual(t *testing.T) {
	sig1 := &Signature{
		R: fromHex("0082235e21a2300022738dabb8e1bbd9d19cfb1e7ab8c30a23b0afbb8d178abcf3"),
		S: fromHex("24bf68e256c534ddfaf966bf908deb944305596f7bdcc38d69acad7f9c868724"),
	}
	sig2 := &Signature{
		R: fromHex("4e45e16932b8af514961a1d3a1a25fdf3f4f7732e9d624c6c61548ab5fb8cd41"),
		S: fromHex("181522ec8eca07de4860a4acdd12909d831cc56cbbac4622082221a8768d1d09"),
	}

	if !sig1.IsEqual(sig1) {
		t.Fatalf("value of IsEqual is incorrect, %v is "+
			"equal to %v", sig1, sig1)
	}

	if sig1.IsEqual(sig2) {
		t.Fatalf("value of IsEqual is incorrect, %v is not "+
			"equal to %v", sig1, sig2)
	}
}
