package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"testing"
)


func _TestString(t *testing.T)  {
	key := []byte("example key 1235asdf")
	//ciphertext, _ := hex.DecodeString("f363f3ccdcb12bb883abf484ba77d9cd7d32b5baecb3d4b1b3e0e4beffdb3ded")
	ciphertext, _ := hex.DecodeString("ada9830b66249a480f58c6bb22547114664b7c3e86d27ec54ed0935d8bab5f973e7cc4870146e830d6ccbe3f823da168c58390bc4769fcb23912a6b18c8075fdce5d0497c3d67bf3c2c2e335dcc73dca")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// IV需要独特，但不安全。 所以这很常见
	// 将其包括在密文的开头。
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC模式总是在整个模块中工作。
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// 如果原始plaintext长度不是块的倍数
	// 大小，填充将不得不在加密时添加，这将是
	// 在这一点删除。 有关示例，请参阅
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. 然而，
	// 至关重要的是要注意密文必须被认证（即通过
	// 使用crypto/hmac）解密之前，以避免创建
	// 一个填充oracle。
	ciphertext = PKCS5UnPadding(ciphertext)
	fmt.Printf("%s\n", ciphertext)
}


func TestData_String(t *testing.T) {
	key := []byte("example key 1235example ")
	plaintext := []byte("exampleplaintext你好asdfasdfasdfasdfasdfasdfasdfasdfasdf")

	// CBC模式在块上工作，所以明文可能需要填充到块
	// 下一个整块。 有关这种填充的示例，请参阅
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. 这里，我们将
	// 假定 plaintext 已经是正确的长度。
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Print("NewCipher")
		panic(err)
	}
	plaintext = PKCS5Padding(plaintext,block.BlockSize())
	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}



	// IV需要独特，但不安全。 所以这很常见
	// 将其包括在密文的开头。
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// 记住密文必须经过认证是很重要的
	// （即通过使用crypto/hmac）以及为了加密而被加密
	// 保持secure。

	fmt.Printf("%x\n", ciphertext)
}
