package request


type TaobaoShopCoverageQueryRequest struct {
    /*
        门店code     */
    ShopCode  *string `json:"shop_code" required:"true" `
}

func (s *TaobaoShopCoverageQueryRequest) SetShopCode(v string) *TaobaoShopCoverageQueryRequest {
    s.ShopCode = &v
    return s
}

func (req *TaobaoShopCoverageQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ShopCode != nil) {
        paramMap["shop_code"] = *req.ShopCode
    }
    return paramMap
}

func (req *TaobaoShopCoverageQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}