package model

type ArticleTag struct {
	*Model
	TagID uint32 `json:"tag_id"`
	ArticleId uint32 `json:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "golb_article_tag"
}
