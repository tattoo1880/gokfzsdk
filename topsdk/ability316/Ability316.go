package ability316

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability316/request"
    "topsdk/ability316/response"
    "topsdk/util"
)

type Ability316 struct {
    Client *topsdk.TopClient
}

func NewAbility316(client *topsdk.TopClient) *Ability316{
    return &Ability316{client}
}

/*
    判断用户是否收藏某个店铺
*/
func (ability *Ability316) AlibabaInteractOpenIsattention(req *request.AlibabaInteractOpenIsattentionRequest) (*response.AlibabaInteractOpenIsattentionResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability316 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.interact.open.isattention",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaInteractOpenIsattentionResponse{}
    if(err != nil){
        log.Println("alibabaInteractOpenIsattention error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    无线开放图片内容安全检查
*/
func (ability *Ability316) TaobaoWirelessPictureCheck(req *request.TaobaoWirelessPictureCheckRequest) (*response.TaobaoWirelessPictureCheckResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability316 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.wireless.picture.check",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoWirelessPictureCheckResponse{}
    if(err != nil){
        log.Println("taobaoWirelessPictureCheck error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    无线开放内容检查
*/
func (ability *Ability316) TaobaoWirelessContentCheck(req *request.TaobaoWirelessContentCheckRequest) (*response.TaobaoWirelessContentCheckResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability316 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.wireless.content.check",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoWirelessContentCheckResponse{}
    if(err != nil){
        log.Println("taobaoWirelessContentCheck error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    无线开放视频内容安全检查
*/
func (ability *Ability316) TaobaoWirelessVideoCheck(req *request.TaobaoWirelessVideoCheckRequest) (*response.TaobaoWirelessVideoCheckResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability316 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.wireless.video.check",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoWirelessVideoCheckResponse{}
    if(err != nil){
        log.Println("taobaoWirelessVideoCheck error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    回传抽奖相关参数
*/
func (ability *Ability316) AlibabaInteractLotteryactivityRegister(req *request.AlibabaInteractLotteryactivityRegisterRequest,session string) (*response.AlibabaInteractLotteryactivityRegisterResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability316 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.interact.lotteryactivity.register",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaInteractLotteryactivityRegisterResponse{}
    if(err != nil){
        log.Println("alibabaInteractLotteryactivityRegister error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    资源位数据推送接口
*/
func (ability *Ability316) AlibabaInteractAopdataRegister(req *request.AlibabaInteractAopdataRegisterRequest,session string) (*response.AlibabaInteractAopdataRegisterResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability316 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.interact.aopdata.register",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaInteractAopdataRegisterResponse{}
    if(err != nil){
        log.Println("alibabaInteractAopdataRegister error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    无线端抽奖接口
*/
func (ability *Ability316) AlibabaInteractLotterydrawDodraw(req *request.AlibabaInteractLotterydrawDodrawRequest,session string) (*response.AlibabaInteractLotterydrawDodrawResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability316 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.interact.lotterydraw.dodraw",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaInteractLotterydrawDodrawResponse{}
    if(err != nil){
        log.Println("alibabaInteractLotterydrawDodraw error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
