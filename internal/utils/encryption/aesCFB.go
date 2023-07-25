package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/rs/zerolog/log"
)

func AesCFBDecryptionString(text, key string) (string, error) {
	if len(key) > 32 {
		log.Err(fmt.Errorf("key > 32")).Msg("key invalid")
		return "", fmt.Errorf("invalid decrypt")
	}

	cipherText, err := hex.DecodeString(text)
	if err != nil {
		log.Err(err).Msg("cannot hex decode text")
		return "", fmt.Errorf("invalid text")
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Err(err).Msg("generate new chiper failed")
		return "", fmt.Errorf("error chiper")
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("ciper text to short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

func AesCFBEncryptionString(text, key string) (string, error) {
	if len(key) > 32 {
		log.Err(fmt.Errorf("key > 32")).Msg("key invalid")
		return "", fmt.Errorf("invalid decrypt")
	}

	plaintext := []byte(text)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Err(err).Msg("generate new chiper failed")
		return "", fmt.Errorf("error chiper")
	}

	cipherText := make([]byte, aes.BlockSize+len(plaintext))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Err(err).Msg("error io readfull iv")
		return "", fmt.Errorf("error read")
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plaintext)

	return fmt.Sprintf("%x", cipherText), nil
}
