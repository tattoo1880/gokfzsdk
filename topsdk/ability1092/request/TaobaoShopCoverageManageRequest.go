package request

import (
        "topsdk/util"
    )

type TaobaoShopCoverageManageRequest struct {
    /*
        门店code     */
    ShopCode  *string `json:"shop_code" required:"true" `
    /*
        覆盖的地址列表。如覆盖整个省则传入省份即可；如只覆盖某个区需要按“省~市~区”格式传入。     */
    Regions  *[]string `json:"regions" required:"true" `
}

func (s *TaobaoShopCoverageManageRequest) SetShopCode(v string) *TaobaoShopCoverageManageRequest {
    s.ShopCode = &v
    return s
}
func (s *TaobaoShopCoverageManageRequest) SetRegions(v []string) *TaobaoShopCoverageManageRequest {
    s.Regions = &v
    return s
}

func (req *TaobaoShopCoverageManageRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ShopCode != nil) {
        paramMap["shop_code"] = *req.ShopCode
    }
    if(req.Regions != nil) {
        paramMap["regions"] = util.ConvertBasicList(*req.Regions)
    }
    return paramMap
}

func (req *TaobaoShopCoverageManageRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}