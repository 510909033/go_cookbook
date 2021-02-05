package xiangha

import "gorm.io/gorm"

type Sitemap struct {
	gorm.Model
	Url string
}

type XianghaSitemap struct {
	List []string
}
