package ability648

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability648/request"
    "topsdk/ability648/response"
    "topsdk/util"
)

type Ability648 struct {
    Client *topsdk.TopClient
}

func NewAbility648(client *topsdk.TopClient) *Ability648{
    return &Ability648{client}
}

/*
    添加SKU
*/
func (ability *Ability648) TaobaoItemSkuAdd(req *request.TaobaoItemSkuAddRequest,session string) (*response.TaobaoItemSkuAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.sku.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemSkuAddResponse{}
    if(err != nil){
        log.Println("taobaoItemSkuAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    更新SKU信息
*/
func (ability *Ability648) TaobaoItemSkuUpdate(req *request.TaobaoItemSkuUpdateRequest,session string) (*response.TaobaoItemSkuUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.sku.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemSkuUpdateResponse{}
    if(err != nil){
        log.Println("taobaoItemSkuUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    限时打折详情查询
*/
func (ability *Ability648) TaobaoPromotionLimitdiscountDetailGet(req *request.TaobaoPromotionLimitdiscountDetailGetRequest,session string) (*response.TaobaoPromotionLimitdiscountDetailGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.promotion.limitdiscount.detail.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPromotionLimitdiscountDetailGetResponse{}
    if(err != nil){
        log.Println("taobaoPromotionLimitdiscountDetailGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    定向优惠活动名称与描述违禁词检查
*/
func (ability *Ability648) TaobaoMarketingPromotionKfc(req *request.TaobaoMarketingPromotionKfcRequest,session string) (*response.TaobaoMarketingPromotionKfcResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.marketing.promotion.kfc",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoMarketingPromotionKfcResponse{}
    if(err != nil){
        log.Println("taobaoMarketingPromotionKfc error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取用户指定运费模板信息
*/
func (ability *Ability648) TaobaoDeliveryTemplateGet(req *request.TaobaoDeliveryTemplateGetRequest,session string) (*response.TaobaoDeliveryTemplateGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.delivery.template.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoDeliveryTemplateGetResponse{}
    if(err != nil){
        log.Println("taobaoDeliveryTemplateGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取用户下所有模板
*/
func (ability *Ability648) TaobaoDeliveryTemplatesGet(req *request.TaobaoDeliveryTemplatesGetRequest,session string) (*response.TaobaoDeliveryTemplatesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.delivery.templates.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoDeliveryTemplatesGetResponse{}
    if(err != nil){
        log.Println("taobaoDeliveryTemplatesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    删除运费模板
*/
func (ability *Ability648) TaobaoDeliveryTemplateDelete(req *request.TaobaoDeliveryTemplateDeleteRequest,session string) (*response.TaobaoDeliveryTemplateDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.delivery.template.delete",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoDeliveryTemplateDeleteResponse{}
    if(err != nil){
        log.Println("taobaoDeliveryTemplateDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    新增运费模板
*/
func (ability *Ability648) TaobaoDeliveryTemplateAdd(req *request.TaobaoDeliveryTemplateAddRequest,session string) (*response.TaobaoDeliveryTemplateAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.delivery.template.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoDeliveryTemplateAddResponse{}
    if(err != nil){
        log.Println("taobaoDeliveryTemplateAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    修改运费模板
*/
func (ability *Ability648) TaobaoDeliveryTemplateUpdate(req *request.TaobaoDeliveryTemplateUpdateRequest,session string) (*response.TaobaoDeliveryTemplateUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.delivery.template.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoDeliveryTemplateUpdateResponse{}
    if(err != nil){
        log.Println("taobaoDeliveryTemplateUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    更新商品SKU的价格
*/
func (ability *Ability648) TaobaoItemSkuPriceUpdate(req *request.TaobaoItemSkuPriceUpdateRequest,session string) (*response.TaobaoItemSkuPriceUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.sku.price.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemSkuPriceUpdateResponse{}
    if(err != nil){
        log.Println("taobaoItemSkuPriceUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
