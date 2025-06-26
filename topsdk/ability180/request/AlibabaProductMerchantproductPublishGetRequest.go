package request


type AlibabaProductMerchantproductPublishGetRequest struct {
    /*
        叶子类目ID     */
    CatId  *int64 `json:"cat_id" required:"true" `
}

func (s *AlibabaProductMerchantproductPublishGetRequest) SetCatId(v int64) *AlibabaProductMerchantproductPublishGetRequest {
    s.CatId = &v
    return s
}

func (req *AlibabaProductMerchantproductPublishGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CatId != nil) {
        paramMap["cat_id"] = *req.CatId
    }
    return paramMap
}

func (req *AlibabaProductMerchantproductPublishGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}