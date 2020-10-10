package strconv

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"testing"
)

func Test_ParseInt(t *testing.T) {
	a := strconv.FormatInt(512, 32)
	t.Log(a)// 生成用户签名
	inviterSignSalt := "2588a03ebf16a12498935ab694ded812cae24969"
	inviter := strconv.FormatInt(75868286058868736, 32)
	b := md5.Sum([]byte(inviter + "." + inviterSignSalt))
	sign := hex.EncodeToString(b[:])
	t.Log(inviter, b, sign)
}
