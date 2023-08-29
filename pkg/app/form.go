package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var resErrs ValidErrors
	err := c.ShouldBindJSON(v)
	if err != nil {
		return false, ValidErrors{&ValidError{Message: err.Error()}}
	}
	// 创建通用的 Translator
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	// 注册翻译器到验证器
	err = zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return false, resErrs
	}
	err = validate.Struct(v)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, value := range errs {
			// can translate each error one at a time.
			//return false, e.Translate(trans)
			resErrs = append(resErrs, &ValidError{
				Key:     value.Field(),
				Message: value.Translate(trans),
			})
		}
		return false, resErrs
	}
	return true, nil

}
