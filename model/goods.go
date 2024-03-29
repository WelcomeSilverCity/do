package model

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/anaskhan96/go-password-encoder"
)

// 定义商品结构体
type Goods struct {
	BaseModel

	//如何关联分类的外键
	CategoryID int32    `gorm:"type:int;not null;comment:'分类的id'"`
	Category   Category //外键关联

	//如何关联分类的外键
	BrandsID int32 `gorm:"type:int;not null;comment:'品牌的id'"`
	Brands   Brands

	Name            string   `gorm:"type:varchar(50);not null;comment:'商品名称'"` //给字符串加索引能命中索引吗？ es
	GoodsSn         string   `gorm:"type:varchar(50);not null;comment:'商品的货号'"`
	ClickNum        int32    `gorm:"type:int;default:0;not null;comment:'商品的浏览量'"` //排行榜
	SoldNum         int32    `gorm:"type:int;default:0;not null;comment:'销售量'"`
	FavNum          int32    `gorm:"type:int;default:0;not null;comment:'点赞量或者收藏'"`
	MarketPrice     float32  `gorm:"not null;comment:'市场价格原价'"`
	ShopPrice       float32  `gorm:"not null;comment:'售价'"`
	GoodsBrief      string   `gorm:"type:varchar(100);not null;comment:'商品介绍'"`
	Images          GormList `gorm:"type:varchar(1000);not null;comment:'商品图片'"` // {"http://1.jpg", "http://2.jpg", "http://3.jpg"}  []{"http://1.jpg","http://2.jpg","http://3.jpg",}
	DescImages      GormList `gorm:"type:varchar(1000);not null;comment:'商品介绍图片'"`
	GoodsFrontImage string   `gorm:"type:varchar(200);not null;comment:'展示封面图'"`

	OnSale   int8 `gorm:"type:tinyint(1);default:0;not null;not null;comment:'上下架状态 0:下架 1:上架'"`
	ShipFree int8 `gorm:"type:tinyint(1);default:1;not null;comment:'是否包邮 0:不包邮 1:包邮'"`
	IsNew    int8 `gorm:"type:tinyint(1);default:0;not null;comment:'是否新品 0:否 1:是'"`
	IsHot    int8 `gorm:"type:tinyint(1);default:0;not null;comment:'是否热销 0:否 1:是'"`
}

// 定义商品分类结构体
type Category struct {
	BaseModel
	Name             string `gorm:"type:varchar(50);not null;comment:'商品分类名称'"`
	ParentCategoryID int32  `gorm:"type:int(11);default:0;not null;comment:'父级id,0代表顶级'"`
	ParentCategory   *Category
	Level            int32 `gorm:"type:int;not null;default:1;comment:'级别,1代表顶级'"`
	IsTab            bool  `gorm:"default:false;not null;comment:'是否是选项，搜索条件'"`
}

// 定义品牌结构体
type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null"`
	Logo string `gorm:"type:varchar(200);default:'';not null;comment:'品牌logo'"`
}

// 定义的是轮播图的结构体
type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null;comment:'轮播图封面'"`
	Url   string `gorm:"type:varchar(200);not null;comment:'活动地址'"`
	Index int32  `gorm:"type:int;default:1;not null;comment:'排序'"`
}

// 定义品牌分类表
type GoodsCategoryBrand struct {
	BaseModel
	//联合索引，组合索引 1，2
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique;comment:'分类的id'"`
	Category   Category

	BrandsID int32 `gorm:"type:int;index:idx_category_brand,unique;comment:'品牌的id'"`
	Brands   Brands
}

// 自定义表名
// TableName 会将 GoodsCategoryBrand 的表名重写为 `goodscategorybrand`
func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

//spu :标准化产品单元 通俗点讲，属性值、特性相同的商品就可以称为一个SPU。 (颜色,版本，内存)
//sku :最小存货单位 ：对一种商品而言，当其品牌、型号、配置、等级、花色、包装容量、单位、生产日期、保质期、用途、价格、产地等属性中任一属性与其他商品存在不同时，可称为一个单品。

// MakePassword 加密密码的方法
func MakePassword(plainPwd string) string {
	options := &password.Options{10, 10000, 50, md5.New}
	salt, encodedPwd := password.Encode(plainPwd, options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd) // pbkdf2-sha512 $符号拼接
	return newPassword
}

// VerifyPassword 校验密码
func VerifyPassword(plainPwd, encodedPwd string) bool {
	options := &password.Options{10, 10000, 50, md5.New}
	split := strings.Split(encodedPwd, "$")
	verify := password.Verify(plainPwd, split[2], split[3], options)
	return verify
}
