package utils

import (
	"crypto/md5"
	"encoding/hex"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
)

// BcryptHash  bcrypt
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: MD5V
//@description: md5
//@param: str []byte
//@return: string

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

var CBC_SECRET_KEY = "b7ff977ced7b625b933871ba54ed3823600aa4b1f40f7910d39ce2e0e5c35c46" // Replace with your actual key

// CBCEncrypt encrypts data using AES-256-CBC
func CBCEncrypt(object interface{}) (map[string]string, error) {
	// Convert object to JSON string
	text, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}

	// Create cipher block
	key, err := hex.DecodeString(CBC_SECRET_KEY)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Generate random IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Create CBC encrypter
	mode := cipher.NewCBCEncrypter(block, iv)

	// Pad the data
	paddedData := pkcs7Padding(text, aes.BlockSize)

	// Encrypt the data
	encrypted := make([]byte, len(paddedData))
	mode.CryptBlocks(encrypted, paddedData)

	return map[string]string{
		"data": hex.EncodeToString(encrypted),
		"iv":   hex.EncodeToString(iv),
	}, nil
}

// CBCDecrypt decrypts data using AES-256-CBC
// 首先定义请求结构体
type DecryptRequest struct {
	Data string `json:"data"`
	IV   string `json:"iv"`
}

func CBCDecrypt(text interface{}) (interface{}, error) {
	var data map[string]string

	// Handle string input
	if str, ok := text.(string); ok {
		if err := json.Unmarshal([]byte(str), &data); err != nil {
			return nil, err
		}
	} else if m, ok := text.(map[string]string); ok {
		data = m
	} else {
		return nil, fmt.Errorf("invalid input type")
	}

	// Decode hex strings
	encryptedText, err := hex.DecodeString(data["data"])
	if err != nil {
		return nil, err
	}
	iv, err := hex.DecodeString(data["iv"])
	if err != nil {
		return nil, err
	}

	// Create cipher block
	key, err := hex.DecodeString(CBC_SECRET_KEY)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Create CBC decrypter
	mode := cipher.NewCBCDecrypter(block, iv)

	// Decrypt the data
	decrypted := make([]byte, len(encryptedText))
	mode.CryptBlocks(decrypted, encryptedText)

	// Remove padding
	unpadded, err := pkcs7Unpadding(decrypted)
	if err != nil {
		return nil, err
	}

	// Parse JSON
	var result interface{}
	if err := json.Unmarshal(unpadded, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// PKCS7 padding
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := make([]byte, len(data)+padding)
	copy(padtext, data)
	for i := len(data); i < len(padtext); i++ {
		padtext[i] = byte(padding)
	}
	return padtext
}

// PKCS7 unpadding
func pkcs7Unpadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, fmt.Errorf("empty data")
	}
	padding := int(data[length-1])
	if padding > length {
		return nil, fmt.Errorf("invalid padding")
	}
	return data[:length-padding], nil
}
