package models

import "github.com/jinzhu/gorm"

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// GetTags gets a list of tags based on paging and constraints
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag, err error) {
	if pageNum > 0 && pageSize > 0 {
		err = db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags).Error

	} else {
		err = db.Where(maps).Find(&tags).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

// GetTags gets a list of tags based on paging and constraints
func GetTagsTotal(maps interface{}) (count int, err error) {
	err = db.Model(&Tag{}).Where(maps).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}
	return
}

// ExistTagByName checks if there is a tag with the same name
func ExistTagByName(name string) (exist bool, err error) {
	var tag Tag
	err = db.Select("id").Where("name = ? AND deleted_on = ?", name, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return
}

// ExistTagByID determines whether a Tag exists based on the ID
func ExistTagByID(id int) (exist bool, err error) {
	var tag Tag
	err = db.Select("id").Where("id =? AND deleted_on = ?", id, 0).First(&tag).Error
	if err != nil {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return
}

// AddTag Add a Tag
func AddTag(name string, state int, createdBy string) (err error) {
	tag := Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}

	err = db.Create(&tag).Error
	if err != nil {
		return err
	}

	return
}

// DeleteTag delete a tag
func DeleteTag(id int) (err error) {
	err = db.Where("id = ?", id).Delete(&Tag{}).Error
	if err != nil {
		return err
	}
	return
}

// EditTag modify a single tag
func EditTag(id int, data interface{}) (err error) {
	err = db.Model(&Tag{}).Where("id = ? AND deleted_on = ?", id, 0).Updates(data).Error
	if err != nil {
		return err
	}
	return
}

//CleanAllTags clean up soft deleted tags regularly
func CleanAllTags() (flag bool, err error) {
	err = db.Unscoped().Where("deleted_on != ?", 0).Delete(&Tag{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// models callbacks 避免其默认值为空的情况
//func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("created_on", time.Now().Unix())
//	return nil
//}
//
//func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
//	scope.SetColumn("modified_on", time.Now().Unix())
//	return nil
//}
