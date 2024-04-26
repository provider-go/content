package models

import (
	"github.com/provider-go/content/global"
	"time"
)

type ContentChannel struct {
	Id          int32     `json:"id" gorm:"auto_increment;primary_key;comment:'主键'"`                                               // 主键ID
	Pid         int32     `json:"pid" gorm:"column:pid;not null;default:0;comment:父id"`                                            // 父类ID
	ChannelName string    `json:"channelName" gorm:"column:channel_name;type:varchar(20);unique;not null;default:'';comment:频道名称"` // 类型名称
	CreateTime  time.Time `json:"create_time" gorm:"autoCreateTime;comment:创建时间"`                                                  // 创建时间
	UpdateTime  time.Time `json:"update_time" gorm:"autoCreateTime;comment:更新时间"`                                                  // 更新时间
}

func CreateChannel(pid int32, name string) error {
	return global.DB.Table("content_channels").Select("pid", "channel_name").
		Create(&ContentChannel{Pid: pid, ChannelName: name}).Error
}

func DeleteChannel(id int32) error {
	return global.DB.Table("content_channels").Where("id = ?", id).Delete(&ContentChannel{}).Error
}

func ListChannel(pageSize, pageNum int) ([]*ContentChannel, int64, error) {
	var rows []*ContentChannel
	//计算列表数量
	var count int64
	global.DB.Table("content_channels").Count(&count)

	if err := global.DB.Table("content_channels").Order("create_time desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, count, nil
}
