package base58

import (
	"bytes"
	"crypto/sha256"
	"math/big"

	"golang.org/x/crypto/ripemd160"
)

const ALPHABET = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

type Version byte

const (
	BITCOIN_PUBKEY_HASH Version = 0x00
	BITCOIN_SCRIPT_HASH         = 0x05
	BITCOIN_PRIVKEY             = 0x80
)

// pubkeys (p2pkh) and scripts (p2sh) are expected to be hashed before being changed into base58check
func PublicHash(b []byte) []byte {
	h1 := sha256.Sum256(b)
	h := ripemd160.New()
	h.Write(h1[0:])
	return h.Sum(nil)
}

// https://en.bitcoin.it/wiki/Base58Check_encoding
func Check(b []byte, v Version, comp bool) string {
	// version
	b = append([]byte{byte(v)}, b...)

	// compressed
	if v == BITCOIN_PRIVKEY && comp {
		b = append(b, byte(0x01))
	}
	if v == BITCOIN_PUBKEY_HASH && len(b) == 33 {
		b = append(b, byte(0x01))
	}

	// checksum
	h1 := sha256.Sum256(b)
	h := sha256.Sum256(h1[0:])
	csum := h[0:4]
	b = append(b, csum...)

	// encode
	i := new(big.Int)
	z := new(big.Int)
	m := new(big.Int)
	i.SetBytes(b)
	var buffer bytes.Buffer
	for i.Cmp(big.NewInt(0)) == 1 {
		i, _ = z.DivMod(i, big.NewInt(58), m)
		buffer.WriteString(string(ALPHABET[m.Uint64()]))
	}

	// leading zeros
	for _, i := range b {
		if i != 0x00 {
			break
		}
		buffer.WriteString(string(ALPHABET[0]))
	}

	// reverse
	c := []rune(buffer.String())
	n := len(c)
	for i := 0; i < n/2; i++ {
		c[i], c[n-1-i] = c[n-1-i], c[i]
	}

	return string(c)
}
