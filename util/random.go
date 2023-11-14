package util

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const number = "0123456789"

//func RandomInt(min, max int) int {
//	return min + rand.Intn(max-min+1)
//}

//func RandomString(n int) string {
//	var sb strings.Builder
//	k := len(alphabet)
//
//	for i := 0; i < n; i++ {
//		c := alphabet[rand.Intn(k)]
//		sb.WriteByte(c)
//	}
//
//	return sb.String()
//}
//
//func RandomPhone(n int) string {
//	var sb strings.Builder
//	k := len(number)
//
//	for i := 0; i < n-1; i++ {
//		c := number[rand.Intn(k)]
//		sb.WriteByte(c)
//	}
//
//	return "0" + sb.String()
//}
//
//func RandomPrice() int32 {
//	return int32(RandomInt(100000, 1000000))
//}
