package request


type AlibabaProductMerchantproductPublishRequest struct {
    /*
        叶子类目ID     */
    CatId  *int64 `json:"cat_id" required:"true" `
    /*
        schema协议值     */
    Product  *string `json:"product" required:"true" `
}

func (s *AlibabaProductMerchantproductPublishRequest) SetCatId(v int64) *AlibabaProductMerchantproductPublishRequest {
    s.CatId = &v
    return s
}
func (s *AlibabaProductMerchantproductPublishRequest) SetProduct(v string) *AlibabaProductMerchantproductPublishRequest {
    s.Product = &v
    return s
}

func (req *AlibabaProductMerchantproductPublishRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CatId != nil) {
        paramMap["cat_id"] = *req.CatId
    }
    if(req.Product != nil) {
        paramMap["product"] = *req.Product
    }
    return paramMap
}

func (req *AlibabaProductMerchantproductPublishRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}