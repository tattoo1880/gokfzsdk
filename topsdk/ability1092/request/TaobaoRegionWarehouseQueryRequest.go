package request


type TaobaoRegionWarehouseQueryRequest struct {
    /*
        仓库编码     */
    StoreCode  *string `json:"store_code" required:"true" `
}

func (s *TaobaoRegionWarehouseQueryRequest) SetStoreCode(v string) *TaobaoRegionWarehouseQueryRequest {
    s.StoreCode = &v
    return s
}

func (req *TaobaoRegionWarehouseQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.StoreCode != nil) {
        paramMap["store_code"] = *req.StoreCode
    }
    return paramMap
}

func (req *TaobaoRegionWarehouseQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}