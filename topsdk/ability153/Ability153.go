package ability153

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability153/request"
    "topsdk/ability153/response"
    "topsdk/util"
)

type Ability153 struct {
    Client *topsdk.TopClient
}

func NewAbility153(client *topsdk.TopClient) *Ability153{
    return &Ability153{client}
}

/*
    派件通知接口
*/
func (ability *Ability153) CainiaoWaybillIiDelivery(req *request.CainiaoWaybillIiDeliveryRequest,session string) (*response.CainiaoWaybillIiDeliveryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability153 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("cainiao.waybill.ii.delivery",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.CainiaoWaybillIiDeliveryResponse{}
    if(err != nil){
        log.Println("cainiaoWaybillIiDelivery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    是否派送可达判定批量查询接口
*/
func (ability *Ability153) CainiaoReachableBatchjudge(req *request.CainiaoReachableBatchjudgeRequest,session string) (*response.CainiaoReachableBatchjudgeResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability153 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("cainiao.reachable.batchjudge",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.CainiaoReachableBatchjudgeResponse{}
    if(err != nil){
        log.Println("cainiaoReachableBatchjudge error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    物流订单确认接口
*/
func (ability *Ability153) CainiaoWaybillIiConfirm(req *request.CainiaoWaybillIiConfirmRequest,session string) (*response.CainiaoWaybillIiConfirmResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability153 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("cainiao.waybill.ii.confirm",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.CainiaoWaybillIiConfirmResponse{}
    if(err != nil){
        log.Println("cainiaoWaybillIiConfirm error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    电子面单物流详情授权url获取
*/
func (ability *Ability153) CainiaoWaybillIiLogisticsdetailUrlGet(req *request.CainiaoWaybillIiLogisticsdetailUrlGetRequest) (*response.CainiaoWaybillIiLogisticsdetailUrlGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability153 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("cainiao.waybill.ii.logisticsdetail.url.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.CainiaoWaybillIiLogisticsdetailUrlGetResponse{}
    if(err != nil){
        log.Println("cainiaoWaybillIiLogisticsdetailUrlGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    隐私面单商家订购查询
*/
func (ability *Ability153) CainiaoWaybillPrivacySubscriptionGet(req *request.CainiaoWaybillPrivacySubscriptionGetRequest,session string) (*response.CainiaoWaybillPrivacySubscriptionGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability153 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("cainiao.waybill.privacy.subscription.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.CainiaoWaybillPrivacySubscriptionGetResponse{}
    if(err != nil){
        log.Println("cainiaoWaybillPrivacySubscriptionGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    隐私面单商家订单查询
*/
func (ability *Ability153) CainiaoWaybillPrivacySellerOrderGet(req *request.CainiaoWaybillPrivacySellerOrderGetRequest,session string) (*response.CainiaoWaybillPrivacySellerOrderGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability153 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("cainiao.waybill.privacy.seller.order.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.CainiaoWaybillPrivacySellerOrderGetResponse{}
    if(err != nil){
        log.Println("cainiaoWaybillPrivacySellerOrderGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
