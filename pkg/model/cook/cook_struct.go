package cook

import (
	"baotian0506.com/go_cookbook/pkg/model/userinfo1"
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
)

type Demo struct {
	Tile    string
	Summary string
}

type Demo3 struct {
	Tile    string
	Summary string
}

type TestEmbedded struct {
	gorm.Model
	Demo   Demo3 `gorm:"embedded"`
	Email2 string
}

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded"`
	Upvotes int32
}

type TestJson struct {
	ID      int
	Author  Author `gorm:"embedded"`
	Upvotes int32
}

func init() {
	return
	userinfo1.GetDB().Debug().AutoMigrate(&XianghaCook{})

	var list = make([]List, 1)
	list[0].Dese = "des1"
	list[0].Img = "img1"

	var x1 = XianghaCook{
		Steps: Steps{
			GormJson: GormJson{List: list},
		},
		Tips: "",
	}
	userinfo1.GetDB().Debug().Where("id=?", 1).Save(&x1)

}

type XianghaCook struct {
	gorm.Model
	Name          string `json:"name" gorm:"not null;size:100"`
	Img           string `json:"img" gorm:"not null;size:500"`
	Video         string `json:"video" gorm:"not null;size:500"`
	HealthClassif string `json:"healthClassif" gorm:"not null;size:100"`
	//Ingredients []Ingredients `json:"ingredients" gorm:"embedded"`
	Steps Steps  `json:"steps" gorm:"type:json"` //注意，不能加 embedded
	Tips  string `json:"tips"`
}

/*
CREATE TABLE `xiangha_cooks` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(100) NOT NULL,
  `img` varchar(500) NOT NULL,
  `video` varchar(500) NOT NULL,
  `health_classif` varchar(100) NOT NULL,
  `steps` json DEFAULT NULL,
  `tips` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_xiangha_cooks_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
*/

type Ingredients struct {
	Name string `json:"name"`
	Num  string `json:"num"`
}
type List struct {
	Img  string `json:"img"`
	Dese string `json:"dese"`
}
type Steps struct {
	GormJson
}

type GormJson struct {
	List []List `json:"list"`
}

func (c GormJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c.List)
	return string(b), err
}

func (c *GormJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c.List)
}

func (c Demo) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Demo) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
