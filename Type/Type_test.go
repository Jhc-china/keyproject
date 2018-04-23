package Type

import (
	"testing"
)

func TestBtcPub2Add(t *testing.T) {
	Pubkey := []byte{
		0x03, 0x25, 0xa2, 0x74, 0xd1, 0x04, 0x8f, 0x00, 0x52, 0x67,
		0xa9, 0x5e, 0x80, 0xe1, 0x99, 0x34, 0xb6, 0x91, 0xf1, 0x34,
		0x67, 0x23, 0x51, 0x9f, 0x4a, 0x02, 0xe2, 0xd1, 0xcb, 0x77,
		0xbc, 0x49, 0x77,
	}
	var Trans Transfunction
	Trans = new(Btc)
	expAdd := "1AYDEUKQeVggshaLmHbZjB4Qp1a8naQaYP"
	ret := Trans.PubkeytoAddress(Pubkey)
	if ret != expAdd {
		t.Error("Expected 1AYDEUKQeVggshaLmHbZjB4Qp1a8naQaYP, got ", ret)
	}
}

func TestBtcBase58hex2Add(t *testing.T) {
	base58Encodedpriv := "L19R5iY4XuSoeMUYqnrURSXwebqX5xq2h2QrMCYSt2emHsLi6ABg"
	var Trans Transfunction
	Trans = new(Btc)
	expAdd := "1AYDEUKQeVggshaLmHbZjB4Qp1a8naQaYP"
	ret := Trans.PrivtoAddress(base58Encodedpriv)
	if ret != expAdd {
		t.Error("Expected 1AYDEUKQeVggshaLmHbZjB4Qp1a8naQaYP, got ", ret)
	}
}

func TestDcrPub2Add(t *testing.T) {
	Pubkey := []byte{
		0x02, 0x8f, 0x53, 0x83, 0x8b, 0x76, 0x39, 0x56, 0x3f, 0x27,
		0xc9, 0x48, 0x45, 0x54, 0x9a, 0x41, 0xe5, 0x14, 0x6b, 0xcd,
		0x52, 0xe7, 0xfe, 0xf0, 0xea, 0x6d, 0xa1, 0x43, 0xa0, 0x2b,
		0x0f, 0xe2, 0xed,
	}
	var Trans Transfunction
	Trans = new(Dcr)
	expAdd := "DsT4FDqBKYG1Xr8aGrT1rKP3kiv6TZ5K5th"
	ret := Trans.PubkeytoAddress(Pubkey)
	if ret != expAdd {
		t.Error("Expected DsT4FDqBKYG1Xr8aGrT1rKP3kiv6TZ5K5th, got ", ret)
	}
}

func TestDcrBase58hex2Add(t *testing.T) {
	base58Encodedpriv := "KzYRq33r1tuvw8MDQaERt1CtWeQF9Spg75xkt2vH19yQkHuMPKaE"
	var Trans Transfunction
	Trans = new(Dcr)
	expAdd := "DsRW6cUCc8xHLRgQ6JRKimY6dm1qbzeFmwv"
	ret := Trans.PrivtoAddress(base58Encodedpriv)
	if ret != expAdd {
		t.Error("Expected DsRW6cUCc8xHLRgQ6JRKimY6dm1qbzeFmwv, got ", ret)
	}
}

func TestEthPub2Add(t *testing.T) {
	Pubkey := []byte{
		0x04, 0xbd, 0xc7, 0xfd, 0x20, 0xe8, 0x2d, 0x01, 0xfe, 0x16,
		0xe4, 0xd3, 0x4b, 0x47, 0x02, 0xa7, 0x9a, 0x4a, 0x56, 0x57,
		0xe3, 0x91, 0x43, 0xdf, 0x83, 0xee, 0x31, 0x26, 0x0a, 0xef,
		0x87, 0xe5, 0x1f, 0x1d, 0x85, 0xb9, 0x0a, 0x6d, 0xd6, 0xb3,
		0xcf, 0x74, 0x83, 0xd3, 0x42, 0x9f, 0x34, 0x41, 0x4a, 0x80,
		0xbc, 0x07, 0x5b, 0x9b, 0x32, 0x1e, 0xfe, 0xd1, 0x6c, 0x66,
		0xd9, 0xa4, 0xa5, 0xe1, 0x42,
	}
	var Trans Transfunction
	Trans = new(Eth)
	expAdd := "0x4E06C4A248916d555c138C7432A9D87c3F0BeC15"
	ret := Trans.PubkeytoAddress(Pubkey)
	if ret != expAdd {
		t.Error("Expected 0x4E06C4A248916d555c138C7432A9D87c3F0BeC15, got ", ret)
	}
}

func TestEthhex2Add(t *testing.T) {
	priv := "b74675a65e33897722a7c466798d218d21840d32db90fa96b944cbba2402a060"
	var Trans Transfunction
	Trans = new(Eth)
	expAdd := "0x4E06C4A248916d555c138C7432A9D87c3F0BeC15"
	ret := Trans.PrivtoAddress(priv)
	if ret != expAdd {
		t.Error("Expected 0x4E06C4A248916d555c138C7432A9D87c3F0BeC15, got ", ret)
	}
}

func TestLtcPub2Add(t *testing.T) {
	Pubkey := []byte{
		0x02, 0x47, 0xf2, 0x8e, 0x22, 0x6d, 0xd9, 0xc3, 0xbd, 0xd1,
		0x2d, 0x4a, 0x70, 0x4d, 0x36, 0xef, 0x87, 0xc7, 0x3e, 0xf2,
		0xcd, 0x9a, 0x0b, 0x0c, 0x7a, 0xd9, 0x76, 0xb0, 0x63, 0xe2,
		0x4f, 0x3b, 0xfe,
	}
	var Trans Transfunction
	Trans = new(Ltc)
	expAdd := "LU98e7YVa9GRxHkQA8qcRihUCJDy8i8Len"
	ret := Trans.PubkeytoAddress(Pubkey)
	if ret != expAdd {
		t.Error("Expected LU98e7YVa9GRxHkQA8qcRihUCJDy8i8Len, got ", ret)
	}
}

func TestLtcBase58hex2Add(t *testing.T) {
	base58Encodedpriv := "T98G8M4vSnmtu8KatsAEh5VAsJk6H5RWrjx1nYKYHQ4AdzS8HeJb"
	var Trans Transfunction
	Trans = new(Ltc)
	expAdd := "LU98e7YVa9GRxHkQA8qcRihUCJDy8i8Len"
	ret := Trans.PrivtoAddress(base58Encodedpriv)
	if ret != expAdd {
		t.Error("Expected LU98e7YVa9GRxHkQA8qcRihUCJDy8i8Len, got ", ret)
	}
}

func TestUsdtPub2Add(t *testing.T) {
	Pubkey := []byte{
		0x03, 0xe2, 0xe9, 0xc0, 0xf8, 0x7e, 0x65, 0xcf, 0xc7, 0x30,
		0x8d, 0xe7, 0x48, 0xec, 0xa8, 0x58, 0x6a, 0xc1, 0x98, 0xfa,
		0x82, 0xa3, 0xa2, 0x95, 0xb1, 0x22, 0x6a, 0xfe, 0x33, 0x7e,
		0xf9, 0xd0, 0x92,
	}
	var Trans Transfunction
	Trans = new(Usdt)
	expAdd := "38f65WFLd6zWuLCMumjZggCiQvqhuxuAeT"
	ret := Trans.PubkeytoAddress(Pubkey)
	if ret != expAdd {
		t.Error("Expected 38f65WFLd6zWuLCMumjZggCiQvqhuxuAeT, got ", ret)
	}
}

func TestUsdtBase58hex2Add(t *testing.T) {
	base58Encodedpriv := "KzYRq33r1tuvw8MDQaERt1CtWeQF9Spg75xkt2vH19yQkHuMPKaE"
	var Trans Transfunction
	Trans = new(Usdt)
	expAdd := "38f65WFLd6zWuLCMumjZggCiQvqhuxuAeT"
	ret := Trans.PrivtoAddress(base58Encodedpriv)
	if ret != expAdd {
		t.Error("Expected 38f65WFLd6zWuLCMumjZggCiQvqhuxuAeT, got ", ret)
	}
}

func TestZecPub2Add(t *testing.T) {
	Pubkey := []byte{
		0x03, 0xef, 0x8d, 0xe9, 0xf1, 0xb0, 0x94, 0x3b, 0xd6, 0x90,
		0x3d, 0x59, 0xd7, 0xcc, 0x64, 0x4a, 0x3c, 0x9b, 0x96, 0x47,
		0x9d, 0x3a, 0x0f, 0x02, 0xef, 0x1d, 0xd1, 0x4d, 0xf6, 0x36,
		0xe5, 0x85, 0x54,
	}
	var Trans Transfunction
	Trans = new(Zec)
	expAdd := "t1VW6RMwkPrMvuxmdzsLMtfejrDxV7yHJUT"
	ret := Trans.PubkeytoAddress(Pubkey)
	if ret != expAdd {
		t.Error("Expected t1VW6RMwkPrMvuxmdzsLMtfejrDxV7yHJUT, got ", ret)
	}
}

func TestZecBase58hex2Add(t *testing.T) {
	base58Encodedpriv := "KzDLA7RD8CNf8SaM64eL9HGmxtDKKXyaigWfbcWPaz4dfxuTsdLK"
	var Trans Transfunction
	Trans = new(Zec)
	expAdd := "t1VW6RMwkPrMvuxmdzsLMtfejrDxV7yHJUT"
	ret := Trans.PrivtoAddress(base58Encodedpriv)
	if ret != expAdd {
		t.Error("Expected t1VW6RMwkPrMvuxmdzsLMtfejrDxV7yHJUT, got ", ret)
	}
}
