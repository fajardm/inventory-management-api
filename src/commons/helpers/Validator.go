package helpers

import (
	"gopkg.in/go-playground/validator.v9"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/locales/en"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

var validate *validator.Validate
var uni *ut.UniversalTranslator
var trans ut.Translator

func SetValidator() *validator.Validate {
	validate = validator.New()

	setTrans()

	return validate
}

func GetValidator() *validator.Validate {
	return validate
}

func GetValidationMessage(err error) map[string]string {
	errs := err.(validator.ValidationErrors)

	result := make(map[string]string)

	for _, e := range errs {
		result[e.Field()] = e.Translate(trans)
	}

	return result
}

func setTrans() {
	en := en.New()
	uni = ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)
}
