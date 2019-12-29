package transformers

import (
	"golang-starter/app/models"
	"golang-starter/config"
)

/**
* stander the single translation response
 */
func TranslationResponse(translation models.Translation) map[string]interface{} {
	var (
		u    = make(map[string]interface{})
		page models.Page
	)

	config.DB.Where("id = ?", translation.PageId).First(&page)

	u["value"] = translation.Value
	u["id"] = translation.ID
	u["slug"] = translation.Slug
	u["lang"] = translation.Lang
	u["page"] = page.Name

	return u
}

/**
* stander the Multi translations response
 */
func TranslationsResponse(translations []models.Translation) []map[string]interface{} {
	var u = make([]map[string]interface{}, 0)
	for _, translation := range translations {
		u = append(u, TranslationResponse(translation))
	}

	return u
}
