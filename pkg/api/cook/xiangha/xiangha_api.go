package xiangha

import (
	"510909033/go_cookbook/pkg/api"
	"510909033/go_cookbook/pkg/model/cook/xiangha"
	"510909033/go_cookbook/pkg/model/cook/xiangha/model"
	"510909033/go_cookbook/pkg/model/userinfo1"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// InitTable godoc
// @Summary 初始化相关表
// @Description 初始化香哈相关表
// @Tags Cook/Xiangha
// @Accept json
// @Produce json
// @Success 200 {string} string "请求成功"
// @Failure 400 {string} string "非200都表示失败"
// @Router /cook/xiangha/init_table [get]
func InitTable(c *gin.Context) {

	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")

	err := userinfo1.GetDB().AutoMigrate(&xiangha.Sitemap{})

	userinfo1.GetDB().AutoMigrate(&model.DishDataDemo2{})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"code":   1,
			"data":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"code":   0,
		"data":   nil,
	})
	return

	//
	//menu := userinfo1.Menu{
	//	UserId:  0,
	//	SStatus: 0,
	//	Title:   "",
	//	Extra:   "",
	//}
	//
	//tx:=userinfo1.GetDB().Create(&menu)
	//if tx.Error !=nil{
	//	c.JSON(http.StatusOK, gin.H{
	//		"status":"fail",
	//		"code":1,
	//		"data":tx.Error,
	//	})
	//	return
	//}
	//
	//ret:=make(map[string]interface{})
	//ret["RowsAffected"] = tx.RowsAffected
	//ret["Name()"] = tx.Name()
	//ret["menu"] = menu.ID
	//c.JSON(http.StatusOK, ret)
	return
}

// UpdateSitemap godoc
// @Summary 更新sitemap表数据
// @Description 更新sitemap表数据
// @Tags Cook/Xiangha
// @Accept json
// @Produce json
// @Success 200 {string} string "请求成功"
// @Failure 400 {string} string "非200都表示失败"
// @Router /cook/xiangha/update_sitemap [get]
func UpdateSitemap(c *gin.Context) {

	url := "https://xml.xiangha.com/xiangha/babytree/dish/sitemap.json"

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "fail",
			"code":    0,
			"message": "获取sitemap数据失败, " + err.Error(),
			"data":    nil,
		})
		return
	}
	defer resp.Body.Close()

	var conByteList []byte
	conByteList, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "fail",
			"code":    0,
			"message": "获取sitemap数据失败, " + err.Error(),
			"data":    nil,
		})
		return
	}

	var conList = xiangha.XianghaSitemap{
		List: make([]string, 0),
	}
	err = json.Unmarshal(conByteList, &(conList.List))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "fail",
			"code":    0,
			"message": "解析sitemap的json数据失败, " + err.Error(),
			"data":    nil,
		})
		return
	}

	for _, v := range conList.List {
		var sitemapModel = xiangha.Sitemap{}
		tx := userinfo1.GetDB().Where("url=?", v).Find(&sitemapModel)
		if tx.Error != nil {
			api.Errorf("Find.db err, err=%+v\n", tx.Error)
			continue
		}
		if sitemapModel.ID > 0 {
			api.Infof("数据已存在, id=%d, url=%s", sitemapModel.ID, sitemapModel.Url)
			continue
		}

		sitemapModel.Url = v
		tx = userinfo1.GetDB().Create(&sitemapModel)
		if tx.Error != nil {
			api.Errorf("Create.db err, err=%+v\n", tx.Error)
			continue
		}
		if tx.RowsAffected != 1 {
			api.Errorf("Create.db err, err=%s\n", "RowsAffected!=1")
			continue
		}
		api.Infof(" url=%s, model=%+v success", v, sitemapModel)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"code":    0,
		"message": "",
		"data":    conList,
	})
	return

}

func UpdateDishData(c *gin.Context) {

}

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query demo.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} demo._ResponsePostList
// @Router /posts2 [get]
func GetPostListHandler2(c *gin.Context) {

}
