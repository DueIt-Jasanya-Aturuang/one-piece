package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	"github.com/rs/zerolog/log"
)

func AesCBCDecryption(text, key, iv string) (string, error) {
	if len(key) > 32 {
		log.Err(fmt.Errorf("key > 32")).Msg("key invalid")
		return "", fmt.Errorf("invalid decrypt")
	}

	cipherText, err := base64.StdEncoding.DecodeString(text)
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

	mode := cipher.NewCBCDecrypter(block, []byte(iv))

	mode.CryptBlocks(cipherText, cipherText)

	cipherText = PKCS7Unpadding(cipherText)

	log.Info().Msgf("%v", cipherText)
	return string(cipherText), nil
}

func PKCS7Unpadding(cipherText []byte) []byte {
	length := len(cipherText)
	unpadding := int(cipherText[length-1])
	return cipherText[:(length - unpadding)]
}

func PKCS7Padding(cipherText []byte) []byte {
	padding := aes.BlockSize - len(cipherText)%aes.BlockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, paddingText...)
}

func AesCBCEncryption(text, key, iv string) (string, error) {
	if len(key) > 32 {
		log.Err(fmt.Errorf("key > 32")).Msg("key invalid")
		return "", fmt.Errorf("invalid decrypt")
	}

	plainText := []byte(text)
	plainText = PKCS7Padding(plainText)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Err(err).Msg("generate new chiper failed")
		return "", fmt.Errorf("error chiper")
	}

	cipherText := make([]byte, len(plainText))

	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(cipherText, plainText)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}
