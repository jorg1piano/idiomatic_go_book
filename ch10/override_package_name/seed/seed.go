package seed

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

// Rand generates a source for random numbers
func Rand() *rand.Rand {
	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		panic("cannot seed with cryptographic random number generator")
	}
	r := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
	return r
}
