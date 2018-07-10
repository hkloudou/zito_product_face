package zitoproductface

//ProductFace 产品接口
type ProductFace interface {
	//GetPrefix 获得前置网址
	GetPrefix() string

	//BindControllers 绑定beego 控制器
	BindControllers()
}
