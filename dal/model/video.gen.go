// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameVideo = "video"

// Video mapped from table <video>
type Video struct {
	ID       int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AuthorID int32  `gorm:"column:author_id;not null" json:"author_id"`
	PlayURL  string `gorm:"column:play_url;not null" json:"play_url"`
	CoverURL string `gorm:"column:cover_url;not null" json:"cover_url"`
	Time     int32  `gorm:"column:time;not null" json:"time"`
	Title    string `gorm:"column:title;not null" json:"title"`
	Removed  int32  `gorm:"column:removed;not null" json:"removed"`
	Deleted  int32  `gorm:"column:deleted;not null" json:"deleted"`
}

// TableName Video's table name
func (*Video) TableName() string {
	return TableNameVideo
}