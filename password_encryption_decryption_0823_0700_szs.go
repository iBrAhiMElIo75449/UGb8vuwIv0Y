// 代码生成时间: 2025-08-23 07:00:50
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "errors"
    "fmt"
    "log"
)

// EncryptionService 封装加密和解密的逻辑
type EncryptionService struct{
    key []byte
}

// NewEncryptionService 创建一个新的EncryptionService实例
func NewEncryptionService(key string) (*EncryptionService, error) {
    if len(key) != 32 {
        return nil, errors.New("key must be 32 bytes long")
    }
    return &EncryptionService{key: []byte(key)}, nil
}

// Encrypt 对数据进行加密
func (e *EncryptionService) Encrypt(plaintext string) (string, error) {
    block, err := aes.NewCipher(e.key)
    if err != nil {
        return "", err
    }

    plaintextBytes := []byte(plaintext)
    blockSize := block.BlockSize()
    padding := blockSize - len(plaintextBytes) % blockSize
    paddedPlaintext := make([]byte, len(plaintextBytes)+padding)
    copy(paddedPlaintext, plaintextBytes)
    copy(paddedPlaintext[len(plaintextBytes):], bytes.Repeat([]byte{byte(padding)}, padding))

    ciphertext := make([]byte, aes.BlockSize+len(plaintextBytes))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], paddedPlaintext)
    return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt 对数据进行解密
func (e *EncryptionService) Decrypt(ciphertext string) (string, error) {
    decodedData, err := base64.URLEncoding.DecodeString(ciphertext)
    if err != nil {
        return "", err
    }

    if len(decodedData) < aes.BlockSize {
        return "", errors.New("ciphertext too short")
    }

    iv := decodedData[:aes.BlockSize]
    decryptedData := make([]byte, len(decodedData)-aes.BlockSize)
    if len(decryptedData)%aes.BlockSize != 0 {
        return "", errors.New("ciphertext is not a multiple of the block size")
    }

    block, err := aes.NewCipher(e.key)
    if err != nil {
        return "", err
    }

    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(decryptedData, decodedData[aes.BlockSize:])

    // 去除填充
    padding := int(decryptedData[len(decryptedData)-1])
    if padding < 1 || padding > aes.BlockSize {
        return "", errors.New("invalid padding")
    }
    for _, v := range decryptedData[len(decryptedData)-padding:] {
        if v != byte(padding) {
            return "", errors.New("invalid padding")
        }
    }
    return string(decryptedData[:len(decryptedData)-padding]), nil
}

func main() {
    key := "your-32-byte-encryption-key" // 替换为32字节的密钥
    encService, err := NewEncryptionService(key)
    if err != nil {
        log.Fatalf("Error creating encryption service: %v", err)
    }

    plaintext := "Hello, this is a secret message!"
    encrypted, err := encService.Encrypt(plaintext)
    if err != nil {
        log.Fatalf("Error encrypting data: %v", err)
    }
    fmt.Printf("Encrypted: %s
", encrypted)

    decrypted, err := encService.Decrypt(encrypted)
    if err != nil {
        log.Fatalf("Error decrypting data: %v", err)
    }
    fmt.Printf("Decrypted: %s
", decrypted)
}
