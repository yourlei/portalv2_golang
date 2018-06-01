package common

import (
	"fmt"
	"regexp"
	"reflect"
	"strings"
)

var (
	MailReg = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)
	Alpha          string = "^[a-zA-Z]+$"
	Alphanumeric   string = "^[a-zA-Z0-9]+$"
	Numeric        string = "^[0-9]+$"
)

func IsEmail(s string) (bool, error) {
	if !MailReg.MatchString(s) {
		return false, fmt.Errorf("邮箱格式不正确")
	}
	return true, nil
}

func IsNumber(n int) bool {
	return true
}

func IsString(s string) bool {
	// if reflect.TypeOf(s) != "string" {
	// 	return false
	// }
	return true
}

const tagName = "validate"

//邮箱验证正则
var mailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

//验证接口
type Validator interface {
    Validate(interface{}) (bool, error)
}

type DefaultValidator struct {
}

func (v DefaultValidator) Validate(val interface{}) (bool, error) {
    return true, nil
}

type StringValidator struct {
    Min int
    Max int
}

func (v StringValidator) Validate(val interface{}) (bool, error) {
    l := len(val.(string))

    if l == 0 {
        return false, fmt.Errorf("cannot be blank")
    }

    if l < v.Min {
        return false, fmt.Errorf("should be at least %v chars long", v.Min)
    }

    if v.Max >= v.Min && l > v.Max {
        return false, fmt.Errorf("should be less than %v chars long", v.Max)
    }

    return true, nil
}

type NumberValidator struct {
    Min int
    Max int
}

func (v NumberValidator) Validate(val interface{}) (bool, error) {
    num := val.(int)

    if num < v.Min {
        return false, fmt.Errorf("should be greater than %v", v.Min)
    }

    if v.Max >= v.Min && num > v.Max {
        return false, fmt.Errorf("should be less than %v", v.Max)
    }

    return true, nil
}

type EmailValidator struct {
}

func (v EmailValidator) Validate(val interface{}) (bool, error) {
    if !mailRe.MatchString(val.(string)) {
        return false, fmt.Errorf("is not a valid email address")
    }
    return true, nil
}

func getValidatorFromTag(tag string) Validator {
    args := strings.Split(tag, ",")

    switch args[0] {
    case "number":
        validator := NumberValidator{}
        //将structTag中的min和max解析到结构体中
        fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
        return validator
    case "string":
        validator := StringValidator{}
        fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
        return validator
    case "email":
        return EmailValidator{}
    }

    return DefaultValidator{}
}

func ValidateStruct(s interface{}) []error {
    errs := []error{}

    v := reflect.ValueOf(s)

    for i := 0; i < v.NumField(); i++ {
        //利用反射获取structTag
        tag := v.Type().Field(i).Tag.Get(tagName)

        if tag == "" || tag == "-" {
            continue
        }

        validator := getValidatorFromTag(tag)

        valid, err := validator.Validate(v.Field(i).Interface())
        if !valid && err != nil {
            errs = append(errs, fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error()))
        }
    }

    return errs
}
