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
	ReleaseViews()
}

//AddProduct AddProduct
func AddProduct(face ProductFace) {
	Products[face.GetPrefix()] = face
}

//ReleaseViews ReleaseViews
func ReleaseViews() {
	for _, v := range Products {
		v.ReleaseViews()
	}
}
