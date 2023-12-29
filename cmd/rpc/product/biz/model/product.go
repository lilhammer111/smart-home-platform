package model

import (
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model/ctype"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type ProductState int8

const (
	OnSale ProductState = 1 << iota
	IsFreeShipping
	IsNew
	IsHot
	IsRecommended
)

type Product struct {
	model.BaseModel
	CategoryId int32 `gorm:"type:int unsigned;not null;uniqueIndex:idx_category_brand_name"`
	BrandId    int32 `gorm:"type:int unsigned;not null;uniqueIndex:idx_category_brand_name"`

	Name    string `gorm:"type:varchar(50);not null;uniqueIndex:idx_category_brand_name"`
	Brief   string `gorm:"type:varchar(100);not null"`
	Picture string `gorm:"type:varchar(200);not null;comment: oss url"`

	State    ProductState `gorm:"type:tinyint unsigned;not null"`
	Price    float32      `gorm:"type:float(8,2) unsigned;index;not null"`
	Rating   float32      `gorm:"type:float(2,1) unsigned;not null;default: 0"`
	Showcase ctype.CList  `gorm:"type:json;not null;comment:provide a range of product illustrations also as product detail"`
}

func (p *Product) TableName() string {
	return "products"
}

func (p *Product) AfterCreate(tx *gorm.DB) (err error) {
	err = initRating(p, tx)
	if err != nil {
		klog.Error(err)
		return err
	}

	err = createCategoryBrandRelation(p, tx)
	if err != nil {
		klog.Error(err)
		return err
	}

	return nil
}

func (p *Product) AfterUpdate(tx *gorm.DB) (err error) {
	err = updateCategoryBrandRelation(p, tx)
	if err != nil {
		klog.Errorf("failed to update category_brand record: %s", err)
		return err
	}
	return err
}

// initRating initialize the rating of the new product as 5.0, a max value.
func initRating(p *Product, tx *gorm.DB) (err error) {
	rating := ProductRating{
		ProductId:     p.ID,
		TotalRating:   0,
		TotalCustomer: 0,
	}

	err = tx.Create(&rating).Error
	if err != nil {
		klog.Errorf("failed to create rating record: %s", err)
		return kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	return nil
}

// createCategoryBrandRelation determines whether to create a new record for category_brand table.
// if the new product make a new relationship between category and brand, And then a new record
// should be inserted into the category_brand table to record it for querying brands with
// the specific category in the next time.
func createCategoryBrandRelation(p *Product, tx *gorm.DB) (err error) {
	categoryBrandInfo := CategoryBrand{
		Id:         p.ID,
		CategoryId: p.CategoryId,
		BrandId:    p.BrandId,
	}
	err = tx.Create(&categoryBrandInfo).Error
	if err != nil {
		klog.Error(err)
		return kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	return nil
}

func updateCategoryBrandRelation(p *Product, tx *gorm.DB) (err error) {

	cbInfo := CategoryBrand{
		Id:         p.ID,
		CategoryId: p.CategoryId,
		BrandId:    p.BrandId,
	}
	err = tx.Updates(&cbInfo).Error
	if err != nil {
		klog.Error(err)
		return kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	return nil
}
