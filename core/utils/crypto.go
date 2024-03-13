package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Xmd5 md5编码
func Xmd5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
