/*
 * Copyright ArxanChain Ltd. 2020 All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

// Validate implement validate the package
func Validate(rawData, ssk, signature string) bool {
	r := sha1.Sum([]byte(rawData + ssk))
	return signature == hex.EncodeToString(r[:])
}

// Decrypt implement decrypt data
func Decrypt(ssk, data, iv string) (bts []byte, err error) {
	key, err := base64.StdEncoding.DecodeString(ssk)
	if err != nil {
		return
	}

	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return
	}

	rawIV, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	size := aes.BlockSize

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < size {
		err = errors.New("cipher too short")
		return
	}

	// CBC mode always works in whole blocks.
	if len(ciphertext)%size != 0 {
		err = errors.New("cipher is not a multiple of the block size")
		return
	}

	mode := cipher.NewCBCDecrypter(block, rawIV[:size])
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	return PKCS5UnPadding(plaintext)
}

// PKCS5UnPadding format the data
func PKCS5UnPadding(plaintext []byte) ([]byte, error) {
	ln := len(plaintext)

	// 去掉最后一个字节 unPadding 次
	unPadding := int(plaintext[ln-1])

	if unPadding > ln {
		return []byte{}, errors.New("数据不正确")
	}

	return plaintext[:(ln - unPadding)], nil
}

// CBCDecrypt encrypt the data
func CBCDecrypt(ssk, data, iv string) (bts []byte, err error) {
	key, err := base64.StdEncoding.DecodeString(ssk)
	if err != nil {
		return
	}

	cipherText, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return
	}

	rawIV, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	size := aes.BlockSize

	if len(cipherText) < size {
		err = errors.New("cipher too short")
		return
	}

	if len(cipherText)%size != 0 {
		err = errors.New("cipher is not a multiple of the block size")
		return
	}

	mode := cipher.NewCBCDecrypter(block, rawIV[:size])
	plaintext := make([]byte, len(cipherText))
	mode.CryptBlocks(plaintext, cipherText)

	return PKCS5UnPadding(plaintext)
}
