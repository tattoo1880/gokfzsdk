package request

import (
        "topsdk/ability1092/domain"
        "topsdk/util"
    )

type TaobaoRegionPriceManageRequest struct {
    /*
        商品id     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        无sku传0 defalutValue��0    */
    SkuId  *int64 `json:"sku_id,omitempty" required:"false" `
    /*
        列表     */
    RegionalPriceDtos  *[]domain.TaobaoRegionPriceManageRegionalPriceDto `json:"regional_price_dtos" required:"true" `
    /*
        true:全量, false:增量 defalutValue��true    */
    IsFull  *bool `json:"is_full,omitempty" required:"false" `
}

func (s *TaobaoRegionPriceManageRequest) SetItemId(v int64) *TaobaoRegionPriceManageRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoRegionPriceManageRequest) SetSkuId(v int64) *TaobaoRegionPriceManageRequest {
    s.SkuId = &v
    return s
}
func (s *TaobaoRegionPriceManageRequest) SetRegionalPriceDtos(v []domain.TaobaoRegionPriceManageRegionalPriceDto) *TaobaoRegionPriceManageRequest {
    s.RegionalPriceDtos = &v
    return s
}
func (s *TaobaoRegionPriceManageRequest) SetIsFull(v bool) *TaobaoRegionPriceManageRequest {
    s.IsFull = &v
    return s
}

func (req *TaobaoRegionPriceManageRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.SkuId != nil) {
        paramMap["sku_id"] = *req.SkuId
    }
    if(req.RegionalPriceDtos != nil) {
        paramMap["regional_price_dtos"] = util.ConvertStructList(*req.RegionalPriceDtos)
    }
    if(req.IsFull != nil) {
        paramMap["is_full"] = *req.IsFull
    }
    return paramMap
}

func (req *TaobaoRegionPriceManageRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}