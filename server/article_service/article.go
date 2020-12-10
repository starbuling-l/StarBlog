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

func (a *Article) Get() (*models.Article, error) {
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

	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}

	go_redis.Set(key, article, 3600)
	return article, nil
}

func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticleByID(a.ID)
}