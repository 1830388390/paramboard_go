package utilss

import (
	"math/rand"
	"time"
	"unsafe"
)

const Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//const Nums = "0000000000000000000011111111111111111111222222222222222222223333333333333333333344444444444444444444" +
//"5555555555555555555566666666666666666666777777777777777777778888888888888888888899999999999999999999"

const Nums = "0123456789"
const LettersNums = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var src = rand.NewSource(time.Now().UnixNano())

const (
	// 支持预设字符串长度(2^8)
	letterIdBits = 8
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

func RandStr(n int, s string) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax Letters!
	for i, cache, remain := n-1, src.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(s) {
			b[i] = s[idx]
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

//func main() {
//	s := randStr(50000, Nums)
//	fmt.Println("方法二生成12位随机字符串: ", s)
//	var n = [10]int{}
//	for i := len(s) - 1; i >= 0; i-- {
//		if s[i] == '0' {
//			n[0] += 1
//		}
//		if s[i] == '1' {
//			n[1] += 1
//		}
//		if s[i] == '2' {
//			n[2] += 1
//		}
//		if s[i] == '3' {
//			n[3] += 1
//		}
//		if s[i] == '4' {
//			n[4] += 1
//		}
//		if s[i] == '5' {
//			n[5] += 1
//		}
//		if s[i] == '6' {
//			n[6] += 1
//		}
//		if s[i] == '7' {
//			n[7] += 1
//		}
//		if s[i] == '8' {
//			n[8] += 1
//		}
//		if s[i] == '9' {
//			n[9] += 1
//		}
//	}
//	print("ok")
//}
