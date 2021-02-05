package demo

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"`   // 可以为空
	Page        int64  `json:"page" form:"page" example:"1"`       // 页码
	Size        int64  `json:"size" form:"size" example:"10"`      // 每页数据量
	Order       string `json:"order" form:"order" example:"score"` // 排序依据
}

// _ResponsePostList 帖子列表接口响应数据
type _ResponsePostList struct {
	Code    int           `json:"code"`    // 业务响应状态码
	Message string        `json:"message"` // 提示信息
	Data    []*DemoDetail `json:"data"`    // 数据
}

type DemoDetail struct {
	Username string //用户名
	Id       int    //用户id
	xiaoxie  int    //小写属性
}
