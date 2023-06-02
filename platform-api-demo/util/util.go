package util

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strings"
)

// Calculate OpenPlatform signature
func GetSign(params map[string]string, apiSecret string) string {
	keys := getKeys(params)
	sort.Strings(keys)

	var build strings.Builder

	for _, v := range keys {
		 kValue := params[v]
		 build.WriteString(v)
		 build.WriteString(kValue)
	}
	build.WriteString(apiSecret)

	preMessage := build.String()
	return md5Sign([]byte(preMessage))
}

// get all keys of map
func getKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// md5 encryption
func md5Sign(message []byte) string {
	ctx := md5.New()
	ctx.Write(message)
	return hex.EncodeToString(ctx.Sum(nil))
}
