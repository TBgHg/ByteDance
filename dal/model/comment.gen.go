// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameComment = "comment"

// Comment mapped from table <comment>
type Comment struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID     int32     `gorm:"column:user_id;not null" json:"user_id"`
	VideoID    int32     `gorm:"column:video_id;not null" json:"video_id"`
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	Removed    int32     `gorm:"column:removed;not null" json:"removed"`
	Deleted    int32     `gorm:"column:deleted;not null" json:"deleted"`
	Content    string    `gorm:"column:content;not null" json:"content"`
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
