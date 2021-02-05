package userinfo1

import (
	"gorm.io/gorm"
	"time"
)

/*
CREATE TABLE `menu` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
`category_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '分类id',
`user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户user_id',
`s_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态0删除，1正常',
`title` varchar(1000) NOT NULL DEFAULT '' COMMENT '名称',
`create_ts` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
`update_ts` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
`extra` varchar(10000) NOT NULL DEFAULT '' COMMENT '扩展字段',
PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=130 DEFAULT CHARSET=utf8mb4 COMMENT='菜单表';
*/

type MenuResponse struct {
	ID        uint      `gorm:"primarykey;not null;autoIncrement"`
	CreatedAt time.Time `gorm:"not null" json:"-"`
	// 在创建时该字段值为零值或者在更新时，使用当前时间戳秒数填充
	UpdatedAt      uint64 `gorm:"not null"`
	UpdatedNanoTs  uint64 `gorm:"autoUpdateTime:nano;not null"`  // 使用时间戳填纳秒数充更新时间
	UpdatedMilliTs uint64 `gorm:"autoUpdateTime:milli;not null"` // 使用时间戳毫秒数填充更新时间
	CreatedTs      uint64 `gorm:"autoCreateTime;not null"`       // 使用时间戳秒数填充创建时间
}

type Menu struct {
	ID        uint      `gorm:"primarykey;not null;autoIncrement"`
	CreatedAt time.Time `gorm:"not null" json:"-"`
	// 在创建时该字段值为零值或者在更新时，使用当前时间戳秒数填充
	UpdatedAt      uint64 `gorm:"not null"`
	UpdatedNanoTs  uint64 `gorm:"autoUpdateTime:nano;not null"`  // 使用时间戳填纳秒数充更新时间
	UpdatedMilliTs uint64 `gorm:"autoUpdateTime:milli;not null"` // 使用时间戳毫秒数填充更新时间
	CreatedTs      uint64 `gorm:"autoCreateTime;not null"`       // 使用时间戳秒数填充创建时间

	gorm.DeletedAt `json:"-"`

	UserId  uint64 `gorm:"not null"`
	SStatus uint8  `gorm:"not null"`
	Title   string `gorm:"not null;size:100"`
	Extra   string `gorm:"not null;size:2000"`
	//Extra ExtraOjb `sql:"TYPE:json"`

	//// Name 姓名
	//Name string `gorm:"not null;size:100;uniqueindex"`
	//// Gender 姓名
	//Gender string `gorm:"not null;size:20;default:;"`
	//// RegInfo 注册信息
	////
	//// 嵌入结构体
	//RegInfo RegInfo `gorm:"embedded"`
	//CountryCn string `gorm:"not null"`
	//CountryEn string `gorm:"not null"`
	//CountryPhoneCode string `gorm:"not null"`

}

type ExtraOjb struct {
	c1 string
	c2 int
	c3 bool
}
