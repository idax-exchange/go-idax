package idax

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
)

//
// IDAX Tool Class
//

// Convert the structure to key = Val &amp; key1 = val1... and add sign parameters
func AddSignToJsonStr(obj interface{}, key string, secret string) string {
	/// Structures to Map
	toMap := StructToMap(obj)
	// Determine whether `key'exists in Map`
	_, existKey := toMap["key"]
	if existKey {
		toMap["key"] = key
	}
	// Determine whether `sign'exists in Map`
	_, existSign := toMap["sign"]
	if existSign {
		// Delete `sign'in Map`
		delete(toMap, "sign")
		// Add a new `sign'`
		toMap["sign"], _ = GetSign(secret, toMap)
	}
	bytes, _ := json.Marshal(toMap)
	fmt.Println("addSignToJsonStr:", string(bytes))
	return string(bytes)
}

// Map To UrlCode and then sha256 encryption based on secret
func GetSign(secret string, params map[string]interface{}) (string, error) {
	payload := ToUrlParam(params)
	sign := hmac.New(sha256.New, []byte(secret))
	sign.Write([]byte(payload))
	return hex.EncodeToString(sign.Sum(nil)), nil
}

// Convert Map to key = Val &amp; key1 = val1...
func ToUrlParam(params map[string]interface{}) string {
	// Initialize all Key lists
	var keys []string
	// Loop to get all keys
	for k := range params {
		keys = append(keys, k)
	}
	// keys sort
	sort.Strings(keys)
	// Initialize the return structure
	var kvs []string
	// Cyclically ordered keys
	for _, k := range keys {
		// The format is: key = Val &amp; key1 = val1...
		kvs = append(kvs, fmt.Sprintf("%s=%v", k, params[k]))
	}
	return strings.Join(kvs, "&")
}

// Structures to Map
func StructToMap(obj interface{}) map[string]interface{} {
	b, err := json.Marshal(obj)
	if err != nil {
		log.Println("json format error:", b)
		return nil
	}
	paramMapIface := make(map[string]interface{})
	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()
	err = d.Decode(&paramMapIface)
	return paramMapIface
}
