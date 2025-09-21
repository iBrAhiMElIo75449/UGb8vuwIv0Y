// 代码生成时间: 2025-09-21 18:49:49
 * Features:
 * - Encrypts and decrypts passwords using AES encryption.
 * - Provides a simple CLI interface for password operations.
 */

package main

import (
    "crypto/aes"
    "crypto/cipher"
    "encoding/hex"
    "flag"
    "io/ioutil"
    "log"
    "os"
    "syscall"
)

// Constants for AES encryption
const (
    plainTextKey  = "your_encryption_key" // Replace with your own encryption key
    plainTextKeySize = 32 // AES key size, 32 bytes for AES-256
)

// encrypt encrypts the plain text using AES encryption
func encrypt(plainText string) (string, error) {
    key := []byte(plainTextKey)

    // Pad the key to the correct size
    if len(key) < plainTextKeySize {
        key = append(key, bytes.Repeat([]byte{0}, plainTextKeySize-len(key))...)
    } else if len(key) > plainTextKeySize {
        return "", fmt.Errorf("encryption key is too long")
    }

    // Convert plain text to bytes
    plainTextBytes := []byte(plainText)

    // Generate a random initialization vector (IV)
    iv := make([]byte, aes.BlockSize)
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }

    // Encrypt the plain text
    cipherBlock, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }
    stream := cipher.NewCFBEncrypter(cipherBlock, iv)
    cipherText := make([]byte, len(plainTextBytes))
    stream.XORKeyStream(cipherText, plainTextBytes)

    // Return the encrypted text as hex string
    return hex.EncodeToString(append(iv, cipherText...)), nil
}

// decrypt decrypts the encrypted text using AES decryption
func decrypt(cipherText string) (string, error) {
    key := []byte(plainTextKey)

    // Pad the key to the correct size
    if len(key) < plainTextKeySize {
        key = append(key, bytes.Repeat([]byte{0}, plainTextKeySize-len(key))...)
    } else if len(key) > plainTextKeySize {
        return "", fmt.Errorf("encryption key is too long")
    }

    // Convert cipher text to bytes
    cipherTextBytes, err := hex.DecodeString(cipherText)
    if err != nil {
        return "", err
    }

    // Extract the IV from the cipher text
    iv := cipherTextBytes[:aes.BlockSize]
    cipherTextBytes = cipherTextBytes[aes.BlockSize:]

    // Decrypt the cipher text
    cipherBlock, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }
    stream := cipher.NewCFBDecrypter(cipherBlock, iv)
    plainTextBytes := make([]byte, len(cipherTextBytes))
    stream.XORKeyStream(plainTextBytes, cipherTextBytes)

    // Return the decrypted text
    return string(plainTextBytes), nil
}

func main() {
    var action string
    var password string
    flag.StringVar(&action, "action", "", "Action to perform: encrypt or decrypt")
    flag.StringVar(&password, "password", "", "Password to encrypt or decrypt")
    flag.Parse()

    if action == "encrypt" {
        encryptedPassword, err := encrypt(password)
        if err != nil {
            log.Fatalf("Error encrypting password: %v", err)
        }
        fmt.Println("Encrypted password: ", encryptedPassword)
    } else if action == "decrypt" {
        decryptedPassword, err := decrypt(password)
        if err != nil {
            log.Fatalf("Error decrypting password: %v", err)
        }
        fmt.Println("Decrypted password: ", decryptedPassword)
    } else {
        log.Fatalf("Invalid action. Please specify 'encrypt' or 'decrypt'.")
    }
}
