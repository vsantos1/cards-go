package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(txt string) string {
	hasher := md5.New()
	hasher.Write([]byte(txt))
	return hex.EncodeToString(hasher.Sum(nil))
}