package crypto

/*
Original Author: alexandrevicenzi
Original Project Link: https://github.com/alexandrevicenzi/unchained
*/

import (
	cryptoRand "crypto/rand"
	"math/big"
	"math/rand"
)

const (
	allowedChars     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	allowedCharsSize = len(allowedChars)
	maxInt           = 1<<63 - 1
)

type source struct{}

func (s *source) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s *source) Uint64() uint64 {
	i, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(maxInt))

	if err != nil {
		panic(err)
	}

	return i.Uint64()
}

func (s *source) Seed(seed int64) {}

// GetRandomString returns a securely generated random string.
func GetRandomString(length int) string {
	b := make([]byte, length)
	rnd := rand.New(&source{})

	for i := range b {
		c := rnd.Intn(allowedCharsSize)
		b[i] = allowedChars[c]
	}

	return string(b)
}
