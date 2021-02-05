package cook

import (
	"baotian0506.com/go_cookbook/pkg/model/userinfo1"
	"fmt"
)

type BookObj struct {
	//....
	FileUrlJson FileJsonObj `json:"file_url_json";gorm:"column:file_url_json;type:json"`
}

type FileJsonObj struct {
	Url  string
	Name string
}

//
//func (f FileJsonObj) Value() (driver.Value, error) {
//	b, err := json.Marshal(f.v)
//	return string(b), err
//}
//
//func (f FileJsonObj) Scan(input interface{}) error {
//	return json.Unmarshal(input.([]byte), f.v)
//}

func init() {
	return
	book := BookObj{
		FileUrlJson: FileJsonObj{Url: "aaa", Name: "bbb"},
	}
	userinfo1.GetDB().Debug().Create(&book)
	var b BookObj
	userinfo1.GetDB().Debug().First(&b)
	fmt.Printf("b=%+v\n", b)
}
