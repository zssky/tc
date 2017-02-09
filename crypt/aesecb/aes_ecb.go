package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// AesEncrypt - The data encryption algorithm in accordance with aes encryption
// key len need 16, 24, 32
func AesDecrypt(origData, key []byte) (crypted []byte, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	var raw []byte
	if raw, err = base64.StdEncoding.DecodeString(string(origData)); err != nil {
		return
	}

	var block cipher.Block
	if block, err = aes.NewCipher(key); err != nil {
		return
	}

	blockMode := NewECBDecrypter(block)
	crypted = make([]byte, len(raw))
	blockMode.CryptBlocks(crypted, raw)
	crypted = pkcs5UnPadding(crypted)
	return
}

// AesDecrypt - Decrypts the aes encrypted string into plaintext
func AesEncrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ecb := NewECBEncrypter(block)
	content := pkcs5Padding([]byte(crypted), block.BlockSize())
	buf := make([]byte, len(content))
	ecb.CryptBlocks(buf, content)

	return []byte(base64.StdEncoding.EncodeToString(buf)), nil
}

// pkcs5Padding - Data Padding
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// pkcs5UnPadding - Data UnPadding
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

// BlockSize - return BlockSize
func (x *ecbEncrypter) BlockSize() int {
	return x.blockSize
}

// CryptBlocks - crypt blocks
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}

	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}

	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

// BlockSize - return BlockSize
func (x *ecbDecrypter) BlockSize() int {
	return x.blockSize
}

// CryptBlocks - crypt blocks
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}

	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}

	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
