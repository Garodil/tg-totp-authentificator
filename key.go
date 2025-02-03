package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"time"
)

// Генерирует TOTP код
func generateTOTP(secret string) string {

	key, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		panic(err)
	}

	// получаем Unix Timestamp и делим на интервал
	timeUnix := time.Now().Unix()
	counter := timeUnix / int64(timeStep)

	// конвертируем в байты
	msg := make([]byte, 8)
	binary.BigEndian.PutUint64(msg, uint64(counter))

	// используем HMAC-SHA1
	h := hmac.New(sha1.New, key)
	h.Write(msg)
	hash := h.Sum(nil)

	// используем "динамическую обрезку"
	offset := hash[len(hash)-1] & 0x0f
	truncatedHash := hash[offset : offset+4]

	// берем последние 31 бит
	code := binary.BigEndian.Uint32(truncatedHash) & 0x7FFFFFFF
	code %= 1000000

	// форматируем в строку нужной длины
	return fmt.Sprintf("%06d", code)
}
