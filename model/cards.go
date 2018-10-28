package model

import null "gopkg.in/guregu/null.v3"

type Card struct {
	ID        int64       `db:"id" json:"id"`
	ImageUrl  null.String `db:"image_url" json:"image_url"`
	Caption   null.String `db:"caption" json:"caption"`
	EmoID     int         `db:"emo_id" json:"emo_id"`
	ArticleId null.Int    `db:"article_id" json:"article_id"`
	Highlight null.String `db:"highlight" json:"highlight"`
	Order     null.Int    `db:"order_" json:"order"`
}
