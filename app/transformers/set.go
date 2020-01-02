package transformers

import "golang-starter/app/models"

func SetResponse(set models.Set) map[string]interface{} {
	var u = make(map[string]interface{})
	u["id"] = set.ID
	u["name"] = set.Name
	u["action_id"] = set.ActionId
	u["trashed_by"] = set.TrashedBy
	u["use_count"] = set.UseCount
	u["used_number"] = set.UsedNumber
	u["user_id"] = set.UserId
	u["who_add"] = set.WhoAdd
	u["words_count"] = set.WordsCount
	u["created_at"] = set.CreatedAt
	u["updated_at"] = set.UpdatedAt

	return u
}

/**
* stander the Multi users response
 */
func SetsResponse(sets []models.Set) []map[string]interface{} {
	var u  = make([]map[string]interface{} , 0)
	for _ , set := range sets {
		u = append(u , SetResponse(set))
	}
	return u
}

