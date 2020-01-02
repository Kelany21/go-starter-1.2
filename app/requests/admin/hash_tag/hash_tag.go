package hash_tag

import (
	"github.com/thedevsaddam/govalidator"
	"golang-starter/app/models"
	"golang-starter/helpers"
	"net/http"
)

func StoreUpdate(r *http.Request, request *models.HashTag) *govalidator.Validator {
	lang := helpers.GetCurrentLangFromHttp(r)
	/// Validation rules
	rules := govalidator.MapData{
		"hash_tag": []string{"required", "min:6", "max:225"},
	}

	messages := govalidator.MapData{
		"hash_tag": []string{helpers.Required(lang), helpers.Min(lang, "6"), helpers.Max(lang, "225")},
	}

	opts := govalidator.Options{
		Request:         r,     // request object
		Rules:           rules, // rules map
		Data:            request,
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true,     // all the field to be pass the rules
	}
	return govalidator.New(opts)
}