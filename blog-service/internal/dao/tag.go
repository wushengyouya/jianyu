package dao

import (
	"github.com/wushengyouya/blog-service/internal/model"
	"github.com/wushengyouya/blog-service/pkg/app"
)

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	// pageOffset 跳过的条数
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createBy string) error {
	tag := model.Tag{Name: name, State: state, Model: &model.Model{CreatedBy: createBy}}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifedBy string) error {
	tag := model.Tag{Name: name, State: state, Model: &model.Model{ID: id, ModifiedBy: modifedBy}}
	return tag.Update(d.engine)
}

func (d *Dao) DeleleTag(id uint32) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Delele(d.engine)
}
