package request

import (
        "topsdk/defaultability/domain"
        "topsdk/util"
    )

type AlibabaItemPublishDistributeSubmitRequest struct {
    /*
        业务扩展参数，需与平台约定好     */
    BizType  *string `json:"biz_type,omitempty" required:"false" `
    /*
        商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版     */
    Market  *string `json:"market" required:"true" `
    /*
        商品类目ID     */
    CatId  *int64 `json:"cat_id" required:"true" `
    /*
        产品ID，天猫市场(market=tmall)时必填     */
    SpuId  *int64 `json:"spu_id,omitempty" required:"false" `
    /*
        商品条码     */
    Barcode  *string `json:"barcode,omitempty" required:"false" `
    /*
        商品schema信息，通过alibaba.item.publish.props.get获取并补全后提交     */
    Schema  *string `json:"schema" required:"true" `
    /*
        上游商品信息     */
    SourceItemInfo  *domain.AlibabaItemPublishDistributeSubmitSourceItemInfoForDistributionDTO `json:"source_item_info" required:"true" `
    /*
        扩展属性     */
    BizExt  *map[string]interface{} `json:"biz_ext,omitempty" required:"false" `
    /*
        产品ID集合，等同于产品ID     */
    SpuIds  *[]int64 `json:"spu_ids,omitempty" required:"false" `
}

func (s *AlibabaItemPublishDistributeSubmitRequest) SetBizType(v string) *AlibabaItemPublishDistributeSubmitRequest {
    s.BizType = &v
    return s
}
func (s *AlibabaItemPublishDistributeSubmitRequest) SetMarket(v string) *AlibabaItemPublishDistributeSubmitRequest {
    s.Market = &v
    return s
}
func (s *AlibabaItemPublishDistributeSubmitRequest) SetCatId(v int64) *AlibabaItemPublishDistributeSubmitRequest {
    s.CatId = &v
    return s
}
func (s *AlibabaItemPublishDistributeSubmitRequest) SetSpuId(v int64) *AlibabaItemPublishDistributeSubmitRequest {
    s.SpuId = &v
    return s
}
func (s *AlibabaItemPublishDistributeSubmitRequest) SetBarcode(v string) *AlibabaItemPublishDistributeSubmitRequest {
    s.Barcode = &v
    return s
}
func (s *AlibabaItemPublishDistributeSubmitRequest) SetSchema(v string) *AlibabaItemPublishDistributeSubmitRequest {
    s.Schema = &v
    return s
}
func (s *AlibabaItemPublishDistributeSubmitRequest) SetSourceItemInfo(v domain.AlibabaItemPublishDistributeSubmitSourceItemInfoForDistributionDTO) *AlibabaItemPublishDistributeSubmitRequest {
    s.SourceItemInfo = &v
    return s
}
func (s *AlibabaItemPublishDistributeSubmitRequest) SetBizExt(v map[string]interface{}) *AlibabaItemPublishDistributeSubmitRequest {
    s.BizExt = &v
    return s
}
func (s *AlibabaItemPublishDistributeSubmitRequest) SetSpuIds(v []int64) *AlibabaItemPublishDistributeSubmitRequest {
    s.SpuIds = &v
    return s
}

func (req *AlibabaItemPublishDistributeSubmitRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BizType != nil) {
        paramMap["biz_type"] = *req.BizType
    }
    if(req.Market != nil) {
        paramMap["market"] = *req.Market
    }
    if(req.CatId != nil) {
        paramMap["cat_id"] = *req.CatId
    }
    if(req.SpuId != nil) {
        paramMap["spu_id"] = *req.SpuId
    }
    if(req.Barcode != nil) {
        paramMap["barcode"] = *req.Barcode
    }
    if(req.Schema != nil) {
        paramMap["schema"] = *req.Schema
    }
    if(req.SourceItemInfo != nil) {
        paramMap["source_item_info"] = util.ConvertStruct(*req.SourceItemInfo)
    }
    if(req.BizExt != nil) {
        paramMap["biz_ext"] = util.ConvertStruct(*req.BizExt)
    }
    if(req.SpuIds != nil) {
        paramMap["spu_ids"] = util.ConvertBasicList(*req.SpuIds)
    }
    return paramMap
}

func (req *AlibabaItemPublishDistributeSubmitRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}