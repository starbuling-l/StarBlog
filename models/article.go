package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles *[]Article, err error) {
	err = db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func GetArticlesTotal(maps interface{}) (count int, err error) {
	err = db.Model(&Article{}).Where(maps).Count(count).Error
	if err != nil {
		return 0, err
	}
	return
}

func GetArticle(id int) (article *Article, err error) {
	err = db.Where("id = ?", id).First(&article).Related(&article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func AddArticle(data map[string]interface{}) (err error) {
	err = db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	}).Error
	if err != nil {
		return err
	}
	return
}

func ExistArticleByID(id int) (flag bool, err error) {
	//var article Article
	//db.Select("id").Where("id =?", id).First(&article)
	//if article.ID > 0 {
	//	return true
	//}
	//return false
	var article Article
	err = db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

/*func ExistArticleByName(name string) bool {
	var Article Article
	db.Select("id").Where("name = ?", name).First(&Article)
	if Article.ID > 0 {
		return true
	}
	return false
}*/

func DeleteArticle(id int) (err error) {
	err = db.Where("id = ?", id).Delete(&Article{}).Error
	if err != nil {
		return err
	}
	return
}

func EditArticle(id int, data interface{}) (err error) {
	err = db.Model(&Article{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}
	return
}

//定时清除软删除的articles
func CleanAllArticles() (err error) {
	err = db.Unscoped().Where("deleted_on != ?", 0).Delete(&Article{}).Error
	if err != nil {
		return err
	}
	return
}

// models callbacks 避免其默认值为空的情况
//func (article *Article) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("created_on", time.Now().Unix())
//	return nil
//}
//
//func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
//	scope.SetColumn("modified_on", time.Now().Unix())
//	return nil
//
