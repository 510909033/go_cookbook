package api

import (
	"baotian0506.com/go_cookbook/pkg/model/userinfo1"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Users(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": "用户列表",
	})
}

func GetUser(c *gin.Context) {
	//tx:=userinfo1.GetDB().Table().Get()
	menu := userinfo1.Menu{}
	var menu1 userinfo1.Menu
	var menu2 userinfo1.Menu
	var menu3 userinfo1.Menu
	db := userinfo1.GetDB().Debug()
	tx := db.Where("id=?", 10).Find(&menu)

	db.Where("id=?", 10).Take(&menu1)
	db.Where("id=?", 10).First(&menu2)
	tx1 := db.Where("id=?", 11110).Last(&menu3)

	//tx.Callback().Query()

	if tx.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"code":   1,
			"data":   tx.Error,
		})
		return
	}

	ret := make(map[string]interface{})
	ret["RowsAffected"] = tx.RowsAffected
	ret["Name()"] = tx.Name()
	ret["menu"] = menu.ID
	ret["tx1"] = tx1.Error
	ret["tx1a"] = errors.Is(tx1.Error, gorm.ErrRecordNotFound)

	//ret["sql"] = tx.
	c.JSON(http.StatusOK, ret)
	return
}

// AddUser godoc
// @Summary 添加一个新菜单
// @Description 通过这个api，可以添加一个全新的菜单
// @Tags Menu
// @Accept json
// @Produce json
// @Param object query userinfo1.MenuResponse false "入参"
// -@Param menu_obj query object false "string enums" userinfo1.MenuResponse
// -@Param enumstring query string false "string enums" Enums(A, B, C)
// -@Param enumint query int false "int enums" Enums(1, 2, 3)
// -@Param enumnumber query number false "int enums" Enums(1.1, 1.2, 1.3)
// -@Param string query string false "string valid" minlength(5) maxlength(10)
// -@Param int query int false "int valid" minimum(1) maximum(10)
// -@Param default query string false "string default" default(A)
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// -@Failure 404 {string} string "ok"
// -@Failure 500 {string} string "ok"
// -@Failure 501 {object} userinfo1.MenuResponse "ok"
// -@Failure 502 {string} userinfo1.MenuResponse "ok"
// @Router /menu/add_menu [get]
func AddMenu(c *gin.Context) {

	menu := userinfo1.Menu{
		UserId:  0,
		SStatus: 0,
		Title:   "",
		Extra:   "",
	}

	tx := userinfo1.GetDB().Create(&menu)
	if tx.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"code":   1,
			"data":   tx.Error,
		})
		return
	}

	ret := make(map[string]interface{})
	ret["RowsAffected"] = tx.RowsAffected
	ret["Name()"] = tx.Name()
	ret["menu"] = menu.ID
	c.JSON(http.StatusOK, ret)
	return
}
