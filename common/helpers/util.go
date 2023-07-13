package helpers

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/alecthomas/log4go"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/nats-io/nats.go"
	"io"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// Config 配置
type Config struct {
	Cluster       string `xml:"cluster"`        // 集群名字,为了同一个nats-server各个集群下互相不影响
	Server        string `xml:"server"`         // nats://127.0.0.1:4222,nats://127.0.0.1:4223
	User          string `xml:"user"`           // 用户名
	Pwd           string `xml:"pwd"`            //密码
	ReconnectWait int64  `xml:"reconnect_wait"` // 重连间隔
	MaxReconnects int32  `xml:"max_reconnects"` // 重连次数
}

// SetupConnOptions 设置启动选项
func SetupConnOptions(name string, cfg *Config) []nats.Option {
	opts := make([]nats.Option, 0)
	opts = append(opts, nats.Name(name))
	if len(cfg.User) > 0 {
		opts = append(opts, nats.UserInfo(cfg.User, cfg.Pwd))
	}
	opts = append(opts, nats.ReconnectWait(time.Second*time.Duration(cfg.ReconnectWait)))
	opts = append(opts, nats.MaxReconnects(int(cfg.MaxReconnects)))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log4go.Warn("[main] nats.Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.DiscoveredServersHandler(func(nc *nats.Conn) {
		log4go.Info("[main] nats.DiscoveredServersHandler", nc.DiscoveredServers())
	}))
	opts = append(opts, nats.DisconnectHandler(func(nc *nats.Conn) {
		log4go.Warn("[main] nats.Disconnect")
	}))
	return opts
}

// joinSubject 把subject用.分割组合
func joinSubject(cluster string, typeName string, subjectPostfix ...interface{}) string {
	sub := fmt.Sprintf("%s.%s", cluster, typeName)
	if nil != subjectPostfix {
		for _, v := range subjectPostfix {
			sub += fmt.Sprintf(".%v", v)
		}
	}
	return sub
}

func JSONToMap(str string) (tempMap map[string]interface{}, err error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal([]byte(str), &tempMap)
	return
}

func GenerateSign(secret string, params map[string]interface{}) (sign string) {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}

	var excludeFields []string
	for _, value := range keys {
		if value == "excludeFields" {
			var temp = fmt.Sprintf("%v", params[value])
			excludeFields = strings.Split(temp, ",")
		}
	}

	sort.Strings(keys)
	paramBuffer := bytes.NewBufferString(secret)
	for _, item := range keys {
		if InArray(item, excludeFields) {
			continue
		}
		paramBuffer.WriteString(item)
		paramBuffer.WriteString(fmt.Sprintf("%v", params[item]))
	}

	sign = Md5(paramBuffer.String())
	return
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return (s[i]) < (s[j])
}

func SortStringSlice(args ...string) string {
	sort.Sort(ByLength(args))
	var str string
	for _, arg := range args {
		str += arg
	}
	return Md5(str)
}

func GetUUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

func ToFeng(in float64) (out float64) {
	return in * 100
}

func ToYuan(in float64) (out float64) {
	return in / 100
}

func GzipEncode(s string) ([]byte, error) {
	jsonStr, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	_, err = w.Write(jsonStr)
	if err != nil {
		w.Close()
		return nil, err
	}
	err = w.Flush()
	if err != nil {
		w.Close()
		return nil, err
	}
	w.Close()

	return b.Bytes(), nil
}

func GzipDecode(b []byte) (s string, err error) {
	r, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return "", err
	}
	jsonStr, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.Unmarshal(jsonStr, s); err != nil {
		return "", err
	}

	return s, nil
}
