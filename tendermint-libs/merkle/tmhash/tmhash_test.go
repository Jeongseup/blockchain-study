package tmhash_test

import (
	"crypto/sha256"
	"testing"

	"example.com/merkle/tmhash"
	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	testVector := []byte("this is a string")

	hasher := tmhash.New()
	hasher.Write(testVector)
	// fmt.Println(hasher.BlockSize())
	// fmt.Println(hasher.Size())

	bz := hasher.Sum(nil)

	hasher = sha256.New()
	hasher.Write(testVector)
	bz2 := hasher.Sum(nil)

	// 그냥 앞에 20까지 사이즈 trunc 해주는 거
	bz2 = bz2[:20]

	t.Logf("tmhash: %x", bz)
	t.Logf("sha256: %x", bz2)

	assert.Equal(t, bz, bz2)
}
