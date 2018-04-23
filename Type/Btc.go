package Type

// BTC,BCH,BTY都是比特币的地址生成方式
// base58check加密的压缩形式私钥 (L，K开头)
// 解码得到私钥，secp256k1计算得到公钥 xy
// 利用压缩形式公钥得到地址 base58(prefix + payload + checksum)

import (
	"github.com/anaskhan96/base58check"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58/base58"
)

type Btc struct {
	prefix [1]byte
}

func (b Btc) PrivtoPubkey(base58hex string) []byte {
	decodeHex, _ := base58check.Decode(base58hex)
	hex := decodeHex[2:66]
	key, _ := crypto.HexToECDSA(hex)
	compressedPubKey := CompressPubKey(key)
	return compressedPubKey
}

func (b *Btc) PubkeytoAddress(pubkey []byte) string {
	b.prefix[0] = 0x00
	pubKeyHash := ComputePubHash(pubkey)
	prefixPubHash := make([]byte, 1+len(pubKeyHash))
	prefixPubHash[0] = b.prefix[0]
	copy(prefixPubHash[1:], pubKeyHash)

	checkSum := DoubleHash(prefixPubHash)
	binaryAddress := make([]byte, 5+len(pubKeyHash))
	copy(binaryAddress[:21], prefixPubHash)
	copy(binaryAddress[21:], checkSum[:4])

	base58Address := base58.Encode(binaryAddress)
	return base58Address
}

func (b *Btc) PrivtoAddress(base58hex string) string {
	b.prefix[0] = 0x00
	pubkey := b.PrivtoPubkey(base58hex)
	address := b.PubkeytoAddress(pubkey)
	return address
}

/*
func btc(version [1]byte, base58hex string) (uncompressbtcAddress string, compressbtcAddress string) {
	//fmt.Printf("=============BTC=or=BCC=or=BTY=======================\n")
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

	uncompressbtcAddress = btcPubToAddress(version, uncompressedPubKey)
	//fmt.Printf("The base58 encode uncompressed Address is: %v\n", uncompressbtcAddress)

	compressbtcAddress = btcPubToAddress(version, compressedPubKey)
	//fmt.Printf("The base58 encode compressed Address is: %v\n", compressbtcAddress)
	//fmt.Printf("=====================================================\n\n")
	return
}

func btcPubToAddress(version [1]byte, pubkey []byte) string {

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
