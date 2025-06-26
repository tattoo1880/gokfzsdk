package request


type AlibabaItemPublishPropsDistributeGetRequest struct {
    /*
        商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版     */
    Market  *string `json:"market" required:"true" `
    /*
        商品类目ID     */
    CatId  *int64 `json:"cat_id" required:"true" `
    /*
        商品条码     */
    Barcode  *string `json:"barcode,omitempty" required:"false" `
    /*
        类目属性渲染schema     */
    Schema  *string `json:"schema" required:"true" `
    /*
        属性ID     */
    PropId  *int64 `json:"prop_id" required:"true" `
}

func (s *AlibabaItemPublishPropsDistributeGetRequest) SetMarket(v string) *AlibabaItemPublishPropsDistributeGetRequest {
    s.Market = &v
    return s
}
func (s *AlibabaItemPublishPropsDistributeGetRequest) SetCatId(v int64) *AlibabaItemPublishPropsDistributeGetRequest {
    s.CatId = &v
    return s
}
func (s *AlibabaItemPublishPropsDistributeGetRequest) SetBarcode(v string) *AlibabaItemPublishPropsDistributeGetRequest {
    s.Barcode = &v
    return s
}
func (s *AlibabaItemPublishPropsDistributeGetRequest) SetSchema(v string) *AlibabaItemPublishPropsDistributeGetRequest {
    s.Schema = &v
    return s
}
func (s *AlibabaItemPublishPropsDistributeGetRequest) SetPropId(v int64) *AlibabaItemPublishPropsDistributeGetRequest {
    s.PropId = &v
    return s
}

func (req *AlibabaItemPublishPropsDistributeGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Market != nil) {
        paramMap["market"] = *req.Market
    }
    if(req.CatId != nil) {
        paramMap["cat_id"] = *req.CatId
    }
    if(req.Barcode != nil) {
        paramMap["barcode"] = *req.Barcode
    }
    if(req.Schema != nil) {
        paramMap["schema"] = *req.Schema
    }
    if(req.PropId != nil) {
        paramMap["prop_id"] = *req.PropId
    }
    return paramMap
}

func (req *AlibabaItemPublishPropsDistributeGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}