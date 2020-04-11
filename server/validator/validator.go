package validator

import (
	"github.com/asaskevich/govalidator"
)

func Add() {
	govalidator.CustomTypeTagMap.Set("trim", func(i interface{}, _ interface{}) bool {
		if i == nil {
			return false
		}
		content := i.(string)
		if content == "" {
			return false
		}
		val := govalidator.Trim(content, " ")
		return val != ""
	})
}
