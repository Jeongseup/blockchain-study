package tmhash

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestSha256Hashing(t *testing.T) {
	data := "hello world"
	hash := sha256.New()

	hash.Write([]byte(data))

	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)

	fmt.Println(mdStr)
}
func TestSha256Hashing2(t *testing.T) {
	s := "sha256 this string"

	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)
	mdStr := hex.EncodeToString(bs)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
	fmt.Println(mdStr)

}
