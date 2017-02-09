package aes

import (
	"encoding/base64"
	"testing"
)

// TestAes - Aes Encrypt testing
func TestAes(t *testing.T) {

	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	key := []byte("sfe023f_9fd&fwfl")
	result, err := AesEncrypt([]byte("davygeek/tc/crypt/aes"), key)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf(base64.StdEncoding.EncodeToString(result))
	origData, err := AesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	t.Logf(string(origData))
}
