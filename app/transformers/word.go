package transformers

import "golang-starter/app/models"

func WordResponse(word models.Word) map[string]interface{} {
	var u = make(map[string]interface{})
	u["id"] = word.ID
	u["type"] = word.Type
	u["word"] = word.Word
	u["source_id"] = word.SourceId
	u["source_type"] = word.SourceType
	u["created_at"] = word.CreatedAt
	u["updated_at"] = word.UpdatedAt
	return u
}

/**
* stander the Multi users response
 */
func WordsResponse(words []models.Word) []map[string]interface{} {
	var u  = make([]map[string]interface{} , 0)
	for _ , word := range words {
		u = append(u , WordResponse(word))
	}
	return u
}
