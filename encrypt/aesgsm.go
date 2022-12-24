package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/industry-netsecurity-solution/ins-security-channel/logger"
	"io"
)

func AES256GSMEncrypt(secretKey []byte, plaintext []byte) ([]byte, error) {

	if len(secretKey) != 32 {
		return nil, fmt.Errorf("secret key is not for AES-256: must be 256 bits")
	}

	// prepare AES-256-GSM cipher
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// make random nonce
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	logger.Debugf("Nonce: %s", hex.EncodeToString(nonce))

	// encrypt plaintext
	ciphertext := aesgcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil // nonce is included in ciphertext. no need to return
}

func AES256GSMDecrypt(secretKey []byte, ciphertext []byte) ([]byte, error) {

	if len(secretKey) != 32 {
		return nil, fmt.Errorf("secret key is not for AES-256: must be 256 bits")
	}

	// prepare AES-256-GSM cipher
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesgcm.NonceSize()
	nonce, pureCiphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	logger.Debugf("Nonce: %s", hex.EncodeToString(nonce))

	// decrypt ciphertext
	plaintext, err := aesgcm.Open(nil, nonce, pureCiphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func GetGCM(secretKey []byte) (cipher.AEAD, error) {

	if len(secretKey) != 32 {
		return nil, fmt.Errorf("secret key is not for AES-256: must be 256 bits")
	}

	// prepare AES-256-GSM cipher
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return aesgcm, nil
}

func GenSecretkeyByPassphrase(passphrase []byte) ([]byte, error) {
	hash := sha256.New()
	_, err := hash.Write(passphrase)
	if err != nil {
		return nil, err
	}

	secretKey := hash.Sum(nil)

	return secretKey, nil
}

func GenRandomData(size int) ([]byte, error) {
	secretKey := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, secretKey); err != nil {
		return nil, err
	}

	return secretKey, nil
}