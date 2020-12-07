package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

var (
	uni   *ut.UniversalTranslator
	trans ut.Translator
)

func init() {
	translator := zh.New()
	uni = ut.New(translator, translator)
	trans, _ = uni.GetTranslator("zh")
	validate := binding.Validator.Engine().(*validator.Validate)
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("display")
	})
	_ = zh_translations.RegisterDefaultTranslations(validate, trans)
}

func Translate(err error) string {
	var result string

	errors := err.(validator.ValidationErrors)

	for _, err := range errors {
		errMessage := err.Translate(trans)
		result += errMessage + ";"
	}
	return result[:len(result)-1]
}
