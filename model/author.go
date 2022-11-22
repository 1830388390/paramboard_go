package model

import (
	"time"
)

type Author struct {
	ID          *int64     `json:"id" gorm:"column:id" form:"id"`
	AuthorToken *string    `json:"author_token" gorm:"column:author_token" form:"author_token"`
	ModelName   *string    `json:"model_name" gorm:"column:model_name" form:"model_name"`
	Describe    *string    `json:"describe" gorm:"column:describe" form:"describe"`
	LabelA1Name *string    `json:"label_a1_name" gorm:"column:label_a1_name" form:"label_a1_name"`
	LabelA2Name *string    `json:"label_a2_name" gorm:"column:label_a2_name" form:"label_a2_name"`
	LabelA3Name *string    `json:"label_a3_name" gorm:"column:label_a3_name" form:"label_a3_name"`
	LabelA4Name *string    `json:"label_a4_name" gorm:"column:label_a4_name" form:"label_a4_name"`
	LabelA5Name *string    `json:"label_a5_name" gorm:"column:label_a5_name" form:"label_a5_name"`
	LabelA6Name *string    `json:"label_a6_name" gorm:"column:label_a6_name" form:"label_a6_name"`
	LabelB1Name *string    `json:"label_b1_name" gorm:"column:label_b1_name" form:"label_b1_name"`
	LabelB2Name *string    `json:"label_b2_name" gorm:"column:label_b2_name" form:"label_b2_name"`
	LabelB3Name *string    `json:"label_b3_name" gorm:"column:label_b3_name" form:"label_b3_name"`
	LabelB4Name *string    `json:"label_b4_name" gorm:"column:label_b4_name" form:"label_b4_name"`
	LabelB5Name *string    `json:"label_b5_name" gorm:"column:label_b5_name" form:"label_b5_name"`
	LabelB6Name *string    `json:"label_b6_name" gorm:"column:label_b6_name" form:"label_b6_name"`
	CreateTime  *time.Time `json:"create_time" gorm:"column:create_time" form:"create_time"`
}

func (m *Author) TableName() string {
	return "t_author"
}

func (m *Author) SetAuthorToken(authorToken *string) {
	m.AuthorToken = authorToken
	return
}
