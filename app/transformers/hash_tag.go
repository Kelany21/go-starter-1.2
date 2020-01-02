package transformers

import "golang-starter/app/models"

/**
* stander the single faq response
 */
func HashTagResponse(hash_tag models.HashTag) map[string]interface{} {
	var u = make(map[string]interface{})

	u["id"] = hash_tag.ID
	u["hash_tag"] = hash_tag.HashTag
	u["good_count"] = hash_tag.GoodCount
	u["bad_count"] = hash_tag.BadCount
	u["tweet_count"] = hash_tag.TweetCount
	u["words_count"] = hash_tag.WordsCount
	u["neutral_count"] = hash_tag.NeutralCount
	u["sets_count"] = hash_tag.SetsCount
	u["bad_word_count"] = hash_tag.BadWordCount
	u["good_word_count"] = hash_tag.GoodWordCount
	u["neutral_word_count"] = hash_tag.NeutralWordCount
	u["user_id"] = hash_tag.UserId
	u["who_add"] = hash_tag.WhoAdd
	u["latest_sync"] = hash_tag.LatestSync
	u["count_move_good"] = hash_tag.CountMoveGood
	u["count_move_bad"] = hash_tag.CountMoveBad
	u["count_move_neutral"] = hash_tag.CountMoveNeutral
	u["stream"] = hash_tag.Stream
	u["action_id"] = hash_tag.ActionId
	u["created_at"] = hash_tag.CreatedAt
	u["updated_at"] = hash_tag.UpdatedAt

	return u
}

/**
* stander the Multi faqs response
 */
func HashTagsResponse(hash_tags []models.HashTag) []map[string]interface{} {
	var u = make([]map[string]interface{}, 0)
	for _, hash_tag := range hash_tags {
		u = append(u, HashTagResponse(hash_tag))
	}
	return u
}