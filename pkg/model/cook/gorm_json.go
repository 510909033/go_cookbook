package cook

import (
	"baotian0506.com/go_cookbook/pkg/model/userinfo1"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	FileUrlJson FileJson `json:"file_url_json";gorm:"column:file_url_json;type:json"`
}

type FileJson struct {
	//Url  string
	//Name string
	v []string //如果有需要，将Url和Name的需求扩展到v里，比如v []map[string]string
}

func (f FileJson) Value() (driver.Value, error) {
	fmt.Printf("fileJson.Value, f.v=%#v\n", f.v)
	b, err := json.Marshal(f.v)
	return string(b), err
}

func (f FileJson) Scan(input interface{}) error {
	fmt.Printf("fleJson.Scan=%+v, str=%s\n", input, string(input.([]byte)))
	return json.Unmarshal(input.([]byte), &f.v)
}

func init() {
	return
	book := Book{
		FileUrlJson: FileJson{v: []string{"file_address1", "file_address2"}},
	}
	userinfo1.GetDB().Debug().AutoMigrate(&book)
	//userinfo1.GetDB().Debug().Create(&book)

	var book1 = Book{}
	userinfo1.GetDB().Debug().First(&book1)
	fmt.Printf("book=%#v\n", book1.FileUrlJson)

	os.Exit(0)

}
