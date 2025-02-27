package model

import "github.com/wushengyouya/blog-service/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

// swagger
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}
