package gomockctx

import (
	"crypto/rand"
	"math/big"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"abcdefghijklmnopqrstuvwxyz" +
	"0123456789"

// randString returns a cryptographically secure random string of alphanumeric
// characters of n length.
//
// Borrowed from github.com/jimeh/rands package.
func randString(n int) (string, error) {
	l := big.NewInt(int64(len(alphabet)))
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		index, err := rand.Int(rand.Reader, l)
		if err != nil {
			return "", err
		}
		b[i] = alphabet[index.Int64()]
	}

	return string(b), nil
}
