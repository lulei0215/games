package utils

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Rules map[string][]string

type RulesMap map[string]Rules

var CustomizeMap = make(map[string]Rules)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: RegisterRule
//@description:
//@param: key string, rule Rules
//@return: err error

func RegisterRule(key string, rule Rules) (err error) {
	if CustomizeMap[key] != nil {
		return errors.New(key + ",")
	} else {
		CustomizeMap[key] = rule
		return nil
	}
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: NotEmpty
//@description:  0
//@return: string

func NotEmpty() string {
	return "notEmpty"
}

// @author: [zooqkl](https://github.com/zooqkl)
// @function: RegexpMatch
// @description:
// @param:  rule string
// @return: string

func RegexpMatch(rule string) string {
	return "regexp=" + rule
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Lt
//@description: (<) string array Slice  int uint float
//@param: mark string
//@return: string

func Lt(mark string) string {
	return "lt=" + mark
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Le
//@description: (<=) string array Slice  int uint float
//@param: mark string
//@return: string

func Le(mark string) string {
	return "le=" + mark
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Eq
//@description: (==) string array Slice  int uint float
//@param: mark string
//@return: string

func Eq(mark string) string {
	return "eq=" + mark
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Ne
//@description: (!=)  string array Slice  int uint float
//@param: mark string
//@return: string

func Ne(mark string) string {
	return "ne=" + mark
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Ge
//@description: (>=) string array Slice  int uint float
//@param: mark string
//@return: string

func Ge(mark string) string {
	return "ge=" + mark
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Gt
//@description: (>) string array Slice  int uint float
//@param: mark string
//@return: string

func Gt(mark string) string {
	return "gt=" + mark
}

//
//@author: [piexlmax](https://github.com/piexlmax)
//@function: Verify
//@description:
//@param: st interface{}, roleMap Rules(，map)
//@return: err error

func Verify(st interface{}, roleMap Rules) (err error) {
	compareMap := map[string]bool{
		"lt": true,
		"le": true,
		"eq": true,
		"ne": true,
		"ge": true,
		"gt": true,
	}

	typ := reflect.TypeOf(st)
	val := reflect.ValueOf(st) // reflect.Type

	kd := val.Kind() // st
	if kd != reflect.Struct {
		return errors.New("expect struct")
	}
	num := val.NumField()
	//
	for i := 0; i < num; i++ {
		tagVal := typ.Field(i)
		val := val.Field(i)
		if tagVal.Type.Kind() == reflect.Struct {
			if err = Verify(val.Interface(), roleMap); err != nil {
				return err
			}
		}
		if len(roleMap[tagVal.Name]) > 0 {
			for _, v := range roleMap[tagVal.Name] {
				switch {
				case v == "notEmpty":
					if isBlank(val) {
						return errors.New(tagVal.Name + "")
					}
				case strings.Split(v, "=")[0] == "regexp":
					if !regexpMatch(strings.Split(v, "=")[1], val.String()) {
						return errors.New(tagVal.Name + "")
					}
				case compareMap[strings.Split(v, "=")[0]]:
					if !compareVerify(val, v) {
						return errors.New(tagVal.Name + "," + v)
					}
				}
			}
		}
	}
	return nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: compareVerify
//@description:
//@param: value reflect.Value, VerifyStr string
//@return: bool

func compareVerify(value reflect.Value, VerifyStr string) bool {
	switch value.Kind() {
	case reflect.String:
		return compare(len([]rune(value.String())), VerifyStr)
	case reflect.Slice, reflect.Array:
		return compare(value.Len(), VerifyStr)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return compare(value.Uint(), VerifyStr)
	case reflect.Float32, reflect.Float64:
		return compare(value.Float(), VerifyStr)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compare(value.Int(), VerifyStr)
	default:
		return false
	}
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: isBlank
//@description:
//@param: value reflect.Value
//@return: bool

func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String, reflect.Slice:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: compare
//@description:
//@param: value interface{}, VerifyStr string
//@return: bool

func compare(value interface{}, VerifyStr string) bool {
	VerifyStrArr := strings.Split(VerifyStr, "=")
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		VInt, VErr := strconv.ParseInt(VerifyStrArr[1], 10, 64)
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Int() < VInt
		case VerifyStrArr[0] == "le":
			return val.Int() <= VInt
		case VerifyStrArr[0] == "eq":
			return val.Int() == VInt
		case VerifyStrArr[0] == "ne":
			return val.Int() != VInt
		case VerifyStrArr[0] == "ge":
			return val.Int() >= VInt
		case VerifyStrArr[0] == "gt":
			return val.Int() > VInt
		default:
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		VInt, VErr := strconv.Atoi(VerifyStrArr[1])
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Uint() < uint64(VInt)
		case VerifyStrArr[0] == "le":
			return val.Uint() <= uint64(VInt)
		case VerifyStrArr[0] == "eq":
			return val.Uint() == uint64(VInt)
		case VerifyStrArr[0] == "ne":
			return val.Uint() != uint64(VInt)
		case VerifyStrArr[0] == "ge":
			return val.Uint() >= uint64(VInt)
		case VerifyStrArr[0] == "gt":
			return val.Uint() > uint64(VInt)
		default:
			return false
		}
	case reflect.Float32, reflect.Float64:
		VFloat, VErr := strconv.ParseFloat(VerifyStrArr[1], 64)
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Float() < VFloat
		case VerifyStrArr[0] == "le":
			return val.Float() <= VFloat
		case VerifyStrArr[0] == "eq":
			return val.Float() == VFloat
		case VerifyStrArr[0] == "ne":
			return val.Float() != VFloat
		case VerifyStrArr[0] == "ge":
			return val.Float() >= VFloat
		case VerifyStrArr[0] == "gt":
			return val.Float() > VFloat
		default:
			return false
		}
	default:
		return false
	}
}

func regexpMatch(rule, matchStr string) bool {
	return regexp.MustCompile(rule).MatchString(matchStr)
}
