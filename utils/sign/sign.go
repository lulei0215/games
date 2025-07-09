package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// SIGN_KEY signature key constant
const SIGN_KEY = "GAME_2025_SIGN_KEY_8F7E6D5C4B3A2918_9A8B7C6D5E4F3210"

// EncryptUtil encryption utility class
type EncryptUtil struct{}

// GenerateSign generate signature
// params: parameter object
// return: signature string (MD5 encrypted and converted to uppercase)
func (e *EncryptUtil) GenerateSign(params map[string]interface{}) string {
	// 1. Sort parameters by ASCII code
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 2. Concatenate parameter string
	var signStr strings.Builder
	for _, key := range keys {
		value := params[key]
		if value != nil && value != "" {
			// For complex objects, use [object Object] format (consistent with TypeScript)
			var valueStr string
			if key == "list" {
				// Generate corresponding number of [object Object] based on array length
				switch v := value.(type) {
				case []interface{}:
					objs := make([]string, len(v))
					for i := range v {
						objs[i] = "[object Object]"
					}
					valueStr = strings.Join(objs, ",")
				default:
					// For other types, use reflection to check if it's a slice
					val := reflect.ValueOf(v)
					if val.Kind() == reflect.Slice {
						objs := make([]string, val.Len())
						for i := range objs {
							objs[i] = "[object Object]"
						}
						valueStr = strings.Join(objs, ",")
					} else {
						valueStr = "[object Object]"
					}
				}
			} else {
				valueStr = fmt.Sprintf("%v", value)
			}
			signStr.WriteString(fmt.Sprintf("%s=%s&", key, valueStr))
		}
	}

	// 3. Add signature key
	signStr.WriteString(fmt.Sprintf("key=%s", SIGN_KEY))

	// 4. MD5 encryption and convert to uppercase
	hash := md5.New()
	hash.Write([]byte(signStr.String()))
	sign := strings.ToUpper(hex.EncodeToString(hash.Sum(nil)))
	return sign
}

// GenerateSignStatic static method to generate signature
func GenerateSign(params map[string]interface{}) string {
	util := &EncryptUtil{}
	return util.GenerateSign(params)
}

// VerifySign verify signature
// params: parameter object
// sign: signature to verify
// return: verification result
func (e *EncryptUtil) VerifySign(params map[string]interface{}, sign string) bool {
	// Copy parameters, remove sign field
	paramsWithoutSign := make(map[string]interface{})
	for key, value := range params {
		if key != "sign" {
			paramsWithoutSign[key] = value
		}
	}

	generatedSign := e.GenerateSign(paramsWithoutSign)
	isValid := strings.ToUpper(generatedSign) == strings.ToUpper(sign)

	return isValid
}

// VerifySignStatic static method to verify signature
func VerifySign(params map[string]interface{}, sign string) bool {
	util := &EncryptUtil{}
	return util.VerifySign(params, sign)
}

// GetCorrectSign get correct signature (for debugging)
func GetCorrectSign(params map[string]interface{}) string {
	util := &EncryptUtil{}
	return util.GenerateSign(params)
}

// MD5 MD5 encryption
// str: string to encrypt
// return: MD5 encrypted string
func (e *EncryptUtil) MD5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// MD5Static static method MD5 encryption
func MD5(str string) string {
	util := &EncryptUtil{}
	return util.MD5(str)
}

// MD5ToUpper MD5 encryption and convert to uppercase
// str: string to encrypt
// return: MD5 encrypted and uppercase converted string
func (e *EncryptUtil) MD5ToUpper(str string) string {
	return strings.ToUpper(e.MD5(str))
}

// MD5ToUpperStatic static method MD5 encryption and convert to uppercase
func MD5ToUpper(str string) string {
	util := &EncryptUtil{}
	return util.MD5ToUpper(str)
}
