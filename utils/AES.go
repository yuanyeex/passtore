package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

//加密过程：
//  1、处理数据，对数据进行填充，采用PKCS7（当密钥长度不够时，缺几位补几个几）的方式。
//  2、对数据进行加密，采用AES加密方法中CBC加密模式
//  3、对得到的加密数据，进行base64加密，得到字符串
// 解密过程相反

//16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法

// the pkcs7Padding
func pkcs7Padding(data []byte, blockSize int) []byte {
	// the padding size, 1 ~ blockSize
	padding := blockSize - len(data)%blockSize
	// padding N byte
	padData := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padData...)
}

// the pkcs7UnPadding
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("invalid Encrypted Data")
	}
	unPadding := int(data[length-1])
	if unPadding > length || unPadding == 0 {
		return nil, errors.New("ErrInvalidPKCS7Padding")
	}
	return data[:(length - unPadding)], nil
}

func EncryptStr(data string, key string) (string, error) {
	encrypted, err := Encrypt([]byte(data), []byte(key))
	if err == nil {
		return base64.StdEncoding.EncodeToString(encrypted), nil
	}
	return "", err
}

func DecryptStr(data string, key string) (string, error) {
	encrypted, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	decrypted, err := Decrypt(encrypted, []byte(key))
	if err == nil {
		return string(decrypted), nil
	}
	return "", err
}

func Encrypt(data []byte, key []byte) ([]byte, error) {
	// create the aesCipher instance
	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// determine the block size
	blockSize := aesCipher.BlockSize()
	// padding
	paddedData := pkcs7Padding(data, blockSize)
	encrypted := make([]byte, len(paddedData))
	// the cbc encryption
	blockMode := cipher.NewCBCEncrypter(aesCipher, key[:blockSize])
	// encryption
	blockMode.CryptBlocks(encrypted, paddedData)
	return encrypted, nil
}

func Decrypt(data []byte, key []byte) ([]byte, error) {
	// create the aesCipher instance
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// determine the block size
	blockSize := block.BlockSize()
	// init decrypted
	decrypted := make([]byte, len(data))
	// decryption
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	blockMode.CryptBlocks(decrypted, data)
	return pkcs7UnPadding(decrypted)
}
