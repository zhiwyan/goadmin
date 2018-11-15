/*
	some common api for change string between int/32/64,uint8/16/32/64,float32/64
*/
package common

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func Atoi(s string) int {
	if s == "" {
		return 0
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func Atouint(s string) uint {
	return uint(Atoi(s))
}

func Atoi64(s string) int64 {
	if s == "" {
		return 0
	}

	i, err := strconv.ParseInt(s, 10, 0)

	if err != nil {
		return 0
	}
	return i
}

func Atoui64(s string) uint64 {
	return (uint64)(Atoi64(s))
}

func Atof32(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0
	}
	return float32(f)
}

func Atof64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

func Atoi32(s string) int32 {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return int32(i)
}

func Atobyte(s string) byte {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return byte(i)
}

func Itoa(n int) string {
	s := strconv.Itoa(n)
	return s
}

func I32toa(n int32) string {
	return fmt.Sprintf("%d", n)
}

func I64toa(i int64) (s string) {
	s = fmt.Sprintf("%d", i)
	return s
}

func U64toa(i uint64) (s string) {
	s = fmt.Sprintf("%d", i)
	return s
}

func String2IntArray(s string) (ar []int) {
	for _, vId := range strings.Split(s, ",") {
		ar = append(ar, Atoi(vId))
	}
	return ar
}

func IntArray2String(ar []int) string {
	vs := []string{}
	for _, a := range ar {
		vs = append(vs, I64toa(int64(a)))
	}
	return strings.Join(vs, ",")
}

func String2Int32Array(s string) (ar []int32) {
	for _, vId := range strings.Split(s, ",") {
		ar = append(ar, Atoi32(vId))
	}
	return ar
}

func Int32Array2String(ar []int32) string {
	vs := []string{}
	for _, a := range ar {
		vs = append(vs, I64toa(int64(a)))
	}
	return strings.Join(vs, ",")
}

func String2Int64Array(s string) (ar []int64) {
	for _, vId := range strings.Split(s, ",") {
		ar = append(ar, Atoi64(vId))
	}
	return ar
}

func StringArray2Int64Array(ss []string) (ar []int64) {
	ar = make([]int64, len(ss))
	for index, s := range ss {
		ar[index] = Atoi64(s)
	}
	return
}

func Int64Array2String(ar []int64) string {
	vs := []string{}
	for _, a := range ar {
		vs = append(vs, I64toa(a))
	}
	return strings.Join(vs, ",")
}

// Int2Fuzzy 3 -> (?,?,?)
func Int2Fuzzy(ir int) string {
	fuzAr := make([]string, ir)
	for i := 0; i < ir; i++ {
		fuzAr[i] = "?"
	}
	return "(" + strings.Join(fuzAr, ",") + ")"
}

func AInt2AInt64(ar []int) (br []int64) {
	br = make([]int64, len(ar))
	for index, r := range ar {
		br[index] = int64(r)
	}
	return
}

func UnderScore2Calm(s string) string {
	words := strings.Split(s, "_")
	var newWords []string

	for _, word := range words {
		a := []rune(word)
		a[0] = unicode.ToUpper(a[0])
		newWords = append(newWords, string(a))
	}

	return strings.Join(newWords, "")
}

func Calm2UnderScore(s string) string {
	var words []string
	for i := 0; s != ""; s = s[i:] {
		i = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
		if i <= 0 {
			i = len(s)
		}
		words = append(words, strings.ToLower(s[:i]))
	}
	return strings.Join(words, "_")
}

func Atob(str string) bool {
	if str == "1" {
		return true
	}

	return false
}

func Atoas(str string) []string {
	return strings.Split(str, "|")
}

func Atoi64s(str string) []int64 {
	as := strings.Split(str, "|")
	is := make([]int64, len(as))

	for i, v := range as {
		is[i] = Atoi64(v)
	}

	return is
}

func DownFirstChar(s string) string {
	cs := []rune(s)
	cs[0] = unicode.ToLower(cs[0])
	return string(cs)
}

func Singular(str string) string {
	if strings.HasSuffix(str, "es") {
		return str[0 : len(str)-2]
	}

	if strings.HasSuffix(str, "s") {
		return str[0 : len(str)-1]
	}

	return str
}

func GetCurrentUinxTimeStamp() string {
	ts := time.Now().Unix()
	stamp := fmt.Sprint(ts)
	return stamp
}

func Inet_addr(ipaddr string) uint32 {
	var (
		segments []string = strings.Split(ipaddr, ".")
		ip       [4]uint64
		ret      uint64
	)

	if len(segments) != 4 {
		return 0
	}

	for i := 0; i < 4; i++ {
		ip[i], _ = strconv.ParseUint(segments[i], 10, 64)
	}
	ret = ip[3]<<24 + ip[2]<<16 + ip[1]<<8 + ip[0]
	return uint32(ret)
}

const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func GetRandomString(n uint32) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = alphanum[rand.Intn(len(alphanum))]
	}
	return string(b)
}

func IsValidString(source string) bool {
	return source != "" && len(strings.TrimSpace(source)) > 0
}

func HideTeacherName(name string) string {
	if !IsValidString(name) {
		return ""
	}

	nameRune := []rune(name)
	return string(nameRune[:1]) + "老师"
}
