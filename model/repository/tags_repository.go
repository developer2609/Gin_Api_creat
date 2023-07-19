package repository

import (
	"errors"

	"gilab.com/pragmaticreviews/golang-gin-poc/helper"
	"gilab.com/pragmaticreviews/golang-gin-poc/model"
	"gorm.io/gorm"
)

type TagsRepositoryImple struct {
	Db *gorm.DB
}

// Delete implements TagsRepository
func (*TagsRepositoryImple) Delete(tagsId int) {

}

// FindAll implements TagsRepository
func (t *TagsRepositoryImple) FindAll() []model.Tags {
	var tags []model.Tags
	result := t.Db.Find(&tags)
	helper.ErorPanic(result.Error)
	return tags

}

// FindById implements TagsRepository
func (t *TagsRepositoryImple) FindById(tagsId int) (tags model.Tags, err error) {
	var tag model.Tags

	result := t.Db.Find(&tag, tagsId)
	if result != nil {
		return tag, nil

	} else {
		return tag, errors.New("tags not found")
	}
}

// Save implements TagsRepository
func (t *TagsRepositoryImple) Save(tags model.Tags) {
	result := t.Db.Create(&tags)
	helper.ErorPanic(result.Error)
}

// Update implements TagsRepository
func (*TagsRepositoryImple) Update(tags model.Tags) {
	panic("unimplemented")
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImple{Db: Db}
}
