package Type

//ETH，ETC有同样的地址生成方法

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Eth struct{}

func (e Eth) PrivtoPubkey(hex string) []byte {
	key, _ := crypto.HexToECDSA(hex)
	pubKey := key.PublicKey
	pubBytes := crypto.FromECDSAPub(&pubKey)
	return pubBytes
}
func (e Eth) PubkeytoAddress(pubkey []byte) string {
	address := common.BytesToAddress(crypto.Keccak256(pubkey[1:])[12:])
	return address.Hex()
}

func (e Eth) PrivtoAddress(hex string) string {
	key, _ := crypto.HexToECDSA(hex)
	pubKey := key.PublicKey
	address := crypto.PubkeyToAddress(pubKey)
	return address.Hex()
}

/*
func Eth(hex string) string {
	//fmt.Printf("======================ETH=or=ETC=====================\n")
	key, _ := crypto.HexToECDSA(hex)
	pubKey := key.PublicKey
	//fmt.Printf("The uncompressed public Key is: %X\n", uncompressedPubKey)
	//compressedPubKey := Util.CompressPubKey(key)
	//fmt.Printf("The compressed public key is :%X\n", compressedPubKey)
	compressedAddress := crypto.PubkeyToAddress(pubKey)
	//fmt.Printf("The address is: %v\n", compressedAddress.Hex())
	//fmt.Printf("=====================================================\n\n")
	return compressedAddress.Hex()
}
*/
