package bech32_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"

	"example.com/bech32"
	btcbech32 "github.com/btcsuite/btcd/btcutil/bech32"
)

// This example demonstrates how to encode data into a bech32 string.
// 비트코인 bech32
func Test_ExampleEncode(*testing.T) {
	testdata := "Test data"
	fmt.Printf("init data: %s\n", testdata)

	data := []byte(testdata)
	fmt.Printf("init byte data: %x\n", data)

	// Convert test data to base32:
	conv, err := btcbech32.ConvertBits(data, 8, 5, true)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("converted data: %x\n", conv)

	encoded, err := btcbech32.Encode("btc", conv)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Show the encoded data.
	fmt.Println("Encoded Data:", encoded)

	// Output:
	// Encoded Data: btc!11111q123jhxapqv3shgcgkxpuhe
}

// This example demonstrates how to decode a bech32 encoded string.
func Test_ExampleDecode(*testing.T) {
	encoded := "btc123jhxapqv3shgcgkc3h0asdaf"
	fmt.Println("encoded:", encoded)
	fmt.Println("encoded length:", len(encoded))
	hrp, decoded, err := btcbech32.Decode(encoded)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Show the decoded data.
	fmt.Println("Decoded human-readable part:", hrp)
	fmt.Println("Decoded Data:", hex.EncodeToString(decoded))

	// Output:
	// Decoded human-readable part: bc
	// Decoded Data: 010e140f070d1a001912060b0d081504140311021d030c1d03040f1814060e1e160e140f070d1a001912060b0d081504140311021d030c1d03040f1814060e1e16
}

func TestEncodeAndDecode(t *testing.T) {

	sum := sha256.Sum256([]byte("hello world\n"))
	fmt.Printf("init data: %x\n", sum)
	bech, err := bech32.ConvertAndEncode("shasum", sum[:])
	fmt.Println("bech32 addr:", bech)

	if err != nil {
		t.Error(err)
	}
	hrp, data, err := bech32.DecodeAndConvert(bech)
	fmt.Println("hrp:", hrp)
	fmt.Printf("data: %x\n", data)

	if err != nil {
		t.Error(err)
	}
	if hrp != "shasum" {
		t.Error("Invalid hrp")
	}
	if bytes.Compare(data, sum[:]) != 0 {
		t.Error("Invalid decode")
	}
}
