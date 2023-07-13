package util

import (
	"bytes"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/davecgh/go-spew/spew"
	"github.com/shopspring/decimal"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var exp = decimal.NewFromInt(1000000000000000000)

// 小数转18位字符串

func DecimalToFutureString(num decimal.Decimal) string {
	return num.Mul(exp).Truncate(0).String()
}

func DecimalToFuture(num decimal.Decimal) decimal.Decimal {
	return num.Mul(exp).Truncate(0)
}

func FutureStringToDecimal(num string) decimal.Decimal {
	dec, err := decimal.NewFromString(num)
	if err != nil {
		return decimal.Zero
	}
	return dec.DivRound(exp, 18)
}

func FutureDeciamlToDecimal(num decimal.Decimal) decimal.Decimal {
	return num.DivRound(exp, 18)
}

func DecimalAddString(num decimal.Decimal, str string) decimal.Decimal {
	strDecimal, _ := decimal.NewFromString(str)
	return num.Add(strDecimal)
}

func StringAdd(one, two string) decimal.Decimal {
	oneDecimal, _ := decimal.NewFromString(one)
	twoDecimal, _ := decimal.NewFromString(two)
	return oneDecimal.Add(twoDecimal)
}

func StringSub(one, two string) decimal.Decimal {
	oneDecimal, _ := decimal.NewFromString(one)
	twoDecimal, _ := decimal.NewFromString(two)
	return oneDecimal.Sub(twoDecimal)
}

// ClientIP 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func GetIpAddr() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	// 192.168.1.20:61085
	ip := strings.Split(localAddr.String(), ":")[0]
	return ip
}

func RemoteIp(req *http.Request) string {
	var remoteAddr string
	// RemoteAddr
	remoteAddr = req.RemoteAddr
	if remoteAddr != "" {
		return remoteAddr
	}
	// ipv4
	remoteAddr = req.Header.Get("ipv4")
	if remoteAddr != "" {
		return remoteAddr
	}
	//
	remoteAddr = req.Header.Get("XForwardedFor")
	if remoteAddr != "" {
		return remoteAddr
	}
	// X-Forwarded-For
	remoteAddr = req.Header.Get("X-Forwarded-For")
	if remoteAddr != "" {
		return remoteAddr
	}
	// X-Real-Ip
	remoteAddr = req.Header.Get("X-Real-Ip")
	if remoteAddr != "" {
		return remoteAddr
	} else {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}

func ConvertToViewTree(treeDatas []map[string]interface{}, labelField, valueField, keyField string) []map[string]interface{} {
	for _, node := range treeDatas {
		node["title"] = node[labelField]
		node["value"] = node[valueField]
		node["key"] = node[keyField]
		child, ok := node["children"]
		if ok {
			node["children"] = ConvertToViewTree(*child.(*[]map[string]interface{}), labelField, valueField, keyField)
		}
	}
	return treeDatas
}

func Slice2Tree(sliceDatas []map[string]interface{}, idField, pidField string) []map[string]interface{} {
	var r []map[string]interface{}
	index := make(map[string]interface{})

	for _, val := range sliceDatas {
		id := fmt.Sprint(val[idField])
		index[id] = val
	}

	for _, val := range sliceDatas {
		pid := fmt.Sprint(val[pidField])
		if _, ok := index[pid]; !ok || pid == "" {
			r = append(r, val)
		} else {
			pval := index[pid].(map[string]interface{})
			if _, ok := pval["children"]; !ok {
				var n []map[string]interface{}
				n = append(n, val)
				pval["children"] = &n
			} else {
				nodes := pval["children"].(*[]map[string]interface{})
				*nodes = append(*nodes, val)
			}
		}
	}
	return r
}

// 内嵌bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*****内存不够了！******")
		}
	}()
	b.WriteString(s)
	return b
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

// 驼峰式写法转为下划线写法
func Camel2Case(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

// 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// Return
//
//	@Description: if实现的三元表达式
//	@param boolExpression: 布尔表达式，最终返回一个布尔值
//	@param trueReturnValue: 当boolExpression返回值为true的时候返回的值
//	@param falseReturnValue: 当boolExpression返回值为false的时候返回的值
//	@return bool: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func Return[T any](boolExpression bool, trueReturnValue, falseReturnValue T) T {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// ReturnByFunc
//
//	@Description: if实现的三元表达式
//	@param boolExpression: 布尔表达式，最终返回一个布尔值
//	@param trueReturnValue: 当boolExpression返回值为true的时候执行此函数并返回值
//	@param falseReturnValue: 当boolExpression返回值为false的时候执行此函数并返回值
//	@return bool: 三元表达式的结果，为trueReturnValue或者falseReturnValue中的一个
func ReturnByFunc[T any](boolExpression bool, trueFuncForReturnValue, falseFuncForReturnValue func() T) T {
	if boolExpression {
		return trueFuncForReturnValue()
	} else {
		return falseFuncForReturnValue()
	}
}

func Dump[T any](params T) {
	spew.Dump(params)
}

// 处理乱码
// 参数1：处理的数据
// 参数2：数据目前的编码
// 参数3：返回的正常数据
// utfStr := ConvertEncoding(gbkStr, "GBK")
func ConvertEncoding(srcStr string, encoding string) (dstStr string) {
	// 创建编码处理器
	enc := mahonia.NewDecoder(encoding)
	// 编码器处理字符串为utf8的字符串
	utfStr := enc.ConvertString(srcStr)
	dstStr = utfStr
	return
}

func GetFileSize(filename string) (int64, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

func GenerateRandInt(min, max int) int {
	rand.Seed(time.Now().Unix()) //随机种子
	return rand.Intn(max-min) + min
}
