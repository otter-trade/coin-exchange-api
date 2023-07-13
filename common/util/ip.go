package util

import (
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"strings"
)

type IpRegionDb struct {
	Path string
}

func IsChinaMainland(path, ip string) bool {
	db, err := xdb.NewWithFileOnly(path)
	if err != nil {
		return false
	}
	defer db.Close()

	region, err := db.SearchByStr(ip)
	res := false
	if strings.Contains(region, "中国") {
		res = true
	}
	if strings.Contains(region, "台湾省") {
		res = false
	}

	return res
}

func Region(path, ip string) (region string) {
	db, err := xdb.NewWithFileOnly(path)
	if err != nil {
		return ""
	}
	defer db.Close()
	region, err = db.SearchByStr(ip)
	return region
}
