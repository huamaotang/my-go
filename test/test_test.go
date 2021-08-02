package test

import (
	"container/list"
	"fmt"
	"hash/crc32"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(1)
}

func TestRand(t *testing.T) {
	t.Log(len(STATICRANDSTRING), RandString(2, 6))
}

const STATICRANDSTRING = "0123456789abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

//生成36进制的邀请码
func RandString(uid int, length int) string {
	var base = []byte(STATICRANDSTRING)
	var quotient = uint64(uid)
	mod := uint64(0)
	l := list.New()
	fmt.Println(base, quotient, mod, l)
	for quotient != 0 {
		mod = quotient % 36
		quotient = quotient / 36
		l.PushFront(base[int(mod)])
	}
	listLen := l.Len()
	if listLen >= 6 {
		res := make([]byte, 0, length)
		for i := l.Front(); i != nil; i = i.Next() {
			res = append(res, i.Value.(byte))
		}
		return string(res)
	} else {
		res := make([]byte, 0, 6)
		for i := 0; i < 6; i++ {
			if i < 6-listLen {
				res = append(res, base[0])
			} else {
				res = append(res, l.Front().Value.(byte))
				l.Remove(l.Front())
			}
		}
		return string(res)
	}
}

func TestCRC(t *testing.T) {
	crc32q := crc32.MakeTable(0xD5828281)
	fmt.Printf("%08x\n", crc32.Checksum([]byte("Hello world"), crc32q))
	fmt.Println(crc32.Checksum([]byte("Hello worlda"), crc32q))
}
