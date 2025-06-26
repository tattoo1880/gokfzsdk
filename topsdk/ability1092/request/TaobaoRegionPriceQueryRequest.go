package request

import (
        "topsdk/ability1092/domain"
        "topsdk/util"
    )

type TaobaoRegionPriceQueryRequest struct {
    /*
        商品id     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        无sku可传0 defalutValue��0    */
    SkuId  *int64 `json:"sku_id,omitempty" required:"false" `
    /*
        不传则返回所有设置的区域价格     */
    RegionalPriceDtos  *[]domain.TaobaoRegionPriceQueryRegionalPriceDto `json:"regional_price_dtos,omitempty" required:"false" `
}

func (s *TaobaoRegionPriceQueryRequest) SetItemId(v int64) *TaobaoRegionPriceQueryRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoRegionPriceQueryRequest) SetSkuId(v int64) *TaobaoRegionPriceQueryRequest {
    s.SkuId = &v
    return s
}
func (s *TaobaoRegionPriceQueryRequest) SetRegionalPriceDtos(v []domain.TaobaoRegionPriceQueryRegionalPriceDto) *TaobaoRegionPriceQueryRequest {
    s.RegionalPriceDtos = &v
    return s
}

func (req *TaobaoRegionPriceQueryRequest) ToMap() map[string]interface{} {
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
    return paramMap
}

func (req *TaobaoRegionPriceQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}