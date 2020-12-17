package article_service

import (
	"encoding/json"
	"github.com/starbuling-l/StarBlog/pkg/go_redis"
	"log"

	"github.com/starbuling-l/StarBlog/models"
	"github.com/starbuling-l/StarBlog/server/cache_service"
)

type Article struct {
	ID            int
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string

	PageNum  int
	PageSize int
}

func (a *Article) Add() (err error) {
	article := map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"created_by":      a.CreatedBy,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
	}
	if err = models.AddArticle(article); err != nil {
		return err
	}
	return
}

func (a *Article) Edit() (err error) {
	article := map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
		"modified_by":     a.ModifiedBy,
	}
	if err = models.EditArticle(a.ID, article); err != nil {
		return err
	}
	return
}

func (a *Article) Get() (article *models.Article, err error) {
	var cacheArticle *models.Article

	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()
	if go_redis.Exists(key) {
		data, err := go_redis.Get(key)
		if err != nil {
			log.Println(err)
		} else {
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}

	article, err = models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}

	go_redis.Set(key, article, 3600)
	return
}

func (a *Article) GetAll() (articles []*models.Article, err error) {
	var cacheArticles []*models.Article

	cache := cache_service.Article{
		TagID:    a.TagID,
		State:    a.State,
		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}
	key := cache.GetArticleKey()
	if go_redis.Exists(key) {
		data, err := go_redis.Get(key)
		if err != nil {
			log.Println(err)
		} else {
			json.Unmarshal(data, &cacheArticles)
			return cacheArticles, nil
		}
	}

	articles, err = models.GetArticles(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	go_redis.Set(key, articles, 3600)
	return
}

func (a *Article) Delete() error {
	return models.DeleteArticle(a.ID)
}

func (a *Article) Count() (int, error) {
	return models.GetArticlesTotal(a.getMaps())
}

func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticleByID(a.ID)
}

func (a *Article) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["delete_on"] = 0
	if a.State != -1 {
		maps["state"] = a.State
	}

	if a.TagID != -1 {
		maps["tag_id"] = a.TagID
	}
	return maps
}
