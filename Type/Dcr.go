package Type

// DCR与比特币也类似，主网络前缀为两位 0x073f
// 256加密算法采用blake256与比特币sha3-256有所区别，同样采用base58加密
// 私钥到公钥也是secp256k1

import (
	"github.com/anaskhan96/base58check"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58/base58"
)

type Dcr struct {
	prefix [2]byte
}

func (d Dcr) PrivtoPubkey(base58hex string) []byte {
	decodeHex, _ := base58check.Decode(base58hex)
	hex := decodeHex[2:66]
	key, _ := crypto.HexToECDSA(hex)
	compressedPubKey := CompressPubKey(key)
	return compressedPubKey
}

func (d *Dcr) PubkeytoAddress(pubkey []byte) string {
	d.prefix[0] = 0x07
	d.prefix[1] = 0x3f
	pubKeyHash := ComputedcrPubHash(pubkey)
	prefixPubHash := make([]byte, 2+len(pubKeyHash))
	prefixPubHash[0] = d.prefix[0]
	prefixPubHash[1] = d.prefix[1]
	copy(prefixPubHash[2:], pubKeyHash)

	checkSum := DcrDoubleHash(prefixPubHash)
	binaryAddress := make([]byte, 6+len(pubKeyHash))
	copy(binaryAddress[:22], prefixPubHash)
	copy(binaryAddress[22:], checkSum[:4])

	base58Address := base58.Encode(binaryAddress)
	return base58Address
}

func (d *Dcr) PrivtoAddress(base58hex string) string {
	d.prefix[0] = 0x07
	d.prefix[1] = 0x3f
	pubkey := d.PrivtoPubkey(base58hex)
	address := d.PubkeytoAddress(pubkey)
	return address
}

/*
func Dcr(version [2]byte, base58hex string) (uncompressdcrAddress string, compressdcrAddress string) {
	//fmt.Printf("=======================DCR===========================\n")
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

	uncompressdcrAddress = dcrPubToAddress(version, uncompressedPubKey)
	//fmt.Printf("The base58 encode uncompressed Address is: %v\n", uncompressdcrAddress)

	compressdcrAddress = dcrPubToAddress(version, compressedPubKey)
	//fmt.Printf("The base58 encode compressed Address is: %v\n", compressdcrAddress)
	//fmt.Printf("=====================================================\n\n")
	return
}

func dcrPubToAddress(version [2]byte, pubkey []byte) string {

	pubKeyHash := Util.ComputedcrPubHash(pubkey)
	prefixPubHash := make([]byte, 2+len(pubKeyHash))
	prefixPubHash[0] = version[0]
	prefixPubHash[1] = version[1]
	copy(prefixPubHash[2:], pubKeyHash)
	//fmt.Printf("The prefixhash is: %v\nthe hex of prefixhash is:%X", prefixPubHash, prefixPubHash)

	checkSum := Util.DcrDoubleHash(prefixPubHash)
	binaryAddress := make([]byte, 6+len(pubKeyHash))
	copy(binaryAddress[:22], prefixPubHash)
	copy(binaryAddress[22:], checkSum[:4])
	//fmt.Printf("The 26-byte binary address is:%X\n",binaryAddress)

	base58Address := base58.Encode(binaryAddress)
	return base58Address
}
*/
