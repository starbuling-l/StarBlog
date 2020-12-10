package models

import "github.com/jinzhu/gorm"

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

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

func GetArticlesTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(count)
	return
}

func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ?", id).First(&article).Related(&article.Tag).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

func ExistArticleByID(id int) (bool, error) {
	//var article Article
	//db.Select("id").Where("id =?", id).First(&article)
	//if article.ID > 0 {
	//	return true
	//}
	//return false
	var article Article
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
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

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(&Article{})
	return true
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)
	return true
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
//}
