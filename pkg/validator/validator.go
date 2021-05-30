package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	"reflect"
	"strings"
)

type CustomValidator struct {
	Validator *validator.Validate
	uni       *ut.UniversalTranslator
	trans     ut.Translator
}

func NewValidate() (*CustomValidator, error) {
	eng := en.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")

	validate := validator.New()
	err := entranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return nil, err
	}

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &CustomValidator{Validator: validate, uni: uni, trans: trans}, nil

}

type ErrorMessage struct {
	Error            int    `json:"code"`
	ErrorDescription string `json:"msg,omitempty"`
}

func (cv *CustomValidator) FormErrorMessage(err error) ErrorMessage {
	var (
		code = 1
		msg  string
	)
	switch err.(type) {
	case validator.ValidationErrors:
		for _, v := range err.(validator.ValidationErrors) {
			msg = v.Translate(cv.trans)
			break
		}
	default:
		msg = err.Error()
	}
	return ErrorMessage{Error: code, ErrorDescription: msg}
}
