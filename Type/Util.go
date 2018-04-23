package Type

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"unsafe"

	"github.com/dchest/blake256"
	"golang.org/x/crypto/ripemd160"
)

func CompressPubKey(key *ecdsa.PrivateKey) []byte {

	testX := key.PublicKey.X
	testY := key.PublicKey.Y
	xLen := len(testX.Bytes())
	yLen := len(testY.Bytes())
	yEnd := testY.Bytes()[yLen-1]
	compressedPubKey := make([]byte, 1+xLen)
	oddOreven := yEnd % 2
	if oddOreven == 0 {
		//fmt.Println("The yEnd is even, so compress with 02")
		compressedPubKey[0] = 2
		copy(compressedPubKey[1:], testX.Bytes())
	} else {
		//fmt.Println("The yEnd is odd, so compress with 03")
		compressedPubKey[0] = 3
		copy(compressedPubKey[1:], testX.Bytes())
	}

	return compressedPubKey
}

func ComputePubHash(pubkey []byte) []byte {
	sha := sha256.New()
	ripemd := ripemd160.New()

	writeBytes := (*string)(unsafe.Pointer(&pubkey))
	io.WriteString(sha, *writeBytes)
	shaPubKey := sha.Sum(nil)
	sha.Reset()

	writeBytes = (*string)(unsafe.Pointer(&shaPubKey))
	io.WriteString(ripemd, *writeBytes)
	pubKeyHash := ripemd.Sum(nil)
	ripemd.Reset()
	//fmt.Printf("The pubKeyHash is :%v,\nthe hex of publickeyhash is :%X\n", pubKeyHash, pubKeyHash)

	return pubKeyHash
}

func ComputedcrPubHash(pubkey []byte) []byte {

	h := blake256.New()
	h.Write(pubkey)
	shaPubKey := h.Sum(nil)
	h.Reset()

	ripemd := ripemd160.New()
	writeBytes := (*string)(unsafe.Pointer(&shaPubKey))
	io.WriteString(ripemd, *writeBytes)
	pubKeyHash := ripemd.Sum(nil)
	ripemd.Reset()
	//fmt.Printf("The pubKeyHash is :%v,\nthe hex of publickeyhash is :%X\n", pubKeyHash, pubKeyHash)

	return pubKeyHash
}

func DoubleHash(prefixHash []byte) []byte {

	sha := sha256.New()

	writeBytes := (*string)(unsafe.Pointer(&prefixHash))
	io.WriteString(sha, *writeBytes)
	firstSum := sha.Sum(nil)
	sha.Reset()
	writeBytes = (*string)(unsafe.Pointer(&firstSum))
	io.WriteString(sha, *writeBytes)
	checkSum := sha.Sum(nil)
	sha.Reset()
	//fmt.Printf("The checkSum is: %X\n",checkSum)

	return checkSum
}

func DcrDoubleHash(prefixPubHash []byte) (cksum [4]byte) {
	h := blake256.New()
	h.Write(prefixPubHash)
	intermediateHash := h.Sum(nil)
	h.Reset()
	h.Write(intermediateHash)
	finalHash := h.Sum(nil)
	copy(cksum[:], finalHash[:])
	return
}

func HexToByte(s string) ([]byte, error) {
	sbyte, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return sbyte, nil
}
