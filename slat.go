package bkit

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"strconv"
)

// SaltWithString 加盐, 盐为字符串格式
func SaltWithString(origin, slat string) string {

	// 如果传入的时空字符串， 就不做处理
	if origin == "" {
		return origin
	}

	if slat == "" {
		slat = "default"
	}

	md5Slat := md5.Sum([]byte(slat))
	sha256Pass := sha256.Sum256([]byte(origin))

	lenSlat := len(md5Slat)
	lenPass := len(sha256Pass)

	md5Bytes := make([]byte, lenSlat+lenPass)

	copy(md5Bytes[:lenPass], md5Slat[:])
	copy(md5Bytes[lenPass:], sha256Pass[:])

	encrypt := md5.Sum(md5Bytes)

	return fmt.Sprintf("%x", encrypt)
}

// SaltWithInt64Hex16 给密码加盐, 盐是 int64 格式的的数字转成 16 进制字符串
func SaltWithInt64Hex16(password string, i64 int64) string {
	return SaltWithString(password, strconv.FormatInt(i64, 16))
}
