package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type DishDataDemo2 struct {
	Name  string `json:"name"`
	Img   string `json:"img"`
	Video string `json:"video"`
	//HealthClassif string `json:"healthClassif" gorm:"type:json"`
	Ingredients IngredientsList `json:"ingredients" gorm:"type:json"`
	//Steps Steps `json:"steps" gorm:"type:json"`
	//Tips string `json:"tips"`
}

type IngredientsList []Ingredients

type Ingredients struct {
	Name string `json:"name"`
	Num  string `json:"num"`
}
type List struct {
	Img  string `json:"img"`
	Dese string `json:"dese"`
}
type Steps struct {
	List []List `json:"list"`
}

func (f IngredientsList) Value() (driver.Value, error) {
	fmt.Printf("IngredientsList.Value, f.v=%#v\n", f)
	b, err := json.Marshal(f)
	return string(b), err
}

func (f IngredientsList) Scan(input interface{}) error {
	fmt.Printf("IngredientsList.Scan=%+v, str=%s\n", input, string(input.([]byte)))
	return json.Unmarshal(input.([]byte), &f)
}

//type DishData struct {
//	Name string `json:"name"`
//	Img string `json:"img"`
//	Video string `json:"video"`
//	HealthClassif string `json:"healthClassif"`
//	Ingredients []Ingredients `json:"ingredients"`
//	Steps Steps `json:"steps"`
//	Tips string `json:"tips"`
//}
//type Ingredients struct {
//	Name string `json:"name"`
//	Num string `json:"num"`
//}
//type List struct {
//	Img string `json:"img"`
//	Dese string `json:"dese"`
//}
//type Steps struct {
//	List []List `json:"list"`
//}
