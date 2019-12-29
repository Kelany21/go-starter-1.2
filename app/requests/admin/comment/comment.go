package comment

import (
	"golang-starter/app/models"
	"golang-starter/helpers"
	"net/http"

	"github.com/thedevsaddam/govalidator"
)

func Store(request *http.Request, data *models.Comment) *govalidator.Validator {
	lang := helpers.GetCurrentLangFromHttp(request)

	rules := govalidator.MapData{
		"post_id": []string{"required", "numeric"},
		"comment": []string{"required", "min:5", "max:255"},
		"status":  []string{"required", "between:1,2"},
	}

	messages := govalidator.MapData{
		"post_id": []string{helpers.Required(lang), helpers.Numeric(lang)},
		"comment": []string{helpers.Required(lang), helpers.Min(lang, "5"), helpers.Max(lang, "255")},
		"status":  []string{"required", "between:1,2"},
	}

	opts := govalidator.Options{
		Request:         request,
		Rules:           rules,
		Data:            data,
		Messages:        messages,
		RequiredDefault: true,
	}

	return govalidator.New(opts)
}

func Update(request *http.Request, data *models.Comment) *govalidator.Validator {
	lang := helpers.GetCurrentLangFromHttp(request)

	rules := govalidator.MapData{
		"comment": []string{"required", "min:5", "max:255"},
		"status":  []string{"required", "between:1,2"},
	}

	messages := govalidator.MapData{
		"comment": []string{helpers.Required(lang), helpers.Min(lang, "5"), helpers.Max(lang, "255")},
		"status":  []string{"required", "between:1,2"},
	}

	opts := govalidator.Options{
		Request:         request,
		Rules:           rules,
		Data:            data,
		Messages:        messages,
		RequiredDefault: true,
	}

	return govalidator.New(opts)
}
