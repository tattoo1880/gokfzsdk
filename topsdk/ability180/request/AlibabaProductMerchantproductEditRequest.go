package request


type AlibabaProductMerchantproductEditRequest struct {
    /*
        产品ID     */
    Id  *int64 `json:"id" required:"true" `
    /*
        schema协议内容     */
    Product  *string `json:"product" required:"true" `
}

func (s *AlibabaProductMerchantproductEditRequest) SetId(v int64) *AlibabaProductMerchantproductEditRequest {
    s.Id = &v
    return s
}
func (s *AlibabaProductMerchantproductEditRequest) SetProduct(v string) *AlibabaProductMerchantproductEditRequest {
    s.Product = &v
    return s
}

func (req *AlibabaProductMerchantproductEditRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Id != nil) {
        paramMap["id"] = *req.Id
    }
    if(req.Product != nil) {
        paramMap["product"] = *req.Product
    }
    return paramMap
}

func (req *AlibabaProductMerchantproductEditRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}