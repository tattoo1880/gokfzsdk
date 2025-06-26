package request

import (
        "topsdk/util"
    )

type TaobaoRegionWarehouseManageRequest struct {
    /*
        仓库编码     */
    StoreCode  *string `json:"store_code" required:"true" `
    /*
        可映射三级地址,例: 广东省     */
    Regions  *[]string `json:"regions" required:"true" `
}

func (s *TaobaoRegionWarehouseManageRequest) SetStoreCode(v string) *TaobaoRegionWarehouseManageRequest {
    s.StoreCode = &v
    return s
}
func (s *TaobaoRegionWarehouseManageRequest) SetRegions(v []string) *TaobaoRegionWarehouseManageRequest {
    s.Regions = &v
    return s
}

func (req *TaobaoRegionWarehouseManageRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.StoreCode != nil) {
        paramMap["store_code"] = *req.StoreCode
    }
    if(req.Regions != nil) {
        paramMap["regions"] = util.ConvertBasicList(*req.Regions)
    }
    return paramMap
}

func (req *TaobaoRegionWarehouseManageRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}