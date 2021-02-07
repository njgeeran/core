package utils

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type Rules map[string][]string

var compareMap = map[string]bool{
	"lt": true,
	"le": true,
	"eq": true,
	"ne": true,
	"ge": true,
	"gt": true,
}

// 非空 不能为其对应类型的0值
func NotEmpty() string {
	return "notEmpty"
}
//验证手机号
func VerPhone() string {
	return "verPhone"
}
//验证邮箱
func VerEmail() string{
	return "verEmail"
}
func VerRegexp(exp string) string {
	return "verRegexp="+exp
}
// 小于入参(<) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
func Lt(mark string) string {
	return "lt=" + mark
}

// 小于等于入参(<=) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
func Le(mark string) string {
	return "le=" + mark
}

// 等于入参(==) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
func Eq(mark string) string {
	return "eq=" + mark
}

// 不等于入参(!=)  如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
func Ne(mark string) string {
	return "ne=" + mark
}

// 大于等于入参(>=) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
func Ge(mark string) string {
	return "ge=" + mark
}

// 大于入参(>) 如果为string array Slice则为长度比较 如果是 int uint float 则为数值比较
func Gt(mark string) string {
	return "gt=" + mark
}
// 校验方法 接收两个参数  入参实例，规则map
func Verify(roleMap Rules,st ...interface{}) (err error) {
	//typ := reflect.TypeOf(st)
	 // 获取reflect.Type类型

	//kd := val.Kind() // 获取到st对应的类别
	//if kd != reflect.Struct {
	//	return errors.New("expect struct")
	//}
	for _,t := range st {
		typ := reflect.TypeOf(t)
		val := reflect.ValueOf(t)
		if err := verify("",typ,val,roleMap);err != nil{
			return err
		}
	}
	return nil
}
func verifyStruct(name string,typ reflect.Type,val reflect.Value,roleMap Rules) error {
	num := val.NumField()
	// 遍历结构体的所有字段
	for i := 0; i < num; i++ {
		val_t := val.Field(i)
		typ_t := typ.Field(i)
		name := IF(name == "",typ_t.Name,name+"."+typ_t.Name).(string)
		if err := verify(name,typ_t.Type, val_t, roleMap); err != nil {
			return err
		}
	}
	return nil
}
func verifyArray(name string,typ reflect.Type,val reflect.Value,roleMap Rules) error {
	for i := 0; i < val.Len(); i++ {
		val_t := val.Index(i)
		if err := verify(name,val_t.Type(),val_t,roleMap);err != nil {
			return err
		}
	}
	return nil
}
func verify(name string,typ reflect.Type,val reflect.Value,roleMap Rules) error {
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		return verifyArray(name,typ,val,roleMap)
	case reflect.Struct:
		return verifyStruct(name,typ,val,roleMap)
	default:
		if len(roleMap[name]) > 0 {
			for _, v := range roleMap[name] {
				switch {
				case v == "notEmpty":
					if isBlank(val) {
						return errors.New(name + "值不能为空")
					}
				case v == "verPhone":
					phone := val.String()
					if !MatchPhone(phone) {
						return errors.New(name + "未通过验证：手机号格式错误")
					}
				case v == "verEmail":
					email := val.String()
					if !MatchEmail(email) {
						return  errors.New(name + "未通过验证：邮箱格式错误")
					}
				case strings.Split(v, "=")[0] == "verRegexp":
					if !verRegexp(val,v) {
						return errors.New(name + "正则校验失败," + v)
					}
				case compareMap[strings.Split(v, "=")[0]]:
					if !compareVerify(val, v) {
						return errors.New(name + "长度或值不在合法范围," + v)
					}
				}
			}
		}
	}
	return nil
}

// 长度和数字的校验方法 根据类型自动校验
func compareVerify(value reflect.Value, VerifyStr string) bool {
	switch value.Kind() {
	case reflect.String, reflect.Slice, reflect.Array:
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

// 非空校验
func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
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

func verRegexp(val reflect.Value, VerifyStr string) bool {
	vs := strings.Split(VerifyStr, "=")
	str := val.String()
	exp := vs[1]
	if !MatchString(str,exp) {
		return false
	}
	return true
}