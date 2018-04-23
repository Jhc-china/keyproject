package Type

// LTC和比特币地址生成方式类似，前缀为一位 0x30

import (
	"github.com/anaskhan96/base58check"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58/base58"
)

type Ltc struct {
	prefix [1]byte
}

func (l Ltc) PrivtoPubkey(base58hex string) []byte {
	decodeHex, _ := base58check.Decode(base58hex)
	hex := decodeHex[2:66]
	key, _ := crypto.HexToECDSA(hex)
	compressedPubKey := CompressPubKey(key)
	return compressedPubKey
}

func (l *Ltc) PubkeytoAddress(pubkey []byte) string {
	l.prefix[0] = 0x30
	pubKeyHash := ComputePubHash(pubkey)
	prefixPubHash := make([]byte, 1+len(pubKeyHash))
	prefixPubHash[0] = l.prefix[0]
	copy(prefixPubHash[1:], pubKeyHash)

	checkSum := DoubleHash(prefixPubHash)
	binaryAddress := make([]byte, 5+len(pubKeyHash))
	copy(binaryAddress[:21], prefixPubHash)
	copy(binaryAddress[21:], checkSum[:4])

	base58Address := base58.Encode(binaryAddress)
	return base58Address
}

func (l *Ltc) PrivtoAddress(base58hex string) string {
	l.prefix[0] = 0x30
	pubkey := l.PrivtoPubkey(base58hex)
	address := l.PubkeytoAddress(pubkey)
	return address
}

/*
func Ltc(version [1]byte, base58hex string) (uncompressltcAddress string, compressltcAddress string) {
	//fmt.Printf("========================LTC==========================\n")
	decodeHex, _ := base58check.Decode(base58hex)
	//fmt.Printf("The decode Hex is: %v\n", decodeHex)

	hex := decodeHex[2:66]
	//fmt.Printf("The decode Hex remove prefix & endfix is :%v\n", hex)

	key, _ := crypto.HexToECDSA(hex)
	pubKey := key.PublicKey
	uncompressedPubKey := crypto.FromECDSAPub(&pubKey)
	//fmt.Printf("The uncompressed public key is :%X\n", uncompressedPubKey)

	compressedPubKey := Util.CompressPubKey(key)
	//fmt.Printf("The compressed public key is :%X\n", compressedPubKey)

	uncompressltcAddress = ltcPubToAddress(version, uncompressedPubKey)
	//fmt.Printf("The base58 encode uncompressed Address is: %v\n", uncompressltcAddress)

	compressltcAddress = ltcPubToAddress(version, compressedPubKey)
	//fmt.Printf("The base58 encode compressed Address is: %v\n", compressltcAddress)
	//fmt.Printf("=====================================================\n\n")
	return
}

func ltcPubToAddress(version [1]byte, pubkey []byte) string {

	pubKeyHash := Util.ComputePubHash(pubkey)
	prefixPubHash := make([]byte, 1+len(pubKeyHash))
	prefixPubHash[0] = version[0]
	copy(prefixPubHash[1:], pubKeyHash)
	//fmt.Printf("The prefixhash is: %v\nthe hex of prefixhash is:%X", prefixPubHash, prefixPubHash)

	checkSum := Util.DoubleHash(prefixPubHash)
	binaryAddress := make([]byte, 5+len(pubKeyHash))
	copy(binaryAddress[:21], prefixPubHash)
	copy(binaryAddress[21:], checkSum[:4])
	//fmt.Printf("The 25-byte binary address is:%X\n",binaryAddress)

	base58Address := base58.Encode(binaryAddress)
	return base58Address
}
*/
