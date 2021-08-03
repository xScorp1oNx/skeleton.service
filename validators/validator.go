package validators

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/en"
	"log"
	"reflect"
	"strings"
)

type StructValidator struct {
	val   *validator.Validate
	trans ut.Translator
}

func (val *StructValidator) Initialize() {
	v := validator.New()
	translator := en.New()

	uni := ut.New(translator, translator)

	trans, _ := uni.GetTranslator("en")

	if err := translations.RegisterDefaultTranslations(v, trans); err != nil {
		log.Fatal(err)
	}

	val.val = v
	val.trans = trans

	val.registerTranslations()
	val.registerTagNameFunctions()
}

func (val *StructValidator) registerTranslations() {
	_ = val.GetValidator().RegisterTranslation("required", val.GetTranslator(), func(ut ut.Translator) error {
		return ut.Add("required", "{0}_is_required", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	_ = val.GetValidator().RegisterTranslation("min", val.GetTranslator(), func(ut ut.Translator) error {
		return ut.Add("min", "{0}_too_short", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field())
		return t
	})

	_ = val.GetValidator().RegisterTranslation("max", val.GetTranslator(), func(ut ut.Translator) error {
		return ut.Add("max", "{0}_too_long", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field())
		return t
	})
}

func (val *StructValidator) registerTagNameFunctions() {
	val.GetValidator().RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func (val *StructValidator) GetValidator() *validator.Validate {
	return val.val
}

func (val *StructValidator) GetTranslator() ut.Translator {
	return val.trans
}
