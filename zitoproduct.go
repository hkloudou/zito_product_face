package zitoproductface

//Products Products
var Products = make(map[string]ProductFace)

//ProductFace 产品接口
type ProductFace interface {
	//GetPrefix 获得前置网址
	GetPrefix() string
	GetTitle() string
	//BindControllers 绑定beego 控制器
	BindControllers()
}

//AddProduct AddProduct
func AddProduct(prefix string, face ProductFace) {
	Products[prefix] = face
}
