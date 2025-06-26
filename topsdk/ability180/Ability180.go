package ability180

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability180/request"
    "topsdk/ability180/response"
    "topsdk/util"
)

type Ability180 struct {
    Client *topsdk.TopClient
}

func NewAbility180(client *topsdk.TopClient) *Ability180{
    return &Ability180{client}
}

/*
    获取编辑商家标准产品规则
*/
func (ability *Ability180) AlibabaProductMerchantproductEditGet(req *request.AlibabaProductMerchantproductEditGetRequest,session string) (*response.AlibabaProductMerchantproductEditGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability180 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.product.merchantproduct.edit.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaProductMerchantproductEditGetResponse{}
    if(err != nil){
        log.Println("alibabaProductMerchantproductEditGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商家标准产品搜索
*/
func (ability *Ability180) AlibabaProductMerchantproductsSearch(req *request.AlibabaProductMerchantproductsSearchRequest,session string) (*response.AlibabaProductMerchantproductsSearchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability180 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.product.merchantproducts.search",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaProductMerchantproductsSearchResponse{}
    if(err != nil){
        log.Println("alibabaProductMerchantproductsSearch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    新发商家标准产品
*/
func (ability *Ability180) AlibabaProductMerchantproductPublish(req *request.AlibabaProductMerchantproductPublishRequest,session string) (*response.AlibabaProductMerchantproductPublishResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability180 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.product.merchantproduct.publish",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaProductMerchantproductPublishResponse{}
    if(err != nil){
        log.Println("alibabaProductMerchantproductPublish error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    编辑商家标准产品
*/
func (ability *Ability180) AlibabaProductMerchantproductEdit(req *request.AlibabaProductMerchantproductEditRequest,session string) (*response.AlibabaProductMerchantproductEditResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability180 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.product.merchantproduct.edit",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaProductMerchantproductEditResponse{}
    if(err != nil){
        log.Println("alibabaProductMerchantproductEdit error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询商家标准产品新发规则
*/
func (ability *Ability180) AlibabaProductMerchantproductPublishGet(req *request.AlibabaProductMerchantproductPublishGetRequest,session string) (*response.AlibabaProductMerchantproductPublishGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability180 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.product.merchantproduct.publish.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaProductMerchantproductPublishGetResponse{}
    if(err != nil){
        log.Println("alibabaProductMerchantproductPublishGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
