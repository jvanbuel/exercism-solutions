package cipher

// Cipher is an interface for a cipher
type Cipher interface {
	Encode(string) string
	Decode(string) string
}
