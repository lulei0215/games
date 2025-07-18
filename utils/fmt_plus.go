package utils

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: StructToMap
//@description: map
//@param: obj interface{}
//@return: map[string]interface{}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ArrayToString
//@description:
//@param: array []interface{}
//@return: string

func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

func Pointer[T any](in T) (out *T) {
	return &in
}

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// MaheHump
func MaheHump(s string) string {
	words := strings.Split(s, "-")

	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}

	return strings.Join(words, "")
}

// HumpToUnderscore
func HumpToUnderscore(s string) string {
	var result strings.Builder

	for i, char := range s {
		if i > 0 && char >= 'A' && char <= 'Z' {
			//
			result.WriteRune('_')
			result.WriteRune(char - 'A' + 'a') //
		} else {
			result.WriteRune(char)
		}
	}

	return strings.ToLower(result.String())
}

// RandomString
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[RandomInt(0, len(letters))]
	}
	return string(b)
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// BuildTree
func BuildTree[T common.TreeNode[T]](nodes []T) []T {
	nodeMap := make(map[int]T)
	// map
	for i := range nodes {
		nodeMap[nodes[i].GetID()] = nodes[i]
	}

	for i := range nodes {
		if nodes[i].GetParentID() != 0 {
			parent := nodeMap[nodes[i].GetParentID()]
			parent.SetChildren(nodes[i])
		}
	}

	var rootNodes []T

	for i := range nodeMap {
		if nodeMap[i].GetParentID() == 0 {
			rootNodes = append(rootNodes, nodeMap[i])
		}
	}
	return rootNodes
}
