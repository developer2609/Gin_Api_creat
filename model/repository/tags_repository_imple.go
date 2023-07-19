package repository

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/model"
)

type TagsRepository interface {
	Save(tags model.Tags)
	Update(tags model.Tags)
	Delete(tagsId int)
	FindAll() []model.Tags
	FindById(tagsId int) (tags model.Tags, err error)
}
