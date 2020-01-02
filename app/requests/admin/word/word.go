package word

import (
	"github.com/thedevsaddam/govalidator"
	"golang-starter/app/models"
	"golang-starter/helpers"
	"net/http"
)

/**
* validate update page request
 */
func StoreUpdate(r *http.Request, request *models.Word) *govalidator.Validator {
	lang := helpers.GetCurrentLangFromHttp(r)
	/// Validation rules
	rules := govalidator.MapData{
		"type": 	   []string{"required", "min:7", "max:8"},
		"word":   	   []string{"required", "min:4", "max:50"},
		"source_id":   []string{"required", "numeric"},
		"source_type": []string{"required", "in:hash_tags,sets"},
	}

	messages := govalidator.MapData{
		"type":   	   []string{helpers.Required(lang), helpers.Min(lang, "7"), helpers.Max(lang, "8")},
		"word":   	   []string{helpers.Required(lang), helpers.Min(lang, "4"), helpers.Max(lang, "50")},
		"source_id":   []string{helpers.Required(lang), helpers.Numeric(lang)},
		"source_type": []string{helpers.Required(lang), helpers.In(lang, "hash_tags", "sets")},
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
