package tag_service

import (
	"encoding/json"
	"log"

	"github.com/starbuling-l/StarBlog/models"
	"github.com/starbuling-l/StarBlog/pkg/go_redis"
	"github.com/starbuling-l/StarBlog/server/cache_service"
)

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

func (t *Tag) ExitByID() (bool, error) {
	return models.ExistTagByID(t.ID)
}

func (t *Tag) Add() error {
	return models.AddTag(t.Name, t.State, t.CreatedBy)
}

func (t *Tag) Edit() error {
	data := make(map[string]interface{})
	data["modified_by"] = t.ModifiedBy
	if t.State > 0 {
		data["state"] = t.State
	}

	return models.EditTag(t.ID, data)
}

func (t *Tag) Delete() error {
	return models.DeleteTag(t.ID)
}

func (t *Tag) Count() (int, error) {
	return models.GetTagsTotal(t.getMaps)
}

func (t *Tag) GetAll() (tags []models.Tag, err error) {
	var cacheTags [] models.Tag
	cache := cache_service.Tag{
		State:    t.State,
		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}

	key := cache.GetTagKey()

	if go_redis.Exists(key) {
		data, err := go_redis.Get(key)
		if err != nil {
			log.Println(err)
			return nil, err
		} else {
			json.Unmarshal(data, &cacheTags)
			return cacheTags, nil
		}
	}

	tags, err = models.GetTags(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}

	go_redis.Set(key, tags, 3600)
	return
}

func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["delete_on"] = 0

	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.State > 0 {
		maps["State"] = t.State
	}

	return maps
}
