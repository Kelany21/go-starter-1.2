package transformers

import "golang-starter/app/models"

func ActionResponse(action models.Action) map[string]interface{} {
	var u = make(map[string]interface{})

	u["id"] = action.ID
	u["title"] = action.Title
	u["count"] = action.Count
	u["module_name"] = action.ModuleName

	return u
}

/**
* stander the Multi Answers response
 */
func ActionsResponse(actions []models.Action) []map[string]interface{} {
	var u  = make([]map[string]interface{} , 0)

	for _ , action := range actions {
		u = append(u , ActionResponse(action))
	}

	return u
}
