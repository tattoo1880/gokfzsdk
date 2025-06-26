package request


type AlibabaProductMerchantproductEditGetRequest struct {
    /*
        产品ID     */
    Id  *int64 `json:"id" required:"true" `
}

func (s *AlibabaProductMerchantproductEditGetRequest) SetId(v int64) *AlibabaProductMerchantproductEditGetRequest {
    s.Id = &v
    return s
}

func (req *AlibabaProductMerchantproductEditGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Id != nil) {
        paramMap["id"] = *req.Id
    }
    return paramMap
}

func (req *AlibabaProductMerchantproductEditGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}