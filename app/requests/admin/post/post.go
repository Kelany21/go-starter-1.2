package post

import (
	"golang-starter/app/models"
	"golang-starter/helpers"
	"net/http"

	"github.com/thedevsaddam/govalidator"
)

func Store(request *http.Request, data *models.Post) *govalidator.Validator {
	lang := helpers.GetCurrentLangFromHttp(request)

	rules := govalidator.MapData{
		"title": []string{"required", "min:6", "max:50"},
		"body":  []string{"required", "min:15"},
		"image": []string{"required"},
		//"dimensions:max_width=650,max_height=450"
		"category_id": []string{"required", "numeric"},
		"status":      []string{"required", "between:1,2"},
	}
	messages := govalidator.MapData{
		"title": []string{helpers.Required(lang), helpers.Min(lang, "6"), helpers.Max(lang, "50")},
		"body":  []string{helpers.Required(lang), helpers.Min(lang, "15")},
		"image": []string{helpers.Required(lang)},
		//helpers.Dimensions(lang, "650", "450")
		"category_id": []string{helpers.Required(lang), helpers.Numeric(lang)},
		"status":      []string{helpers.Required(lang), helpers.Between(lang, "1,2")},
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

func Update(request *http.Request, data *models.Post) *govalidator.Validator {
	lang := helpers.GetCurrentLangFromHttp(request)

	rules := govalidator.MapData{
		"title": []string{"required", "min:6", "max:50"},
		"body":  []string{"required", "min:15"},
		// "image":  []string{"dimensions:max_width=650,max_height=450"},
		"status": []string{"required", "between:1,2"},
	}
	messages := govalidator.MapData{
		"title": []string{helpers.Required(lang), helpers.Min(lang, "6"), helpers.Max(lang, "50")},
		"body":  []string{helpers.Required(lang), helpers.Min(lang, "15")},
		// "image":  []string{helpers.Dimensions(lang, "650", "450")},
		"status": []string{helpers.Required(lang), helpers.Between(lang, "1,2")},
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
