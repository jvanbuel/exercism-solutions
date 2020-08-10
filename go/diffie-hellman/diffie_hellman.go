package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

// PrivateKey creates a random private key smaller than a large prime number p
func PrivateKey(p *big.Int) *big.Int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	private := big.NewInt(0).Rand(rnd, p)
	if private.Cmp(big.NewInt(1)) == 1 {
		return private
	}
	return private.Add(private, big.NewInt(2))
}

// PublicKey creates a public key from a private key, and a pair of prime numbers p and g
func PublicKey(private, p *big.Int, g int64) *big.Int {
	gbig := big.NewInt(g)
	return gbig.Exp(gbig, private, p)
}

// NewPair creates a private and public key from a pair of prime numbers p and g
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return private, public
}

// SecretKey generates a secret key from a private and public key and a large prime number p
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return big.NewInt(0).Exp(public2, private1, p)
}
