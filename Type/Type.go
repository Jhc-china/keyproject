package Type

type Transfunction interface {
	PrivtoPubkey(string) []byte
	PubkeytoAddress([]byte) string
	PrivtoAddress(string) string
}
