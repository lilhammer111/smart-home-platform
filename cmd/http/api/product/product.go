package product

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/api/product/handler"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/product"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"net/http"
)

// GetProductList .
// @id			GetProductList
// @Summary		get product list
// @Tags		products
// @Produce		json
// @Param       Authorization  header    string  true  "Bearer User's access token"
// @Param    page             query    int16    false    "Page number for pagination"
// @Param    limit            query    int16    false    "Number of items per page"
// @Param    search           query    string   false    "Search term"
// @Param    category_id      query    int32    false    "Category ID to filter products"
// @Param    brand_id_list    query    []int32  false    "List of Brand IDs to filter products"
// @Param    is_price_asc     query    bool     false    "Set to true for ascending price sorting, false for descending"
// @Param    is_rating_asc    query    bool     false    "Set to true for ascending rating sorting, false for descending"
// @Param    on_sale          query    bool     false    "Set to true to filter products that are on sale"
// @Param    is_free_shipping query    bool     false    "Set to true to filter products with free shipping"
// @Param    is_new           query    bool     false    "Set to true to filter new products"
// @Param    is_hot           query    bool     false    "Set to true to filter hot products"
// @Param    is_recommended   query    bool     false    "Set to true to filter recommended products"
// @Success		200				{object}		example.RespOk{data=[]example.UpdateProductData{state=example.ProductState}} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/list [GET]
func GetProductList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ProductFilter
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewGetProductListService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetProductDetail .
// @id			GetProductDetail
// @Summary		get product detail
// @Tags		products
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		id		query	int	true	"id"
// @Success		200				{object}		example.RespOk{data=example.UpdateProductData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/detail [GET]
func GetProductDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewGetProductDetailService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// AddNewProduct .
// @id			AddNewProduct
// @Summary		add a new product
// @Tags		products
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		product	body	example.AddProductData{state=example.ProductState}	true	"product data"
// @Success		201				{object}		example.RespCreated{data=example.UpdateProductData{state=example.ProductState}} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/add [POST]
func AddNewProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.AddProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewAddNewProductService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteProduct .
// @id			DeleteProduct
// @Summary		delete product
// @Tags		products
// @Produce		json
// @Param       Authorization  header    string  true  "Bearer User's access token"
// @Param		id	query	string	true	"id"
// @Success		204				{object}		example.RespNoContent				"success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/delete [DELETE]
func DeleteProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Req
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewDeleteProductService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateProduct .
// @id			UpdateProduct
// @Summary		update product info
// @Tags		products
// @Access		json
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		product	body	example.UpdateProductData{state=example.ProductState}	true	"product data"
// @Success		200				{object}		example.RespOk{data=example.UpdateProductData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/update [PUT]
func UpdateProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ProductInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewUpdateProductService(ctx, c).Do(&req)

	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// RateProduct .
// @id			UpdateRating
// @Summary		update product rating info
// @Tags		products
// @Access		json
// @Produce		json
// @Param        Authorization  header    string  true  "Bearer User's access token"
// @Param		product	rating body	example.UpdateRatingData	true	"product rating data"
// @Success		200				{object}		example.RespOk{data=example.RatingData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/products/rating/update [PUT]
func RateProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.RatingReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewRateProductService(ctx, c).Do(&req)

	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
