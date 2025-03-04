package services

import (
	"github.com/wushengyouya/blog-service/internal/model"
	"github.com/wushengyouya/blog-service/pkg/app"
)

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TageListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100" json:"name"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100" json:"created_by"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1" json:"state"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" json:"id" binding:"required,gte=1"`
	Name       string `form:"name" json:"name"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1" json:"state"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100" json:"modified_by"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1" json:"id"`
}

func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *TageListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}
func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleleTag(param.ID)
}
