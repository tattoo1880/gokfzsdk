package request


type AlibabaProductMerchantproductsSearchRequest struct {
    /*
        关键词，建议填条形码     */
    Keyword  *string `json:"keyword" required:"true" `
    /*
        类目id     */
    CategoryId  *int64 `json:"category_id,omitempty" required:"false" `
    /*
        品牌ID     */
    BrandId  *int64 `json:"brand_id,omitempty" required:"false" `
}

func (s *AlibabaProductMerchantproductsSearchRequest) SetKeyword(v string) *AlibabaProductMerchantproductsSearchRequest {
    s.Keyword = &v
    return s
}
func (s *AlibabaProductMerchantproductsSearchRequest) SetCategoryId(v int64) *AlibabaProductMerchantproductsSearchRequest {
    s.CategoryId = &v
    return s
}
func (s *AlibabaProductMerchantproductsSearchRequest) SetBrandId(v int64) *AlibabaProductMerchantproductsSearchRequest {
    s.BrandId = &v
    return s
}

func (req *AlibabaProductMerchantproductsSearchRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Keyword != nil) {
        paramMap["keyword"] = *req.Keyword
    }
    if(req.CategoryId != nil) {
        paramMap["category_id"] = *req.CategoryId
    }
    if(req.BrandId != nil) {
        paramMap["brand_id"] = *req.BrandId
    }
    return paramMap
}

func (req *AlibabaProductMerchantproductsSearchRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}