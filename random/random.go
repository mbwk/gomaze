package random

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	math_rand "math/rand"
)

type Rng struct {
	r *math_rand.Rand
	Seed int64
}

func NewRng(seed int64) *Rng {
	var num int64 = 0
	if seed == 0 {
		buf := make([]byte, 8)
		crypto_rand.Read(buf[:])
		num = int64(binary.LittleEndian.Uint64(buf))
	} else {
		num = seed
	}
	return &Rng{r: math_rand.New(math_rand.NewSource(num)), Seed: num}
}

func (rng *Rng) IntN(n int) int {
	return rng.r.Intn(n)
}

func RandomSelection[T any](rng *Rng, slice []T) T {
	possibilities := len(slice)
	return slice[rng.IntN(possibilities)]
}
