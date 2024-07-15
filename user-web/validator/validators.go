package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	//使用正则表达式判断是否合法
	ok, _ := regexp.Match(`^(\\+?86)? (1[3-9]\\d{9})$`, []byte(mobile))
	if !ok {
		return false
	}
	return true
}
