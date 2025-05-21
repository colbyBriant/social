package video

import "time"

type XkVideoCategory struct {
	ID           uint      `gorm:"primarykey;comment:主键ID" json:"id" form:"id"`
	CategoryName string    `json:"categoryName" gorm:"not null;default:'';comment:分类名称"`
	Description  string    `json:"description" gorm:"not null;default:'';comment:分类描述"`
	CreateTime   time.Time `gorm:"type:datetime(0);comment:创建时间" json:"createTime"`
	UpdatedTime  time.Time `gorm:"type:datetime(0);comment:更新时间" json:"updatedTime"`
	ParentId     uint      `json:"parentId" gorm:"not null;default:0;comment:分类的主ID"`
	Status       int8      `json:"status" gorm:"not null;default:1;comment:0 未发布 1 发布"`
	Sorted       int8      `json:"sorted" gorm:"not null;default:1;comment:0 排序"`
	IsDelete     int8      `json:"isDelete" gorm:"not null;default:0;comment:0 未删除 1 删除"`
	// 忽略该字段，- 表示无读写，-:migration 表示无迁移权限，-:all 表示无读写迁移权限
	Children []*XkVideoCategory `gorm:"-" json:"children"`
	TopObj   *XkVideoCategory   `gorm:"-" json:"-"`
}

func (XkVideoCategory) TableName() string {
	return "xk_video_category"
}
