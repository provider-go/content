package models

import (
	"github.com/provider-go/content/global"
	"time"
)

type ContentSingle struct {
	Id         int32     `json:"id" gorm:"auto_increment;primary_key;comment:'主键'"`                                     // 主键ID
	Title      string    `json:"title" gorm:"column:title;type:varchar(50);not null;default:'';comment:文章标题"`           // 文章标题
	Content    string    `json:"content" gorm:"column:content;type:text;not null;comment:文章内容"`                         // 文章内容
	Status     int       `json:"status" gorm:"column:status;type:tinyint(1);not null;default:0;comment:文章状态：0：正常；1：隐藏"` // 文章状态：0：正常；1：隐藏
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime;comment:创建时间"`                                        // 创建时间
	UpdateTime time.Time `json:"update_time" gorm:"autoCreateTime;comment:更新时间"`                                        // 更新时间
}

func CreateSingle(title, content string) error {
	return global.DB.Table("content_singles").Select("title", "content").
		Create(&ContentSingle{Title: title, Content: content}).Error
}

func DeleteSingle(id int32) error {
	return global.DB.Table("content_singles").Where("id = ?", id).Delete(&ContentSingle{}).Error
}

func ListSingle(pageSize, pageNum int) ([]*ContentSingle, int64, error) {
	var rows []*ContentSingle
	//计算列表数量
	var count int64
	global.DB.Table("content_singles").Count(&count)

	if err := global.DB.Table("content_singles").Order("create_time desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, count, nil
}

func ViewSingle(id int32) (*ContentSingle, error) {
	row := new(ContentSingle)
	if err := global.DB.Table("content_singles").Where("id = ?", id).First(&row).Error; err != nil {
		return nil, err
	}
	return row, nil
}
