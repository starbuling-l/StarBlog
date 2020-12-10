package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/starbuling-l/StarBlog/pkg/setting"
	"log"
	"time"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))
	if err != nil {
		log.Printf("models.Setup err %v", err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}
	db.SingularTable(true)
	//db.LogMode(true)
	//callback方法注册进gorm的钩子里
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimestampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createOn, ok := scope.FieldByName("CreatedOn"); ok {
			if createOn.IsBlank {
				createOn.Set(nowTime)
				log.Println(createOn)
			}
		}
		if modifiedOn, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifiedOn.IsBlank {
				modifiedOn.Set(nowTime)
				log.Println(modifiedOn)
			}
		}
	}
}

// updateTimestampForUpdateCallback will set  `ModifiedOn` when updating
func updateTimestampForUpdateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		if _, ok := scope.Get("gorm:update_column"); !ok {
			scope.SetColumn("ModifiedOn", time.Now().Unix())
		}
	}
}

// callback 实现软删除
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOn, ok := scope.FieldByName("DeletedOn")
		if ok && !scope.Search.Unscoped {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v = %v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOn.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}

	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return "" + str
	}
	return ""
}

//定时清理软删除的tag
func CleanAllTags() bool {
	db.Unscoped().Where("deleted_on != ?", 0).Delete(&Tag{})
	return true
}

//定时清除软删除的articles
func CleanAllArticles() bool {
	db.Unscoped().Where("deleted_on != ?", 0).Delete(&Article{})
	return true
}

func closeDb() {
	defer db.Close()
}
